package badger

import (
	"context"
	badgerDb "github.com/dgraph-io/badger/v2"
	badgerDbOptions "github.com/dgraph-io/badger/v2/options"
	"github.com/mkawserm/flamed/pkg/pb"
	"github.com/mkawserm/flamed/pkg/uidutil"
	"github.com/mkawserm/flamed/pkg/x"
	"time"
)

type Badger struct {
	mDbPath          string
	mDb              *badgerDb.DB
	mDbConfiguration Configuration
}

func (b *Badger) Open(path string, secretKey []byte, readOnly bool, configuration interface{}) (bool, error) {
	if b.mDb != nil {
		return true, nil
	}

	b.mDbPath = path
	b.mDbConfiguration.BadgerOptions = badgerDb.DefaultOptions(b.mDbPath)

	if configuration == nil {
		b.mDbConfiguration.BadgerOptions.Truncate = true
		b.mDbConfiguration.BadgerOptions.TableLoadingMode = badgerDbOptions.LoadToRAM
		b.mDbConfiguration.BadgerOptions.ValueLogLoadingMode = badgerDbOptions.MemoryMap
		b.mDbConfiguration.BadgerOptions.Compression = badgerDbOptions.Snappy
	} else {
		if opts, ok := configuration.(badgerDb.Options); ok {
			b.mDbConfiguration.BadgerOptions = opts
			b.mDbConfiguration.BadgerOptions.ValueDir = path
			b.mDbConfiguration.BadgerOptions.Dir = path
		}
	}

	b.mDbConfiguration.BadgerOptions.ReadOnly = readOnly
	b.mDbConfiguration.BadgerOptions.EncryptionKey = secretKey

	if b.mDbConfiguration.GoroutineNumber <= 0 {
		b.mDbConfiguration.GoroutineNumber = 16
	}

	if b.mDbConfiguration.LogPrefix == "" {
		b.mDbConfiguration.LogPrefix = "Flamed:Async.Snapshot"
	}

	if b.mDbConfiguration.SliceCap <= 0 {
		b.mDbConfiguration.SliceCap = 100
	}

	if b.mDbConfiguration.EncryptionKeyRotationDuration <= 0 {
		b.mDbConfiguration.EncryptionKeyRotationDuration = 10 * 24 * time.Hour
	}

	db, err := badgerDb.Open(b.mDbConfiguration.BadgerOptions)
	if err != nil {
		return false, x.ErrFailedToOpenStorage
	}
	b.mDb = db
	return true, nil
}

func (b *Badger) Close() error {
	b.mDbPath = ""

	if b.mDb == nil {
		return nil
	}

	err := b.mDb.Close()
	b.mDb = nil
	if err != nil {
		return x.ErrFailedToCloseStorage
	}

	return nil
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

func (b *Badger) ChangeSecretKey(oldSecretKey []byte, newSecretKey []byte) (bool, error) {
	if b.mDb == nil {
		return false, x.ErrFailedToChangeSecretKey
	}

	//if len(oldSecretKey) == 0 || len(newSecretKey) == 0 {
	//	return false, x.ErrFailedToChangeSecretKey
	//}

	opt := badgerDb.KeyRegistryOptions{
		Dir:                           b.mDbPath,
		ReadOnly:                      true,
		EncryptionKey:                 oldSecretKey,
		EncryptionKeyRotationDuration: b.mDbConfiguration.EncryptionKeyRotationDuration,
	}

	kr, err := badgerDb.OpenKeyRegistry(opt)
	if err != nil {
		return false, x.ErrFailedToChangeSecretKey
	}

	opt.EncryptionKey = newSecretKey

	err = badgerDb.WriteKeyRegistry(kr, opt)
	if err != nil {
		return false, x.ErrFailedToChangeSecretKey
	}

	return true, nil
}

func (b *Badger) ReadUsingPrefix(prefix []byte) ([]*pb.FlameEntry, error) {
	if b.mDb == nil {
		return nil, x.ErrFailedToReadDataFromStorage
	}

	data := make([]*pb.FlameEntry, 0, b.mDbConfiguration.SliceCap)

	err := b.mDb.View(func(txn *badgerDb.Txn) error {
		it := txn.NewIterator(badgerDb.DefaultIteratorOptions)
		defer it.Close()

		for it.Seek(prefix); it.ValidForPrefix(prefix); it.Next() {
			item := it.Item()
			uid := item.Key()

			if value, err := item.ValueCopy(nil); err == nil {
				namespace, key := uidutil.SplitUid(uid)
				entry := &pb.FlameEntry{
					Namespace: namespace,
					Key:       key,
					Value:     value,
				}
				data = append(data, entry)
			} else {
				return err
			}
		}
		return nil
	})

	if err != nil {
		return nil, x.ErrFailedToReadDataFromStorage
	}

	return data, nil
}

func (b *Badger) Read(namespace []byte, key []byte) ([]byte, error) {
	if b.mDb == nil {
		return nil, x.ErrFailedToReadDataFromStorage
	}

	uid := uidutil.GetUid(namespace, key)
	var data []byte

	err := b.mDb.View(func(txn *badgerDb.Txn) error {
		if item, err := txn.Get(uid); err != nil {
			return err
		} else {
			value, err2 := item.ValueCopy(nil)
			data = value
			return err2
		}
	})

	if err != nil {
		if err == badgerDb.ErrKeyNotFound {
			return nil, x.ErrUidDoesNotExists
		} else {
			return nil, x.ErrFailedToReadDataFromStorage
		}
	}

	return data, nil
}

func (b *Badger) Delete(namespace []byte, key []byte) (bool, error) {
	if b.mDb == nil {
		return false, x.ErrFailedToDeleteDataFromStorage
	}

	uid := uidutil.GetUid(namespace, key)

	err := b.mDb.Update(func(txn *badgerDb.Txn) error {
		err := txn.Delete(uid)
		return err
	})

	if err != nil {
		return false, x.ErrFailedToDeleteDataFromStorage
	}

	return true, nil
}

func (b *Badger) Create(namespace []byte, key []byte, value []byte) (bool, error) {
	if b.mDb == nil {
		return false, x.ErrFailedToCreateDataToStorage
	}

	uid := uidutil.GetUid(namespace, key)

	err := b.mDb.Update(func(txn *badgerDb.Txn) error {
		err := txn.Set(uid, value)
		return err
	})

	if err != nil {
		return false, x.ErrFailedToCreateDataToStorage
	}

	return true, nil
}

func (b *Badger) Update(namespace []byte, key []byte, value []byte) (bool, error) {
	if b.mDb == nil {
		return false, x.ErrFailedToUpdateDataToStorage
	}

	uid := uidutil.GetUid(namespace, key)

	err := b.mDb.Update(func(txn *badgerDb.Txn) error {
		err := txn.Set(uid, value)
		return err
	})

	if err != nil {
		return false, x.ErrFailedToUpdateDataToStorage
	}

	return true, nil
}

func (b *Badger) IsExists(namespace []byte, key []byte) bool {
	if b.mDb == nil {
		return false
	}

	uid := uidutil.GetUid(namespace, key)

	err := b.mDb.View(func(txn *badgerDb.Txn) error {
		_, err := txn.Get(uid)
		return err
	})

	if err == badgerDb.ErrKeyNotFound {
		return false
	} else if err == nil {
		return true
	} else {
		return false
	}
}

func (b *Badger) ApplyBatch(batch *pb.FlameBatch) (bool, error) {
	if b.mDb == nil {
		return false, x.ErrFailedToApplyBatchToStorage
	}

	txn := b.mDb.NewTransaction(true)
	defer txn.Discard()

	for _, action := range batch.FlameActionList {
		if action.FlameActionType == pb.FlameAction_CREATE || action.FlameActionType == pb.FlameAction_UPDATE {
			uid := uidutil.GetUid(action.FlameEntry.Namespace, action.FlameEntry.Key)
			if err := txn.Set(uid, action.FlameEntry.Value); err != nil {
				return false, x.ErrFailedToApplyBatchToStorage
			}
		} else if action.FlameActionType == pb.FlameAction_DELETE {
			uid := uidutil.GetUid(action.FlameEntry.Namespace, action.FlameEntry.Key)
			if err := txn.Delete(uid); err != nil {
				return false, x.ErrFailedToApplyBatchToStorage
			}
		}
	}

	if err := txn.Commit(); err != nil {
		return false, x.ErrFailedToApplyBatchToStorage
	}

	return true, nil
}

func (b *Badger) ApplyAction(action *pb.FlameAction) (bool, error) {
	if action.FlameActionType == pb.FlameAction_CREATE {
		return b.Create(action.FlameEntry.Namespace, action.FlameEntry.Key, action.FlameEntry.Value)
	} else if action.FlameActionType == pb.FlameAction_UPDATE {
		return b.Update(action.FlameEntry.Namespace, action.FlameEntry.Key, action.FlameEntry.Value)
	} else if action.FlameActionType == pb.FlameAction_DELETE {
		return b.Delete(action.FlameEntry.Namespace, action.FlameEntry.Key)
	}

	return false, x.ErrFailedToApplyActionToStorage
}

func (b *Badger) AsyncSnapshot(snapshot chan<- *pb.FlameSnapshot) error {
	if b.mDb == nil {
		return x.ErrFailedToGenerateAsyncSnapshotFromStorage
	}

	stream := b.mDb.NewStream()
	stream.NumGo = b.mDbConfiguration.GoroutineNumber
	stream.LogPrefix = b.mDbConfiguration.LogPrefix
	stream.KeyToList = nil

	stream.Send = func(list *badgerDb.KVList) error {
		data := &pb.FlameSnapshot{
			Version:                0,
			Length:                 0,
			FlameSnapshotEntryList: make([]*pb.FlameSnapshotEntry, 0, b.mDbConfiguration.SliceCap),
		}

		for _, kv := range list.Kv {
			entry := &pb.FlameSnapshotEntry{
				Uid:  kv.Key,
				Data: kv.Value,
			}
			data.FlameSnapshotEntryList = append(data.FlameSnapshotEntryList, entry)
		}

		data.Length = uint64(len(data.FlameSnapshotEntryList))
		snapshot <- data
		return nil
	}

	if err := stream.Orchestrate(context.Background()); err != nil {
		return err
	}

	return nil
}

func (b *Badger) ApplyAsyncSnapshot(snapshot <-chan *pb.FlameSnapshot) (bool, error) {
	if b.mDb == nil {
		return false, x.ErrFailedToApplyAsyncSnapshotToStorage
	}

	for {
		ss := <-snapshot
		if len(ss.FlameSnapshotEntryList) == 0 {
			break
		}

		txn := b.mDb.NewTransaction(true)
		for _, entry := range ss.FlameSnapshotEntryList {
			if err := txn.Set(entry.Uid, entry.Data); err == badgerDb.ErrTxnTooBig {
				if err := txn.Commit(); err != nil {
					return false, x.ErrFailedToApplyAsyncSnapshotToStorage
				}

				txn = b.mDb.NewTransaction(true)
				if err := txn.Set(entry.Uid, entry.Data); err != nil {
					return false, x.ErrFailedToApplyAsyncSnapshotToStorage
				}

			} else if err != nil {
				return false, x.ErrFailedToApplyAsyncSnapshotToStorage
			}
		}

		if err := txn.Commit(); err != nil {
			txn.Discard()
			return false, x.ErrFailedToApplyAsyncSnapshotToStorage
		}

		txn.Discard()
	}

	return true, nil
}

func (b *Badger) SyncSnapshot() (*pb.FlameSnapshot, error) {
	if b.mDb == nil {
		return nil, x.ErrFailedToGenerateSyncSnapshotFromStorage
	}

	snapshot := &pb.FlameSnapshot{
		Version:                0,
		Length:                 0,
		FlameSnapshotEntryList: make([]*pb.FlameSnapshotEntry, 0, b.mDbConfiguration.SliceCap),
	}

	err := b.mDb.View(func(txn *badgerDb.Txn) error {
		opts := badgerDb.DefaultIteratorOptions
		opts.PrefetchSize = 100
		it := txn.NewIterator(opts)
		defer it.Close()

		for it.Rewind(); it.Valid(); it.Next() {
			item := it.Item()
			if value, err := item.ValueCopy(nil); err != nil {
				return err
			} else {
				snapshot.FlameSnapshotEntryList = append(snapshot.FlameSnapshotEntryList, &pb.FlameSnapshotEntry{
					Uid:  item.Key(),
					Data: value,
				})
			}
		}
		return nil
	})

	if err != nil {
		return nil, x.ErrFailedToGenerateSyncSnapshotFromStorage
	}
	snapshot.Length = uint64(len(snapshot.FlameSnapshotEntryList))
	return snapshot, nil
}

func (b *Badger) ApplySyncSnapshot(snapshot *pb.FlameSnapshot) (bool, error) {
	if b.mDb == nil {
		return false, x.ErrFailedToApplySyncSnapshotToStorage
	}

	txn := b.mDb.NewTransaction(true)
	defer txn.Discard()

	for _, entry := range snapshot.FlameSnapshotEntryList {
		if err := txn.Set(entry.Uid, entry.Data); err == badgerDb.ErrTxnTooBig {
			if err := txn.Commit(); err != nil {
				return false, x.ErrFailedToApplySyncSnapshotToStorage
			}

			txn = b.mDb.NewTransaction(true)

			if err := txn.Set(entry.Uid, entry.Data); err != nil {
				return false, x.ErrFailedToApplySyncSnapshotToStorage
			}

		} else if err != nil {
			return false, x.ErrFailedToApplySyncSnapshotToStorage
		}
	}

	if err := txn.Commit(); err != nil {
		return false, x.ErrFailedToApplySyncSnapshotToStorage
	}

	return true, nil
}
