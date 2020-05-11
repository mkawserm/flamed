package variant

type TransactionProcessorResponse struct {
	Status    uint8  /* 0 - failed, 1 - success */
	ErrorCode uint32 /* error code */
	ErrorText string /* failure reason short message */
}

func (t *TransactionProcessorResponse) GetStatus() uint8 {
	return t.Status
}

func (t *TransactionProcessorResponse) GetErrorCode() uint32 {
	return t.ErrorCode
}

func (t *TransactionProcessorResponse) GetErrorText() string {
	return t.ErrorText
}
