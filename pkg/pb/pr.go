package pb

func (m *ProposalResponse) Append(tpr *TransactionResponse) {
	if m.TransactionProcessorResponses == nil {
		m.TransactionProcessorResponses = make([]*TransactionResponse, 0)
	}
	m.TransactionProcessorResponses = append(m.TransactionProcessorResponses, tpr)
}

func NewProposalResponse(status uint32) *ProposalResponse {
	return &ProposalResponse{
		Status:                        status,
		TransactionProcessorResponses: make([]*TransactionResponse, 0)}
}
