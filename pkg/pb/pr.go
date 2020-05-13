package pb

func (m *ProposalResponse) Append(tpr *TransactionProcessorResponse) {
	if m.TransactionProcessorResponseList == nil {
		m.TransactionProcessorResponseList = make([]*TransactionProcessorResponse, 0)
	}
	m.TransactionProcessorResponseList = append(m.TransactionProcessorResponseList, tpr)
}

func NewProposalResponse(status uint32) *ProposalResponse {
	return &ProposalResponse{
		Status:                           status,
		TransactionProcessorResponseList: make([]*TransactionProcessorResponse, 0)}
}
