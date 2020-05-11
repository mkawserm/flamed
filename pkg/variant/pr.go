package variant

type ProposalResponse struct {
	ProposalStatus                   uint8                           `json:"proposalStatus"` /*0 - rejected, 1- accepted*/
	TransactionProcessorResponseList []*TransactionProcessorResponse `json:"transactionProcessorResponseList"`
}

func (p *ProposalResponse) Append(tpr *TransactionProcessorResponse) {
	if p.TransactionProcessorResponseList == nil {
		p.TransactionProcessorResponseList = make([]*TransactionProcessorResponse, 0)
	}
	p.TransactionProcessorResponseList = append(p.TransactionProcessorResponseList, tpr)
}

func NewProposalResponse(proposalStatus uint8) *ProposalResponse {
	return &ProposalResponse{
		ProposalStatus:                   proposalStatus,
		TransactionProcessorResponseList: make([]*TransactionProcessorResponse, 0)}
}
