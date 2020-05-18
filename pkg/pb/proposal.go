package pb

import (
	"github.com/google/uuid"
	"time"
)

func NewProposal() *Proposal {
	uuidValue := uuid.New()
	proposal := &Proposal{
		Uuid:      uuidValue[:],
		CreatedAt: uint64(time.Now().UnixNano()),
	}

	return proposal
}

func (m *Proposal) AddTransactionObject(txn *Transaction) {
	if m.Transactions == nil {
		m.Transactions = make([]*Transaction, 0, 1)
	}

	m.Transactions = append(m.Transactions, txn)
}

func (m *Proposal) AddTransaction(namespace []byte, familyName string, familyVersion string, payload []byte) {
	if m.Transactions == nil {
		m.Transactions = make([]*Transaction, 0, 1)
	}

	txn := &Transaction{
		Payload:       payload,
		Namespace:     namespace,
		FamilyName:    familyName,
		FamilyVersion: familyVersion,
	}

	m.Transactions = append(m.Transactions, txn)
}
