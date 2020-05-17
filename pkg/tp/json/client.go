package json

import (
	"encoding/json"
	"errors"
	"github.com/golang/protobuf/proto"
	"github.com/mkawserm/flamed/pkg/iface"
	"github.com/mkawserm/flamed/pkg/pb"
	"github.com/mkawserm/flamed/pkg/utility"
	"github.com/mkawserm/flamed/pkg/x"
	"time"
)

var ErrInvalidID = errors.New("invalid id")
var ErrEmptyBatch = errors.New("empty batch")
var ErrIdIsNotAvailable = errors.New("id is not available")

type Batch struct {
	mNamespace      []byte
	mMaxBatchLength int
	mTransactions   []*pb.Transaction
}

func (b *Batch) validateIDFromMap(data map[string]interface{}) error {
	if id, found := data["id"]; found {
		if _, ok := id.(string); ok {
			return nil
		} else {
			return ErrInvalidID
		}
	} else {
		return ErrIdIsNotAvailable
	}
}

func (b *Batch) addTransaction(action Action, data map[string]interface{}) error {
	if err := b.validateIDFromMap(data); err != nil {
		return err
	}

	payloadBytes, err := json.Marshal(data)

	if err != nil {
		return err
	}

	jsonPayload := &JSONPayload{
		Action:  action,
		Payload: payloadBytes,
	}

	payload, err := json.Marshal(jsonPayload)
	if err != nil {
		return err
	}

	txn := &pb.Transaction{
		Payload:       payload,
		Namespace:     b.mNamespace,
		FamilyName:    Name,
		FamilyVersion: Version,
	}

	if b.mTransactions == nil {
		b.mTransactions = make([]*pb.Transaction, 0, b.mMaxBatchLength)
	}
	b.mTransactions = append(b.mTransactions, txn)

	return nil
}

func (b *Batch) Reset() {
	b.mTransactions = make([]*pb.Transaction, 0, b.mMaxBatchLength)
}

func (b *Batch) MergeJSONMap(data map[string]interface{}) error {
	return b.addTransaction(Action_MERGE, data)
}

func (b *Batch) InsertJSONMap(data map[string]interface{}) error {
	return b.addTransaction(Action_INSERT, data)
}

func (b *Batch) UpdateJSONMap(data map[string]interface{}) error {
	return b.addTransaction(Action_UPDATE, data)
}

func (b *Batch) UpsertJSONMap(data map[string]interface{}) error {
	return b.addTransaction(Action_UPSERT, data)
}

func (b *Batch) DeleteJSONMap(data map[string]interface{}) error {
	return b.addTransaction(Action_DELETE, data)
}

type Client struct {
	mClusterID      uint64
	mNamespace      string
	mMaxBatchLength int
	mTimeout        time.Duration

	mRW iface.IReaderWriter
}

func (c *Client) UpdateMaxBatchLength(length int) {
	c.mMaxBatchLength = length
}

func (c *Client) UpdateTimeout(timeout time.Duration) {
	c.mTimeout = timeout
}

func (c *Client) Setup(clusterID uint64,
	namespace string,
	rw iface.IReaderWriter,
	timeout time.Duration) error {
	if !utility.IsNamespaceValid([]byte(namespace)) {
		return x.ErrInvalidNamespace
	}

	if rw == nil {
		return x.ErrUnexpectedNilValue
	}

	c.mClusterID = clusterID
	c.mNamespace = namespace
	c.mRW = rw
	c.mTimeout = timeout
	c.mMaxBatchLength = 100
	return nil
}

func (c *Client) NewBatch() *Batch {
	return &Batch{
		mNamespace:      []byte(c.mNamespace),
		mMaxBatchLength: c.mMaxBatchLength,
		mTransactions:   nil,
	}
}

func (c *Client) ApplyBatch(b *Batch) (*pb.ProposalResponse, error) {
	if len(b.mTransactions) == 0 {
		return nil, ErrEmptyBatch
	}

	proposal := pb.NewProposal()
	proposal.Transactions = b.mTransactions
	r, err := c.mRW.Write(c.mClusterID, proposal, c.mTimeout)

	if err != nil {
		return nil, err
	}

	pr := &pb.ProposalResponse{}

	if err := proto.Unmarshal(r.Data, pr); err != nil {
		return nil, err
	}

	return pr, nil
}
