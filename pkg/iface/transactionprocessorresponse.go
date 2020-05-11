package iface

type ITransactionProcessorResponse interface {
	Status() uint8 /* 0 - failed, 1 - success */

	ErrorCode() uint32 /* error code */
	ErrorText() string /* failure reason short message */
}
