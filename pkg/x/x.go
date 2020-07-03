package x

import "errors"

// ErrUnknownValue defines unknown value error.
var ErrUnknownValue = errors.New("e1: unknown value")

// ErrClusterNotFound defines cluster is not available error.
// var ErrClusterNotFound = errors.New("e2: cluster not found")

// ErrStateNotFound defines state is not available error.
var ErrStateNotFound = errors.New("e3: state is not available")

// ErrPathCanNotBeEmpty defines path empty error.
var ErrPathCanNotBeEmpty = errors.New("e4: path can not be empty")

// ErrTPNotFound defines transaction processor not found error.
var ErrTPNotFound = errors.New("e5: transaction processor not found")

// ErrNotImplemented defines not implemented error.
var ErrNotImplemented = errors.New("e6: not implemented")

// ErrInvalidInput defines invalid input error.
var ErrInvalidInput = errors.New("e7: invalid input")

// ErrAccessViolation defines access violation error. This error must
// be raised when user does not have proper access permission
// to access the resource
var ErrAccessViolation = errors.New("e8: access violation")

// ErrAccessDenied defines access denied error.
var ErrAccessDenied = errors.New("e9: access denied")

// ErrClusterIsNotAvailable defines cluster unavailability error.
var ErrClusterIsNotAvailable = errors.New("e10: cluster is not available")

// ErrAddressNotFound defines key unavailability error.
var ErrAddressNotFound = errors.New("e11: key is not available")

// ErrInvalidPassword defines invalid password error.
var ErrInvalidPassword = errors.New("e12: password length must be greater than 5")

// ErrUnexpectedNilValue defines unexpected nil value error.
var ErrUnexpectedNilValue = errors.New("e13: unexpected nil value")

// ErrInvalidConfiguration defines invalid configuration error.
var ErrInvalidConfiguration = errors.New("e14: invalid configuration")

// ErrStorageIsAlreadyOpen defines storage already open error.
var ErrStorageIsAlreadyOpen = errors.New("e15: storage is already open")

// ErrFailedToOpenStorage defines storage open failure error.
var ErrFailedToOpenStorage = errors.New("e16: failed to open the storage")

// ErrFailedToCloseStorage defines storage closing failure error.
var ErrFailedToCloseStorage = errors.New("e17: failed to close the storage")

// ErrFailedToChangeSecretKey defines secret key change error.
var ErrFailedToChangeSecretKey = errors.New("e18: failed to change secret key")

// ErrFailedToReadDataFromStorage defines reading from storage failure error.
var ErrFailedToReadDataFromStorage = errors.New("e19: failed to read data from the storage")

// ErrFailedToCreateDataToStorage defines data creation to storage failure error.
var ErrFailedToCreateDataToStorage = errors.New("e20: failed to create data to the storage")

// ErrFailedToDeleteDataFromStorage defines data deletion from storage failure error.
var ErrFailedToDeleteDataFromStorage = errors.New("e21: failed to delete data from the storage")

// ErrFailedToCreateIndexMeta defines index meta creation failure error.
var ErrFailedToCreateIndexMeta = errors.New("e22: failed to create indexmeta meta")

// ErrFailedToAddCustomIndexRule defines custom index rule addition failure error.
var ErrFailedToAddCustomIndexRule = errors.New("e23: failed to add custom index rul")

// ErrInvalidUsername defines invalid username error.
var ErrInvalidUsername = errors.New("e24: username length must be greater than 2")

// ErrFailedToApplyIndex defines apply index failure error.
var ErrFailedToApplyIndex = errors.New("e25: failed to apply index")

// ErrFailedToCreateIndex defines index creation failure error.
var ErrFailedToCreateIndex = errors.New("e26: failed to create index")

// ErrInvalidNamespace defines invalid namespace failure error.
var ErrInvalidNamespace = errors.New("e27: invalid namespace: namespace should start with a letter" +
	" and minimum 3 characters and can not contain `::`")

// ErrInvalidLookupInput defines invalid lookup input error.
var ErrInvalidLookupInput = errors.New("e28: invalid lookup input")

//var ErrFailedToPrepareSnapshot = errors.New("failed to prepare snapshot")

// ErrFailedToSaveSnapshot defines snapshot save failure error.
var ErrFailedToSaveSnapshot = errors.New("e29: failed to save snapshot")

// ErrFailedToRecoverFromSnapshot defines recover from snapshot failure error.
var ErrFailedToRecoverFromSnapshot = errors.New("e30: failed to recover from snapshot")

//var ErrInvalidSnapshotContext = errors.New("invalid snapshot context")

// ErrLastIndexIsNotMovingForward defines last index moving forward failure error.
var ErrLastIndexIsNotMovingForward = errors.New("e31: last index is not moving forward")

// ErrNodeIsNotReady defines node is not ready error.
var ErrNodeIsNotReady = errors.New("e32: node is not ready")

// ErrStorageIsNotReady defines storage is not ready error.
var ErrStorageIsNotReady = errors.New("e33: storage is not ready")

// ErrFailedToStopCluster defines cluster stop failure error.
var ErrFailedToStopCluster = errors.New("e34: failed to stop cluster")

// ErrFailedToCreateWALDir defines wal dir create failure error.
var ErrFailedToCreateWALDir = errors.New("e35: failed to create wal dir")

// ErrNodeAlreadyConfigured defines node is not configured failure error.
var ErrNodeAlreadyConfigured = errors.New("e36: node is already configured")

// ErrFailedToCreateNodeHostDir defines node host dir create failure error.
var ErrFailedToCreateNodeHostDir = errors.New("e37: failed to create node host dir")

// ErrInvalidStoragedConfiguration defines invalid storaged configuration error.
var ErrInvalidStoragedConfiguration = errors.New("e38: invalid storaged configuration")

// ErrIndexStorageIsNotReady defines index storage not ready error.
var ErrIndexStorageIsNotReady = errors.New("e39: index storage is not ready")

// ErrInvalidOperation defines invalid operation failure error.
var ErrInvalidOperation = errors.New("e40: invalid operation")

//var ErrResourceIsNotAvailable = errors.New("e41: resource is not available")

// ErrPasswordHashAlgorithmIsNotAvailable defines password hash algorithm unavailable error.
var ErrPasswordHashAlgorithmIsNotAvailable = errors.New("e42: password hash algorithm is not available")

// ErrFailedToGeneratePassword defines password generation failure error.
var ErrFailedToGeneratePassword = errors.New("e43: failed to generate password")

// ErrMetaNamespaceIsReserved defines meta namespace reserved error.
var ErrMetaNamespaceIsReserved = errors.New("e44: meta namespace is reserved")

// ErrAccessControlNotFound defines access control not found error.
var ErrAccessControlNotFound = errors.New("e45: access control not found")

// ErrDecodingError defines decoding failure error.
var ErrDecodingError = errors.New("e46: decoding error")

// ErrPayloadCanNotBeEmpty defines payload can not be empty error.
var ErrPayloadCanNotBeEmpty = errors.New("e47: payload can not be empty")

// ErrFamilyNameCanNotBeEmpty defines family name can not be empty error.
var ErrFamilyNameCanNotBeEmpty = errors.New("e48: family name can not be empty")

// ErrFamilyVersionCanNotBeEmpty defines family version can not be empty error.
var ErrFamilyVersionCanNotBeEmpty = errors.New("e49: family version can not be empty")

// ErrInvalidSourceType defines invalid source type error.
var ErrInvalidSourceType = errors.New("e50: invalid source type")

// ErrPasswordCanNotBeEmpty defines password can not be empty error.
var ErrPasswordCanNotBeEmpty = errors.New("e51: password can not be empty")

// ErrInvalidProposal defines invalid proposal error.
var ErrInvalidProposal = errors.New("e52: invalid proposal")

// ErrGlobalSearchPermissionRequired defines global search permission required error.
var ErrGlobalSearchPermissionRequired = errors.New("e101: global search permission required")

// ErrGlobalIteratePermissionRequired defines global iterate permission required error.
var ErrGlobalIteratePermissionRequired = errors.New("e102: global iterate permission required")

// ErrGlobalRetrievePermissionRequired defines global retrieve permission required error.
var ErrGlobalRetrievePermissionRequired = errors.New("e103: global retrieve permission required")

// ErrGlobalCRUDPermissionRequired defines global CRUD permission required error.
var ErrGlobalCRUDPermissionRequired = errors.New("e104: global CRUD permission required")

// ErrReadPermissionRequired defines read permission required error.
var ErrReadPermissionRequired = errors.New("e105: read permission required")

// ErrWritePermissionRequired defines write permission required error.
var ErrWritePermissionRequired = errors.New("e106: write permission required")

// ErrUpdatePermissionRequired defines update permission required error.
var ErrUpdatePermissionRequired = errors.New("e107: update permission required")

// ErrDeletePermissionRequired defines delete permission required error.
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
