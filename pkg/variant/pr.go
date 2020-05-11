package variant

type ProposalResponse struct {
	Status    uint8                           `json:"status"`    /* 0 - rejected, 1 - accepted */
	ErrorCode uint32                          `json:"errorCode"` /* error code */
	ErrorText string                          `json:"errorText"` /* failure reason short message */
	TPRList   []*TransactionProcessorResponse `json:"tprList"`
}

func (p *ProposalResponse) Append(tpr *TransactionProcessorResponse) {
	if p.TPRList == nil {
		p.TPRList = make([]*TransactionProcessorResponse, 0)
	}
	p.TPRList = append(p.TPRList, tpr)
}

func NewProposalResponse(status uint8) *ProposalResponse {
	return &ProposalResponse{
		Status:  status,
		TPRList: make([]*TransactionProcessorResponse, 0)}
}
