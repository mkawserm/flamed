package pb

func (m *ProposalResponse) Append(tpr *TransactionResponse) {
	if m.TransactionResponses == nil {
		m.TransactionResponses = make([]*TransactionResponse, 0)
	}
	m.TransactionResponses = append(m.TransactionResponses, tpr)
}

func NewProposalResponse(status Status) *ProposalResponse {
	return &ProposalResponse{
		Status:               status,
		TransactionResponses: make([]*TransactionResponse, 0)}
}
