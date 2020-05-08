package badger

import (
	badgerDb "github.com/dgraph-io/badger/v2"
	badgerDbOptions "github.com/dgraph-io/badger/v2/options"
	"github.com/golang/protobuf/proto"
	"github.com/mkawserm/flamed/pkg/logger"
	"github.com/mkawserm/flamed/pkg/pb"
	"github.com/mkawserm/flamed/pkg/uidutil"
	"github.com/mkawserm/flamed/pkg/x"
	"go.uber.org/zap"
	"io"
	"time"
)

type Badger struct {
	mDbPath          string
	mSecretKey       []byte
	mDb              *badgerDb.DB
	mDbConfiguration Configuration
}

func (b *Badger) Open(path string, secretKey []byte, readOnly bool, configuration interface{}) error {
	internalLogger.Debug("opening badger database")
	defer func() {
		_ = internalLogger.Sync()
	}()

	if b.mDb != nil {
		return x.ErrStorageIsAlreadyOpen
	}

	b.mDbPath = path
	b.mSecretKey = secretKey
	b.mDbConfiguration.BadgerOptions = badgerDb.DefaultOptions(b.mDbPath)

	if configuration == nil {
		b.mDbConfiguration.BadgerOptions.Truncate = true
		b.mDbConfiguration.BadgerOptions.TableLoadingMode = badgerDbOptions.LoadToRAM
		b.mDbConfiguration.BadgerOptions.ValueLogLoadingMode = badgerDbOptions.MemoryMap
		b.mDbConfiguration.BadgerOptions.Compression = badgerDbOptions.Snappy
		b.mDbConfiguration.BadgerOptions.Logger = logger.S("badger")
	} else {
		if opts, ok := configuration.(Configuration); ok {
			b.mDbConfiguration.BadgerOptions = opts.BadgerOptions
			if !b.mDbConfiguration.BadgerOptions.InMemory {
				b.mDbConfiguration.BadgerOptions.ValueDir = path
				b.mDbConfiguration.BadgerOptions.Dir = path
			}
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
		internalLogger.Error("failed to open badger db", zap.Error(err))
		return x.ErrFailedToOpenStorage
	}
	b.mDb = db
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
		internalLogger.Error("failed to close badger db", zap.Error(err))
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

func (b *Badger) ChangeSecretKey(oldSecretKey []byte, newSecretKey []byte) error {
	defer func() {
		_ = internalLogger.Sync()
	}()

	if b.mDb == nil {
		return x.ErrStorageIsNotReady
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
		internalLogger.Error("open key registry failure", zap.Error(err))
		return x.ErrFailedToChangeSecretKey
	}

	opt.EncryptionKey = newSecretKey

	err = badgerDb.WriteKeyRegistry(kr, opt)
	if err != nil {
		internalLogger.Error("write to the key registry failure", zap.Error(err))
		return x.ErrFailedToChangeSecretKey
	}

	return nil
}

func (b *Badger) Iterate(seek, prefix []byte, limit int, receiver func(entry *pb.FlameEntry) bool) error {
	defer func() {
		_ = internalLogger.Sync()
	}()

	if b.mDb == nil {
		return x.ErrStorageIsNotReady
	}

	err := b.mDb.View(func(txn *badgerDb.Txn) error {
		it := txn.NewIterator(badgerDb.DefaultIteratorOptions)
		defer it.Close()

		counter := 0
		if limit == 0 {
			counter = -1
		}
		for it.Seek(seek); it.ValidForPrefix(prefix) && counter < limit; it.Next() {
			item := it.Item()
			uid := item.Key()

			if value, err := item.ValueCopy(nil); err == nil {
				namespace, key := uidutil.SplitUid(uid)
				entry := &pb.FlameEntry{
					Namespace: namespace,
					Key:       key,
					Value:     value,
				}

				if limit != 0 {
					counter = counter + 1
				}

				next := receiver(entry)
				if !next {
					break
				}
			} else {
				return err
			}
		}
		return nil
	})

	if err != nil {
		internalLogger.Error("iteration failure", zap.Error(err))
		return x.ErrFailedToIterate
	}

	return nil
}

//func (b *Badger) IterateKeyOnly(seek, prefix []byte, limit int, receiver func(entry *pb.FlameEntry) bool) error {
//	defer func() {
//		_ = internalLogger.Sync()
//	}()
//
//	if b.mDb == nil {
//		return x.ErrStorageIsNotReady
//	}
//
//	err := b.mDb.View(func(txn *badgerDb.Txn) error {
//		opts := badgerDb.DefaultIteratorOptions
//		opts.PrefetchValues = false
//		it := txn.NewIterator(opts)
//		defer it.Close()
//
//		counter := 0
//		if limit == 0 {
//			counter = -1
//		}
//		for it.Seek(seek); it.ValidForPrefix(prefix) && counter < limit; it.Next() {
//			item := it.Item()
//			uid := item.Key()
//			namespace, key := uidutil.SplitUid(uid)
//			entry := &pb.FlameEntry{
//				Namespace: namespace,
//				Key:       key,
//			}
//
//			if limit != 0 {
//				counter = counter + 1
//			}
//
//			next := receiver(entry)
//			if !next {
//				break
//			}
//		}
//		return nil
//	})
//
//	if err != nil {
//		internalLogger.Error("iteration failure", zap.Error(err))
//		return x.ErrFailedToIterate
//	}
//
//	return nil
//}

func (b *Badger) Read(namespace []byte, key []byte) ([]byte, error) {
	defer func() {
		_ = internalLogger.Sync()
	}()

	if b.mDb == nil {
		return nil, x.ErrStorageIsNotReady
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
			internalLogger.Error("failed to read", zap.Error(err))
			return nil, x.ErrFailedToReadDataFromStorage
		}
	}

	return data, nil
}

func (b *Badger) Delete(namespace []byte, key []byte) error {
	defer func() {
		_ = internalLogger.Sync()
	}()

	if b.mDb == nil {
		return x.ErrStorageIsNotReady
	}

	uid := uidutil.GetUid(namespace, key)

	err := b.mDb.Update(func(txn *badgerDb.Txn) error {
		err := txn.Delete(uid)
		return err
	})

	if err != nil {
		internalLogger.Error("failed to delete", zap.Error(err))
		return x.ErrFailedToDeleteDataFromStorage
	}

	return nil
}

func (b *Badger) Create(namespace []byte, key []byte, value []byte) error {
	defer func() {
		_ = internalLogger.Sync()
	}()

	if b.mDb == nil {
		return x.ErrStorageIsNotReady
	}

	uid := uidutil.GetUid(namespace, key)

	err := b.mDb.Update(func(txn *badgerDb.Txn) error {
		err := txn.Set(uid, value)
		return err
	})

	if err != nil {
		internalLogger.Error("failed to create", zap.Error(err))
		return x.ErrFailedToCreateDataToStorage
	}

	return nil
}

func (b *Badger) Update(namespace []byte, key []byte, value []byte) error {
	defer func() {
		_ = internalLogger.Sync()
	}()

	if b.mDb == nil {
		return x.ErrStorageIsNotReady
	}

	uid := uidutil.GetUid(namespace, key)

	err := b.mDb.Update(func(txn *badgerDb.Txn) error {
		err := txn.Set(uid, value)
		return err
	})

	if err != nil {
		internalLogger.Error("failed to update", zap.Error(err))
		return x.ErrFailedToUpdateDataToStorage
	}

	return nil
}

func (b *Badger) Append(namespace []byte, key []byte, value []byte) error {
	defer func() {
		_ = internalLogger.Sync()
	}()

	if b.mDb == nil {
		return x.ErrStorageIsNotReady
	}

	uid := uidutil.GetUid(namespace, key)

	err := b.mDb.Update(func(txn *badgerDb.Txn) error {
		var data = b.getValue(txn, uid)
		data = append(data, value...)
		return txn.Set(uid, data)
	})

	if err != nil {
		internalLogger.Error("failed to append", zap.Error(err))
		return x.ErrFailedToAppendDataToStorage
	}

	return nil
}

func (b *Badger) IsExists(namespace []byte, key []byte) bool {
	defer func() {
		_ = internalLogger.Sync()
	}()

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
		internalLogger.Error("is exist check failed", zap.Error(err))
		return false
	}
}

func (b *Badger) getValue(txn *badgerDb.Txn, uid []byte) []byte {
	var data []byte
	item, err := txn.Get(uid)
	if err != nil {
		return nil
	}

	if v, err := item.ValueCopy(nil); err == nil {
		data = v
	}

	return data
}

func (b *Badger) ReadBatch(batch *pb.FlameBatchRead) error {
	defer func() {
		_ = internalLogger.Sync()
	}()

	if b.mDb == nil {
		return x.ErrStorageIsNotReady
	}

	err := b.mDb.View(func(txn *badgerDb.Txn) error {
		for idx := range batch.FlameEntryList {
			uid := uidutil.GetUid(batch.FlameEntryList[idx].Namespace, batch.FlameEntryList[idx].Key)
			batch.FlameEntryList[idx].Value = b.getValue(txn, uid)
		}
		return nil
	})

	if err != nil {
		internalLogger.Error("read batch failure", zap.Error(err))
		return x.ErrFailedToReadBatchFromStorage
	}

	return nil
}

func (b *Badger) ApplyBatchAction(batch *pb.FlameBatchAction) error {
	defer func() {
		_ = internalLogger.Sync()
	}()

	if b.mDb == nil {
		return x.ErrStorageIsNotReady
	}

	txn := b.mDb.NewTransaction(true)
	defer txn.Discard()

	for _, action := range batch.FlameActionList {
		if action.FlameActionType == pb.FlameAction_CREATE || action.FlameActionType == pb.FlameAction_UPDATE {
			uid := uidutil.GetUid(action.FlameEntry.Namespace, action.FlameEntry.Key)
			if err := txn.Set(uid, action.FlameEntry.Value); err != nil {
				return x.ErrFailedToApplyBatchToStorage
			}
		} else if action.FlameActionType == pb.FlameAction_DELETE {
			uid := uidutil.GetUid(action.FlameEntry.Namespace, action.FlameEntry.Key)
			if err := txn.Delete(uid); err != nil {
				return x.ErrFailedToApplyBatchToStorage
			}
		} else if action.FlameActionType == pb.FlameAction_APPEND {
			uid := uidutil.GetUid(action.FlameEntry.Namespace, action.FlameEntry.Key)

			var data = b.getValue(txn, uid)
			data = append(data, action.FlameEntry.Value...)
			if err := txn.Set(uid, data); err != nil {
				return x.ErrFailedToApplyBatchToStorage
			}
		}
	}

	if err := txn.Commit(); err != nil {
		internalLogger.Error("txn commit error", zap.Error(err))
		return x.ErrFailedToApplyBatchToStorage
	}

	return nil
}

//func (b *Badger) ApplyAction(action *pb.FlameAction) error {
//	if action.FlameActionType == pb.FlameAction_CREATE {
//		return b.Create(action.FlameEntry.Namespace, action.FlameEntry.Key, action.FlameEntry.Value)
//	} else if action.FlameActionType == pb.FlameAction_UPDATE {
//		return b.Update(action.FlameEntry.Namespace, action.FlameEntry.Key, action.FlameEntry.Value)
//	} else if action.FlameActionType == pb.FlameAction_APPEND {
//		return b.Append(action.FlameEntry.Namespace, action.FlameEntry.Key, action.FlameEntry.Value)
//	} else if action.FlameActionType == pb.FlameAction_DELETE {
//		return b.Delete(action.FlameEntry.Namespace, action.FlameEntry.Key)
//	}
//
//	return x.ErrFailedToApplyActionToStorage
//}

func (b *Badger) PrepareSnapshot() (interface{}, error) {
	defer func() {
		_ = internalLogger.Sync()
	}()
	if b.mDb == nil {
		return nil, x.ErrStorageIsNotReady
	}

	internalLogger.Debug("badger db snapshot prepared")
	return b.mDb.NewTransaction(false), nil
}

func (b *Badger) SaveSnapshot(snapshotContext interface{}, w io.Writer) error {
	defer func() {
		_ = internalLogger.Sync()
	}()

	internalLogger.Debug("badger db saving snapshot")
	if b.mDb == nil {
		return x.ErrStorageIsNotReady
	}

	if snapshotContext == nil {
		return x.ErrFailedToSaveSnapshot
	}

	var txn *badgerDb.Txn
	if v, ok := snapshotContext.(*badgerDb.Txn); ok {
		txn = v
	} else {
		return x.ErrFailedToSaveSnapshot
	}

	defer txn.Discard()

	total := uint64(0)
	opts := badgerDb.DefaultIteratorOptions
	opts.PrefetchValues = false
	it := txn.NewIterator(opts)
	for it.Rewind(); it.Valid(); it.Next() {
		total = total + 1
	}
	it.Close()

	if _, err := w.Write(uidutil.Uint64ToByteSlice(total)); err != nil {
		internalLogger.Error("write error", zap.Error(err))
		return x.ErrFailedToSaveSnapshot
	}

	opts = badgerDb.DefaultIteratorOptions
	opts.PrefetchSize = 100

	it = txn.NewIterator(opts)
	defer it.Close()

	for it.Rewind(); it.Valid(); it.Next() {
		item := it.Item()
		if value, err := item.ValueCopy(nil); err != nil {
			internalLogger.Error("value copy error", zap.Error(err))
			return x.ErrFailedToSaveSnapshot
		} else {
			entry := &pb.FlameSnapshotEntry{
				Uid:  item.Key(),
				Data: value,
			}

			if data, err := proto.Marshal(entry); err != nil {
				internalLogger.Error("marshal error", zap.Error(err))
				return x.ErrFailedToSaveSnapshot
			} else {
				dataLength := uint64(len(data))
				if _, err := w.Write(uidutil.Uint64ToByteSlice(dataLength)); err != nil {
					internalLogger.Error("write error", zap.Error(err))
					return x.ErrFailedToSaveSnapshot
				}
				if _, err := w.Write(data); err != nil {
					return x.ErrFailedToSaveSnapshot
				}
			}
		}
	}

	internalLogger.Debug("badger db snapshot saved")

	return nil
}

func (b *Badger) RecoverFromSnapshot(r io.Reader) error {
	defer func() {
		_ = internalLogger.Sync()
	}()

	internalLogger.Debug("badger db recovering from snapshot")

	if b.mDb == nil {
		return x.ErrStorageIsNotReady
	}

	sz := make([]byte, 8)
	if _, err := io.ReadFull(r, sz); err != nil {
		internalLogger.Error("read error", zap.Error(err))
		return x.ErrFailedToRecoverFromSnapshot
	}

	total := uidutil.ByteSliceToUint64(sz)

	txn := b.mDb.NewTransaction(true)
	defer txn.Discard()

	for i := uint64(0); i < total; i++ {
		if _, err := io.ReadFull(r, sz); err != nil {
			internalLogger.Error("read error", zap.Error(err))
			return x.ErrFailedToRecoverFromSnapshot
		}

		toRead := uidutil.ByteSliceToUint64(sz)
		data := make([]byte, toRead)
		if _, err := io.ReadFull(r, data); err != nil {
			internalLogger.Error("read error", zap.Error(err))
			return x.ErrFailedToRecoverFromSnapshot
		}

		entry := &pb.FlameSnapshotEntry{}
		if err := proto.Unmarshal(data, entry); err != nil {
			internalLogger.Error("unmarshal error", zap.Error(err))
			return x.ErrFailedToRecoverFromSnapshot
		}

		if err := txn.Set(entry.Uid, entry.Data); err == badgerDb.ErrTxnTooBig {
			if err := txn.Commit(); err != nil {
				internalLogger.Error("txn commit error", zap.Error(err))
				return x.ErrFailedToRecoverFromSnapshot
			}

			txn = b.mDb.NewTransaction(true)

			if err := txn.Set(entry.Uid, entry.Data); err != nil {
				internalLogger.Error("txn set error", zap.Error(err))
				return x.ErrFailedToRecoverFromSnapshot
			}

		} else if err != nil {
			internalLogger.Error("txn set error", zap.Error(err))
			return x.ErrFailedToRecoverFromSnapshot
		}
	}

	if err := txn.Commit(); err != nil {
		internalLogger.Error("txn commit error", zap.Error(err))
		return x.ErrFailedToRecoverFromSnapshot
	}

	internalLogger.Debug("badger db recovered from snapshot")

	return nil
}

//func (b *Badger) CreateIndexMeta(meta *pb.FlameIndexMeta) error {
//	defer func() {
//		_ = internalLogger.Sync()
//	}()
//
//	if b.mDb == nil {
//		return x.ErrStorageIsNotReady
//	}
//
//	data, err := proto.Marshal(meta)
//	if err != nil {
//		internalLogger.Error("marshal error", zap.Error(err))
//		return x.ErrDataMarshalError
//	}
//
//	uid := uidutil.GetUid([]byte(constant.IndexMetaNamespace), meta.Namespace)
//	err = b.mDb.Update(func(txn *badgerDb.Txn) error {
//		err := txn.Set(uid, data)
//		return err
//	})
//
//	if err != nil {
//		internalLogger.Error("update error", zap.Error(err))
//		return x.ErrFailedToCreateIndexMeta
//	}
//
//	return nil
//}

//func (b *Badger) IsIndexMetaExists(meta *pb.FlameIndexMeta) bool {
//	return b.IsExists([]byte(constant.IndexMetaNamespace), meta.Namespace)
//}

//func (b *Badger) GetIndexMeta(meta *pb.FlameIndexMeta) error {
//	defer func() {
//		_ = internalLogger.Sync()
//	}()
//
//	if b.mDb == nil {
//		return x.ErrStorageIsNotReady
//	}
//
//	uid := uidutil.GetUid([]byte(constant.IndexMetaNamespace), meta.Namespace)
//	err := b.mDb.View(func(txn *badgerDb.Txn) error {
//		data := b.getValue(txn, uid)
//		if data == nil {
//			return x.ErrFailedToGetIndexMeta
//		}
//
//		return proto.Unmarshal(data, meta)
//	})
//
//	if err != nil {
//		internalLogger.Error("update error", zap.Error(err))
//		return x.ErrFailedToGetIndexMeta
//	}
//
//	return nil
//}

//func (b *Badger) GetAllIndexMeta() ([]*pb.FlameIndexMeta, error) {
//	defer func() {
//		_ = internalLogger.Sync()
//	}()
//
//	if b.mDb == nil {
//		return nil, x.ErrStorageIsNotReady
//	}
//
//	data := make([]*pb.FlameIndexMeta, 0, b.mDbConfiguration.SliceCap)
//
//	err := b.mDb.View(func(txn *badgerDb.Txn) error {
//		it := txn.NewIterator(badgerDb.DefaultIteratorOptions)
//		defer it.Close()
//
//		for it.Seek([]byte(constant.IndexMetaNamespace)); it.ValidForPrefix([]byte(constant.IndexMetaNamespace)); it.Next() {
//			item := it.Item()
//			if value, err := item.ValueCopy(nil); err == nil {
//				fim := &pb.FlameIndexMeta{}
//				if err := proto.Unmarshal(value, fim); err == nil {
//					data = append(data, fim)
//				}
//			} else {
//				return err
//			}
//		}
//		return nil
//	})
//
//	if err != nil {
//		internalLogger.Error("view error", zap.Error(err))
//		return nil, x.ErrFailedToGetAllIndexMeta
//	}
//
//	return data, nil
//}

//func (b *Badger) UpdateIndexMeta(meta *pb.FlameIndexMeta) error {
//	defer func() {
//		_ = internalLogger.Sync()
//	}()
//
//	if b.mDb == nil {
//		return x.ErrStorageIsNotReady
//	}
//
//	data, err := proto.Marshal(meta)
//	if err != nil {
//		internalLogger.Error("proto marshal error", zap.Error(err))
//		return x.ErrDataMarshalError
//	}
//
//	uid := uidutil.GetUid([]byte(constant.IndexMetaNamespace), meta.Namespace)
//	err = b.mDb.Update(func(txn *badgerDb.Txn) error {
//		err := txn.Set(uid, data)
//		return err
//	})
//
//	if err != nil {
//		internalLogger.Error("badger db update error", zap.Error(err))
//		return x.ErrFailedToUpdateIndexMeta
//	}
//
//	return nil
//}

//func (b *Badger) DeleteIndexMeta(meta *pb.FlameIndexMeta) error {
//	defer func() {
//		_ = internalLogger.Sync()
//	}()
//
//	if b.mDb == nil {
//		return x.ErrStorageIsNotReady
//	}
//
//	uid := uidutil.GetUid([]byte(constant.IndexMetaNamespace), meta.Namespace)
//	err := b.mDb.Update(func(txn *badgerDb.Txn) error {
//		err := txn.Delete(uid)
//		return err
//	})
//
//	if err != nil {
//		internalLogger.Error("badger db delete error", zap.Error(err))
//		return x.ErrFailedToDeleteIndexMeta
//	}
//
//	return nil
//}

//func (b *Badger) CreateUser(user *pb.FlameUser) error {
//	defer func() {
//		_ = internalLogger.Sync()
//	}()
//
//	if b.mDb == nil {
//		return x.ErrStorageIsNotReady
//	}
//
//	data, err := proto.Marshal(user)
//	if err != nil {
//		internalLogger.Error("proto marshal error", zap.Error(err))
//		return x.ErrDataMarshalError
//	}
//
//	uid := uidutil.GetUid([]byte(constant.UserNamespace), []byte(user.Username))
//	err = b.mDb.Update(func(txn *badgerDb.Txn) error {
//		err := txn.Set(uid, data)
//		return err
//	})
//
//	if err != nil {
//		internalLogger.Error("badger db delete error", zap.Error(err))
//		return x.ErrFailedToCreateUser
//	}
//
//	return nil
//}

//func (b *Badger) IsUserExists(user *pb.FlameUser) bool {
//	return b.IsExists([]byte(constant.UserNamespace), []byte(user.Username))
//}

//func (b *Badger) GetUser(user *pb.FlameUser) error {
//	defer func() {
//		_ = internalLogger.Sync()
//	}()
//
//	if b.mDb == nil {
//		return x.ErrStorageIsNotReady
//	}
//
//	uid := uidutil.GetUid([]byte(constant.UserNamespace), []byte(user.Username))
//	err := b.mDb.View(func(txn *badgerDb.Txn) error {
//		data := b.getValue(txn, uid)
//		if data == nil {
//			return x.ErrFailedToGetUser
//		}
//
//		return proto.Unmarshal(data, user)
//	})
//
//	if err != nil {
//		internalLogger.Error("badger db update error", zap.Error(err))
//		return x.ErrFailedToGetUser
//	}
//
//	return nil
//}

//func (b *Badger) GetAllUser() ([]*pb.FlameUser, error) {
//	defer func() {
//		_ = internalLogger.Sync()
//	}()
//
//	if b.mDb == nil {
//		return nil, x.ErrStorageIsNotReady
//	}
//
//	data := make([]*pb.FlameUser, 0, b.mDbConfiguration.SliceCap)
//
//	err := b.mDb.View(func(txn *badgerDb.Txn) error {
//		it := txn.NewIterator(badgerDb.DefaultIteratorOptions)
//		defer it.Close()
//
//		for it.Seek([]byte(constant.UserNamespace)); it.ValidForPrefix([]byte(constant.UserNamespace)); it.Next() {
//			item := it.Item()
//			if value, err := item.ValueCopy(nil); err == nil {
//				fu := &pb.FlameUser{}
//				if err := proto.Unmarshal(value, fu); err == nil {
//					data = append(data, fu)
//				}
//			} else {
//				return err
//			}
//		}
//		return nil
//	})
//
//	if err != nil {
//		internalLogger.Error("badger db view error", zap.Error(err))
//		return nil, x.ErrFailedToGetAllUser
//	}
//
//	return data, nil
//}

//func (b *Badger) UpdateUser(user *pb.FlameUser) error {
//	defer func() {
//		_ = internalLogger.Sync()
//	}()
//
//	if b.mDb == nil {
//		return x.ErrStorageIsNotReady
//	}
//
//	data, err := proto.Marshal(user)
//	if err != nil {
//		internalLogger.Error("proto marshal error", zap.Error(err))
//		return x.ErrDataMarshalError
//	}
//
//	uid := uidutil.GetUid([]byte(constant.UserNamespace), []byte(user.Username))
//	err = b.mDb.Update(func(txn *badgerDb.Txn) error {
//		err := txn.Set(uid, data)
//		return err
//	})
//
//	if err != nil {
//		internalLogger.Error("badger db update error", zap.Error(err))
//		return x.ErrFailedToUpdateUser
//	}
//
//	return nil
//}

//func (b *Badger) DeleteUser(user *pb.FlameUser) error {
//	defer func() {
//		_ = internalLogger.Sync()
//	}()
//
//	if b.mDb == nil {
//		return x.ErrStorageIsNotReady
//	}
//
//	uid := uidutil.GetUid([]byte(constant.UserNamespace), []byte(user.Username))
//	err := b.mDb.Update(func(txn *badgerDb.Txn) error {
//		err := txn.Delete(uid)
//		return err
//	})
//
//	if err != nil {
//		internalLogger.Error("badger db update error", zap.Error(err))
//		return x.ErrFailedToDeleteUser
//	}
//
//	return nil
//}

//func (b *Badger) CreateAccessControl(ac *pb.FlameAccessControl) error {
//	defer func() {
//		_ = internalLogger.Sync()
//	}()
//
//	if b.mDb == nil {
//		return x.ErrStorageIsNotReady
//	}
//
//	data, err := proto.Marshal(ac)
//	if err != nil {
//		internalLogger.Error("proto marshal error", zap.Error(err))
//		return x.ErrDataMarshalError
//	}
//
//	uid := uidutil.GetUid([]byte(constant.AccessControlNamespace),
//		uidutil.GetUid(ac.Namespace, []byte(ac.Username)))
//
//	err = b.mDb.Update(func(txn *badgerDb.Txn) error {
//		err := txn.Set(uid, data)
//		return err
//	})
//
//	if err != nil {
//		internalLogger.Error("badger db update error", zap.Error(err))
//		return x.ErrFailedToCreateAccessControl
//	}
//
//	return nil
//}

//func (b *Badger) IsAccessControlExists(ac *pb.FlameAccessControl) bool {
//	return b.IsExists([]byte(constant.UserNamespace), uidutil.GetUid(ac.Namespace, []byte(ac.Username)))
//}

//func (b *Badger) GetAccessControl(ac *pb.FlameAccessControl) error {
//	defer func() {
//		_ = internalLogger.Sync()
//	}()
//
//	if b.mDb == nil {
//		return x.ErrStorageIsNotReady
//	}
//
//	uid := uidutil.GetUid([]byte(constant.AccessControlNamespace),
//		uidutil.GetUid(ac.Namespace, []byte(ac.Username)))
//	err := b.mDb.View(func(txn *badgerDb.Txn) error {
//		data := b.getValue(txn, uid)
//		if data == nil {
//			return x.ErrFailedToGetAccessControl
//		}
//
//		return proto.Unmarshal(data, ac)
//	})
//
//	if err != nil {
//		internalLogger.Error("badger db view error", zap.Error(err))
//		return x.ErrFailedToGetAccessControl
//	}
//
//	return nil
//}

//func (b *Badger) GetAllAccessControl() ([]*pb.FlameAccessControl, error) {
//	defer func() {
//		_ = internalLogger.Sync()
//	}()
//
//	if b.mDb == nil {
//		return nil, x.ErrStorageIsNotReady
//	}
//
//	data := make([]*pb.FlameAccessControl, 0, b.mDbConfiguration.SliceCap)
//
//	err := b.mDb.View(func(txn *badgerDb.Txn) error {
//		it := txn.NewIterator(badgerDb.DefaultIteratorOptions)
//		defer it.Close()
//
//		for it.Seek([]byte(constant.AccessControlNamespace)); it.ValidForPrefix([]byte(constant.AccessControlNamespace)); it.Next() {
//			item := it.Item()
//			if value, err := item.ValueCopy(nil); err == nil {
//				fac := &pb.FlameAccessControl{}
//				if err := proto.Unmarshal(value, fac); err == nil {
//					data = append(data, fac)
//				}
//			} else {
//				return err
//			}
//		}
//		return nil
//	})
//
//	if err != nil {
//		internalLogger.Error("badger db view error", zap.Error(err))
//		return nil, x.ErrFailedToGetAllAccessControl
//	}
//
//	return data, nil
//}

//func (b *Badger) UpdateAccessControl(ac *pb.FlameAccessControl) error {
//	defer func() {
//		_ = internalLogger.Sync()
//	}()
//
//	if b.mDb == nil {
//		return x.ErrStorageIsNotReady
//	}
//
//	data, err := proto.Marshal(ac)
//	if err != nil {
//		internalLogger.Error("proto marshal error", zap.Error(err))
//		return x.ErrDataMarshalError
//	}
//
//	uid := uidutil.GetUid([]byte(constant.AccessControlNamespace),
//		uidutil.GetUid(ac.Namespace, []byte(ac.Username)))
//	err = b.mDb.Update(func(txn *badgerDb.Txn) error {
//		err := txn.Set(uid, data)
//		return err
//	})
//
//	if err != nil {
//		internalLogger.Error("badger db update error", zap.Error(err))
//		return x.ErrFailedToUpdateAccessControl
//	}
//
//	return nil
//}

//func (b *Badger) DeleteAccessControl(ac *pb.FlameAccessControl) error {
//	defer func() {
//		_ = internalLogger.Sync()
//	}()
//
//	if b.mDb == nil {
//		return x.ErrStorageIsNotReady
//	}
//
//	uid := uidutil.GetUid([]byte(constant.AccessControlNamespace),
//		uidutil.GetUid(ac.Namespace, []byte(ac.Username)))
//	err := b.mDb.Update(func(txn *badgerDb.Txn) error {
//		err := txn.Delete(uid)
//		return err
//	})
//
//	if err != nil {
//		internalLogger.Error("badger db update error", zap.Error(err))
//		return x.ErrFailedToDeleteAccessControl
//	}
//
//	return nil
//}
