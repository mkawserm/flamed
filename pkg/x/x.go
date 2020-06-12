package x

import "errors"

var ErrUnknownValue = errors.New("unknown value")
var ErrClusterNotFound = errors.New("cluster not found")
var ErrStateNotFound = errors.New("state is not available")
var ErrPathCanNotBeEmpty = errors.New("path can not be empty")
var ErrTPNotFound = errors.New("transaction processor not found")
var ErrNotImplemented = errors.New("not implemented")
var ErrInvalidInput = errors.New("invalid input")

var ErrAccessViolation = errors.New("access violation")

var ErrAddressNotFound = errors.New("key is not available")

var ErrInvalidPassword = errors.New("password length must be greater than 5")
var ErrUnexpectedNilValue = errors.New("unexpected nil value")

var ErrInvalidConfiguration = errors.New("invalid configuration")
var ErrStorageIsAlreadyOpen = errors.New("storage is already open")
var ErrFailedToOpenStorage = errors.New("failed to open the storage")
var ErrFailedToCloseStorage = errors.New("failed to close the storage")
var ErrFailedToChangeSecretKey = errors.New("failed to change secret key")
var ErrFailedToReadDataFromStorage = errors.New("failed to read data from the storage")

var ErrFailedToCreateDataToStorage = errors.New("failed to create data to the storage")
var ErrFailedToDeleteDataFromStorage = errors.New("failed to delete data from the storage")

var ErrFailedToCreateIndexMeta = errors.New("failed to create indexmeta meta")

var ErrFailedToAddCustomIndexRule = errors.New("failed to add custom index rul")

var ErrInvalidUsername = errors.New("username length must be greater than 2")

var ErrFailedToApplyIndex = errors.New("failed to apply indexmeta")
var ErrFailedToCreateIndex = errors.New("failed to create indexmeta")

var ErrInvalidNamespace = errors.New("invalid namespace: namespace should start with a letter" +
	" and minimum 3 characters and can not contain `::`")

var ErrInvalidLookupInput = errors.New("invalid lookup input")

//var ErrFailedToPrepareSnapshot = errors.New("failed to prepare snapshot")
var ErrFailedToSaveSnapshot = errors.New("failed to save snapshot")
var ErrFailedToRecoverFromSnapshot = errors.New("failed to recover from snapshot")

//var ErrInvalidSnapshotContext = errors.New("invalid snapshot context")
var ErrLastIndexIsNotMovingForward = errors.New("last indexmeta is not moving forward")

var ErrNodeIsNotReady = errors.New("node is not ready")
var ErrStorageIsNotReady = errors.New("storage is not ready")
var ErrFailedToStopCluster = errors.New("failed to stop cluster")
var ErrFailedToCreateWALDir = errors.New("failed to create wal dir")
var ErrNodeAlreadyConfigured = errors.New("node is already configured")
var ErrFailedToCreateNodeHostDir = errors.New("failed to create node host dir")
var ErrInvalidStoragedConfiguration = errors.New("invalid storaged configuration")

var ErrIndexStorageIsNotReady = errors.New("index storage is not ready")

//func IsInvalidConfiguration(err error) bool {
//	return err == ErrInvalidConfiguration
//}
//
//func IsUidDoesNotExists(err error) bool {
//	return err == ErrUidDoesNotExists
//}
//
//func IsFailedToOpenStorage(err error) bool {
//	return err == ErrFailedToOpenStorage
//}
//
//func IsFailedToCloseStorage(err error) bool {
//	return err == ErrFailedToCloseStorage
//}
//
//func IsFailedToChangeSecretKey(err error) bool {
//	return err == ErrFailedToChangeSecretKey
//}
//
//func IsFailedToReadDataFromStorage(err error) bool {
//	return err == ErrFailedToReadDataFromStorage
//}
//
//func IsFailedToDeleteDataFromStorage(err error) bool {
//	return err == ErrFailedToDeleteDataFromStorage
//}
//
//func IsFailedToCreateDataToStorage(err error) bool {
//	return err == ErrFailedToCreateDataToStorage
//}
//
//func IsFailedToUpdateDataToStorage(err error) bool {
//	return err == ErrFailedToUpdateDataToStorage
//}
//
//func IsFailedToApplyBatchToStorage(err error) bool {
//	return err == ErrFailedToApplyBatchToStorage
//}
//
//func IsFailedToApplyActionToStorage(err error) bool {
//	return err == ErrFailedToApplyActionToStorage
//}
//
//func IsFailedToGenerateAsyncSnapshotFromStorage(err error) bool {
//	return err == ErrFailedToGenerateAsyncSnapshotFromStorage
//}
//
//func IsFailedToApplyAsyncSnapshotToStorage(err error) bool {
//	return err == ErrFailedToApplyAsyncSnapshotToStorage
//}
//
//func IsFailedToGenerateSyncSnapshotFromStorage(err error) bool {
//	return err == ErrFailedToGenerateSyncSnapshotFromStorage
//}
//
//func IsFailedToApplySyncSnapshotToStorage(err error) bool {
//	return err == ErrFailedToApplySyncSnapshotToStorage
//}
//
//func IsLastIndexIsNotMovingForward(err error) bool {
//	return err == ErrLastIndexIsNotMovingForward
//}
//
//func IsInvalidLookupInput(err error) bool {
//	return err == ErrInvalidLookupInput
//}
//
//func IsFailedToPrepareSnapshot(err error) bool {
//	return err == ErrFailedToPrepareSnapshot
//}
//
//func IsFailedToSaveSnapshot(err error) bool {
//	return err == ErrFailedToSaveSnapshot
//}
//func IsFailedToRecoverFromSnapshot(err error) bool {
//	return err == ErrFailedToRecoverFromSnapshot
//}
//
//func IsInvalidSnapshotContext(err error) bool {
//	return err == ErrInvalidSnapshotContext
//}
