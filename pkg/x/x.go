package x

import "errors"

var ErrInvalidConfiguration = errors.New("invalid configuration")
var ErrStorageIsAlreadyOpen = errors.New("storage is already open")
var ErrFailedToOpenStorage = errors.New("failed to open the storage")
var ErrFailedToCloseStorage = errors.New("failed to close the storage")
var ErrFailedToChangeSecretKey = errors.New("failed to change secret key")
var ErrFailedToReadDataFromStorage = errors.New("failed to read data from the storage")
var ErrUidDoesNotExists = errors.New("uid does not exists")
var ErrFailedToDeleteDataFromStorage = errors.New("failed to delete data from the storage")
var ErrFailedToCreateDataToStorage = errors.New("failed to create data to the storage")
var ErrFailedToUpdateDataToStorage = errors.New("failed to update data to the storage")
var ErrFailedToAppendDataToStorage = errors.New("failed to append data to the storage")
var ErrFailedToApplyBatchToStorage = errors.New("failed to apply batch to the storage")
var ErrFailedToApplyActionToStorage = errors.New("failed to apply action to the storage")
var ErrFailedToGenerateAsyncSnapshotFromStorage = errors.New("failed to generate async snapshot from the storage")
var ErrFailedToApplyAsyncSnapshotToStorage = errors.New("failed to apply async snapshot to the storage")
var ErrFailedToGenerateSyncSnapshotFromStorage = errors.New("failed to generate sync snapshot from the storage")
var ErrFailedToApplySyncSnapshotToStorage = errors.New("failed to apply sync snapshot to the storage")

var ErrInvalidLookupInput = errors.New("invalid lookup input")

var ErrFailedToPrepareSnapshot = errors.New("failed to prepare snapshot")
var ErrFailedToSaveSnapshot = errors.New("failed to save snapshot")
var ErrFailedToRecoverFromSnapshot = errors.New("failed to recover from snapshot")
var ErrInvalidSnapshotContext = errors.New("invalid snapshot context")
var ErrLastIndexIsNotMovingForward = errors.New("last index is not moving forward")

var ErrStorageIsNotReady = errors.New("storage is not ready")
var ErrNodeIsNotReady = errors.New("node is not ready")
var ErrNodeAlreadyConfigured = errors.New("node is already configured")
var ErrFailedToStopCluster = errors.New("failed to stop cluster")
var ErrFailedToCreateNodeHostDir = errors.New("failed to create node host dir")
var ErrFailedToCreateWALDir = errors.New("failed to create wal dir")
var ErrInvalidStoragedConfiguration = errors.New("invalid storaged configuration")

func IsInvalidConfiguration(err error) bool {
	return err == ErrInvalidConfiguration
}

func IsUidDoesNotExists(err error) bool {
	return err == ErrUidDoesNotExists
}

func IsFailedToOpenStorage(err error) bool {
	return err == ErrFailedToOpenStorage
}

func IsFailedToCloseStorage(err error) bool {
	return err == ErrFailedToCloseStorage
}

func IsFailedToChangeSecretKey(err error) bool {
	return err == ErrFailedToChangeSecretKey
}

func IsFailedToReadDataFromStorage(err error) bool {
	return err == ErrFailedToReadDataFromStorage
}

func IsFailedToDeleteDataFromStorage(err error) bool {
	return err == ErrFailedToDeleteDataFromStorage
}

func IsFailedToCreateDataToStorage(err error) bool {
	return err == ErrFailedToCreateDataToStorage
}

func IsFailedToUpdateDataToStorage(err error) bool {
	return err == ErrFailedToUpdateDataToStorage
}

func IsFailedToApplyBatchToStorage(err error) bool {
	return err == ErrFailedToApplyBatchToStorage
}

func IsFailedToApplyActionToStorage(err error) bool {
	return err == ErrFailedToApplyActionToStorage
}

func IsFailedToGenerateAsyncSnapshotFromStorage(err error) bool {
	return err == ErrFailedToGenerateAsyncSnapshotFromStorage
}

func IsFailedToApplyAsyncSnapshotToStorage(err error) bool {
	return err == ErrFailedToApplyAsyncSnapshotToStorage
}

func IsFailedToGenerateSyncSnapshotFromStorage(err error) bool {
	return err == ErrFailedToGenerateSyncSnapshotFromStorage
}

func IsFailedToApplySyncSnapshotToStorage(err error) bool {
	return err == ErrFailedToApplySyncSnapshotToStorage
}

func IsLastIndexIsNotMovingForward(err error) bool {
	return err == ErrLastIndexIsNotMovingForward
}

func IsInvalidLookupInput(err error) bool {
	return err == ErrInvalidLookupInput
}

func IsFailedToPrepareSnapshot(err error) bool {
	return err == ErrFailedToPrepareSnapshot
}

func IsFailedToSaveSnapshot(err error) bool {
	return err == ErrFailedToSaveSnapshot
}
func IsFailedToRecoverFromSnapshot(err error) bool {
	return err == ErrFailedToRecoverFromSnapshot
}

func IsInvalidSnapshotContext(err error) bool {
	return err == ErrInvalidSnapshotContext
}
