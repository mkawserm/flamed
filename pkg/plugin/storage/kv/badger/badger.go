package badger

import (
	badgerDb "github.com/dgraph-io/badger/v2"
	badgerDbOptions "github.com/dgraph-io/badger/v2/options"
	"github.com/golang/protobuf/proto"
	"github.com/mkawserm/flamed/pkg/constant"
	"github.com/mkawserm/flamed/pkg/pb"
	"github.com/mkawserm/flamed/pkg/uidutil"
	"github.com/mkawserm/flamed/pkg/x"
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
		return x.ErrFailedToOpenStorage
	}
	b.mDb = db
	return nil
}

func (b *Badger) Close() error {
	if b.mDb == nil {
		return nil
	}

	b.mDbPath = ""
	b.mSecretKey = nil
	b.mDbConfiguration = Configuration{}

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

func (b *Badger) ChangeSecretKey(oldSecretKey []byte, newSecretKey []byte) error {
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
		return x.ErrFailedToChangeSecretKey
	}

	opt.EncryptionKey = newSecretKey

	err = badgerDb.WriteKeyRegistry(kr, opt)
	if err != nil {
		return x.ErrFailedToChangeSecretKey
	}

	return nil
}

func (b *Badger) ReadUsingPrefix(prefix []byte) ([]*pb.FlameEntry, error) {
	if b.mDb == nil {
		return nil, x.ErrStorageIsNotReady
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
			return nil, x.ErrFailedToReadDataFromStorage
		}
	}

	return data, nil
}

func (b *Badger) Delete(namespace []byte, key []byte) error {
	if b.mDb == nil {
		return x.ErrStorageIsNotReady
	}

	uid := uidutil.GetUid(namespace, key)

	err := b.mDb.Update(func(txn *badgerDb.Txn) error {
		err := txn.Delete(uid)
		return err
	})

	if err != nil {
		return x.ErrFailedToDeleteDataFromStorage
	}

	return nil
}

func (b *Badger) Create(namespace []byte, key []byte, value []byte) error {
	if b.mDb == nil {
		return x.ErrStorageIsNotReady
	}

	uid := uidutil.GetUid(namespace, key)

	err := b.mDb.Update(func(txn *badgerDb.Txn) error {
		err := txn.Set(uid, value)
		return err
	})

	if err != nil {
		return x.ErrFailedToCreateDataToStorage
	}

	return nil
}

func (b *Badger) Update(namespace []byte, key []byte, value []byte) error {
	if b.mDb == nil {
		return x.ErrStorageIsNotReady
	}

	uid := uidutil.GetUid(namespace, key)

	err := b.mDb.Update(func(txn *badgerDb.Txn) error {
		err := txn.Set(uid, value)
		return err
	})

	if err != nil {
		return x.ErrFailedToUpdateDataToStorage
	}

	return nil
}

func (b *Badger) Append(namespace []byte, key []byte, value []byte) error {
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
		return x.ErrFailedToAppendDataToStorage
	}

	return nil
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
		return x.ErrFailedToReadBatchFromStorage
	}

	return nil
}

func (b *Badger) ApplyBatchAction(batch *pb.FlameBatchAction) error {
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
		return x.ErrFailedToApplyBatchToStorage
	}

	return nil
}

func (b *Badger) ApplyAction(action *pb.FlameAction) error {
	if action.FlameActionType == pb.FlameAction_CREATE {
		return b.Create(action.FlameEntry.Namespace, action.FlameEntry.Key, action.FlameEntry.Value)
	} else if action.FlameActionType == pb.FlameAction_UPDATE {
		return b.Update(action.FlameEntry.Namespace, action.FlameEntry.Key, action.FlameEntry.Value)
	} else if action.FlameActionType == pb.FlameAction_APPEND {
		return b.Append(action.FlameEntry.Namespace, action.FlameEntry.Key, action.FlameEntry.Value)
	} else if action.FlameActionType == pb.FlameAction_DELETE {
		return b.Delete(action.FlameEntry.Namespace, action.FlameEntry.Key)
	}

	return x.ErrFailedToApplyActionToStorage
}

//func (b *Badger) AsyncSnapshot(snapshot chan<- *pb.FlameSnapshot) error {
//	if b.mDb == nil {
//		return x.ErrFailedToGenerateAsyncSnapshotFromStorage
//	}
//
//	stream := b.mDb.NewStream()
//	stream.NumGo = b.mDbConfiguration.GoroutineNumber
//	stream.LogPrefix = b.mDbConfiguration.LogPrefix
//	stream.KeyToList = nil
//
//	stream.Send = func(list *badgerDb.KVList) error {
//		data := &pb.FlameSnapshot{
//			Version:                0,
//			Length:                 0,
//			FlameSnapshotEntryList: make([]*pb.FlameSnapshotEntry, 0, b.mDbConfiguration.SliceCap),
//		}
//
//		for _, kv := range list.Kv {
//			entry := &pb.FlameSnapshotEntry{
//				Uid:  kv.Key,
//				Data: kv.Value,
//			}
//			data.FlameSnapshotEntryList = append(data.FlameSnapshotEntryList, entry)
//		}
//
//		data.Length = uint64(len(data.FlameSnapshotEntryList))
//		snapshot <- data
//		return nil
//	}
//
//	if err := stream.Orchestrate(context.Background()); err != nil {
//		return err
//	}
//
//	return nil
//}
//
//func (b *Badger) ApplyAsyncSnapshot(snapshot <-chan *pb.FlameSnapshot) (bool, error) {
//	if b.mDb == nil {
//		return false, x.ErrFailedToApplyAsyncSnapshotToStorage
//	}
//
//	for {
//		ss := <-snapshot
//		if len(ss.FlameSnapshotEntryList) == 0 {
//			break
//		}
//
//		txn := b.mDb.NewTransaction(true)
//		for _, entry := range ss.FlameSnapshotEntryList {
//			if err := txn.Set(entry.Uid, entry.Data); err == badgerDb.ErrTxnTooBig {
//				if err := txn.Commit(); err != nil {
//					return false, x.ErrFailedToApplyAsyncSnapshotToStorage
//				}
//
//				txn = b.mDb.NewTransaction(true)
//				if err := txn.Set(entry.Uid, entry.Data); err != nil {
//					return false, x.ErrFailedToApplyAsyncSnapshotToStorage
//				}
//
//			} else if err != nil {
//				return false, x.ErrFailedToApplyAsyncSnapshotToStorage
//			}
//		}
//
//		if err := txn.Commit(); err != nil {
//			txn.Discard()
//			return false, x.ErrFailedToApplyAsyncSnapshotToStorage
//		}
//
//		txn.Discard()
//	}
//
//	return true, nil
//}
//
//func (b *Badger) SyncSnapshot() (*pb.FlameSnapshot, error) {
//	if b.mDb == nil {
//		return nil, x.ErrFailedToGenerateSyncSnapshotFromStorage
//	}
//
//	snapshot := &pb.FlameSnapshot{
//		Version:                0,
//		Length:                 0,
//		FlameSnapshotEntryList: make([]*pb.FlameSnapshotEntry, 0, b.mDbConfiguration.SliceCap),
//	}
//
//	err := b.mDb.View(func(txn *badgerDb.Txn) error {
//		opts := badgerDb.DefaultIteratorOptions
//		opts.PrefetchSize = 100
//		it := txn.NewIterator(opts)
//		defer it.Close()
//
//		for it.Rewind(); it.Valid(); it.Next() {
//			item := it.Item()
//			if value, err := item.ValueCopy(nil); err != nil {
//				return err
//			} else {
//				snapshot.FlameSnapshotEntryList = append(snapshot.FlameSnapshotEntryList, &pb.FlameSnapshotEntry{
//					Uid:  item.Key(),
//					Data: value,
//				})
//			}
//		}
//		return nil
//	})
//
//	if err != nil {
//		return nil, x.ErrFailedToGenerateSyncSnapshotFromStorage
//	}
//	snapshot.Length = uint64(len(snapshot.FlameSnapshotEntryList))
//	return snapshot, nil
//}
//
//func (b *Badger) ApplySyncSnapshot(snapshot *pb.FlameSnapshot) (bool, error) {
//	if b.mDb == nil {
//		return false, x.ErrFailedToApplySyncSnapshotToStorage
//	}
//
//	txn := b.mDb.NewTransaction(true)
//	defer txn.Discard()
//
//	for _, entry := range snapshot.FlameSnapshotEntryList {
//		if err := txn.Set(entry.Uid, entry.Data); err == badgerDb.ErrTxnTooBig {
//			if err := txn.Commit(); err != nil {
//				return false, x.ErrFailedToApplySyncSnapshotToStorage
//			}
//
//			txn = b.mDb.NewTransaction(true)
//
//			if err := txn.Set(entry.Uid, entry.Data); err != nil {
//				return false, x.ErrFailedToApplySyncSnapshotToStorage
//			}
//
//		} else if err != nil {
//			return false, x.ErrFailedToApplySyncSnapshotToStorage
//		}
//	}
//
//	if err := txn.Commit(); err != nil {
//		return false, x.ErrFailedToApplySyncSnapshotToStorage
//	}
//
//	return true, nil
//}

func (b *Badger) PrepareSnapshot() (interface{}, error) {
	if b.mDb == nil {
		return nil, x.ErrStorageIsNotReady
	}

	return b.mDb.NewTransaction(false), nil
}

func (b *Badger) SaveSnapshot(snapshotContext interface{}, w io.Writer) error {
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
		return x.ErrFailedToSaveSnapshot
	}

	opts = badgerDb.DefaultIteratorOptions
	opts.PrefetchSize = 100

	it = txn.NewIterator(opts)
	defer it.Close()

	for it.Rewind(); it.Valid(); it.Next() {
		item := it.Item()
		if value, err := item.ValueCopy(nil); err != nil {
			return x.ErrFailedToSaveSnapshot
		} else {
			entry := &pb.FlameSnapshotEntry{
				Uid:  item.Key(),
				Data: value,
			}

			if data, err := proto.Marshal(entry); err != nil {
				return x.ErrFailedToSaveSnapshot
			} else {
				dataLength := uint64(len(data))
				if _, err := w.Write(uidutil.Uint64ToByteSlice(dataLength)); err != nil {
					return x.ErrFailedToSaveSnapshot
				}
				if _, err := w.Write(data); err != nil {
					return x.ErrFailedToSaveSnapshot
				}
			}
		}
	}

	return nil
}

func (b *Badger) RecoverFromSnapshot(r io.Reader) error {
	if b.mDb == nil {
		return x.ErrStorageIsNotReady
	}

	sz := make([]byte, 8)
	if _, err := io.ReadFull(r, sz); err != nil {
		return x.ErrFailedToRecoverFromSnapshot
	}

	total := uidutil.ByteSliceToUint64(sz)

	txn := b.mDb.NewTransaction(true)
	defer txn.Discard()

	for i := uint64(0); i < total; i++ {
		if _, err := io.ReadFull(r, sz); err != nil {
			return x.ErrFailedToRecoverFromSnapshot
		}

		toRead := uidutil.ByteSliceToUint64(sz)
		data := make([]byte, toRead)
		if _, err := io.ReadFull(r, data); err != nil {
			return x.ErrFailedToRecoverFromSnapshot
		}

		entry := &pb.FlameSnapshotEntry{}
		if err := proto.Unmarshal(data, entry); err != nil {
			return x.ErrFailedToRecoverFromSnapshot
		}

		if err := txn.Set(entry.Uid, entry.Data); err == badgerDb.ErrTxnTooBig {
			if err := txn.Commit(); err != nil {
				return x.ErrFailedToRecoverFromSnapshot
			}

			txn = b.mDb.NewTransaction(true)

			if err := txn.Set(entry.Uid, entry.Data); err != nil {
				return x.ErrFailedToRecoverFromSnapshot
			}

		} else if err != nil {
			return x.ErrFailedToRecoverFromSnapshot
		}
	}

	if err := txn.Commit(); err != nil {
		return x.ErrFailedToRecoverFromSnapshot
	}

	return nil
}

func (b *Badger) SaveAppliedIndex(u uint64) error {
	if b.mDb == nil {
		return x.ErrStorageIsNotReady
	}

	return b.Create(
		[]byte(constant.AppliedIndexNamespace),
		[]byte(constant.AppliedIndexKey),
		uidutil.Uint64ToByteSlice(u))
}

func (b *Badger) QueryAppliedIndex() (uint64, error) {
	if b.mDb == nil {
		return 0, x.ErrStorageIsNotReady
	}

	data, err := b.Read(
		[]byte(constant.AppliedIndexNamespace),
		[]byte(constant.AppliedIndexKey))

	if err == x.ErrUidDoesNotExists {
		return 0, nil
	}

	if err != nil {
		return 0, err
	}
	return uidutil.ByteSliceToUint64(data), nil
}
