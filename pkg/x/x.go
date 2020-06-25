package x

import "errors"

var ErrUnknownValue = errors.New("e1: unknown value")
var ErrClusterNotFound = errors.New("e2: cluster not found")
var ErrStateNotFound = errors.New("e3: state is not available")
var ErrPathCanNotBeEmpty = errors.New("e4: path can not be empty")
var ErrTPNotFound = errors.New("e5: transaction processor not found")
var ErrNotImplemented = errors.New("e6: not implemented")
var ErrInvalidInput = errors.New("e7: invalid input")

var ErrAccessViolation = errors.New("e8: access violation")
var ErrAccessDenied = errors.New("e9: access denied")
var ErrClusterIsNotAvailable = errors.New("e10: cluster is not available")

var ErrAddressNotFound = errors.New("e11: key is not available")

var ErrInvalidPassword = errors.New("e12: password length must be greater than 5")
var ErrUnexpectedNilValue = errors.New("e13: unexpected nil value")

var ErrInvalidConfiguration = errors.New("e14: invalid configuration")
var ErrStorageIsAlreadyOpen = errors.New("e15: storage is already open")
var ErrFailedToOpenStorage = errors.New("e16: failed to open the storage")
var ErrFailedToCloseStorage = errors.New("e17: failed to close the storage")
var ErrFailedToChangeSecretKey = errors.New("e18: failed to change secret key")
var ErrFailedToReadDataFromStorage = errors.New("e19: failed to read data from the storage")

var ErrFailedToCreateDataToStorage = errors.New("e20: failed to create data to the storage")
var ErrFailedToDeleteDataFromStorage = errors.New("e21: failed to delete data from the storage")

var ErrFailedToCreateIndexMeta = errors.New("e22: failed to create indexmeta meta")

var ErrFailedToAddCustomIndexRule = errors.New("e23: failed to add custom index rul")

var ErrInvalidUsername = errors.New("e24: username length must be greater than 2")

var ErrFailedToApplyIndex = errors.New("e25: failed to apply indexmeta")
var ErrFailedToCreateIndex = errors.New("e26: failed to create indexmeta")

var ErrInvalidNamespace = errors.New("e27: invalid namespace: namespace should start with a letter" +
	" and minimum 3 characters and can not contain `::`")

var ErrInvalidLookupInput = errors.New("e28: invalid lookup input")

//var ErrFailedToPrepareSnapshot = errors.New("failed to prepare snapshot")
var ErrFailedToSaveSnapshot = errors.New("e29: failed to save snapshot")
var ErrFailedToRecoverFromSnapshot = errors.New("e30: failed to recover from snapshot")

//var ErrInvalidSnapshotContext = errors.New("invalid snapshot context")
var ErrLastIndexIsNotMovingForward = errors.New("e31: last indexmeta is not moving forward")

var ErrNodeIsNotReady = errors.New("e32: node is not ready")
var ErrStorageIsNotReady = errors.New("e33: storage is not ready")
var ErrFailedToStopCluster = errors.New("e34: failed to stop cluster")
var ErrFailedToCreateWALDir = errors.New("e35: failed to create wal dir")
var ErrNodeAlreadyConfigured = errors.New("e36: node is already configured")
var ErrFailedToCreateNodeHostDir = errors.New("e37: failed to create node host dir")
var ErrInvalidStoragedConfiguration = errors.New("e38: invalid storaged configuration")

var ErrIndexStorageIsNotReady = errors.New("e39: index storage is not ready")

var ErrInvalidOperation = errors.New("e40: invalid operation")

//var ErrResourceIsNotAvailable = errors.New("e41: resource is not available")
var ErrPasswordHashAlgorithmIsNotAvailable = errors.New("e42: password hash algorithm is not available")
var ErrFailedToGeneratePassword = errors.New("e43: failed to generate password")
var ErrMetaNamespaceIsReserved = errors.New("e44: meta namespace is reserved")
var ErrAccessControlNotFound = errors.New("e45: access control not found")
var ErrDecodingError = errors.New("e46: decoding error")
var ErrPayloadCanNotBeEmpty = errors.New("e47: payload can not be empty")
var ErrFamilyNameCanNotBeEmpty = errors.New("e48: family name can not be empty")
var ErrFamilyVersionCanNotBeEmpty = errors.New("e49: family version can not be empty")
var ErrInvalidSourceType = errors.New("e50: invalid source type")
var ErrPasswordCanNotBeEmpty = errors.New("e51: password can not be empty")

var ErrGlobalSearchPermissionRequired = errors.New("e101: global search permission required")
var ErrGlobalIteratePermissionRequired = errors.New("e102: global iterate permission required")
var ErrGlobalRetrievePermissionRequired = errors.New("e103: global retrieve permission required")
var ErrGlobalCRUDPermissionRequired = errors.New("e104: global CRUD permission required")
var ErrReadPermissionRequired = errors.New("e105: read permission required")
var ErrWritePermissionRequired = errors.New("e106: write permission required")
var ErrUpdatePermissionRequired = errors.New("e107: update permission required")
var ErrDeletePermissionRequired = errors.New("e108: delete permission required")

//var ErrGenericNotFound = errors.New("e201: generic not found")
//var ErrGenericCanNotProcess = errors.New("e202: generic can not process")

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
