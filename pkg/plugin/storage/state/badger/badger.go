package badger

import (
	badgerDb "github.com/dgraph-io/badger/v2"
	badgerDbOptions "github.com/dgraph-io/badger/v2/options"
	"github.com/mkawserm/flamed/pkg/iface"
	"github.com/mkawserm/flamed/pkg/logger"
	"github.com/mkawserm/flamed/pkg/pb"
	"github.com/mkawserm/flamed/pkg/x"
	"go.uber.org/zap"
	"time"
)

type Iterator struct {
	mBadgerIterator *badgerDb.Iterator
}

func (i *Iterator) Next() {
	i.mBadgerIterator.Next()
}

func (i *Iterator) Rewind() {
	i.mBadgerIterator.Rewind()
}

func (i *Iterator) Close() {
	i.mBadgerIterator.Close()
}

func (i *Iterator) Valid() bool {
	return i.mBadgerIterator.Valid()
}

func (i *Iterator) Seek(address []byte) {
	i.mBadgerIterator.Seek(address)
}

func (i *Iterator) StateSnapshot() *pb.StateSnapshot {
	item := i.mBadgerIterator.Item()

	if item == nil {
		return nil
	}

	ss := &pb.StateSnapshot{}
	ss.Address = item.Key()

	if value, err := item.ValueCopy(nil); err == nil {
		ss.Data = value
	} else {
		internalLogger.Error(Name+" value copy error", zap.Error(err))
	}

	return ss
}

func (i *Iterator) ValidForPrefix(prefix []byte) bool {
	return i.mBadgerIterator.ValidForPrefix(prefix)
}

type Transaction struct {
	mBadgerTxn *badgerDb.Txn
	mCacheSize int
}

func (t *Transaction) Discard() {
	t.mBadgerTxn.Discard()
}

func (t *Transaction) Commit() error {
	return t.mBadgerTxn.Commit()
}

func (t *Transaction) Delete(address []byte) error {
	return t.mBadgerTxn.Delete(address)
}

func (t *Transaction) Get(address []byte) ([]byte, error) {
	item, err := t.mBadgerTxn.Get(address)

	if err == badgerDb.ErrKeyNotFound {
		return nil, x.ErrAddressNotFound
	}

	if item == nil {
		return nil, x.ErrAddressNotFound
	}

	if val, err := item.ValueCopy(nil); err == nil {
		return val, nil
	} else {
		return nil, err
	}
}

func (t *Transaction) Set(address []byte, data []byte) error {
	return t.mBadgerTxn.Set(address, data)
}

func (t *Transaction) ForwardIterator() iface.IStateStorageIterator {
	opts := badgerDb.DefaultIteratorOptions
	opts.PrefetchSize = t.mCacheSize
	opts.Reverse = false
	it := t.mBadgerTxn.NewIterator(opts)
	return &Iterator{mBadgerIterator: it}
}

func (t *Transaction) ReverseIterator() iface.IStateStorageIterator {
	opts := badgerDb.DefaultIteratorOptions
	opts.PrefetchSize = t.mCacheSize
	opts.Reverse = true
	it := t.mBadgerTxn.NewIterator(opts)
	return &Iterator{mBadgerIterator: it}
}

func (t *Transaction) KeyOnlyForwardIterator() iface.IStateStorageIterator {
	opts := badgerDb.DefaultIteratorOptions
	opts.PrefetchValues = false
	opts.Reverse = false
	it := t.mBadgerTxn.NewIterator(opts)
	return &Iterator{mBadgerIterator: it}
}

func (t *Transaction) KeyOnlyReverseIterator() iface.IStateStorageIterator {
	opts := badgerDb.DefaultIteratorOptions
	opts.PrefetchValues = false
	opts.Reverse = true
	it := t.mBadgerTxn.NewIterator(opts)
	return &Iterator{mBadgerIterator: it}
}

type Badger struct {
	mDbPath          string
	mSecretKey       []byte
	mDb              *badgerDb.DB
	mDbConfiguration Configuration
}

func (b *Badger) Setup(path string, secretKey []byte, configuration interface{}) {
	if b.mDb != nil {
		internalLogger.Debug("state::badger db already open")
		_ = internalLogger.Sync()
		return
	}

	b.mDbPath = path
	b.mSecretKey = secretKey
	b.mDbConfiguration.BadgerOptions = badgerDb.DefaultOptions(b.mDbPath)

	if configuration == nil {
		b.mDbConfiguration.BadgerOptions.Truncate = true
		b.mDbConfiguration.BadgerOptions.TableLoadingMode = badgerDbOptions.MemoryMap
		b.mDbConfiguration.BadgerOptions.ValueLogLoadingMode = badgerDbOptions.MemoryMap
		b.mDbConfiguration.BadgerOptions.Compression = badgerDbOptions.Snappy
	} else {
		if opts, ok := configuration.(Configuration); ok {
			b.mDbConfiguration.BadgerOptions = opts.BadgerOptions
			if !b.mDbConfiguration.BadgerOptions.InMemory {
				b.mDbConfiguration.BadgerOptions.ValueDir = path
				b.mDbConfiguration.BadgerOptions.Dir = path
			}
		}
	}

	b.mDbConfiguration.BadgerOptions.Logger = logger.CS(Name)
	b.mDbConfiguration.BadgerOptions.EncryptionKey = secretKey

	if b.mDbConfiguration.GoroutineNumber <= 0 {
		b.mDbConfiguration.GoroutineNumber = 16
	}

	if b.mDbConfiguration.GCDiscardRatio <= 0.0 {
		b.mDbConfiguration.GCDiscardRatio = 0.5
	}

	if b.mDbConfiguration.LogPrefix == "" {
		b.mDbConfiguration.LogPrefix = Name
	}

	if b.mDbConfiguration.SliceCap <= 0 {
		b.mDbConfiguration.SliceCap = 100
	}

	if b.mDbConfiguration.BadgerOptions.EncryptionKeyRotationDuration <= 0 {
		b.mDbConfiguration.BadgerOptions.EncryptionKeyRotationDuration = 10 * 24 * time.Hour
	}
}

func (b *Badger) Open() error {
	internalLogger.Debug("state::badger opening db")
	defer func() {
		_ = internalLogger.Sync()
	}()

	if b.mDb != nil {
		return x.ErrStorageIsAlreadyOpen
	}

	db, err := badgerDb.Open(b.mDbConfiguration.BadgerOptions)

	if err != nil {
		internalLogger.Error("state::badger failed to open db", zap.Error(err))
		return x.ErrFailedToOpenStorage
	}

	b.mDb = db
	internalLogger.Debug("state::badger db opened")
	return nil
}

func (b *Badger) Close() error {
	defer func() {
		_ = internalLogger.Sync()
	}()

	if b.mDb == nil {
		return nil
	}

	b.mDbPath = ""
	b.mSecretKey = nil
	b.mDbConfiguration = Configuration{}

	err := b.mDb.Close()
	b.mDb = nil
	if err != nil {
		internalLogger.Error("state::badger failed to close db", zap.Error(err))
		return x.ErrFailedToCloseStorage
	}

	return nil
}

func (b *Badger) RunGC() {
	for {
		if b.mDb == nil {
			return
		}
		err := b.mDb.RunValueLogGC(b.mDbConfiguration.GCDiscardRatio)
		if err != nil {
			break
		}
	}
}

func (b *Badger) ChangeSecretKey(path string, oldSecretKey []byte, newSecretKey []byte, encryptionKeyRotationDuration time.Duration) error {
	defer func() {
		_ = internalLogger.Sync()
	}()

	opt := badgerDb.KeyRegistryOptions{
		Dir:                           path,
		ReadOnly:                      true,
		EncryptionKey:                 oldSecretKey,
		EncryptionKeyRotationDuration: encryptionKeyRotationDuration,
	}

	if opt.EncryptionKeyRotationDuration <= 0 {
		opt.EncryptionKeyRotationDuration = 10 * 24 * time.Hour
	}

	kr, err := badgerDb.OpenKeyRegistry(opt)
	if err != nil {
		internalLogger.Error("state::badger open key registry failure", zap.Error(err))
		return x.ErrFailedToChangeSecretKey
	}

	opt.EncryptionKey = newSecretKey

	err = badgerDb.WriteKeyRegistry(kr, opt)
	if err != nil {
		internalLogger.Error("state::badger write to the key registry failure", zap.Error(err))
		return x.ErrFailedToChangeSecretKey
	}

	return nil
}

func (b *Badger) NewTransaction() iface.IStateStorageTransaction {
	return &Transaction{
		mBadgerTxn: b.mDb.NewTransaction(true),
		mCacheSize: b.mDbConfiguration.SliceCap,
	}
}

func (b *Badger) NewReadOnlyTransaction() iface.IStateStorageTransaction {
	return &Transaction{
		mBadgerTxn: b.mDb.NewTransaction(false),
		mCacheSize: b.mDbConfiguration.SliceCap,
	}
}
