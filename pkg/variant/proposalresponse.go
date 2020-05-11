package variant

type ProposalStatus int8

const (
	REJECTED ProposalStatus = 0
	ACCEPTED ProposalStatus = 1
)

type ProposalResponse struct {
	ProposalStatus                   ProposalStatus                  `json:"proposalStatus"`
	TransactionProcessorResponseList []*TransactionProcessorResponse `json:"transactionProcessorResponseList"`
}

func NewProposalResponse(proposalStatus ProposalStatus) *ProposalResponse {
	return &ProposalResponse{
		ProposalStatus:                   proposalStatus,
		TransactionProcessorResponseList: make([]*TransactionProcessorResponse, 0)}
}
