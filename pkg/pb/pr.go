package pb

func (m *ProposalResponse) Append(tpr *TransactionResponse) {
	if m.TransactionResponses == nil {
		m.TransactionResponses = make([]*TransactionResponse, 0)
	}
	m.TransactionResponses = append(m.TransactionResponses, tpr)
}

func NewProposalResponse(status uint32) *ProposalResponse {
	return &ProposalResponse{
		Status:               status,
		TransactionResponses: make([]*TransactionResponse, 0)}
}
