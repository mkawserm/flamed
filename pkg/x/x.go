package x

import "errors"

var ErrFailedToOpenStorage = errors.New("failed to open the storage")
var ErrFailedToCloseStorage = errors.New("failed to close the storage")
var ErrFailedToChangeSecretKey = errors.New("failed to change secret key")
var ErrFailedToReadDataFromStorage = errors.New("failed to read data from the storage")
var ErrUIDDoesNotExists = errors.New("uid does not exists")
var ErrFailedToDeleteDataFromStorage = errors.New("failed to delete data from the storage")
var ErrFailedToCreateDataToStorage = errors.New("failed to create data to the storage")
var ErrFailedToUpdateDataToStorage = errors.New("failed to update data to the storage")
var ErrFailedToApplyBatchToStorage = errors.New("failed to apply batch to the storage")
var ErrFailedToApplyActionToStorage = errors.New("failed to apply action to the storage")
var ErrFailedToGenerateAsyncSnapshotFromStorage = errors.New("failed to generate async snapshot from the storage")
var ErrFailedToApplyAsyncSnapshotToStorage = errors.New("failed to apply async snapshot to the storage")
var ErrFailedToGenerateSyncSnapshotFromStorage = errors.New("failed to generate sync snapshot from the storage")
var ErrFailedToApplySyncSnapshotToStorage = errors.New("failed to apply sync snapshot to the storage")

func IsUIDDoesNotExists(err error) bool {
	return err == ErrUIDDoesNotExists
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
