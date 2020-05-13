package pb

func (m *ProposalResponse) Append(tpr *TransactionProcessorResponse) {
	if m.TransactionProcessorResponses == nil {
		m.TransactionProcessorResponses = make([]*TransactionProcessorResponse, 0)
	}
	m.TransactionProcessorResponses = append(m.TransactionProcessorResponses, tpr)
}

func NewProposalResponse(status uint32) *ProposalResponse {
	return &ProposalResponse{
		Status:                        status,
		TransactionProcessorResponses: make([]*TransactionProcessorResponse, 0)}
}
