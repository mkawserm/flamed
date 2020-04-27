package badger

import (
	badgerDb "github.com/dgraph-io/badger/v2"
	badgerDbOptions "github.com/dgraph-io/badger/v2/options"
	"github.com/mkawserm/flamed/pkg/pb"
	"github.com/mkawserm/flamed/pkg/uidutil"
	"github.com/mkawserm/flamed/pkg/x"
	"time"
)

type Badger struct {
	mDbPath        string
	mDb            *badgerDb.DB
	mDbOpenOptions badgerDb.Options
	mSnapshotConf  SnapshotConfiguration
}

func (b *Badger) Open(path string, secretKey []byte, configuration interface{}) (bool, error) {
	if b.mDb != nil {
		return true, nil
	}

	b.mDbPath = path
	b.mDbOpenOptions = badgerDb.DefaultOptions(b.mDbPath)

	if configuration == nil {
		b.mDbOpenOptions.ReadOnly = false
		b.mDbOpenOptions.Truncate = true
		b.mDbOpenOptions.TableLoadingMode = badgerDbOptions.LoadToRAM
		b.mDbOpenOptions.ValueLogLoadingMode = badgerDbOptions.MemoryMap
		b.mDbOpenOptions.Compression = badgerDbOptions.Snappy
	} else {
		if opts, ok := configuration.(badgerDb.Options); ok {
			b.mDbOpenOptions = opts
			b.mDbOpenOptions.ValueDir = path
			b.mDbOpenOptions.Dir = path
		}
	}

	b.mDbOpenOptions.EncryptionKey = secretKey
	db, err := badgerDb.Open(b.mDbOpenOptions)
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
		EncryptionKeyRotationDuration: 10 * 24 * time.Hour,
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

	return false, nil
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

func (b *Badger) SetSnapshotConfiguration(configuration interface{}) {
	if conf, ok := configuration.(SnapshotConfiguration); ok {
		b.mSnapshotConf = conf
	}
}

func (b *Badger) AsyncSnapshot(snapshot chan *pb.FlameSnapshot, maxItem int) error {
	if b.mDb == nil {
		return x.ErrFailedToGenerateAsyncSnapshotFromStorage
	}

	return nil
}

func (b *Badger) ApplyAsyncSnapshot(snapshot chan *pb.FlameSnapshot) (bool, error) {
	if b.mDb == nil {
		return false, x.ErrFailedToApplyAsyncSnapshotToStorage
	}

	return false, nil
}

func (b *Badger) SyncSnapshot() (*pb.FlameSnapshot, error) {
	if b.mDb == nil {
		return nil, x.ErrFailedToGenerateSyncSnapshotFromStorage
	}

	return nil, nil
}

func (b *Badger) ApplySyncSnapshot(snapshot *pb.FlameSnapshot) (bool, error) {
	if b.mDb == nil {
		return false, x.ErrFailedToApplySyncSnapshotToStorage
	}

	return false, nil
}
