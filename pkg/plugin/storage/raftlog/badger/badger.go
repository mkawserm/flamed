package badger

import (
	"bytes"
	badgerDb "github.com/dgraph-io/badger/v2"
	badgerDbOptions "github.com/dgraph-io/badger/v2/options"
	"github.com/mkawserm/flamed/pkg/iface"
	"github.com/mkawserm/flamed/pkg/logger"
	"time"
)

const MaxKeyLength = 1024

type Transaction struct {
	mDb    *badgerDb.DB
	mTxn   *badgerDb.Txn
	mCount int
}

func (t *Transaction) Clear() error {
	t.mTxn.Discard()
	t.mTxn = t.mDb.NewTransaction(true)
	t.mCount = 0
	return nil
}

func (t *Transaction) Count() int {
	return t.mCount
}

func (t *Transaction) Commit() error {
	return t.mTxn.Commit()
}

func (t *Transaction) Put(key []byte, val []byte) error {
	if err := t.mTxn.Set(key, val); err == badgerDb.ErrTxnTooBig {
		if err := t.mTxn.Commit(); err != nil {
			return err
		} else {
			t.mCount = 0
			return nil
		}
	} else if err == nil {
		t.mCount = t.mCount + 1
		return nil
	} else {
		return err
	}
}

func (t *Transaction) Delete(key []byte) error {
	if err := t.mTxn.Delete(key); err == badgerDb.ErrTxnTooBig {
		if err := t.mTxn.Commit(); err != nil {
			return err
		} else {
			t.mCount = 0
			return nil
		}
	} else if err == nil {
		t.mCount = t.mCount + 1
		return nil
	} else {
		return err
	}
}

func (t *Transaction) Destroy() error {
	t.mTxn.Discard()
	return nil
}

type Badger struct {
	mDb *badgerDb.DB

	//mDir       string
	//mWalDir    string
	//mSecretKey []byte

	//mOpts badgerDb.Options
}

func (b *Badger) Open(dir, valueDir string, secretKey []byte, _ interface{}) error {
	if b.mDb != nil {
		return nil
	}

	opts := badgerDb.DefaultOptions(dir)
	opts.Dir = dir
	opts.ValueDir = valueDir
	opts.EncryptionKey = secretKey
	opts.Truncate = true
	opts.TableLoadingMode = badgerDbOptions.LoadToRAM
	opts.ValueLogLoadingMode = badgerDbOptions.MemoryMap
	opts.Compression = badgerDbOptions.Snappy
	opts.Logger = logger.S("raft-log-badger")

	db, err := badgerDb.Open(opts)

	if err != nil {
		return nil
	} else {
		b.mDb = db
		return nil
	}
}

func (b *Badger) Name() string {
	return "badger"
}

func (b *Badger) Close() error {
	if b.mDb == nil {
		return nil
	}
	_ = b.mDb.Close()
	b.mDb = nil
	return nil
}

func (b *Badger) IterateValue(fk []byte, lk []byte, inc bool, op func(key []byte, data []byte) (bool, error)) error {
	err := b.mDb.View(func(txn *badgerDb.Txn) error {
		it := txn.NewIterator(badgerDb.DefaultIteratorOptions)
		defer it.Close()

		for it.Seek(fk); it.Valid(); it.Next() {
			item := it.Item()
			key := item.Key()
			val, err := item.ValueCopy(nil)
			if err != nil {
				return err
			}

			if inc {
				if bytes.Compare(key, lk) > 0 {
					return nil
				}
			} else {
				if bytes.Compare(key, lk) >= 0 {
					return nil
				}
			}

			cont, err := op(key, val)
			if err != nil {
				return err
			}
			if !cont {
				break
			}
		}
		return nil
	})

	return err
}

func (b *Badger) GetValue(key []byte, op func([]byte) error) error {
	err := b.mDb.View(func(txn *badgerDb.Txn) error {
		item, err := txn.Get(key)

		if err != nil && err != badgerDb.ErrKeyNotFound {
			return err
		}
		if item == nil {
			return nil
		}

		val, err := item.ValueCopy(nil)
		if err != nil {
			return err
		}
		return op(val)
	})

	return err
}

func (b *Badger) SaveValue(key []byte, value []byte) error {
	err := b.mDb.Update(func(txn *badgerDb.Txn) error {
		return txn.Set(key, value)
	})

	return err
}

func (b *Badger) DeleteValue(key []byte) error {
	err := b.mDb.Update(func(txn *badgerDb.Txn) error {
		return txn.Delete(key)
	})

	return err
}

func (b *Badger) GetWriteBatch() iface.ITransaction {
	return &Transaction{
		mDb:    b.mDb,
		mTxn:   b.mDb.NewTransaction(true),
		mCount: 0,
	}
}

func (b *Badger) CommitWriteBatch(wb iface.ITransaction) error {
	return wb.Commit()
}

func (b *Badger) BulkRemoveEntries(firstKey []byte, lastKey []byte) error {
	return b.CompactEntries(firstKey, lastKey)
}

func (b *Badger) CompactEntries(firstKey []byte, lastKey []byte) error {
	err := b.mDb.Update(func(txn *badgerDb.Txn) error {
		opts := badgerDb.DefaultIteratorOptions
		opts.PrefetchValues = false
		it := txn.NewIterator(opts)
		defer it.Close()

		for it.Seek(firstKey); it.Valid(); it.Next() {
			item := it.Item()
			key := item.Key()
			if bytes.Compare(key, lastKey) >= 0 {
				break
			}
			if err := txn.Delete(key); err == badgerDb.ErrTxnTooBig {
				if err := txn.Commit(); err != nil {
					return err
				}
			}
		}
		return nil
	})

	if err != nil {
		b.RunGC()
	}

	return err
}

func (b *Badger) FullCompaction() error {
	fk := make([]byte, MaxKeyLength)
	lk := make([]byte, MaxKeyLength)
	for i := uint64(0); i < MaxKeyLength; i++ {
		fk[i] = 0
		lk[i] = 0xFF
	}

	return b.CompactEntries(fk, lk)
}

func (b *Badger) RunGC() {
	if b.mDb == nil {
		return
	}

	ticker := time.NewTicker(5 * time.Minute)
	defer ticker.Stop()

	for range ticker.C {
	again:
		if b.mDb == nil {
			return
		}

		err := b.mDb.RunValueLogGC(0.5)
		if err == nil {
			goto again
		}
	}
}
