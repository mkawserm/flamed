package x

//
//import "testing"
//
//func TestIsInvalidConfiguration(t *testing.T) {
//	t.Helper()
//
//	if !IsInvalidConfiguration(ErrInvalidConfiguration) {
//		t.Fatalf("error mismatch")
//	}
//}
//
//func TestIsUidDoesNotExists(t *testing.T) {
//	t.Helper()
//
//	if !IsUidDoesNotExists(ErrUidDoesNotExists) {
//		t.Fatalf("error mismatch")
//	}
//}
//
//func TestIsFailedToApplyActionToStorage(t *testing.T) {
//	t.Helper()
//
//	if !IsFailedToApplyActionToStorage(ErrFailedToApplyActionToStorage) {
//		t.Fatalf("error mismatch")
//	}
//}
//
//func TestIsFailedToApplyAsyncSnapshotToStorage(t *testing.T) {
//	t.Helper()
//
//	if !IsFailedToApplyAsyncSnapshotToStorage(ErrFailedToApplyAsyncSnapshotToStorage) {
//		t.Fatalf("error mismatch")
//	}
//}
//
//func TestIsFailedToApplyBatchToStorage(t *testing.T) {
//	t.Helper()
//
//	if !IsFailedToApplyBatchToStorage(ErrFailedToApplyBatchToStorage) {
//		t.Fatalf("error mismatch")
//	}
//}
//
//func TestIsFailedToApplySyncSnapshotToStorage(t *testing.T) {
//	t.Helper()
//
//	if !IsFailedToApplySyncSnapshotToStorage(ErrFailedToApplySyncSnapshotToStorage) {
//		t.Fatalf("error mismatch")
//	}
//}
//
//func TestIsFailedToChangeSecretKey(t *testing.T) {
//	t.Helper()
//
//	if !IsFailedToChangeSecretKey(ErrFailedToChangeSecretKey) {
//		t.Fatalf("error mismatch")
//	}
//}
//
//func TestIsFailedToCloseStorage(t *testing.T) {
//	t.Helper()
//
//	if !IsFailedToCloseStorage(ErrFailedToCloseStorage) {
//		t.Fatalf("error mismatch")
//	}
//}
//
//func TestIsFailedToCreateDataToStorage(t *testing.T) {
//	t.Helper()
//
//	if !IsFailedToCreateDataToStorage(ErrFailedToCreateDataToStorage) {
//		t.Fatalf("error mismatch")
//	}
//}
//
//func TestIsFailedToDeleteDataFromStorage(t *testing.T) {
//	t.Helper()
//
//	if !IsFailedToDeleteDataFromStorage(ErrFailedToDeleteDataFromStorage) {
//		t.Fatalf("error mismatch")
//	}
//}
//
//func TestIsFailedToGenerateAsyncSnapshotFromStorage(t *testing.T) {
//	t.Helper()
//
//	if !IsFailedToGenerateAsyncSnapshotFromStorage(ErrFailedToGenerateAsyncSnapshotFromStorage) {
//		t.Fatalf("error mismatch")
//	}
//}
//
//func TestIsFailedToGenerateSyncSnapshotFromStorage(t *testing.T) {
//	t.Helper()
//
//	if !IsFailedToGenerateSyncSnapshotFromStorage(ErrFailedToGenerateSyncSnapshotFromStorage) {
//		t.Fatalf("error mismatch")
//	}
//}
//
//func TestIsFailedToOpenStorage(t *testing.T) {
//	t.Helper()
//
//	if !IsFailedToOpenStorage(ErrFailedToOpenStorage) {
//		t.Fatalf("error mismatch")
//	}
//}
//
//func TestIsFailedToReadDataFromStorage(t *testing.T) {
//	t.Helper()
//
//	if !IsFailedToReadDataFromStorage(ErrFailedToReadDataFromStorage) {
//		t.Fatalf("error mismatch")
//	}
//}
//
//func TestIsFailedToUpdateDataToStorage(t *testing.T) {
//	t.Helper()
//
//	if !IsFailedToUpdateDataToStorage(ErrFailedToUpdateDataToStorage) {
//		t.Fatalf("error mismatch")
//	}
//}
//
//func TestIsLastIndexIsNotMovingForward(t *testing.T) {
//	t.Helper()
//
//	if !IsLastIndexIsNotMovingForward(ErrLastIndexIsNotMovingForward) {
//		t.Fatalf("error mismatch")
//	}
//}
//
//func TestIsInvalidLookupInput(t *testing.T) {
//	t.Helper()
//
//	if !IsInvalidLookupInput(ErrInvalidLookupInput) {
//		t.Fatalf("error mismatch")
//	}
//}
//
//func TestIsFailedToPrepareSnapshot(t *testing.T) {
//	t.Helper()
//
//	if !IsFailedToPrepareSnapshot(ErrFailedToPrepareSnapshot) {
//		t.Fatalf("error mismatch")
//	}
//}
//
//func TestIsFailedToRecoverFromSnapshot(t *testing.T) {
//	t.Helper()
//
//	if !IsFailedToRecoverFromSnapshot(ErrFailedToRecoverFromSnapshot) {
//		t.Fatalf("error mismatch")
//	}
//}
//
//func TestIsFailedToSaveSnapshot(t *testing.T) {
//	t.Helper()
//
//	if !IsFailedToSaveSnapshot(ErrFailedToSaveSnapshot) {
//		t.Fatalf("error mismatch")
//	}
//}
//
//func TestIsInvalidSnapshotContext(t *testing.T) {
//	t.Helper()
//
//	if !IsInvalidSnapshotContext(ErrInvalidSnapshotContext) {
//		t.Fatalf("error mismatch")
//	}
//}
