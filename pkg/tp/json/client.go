package json

import (
	"encoding/json"
	"errors"
	"github.com/golang/protobuf/proto"
	"github.com/mkawserm/flamed/pkg/iface"
	"github.com/mkawserm/flamed/pkg/pb"
	"github.com/mkawserm/flamed/pkg/utility"
	"github.com/mkawserm/flamed/pkg/variant"
	"github.com/mkawserm/flamed/pkg/x"
	"time"
)

var ErrInvalidID = errors.New("invalid id")
var ErrEmptyBatch = errors.New("empty batch")
var ErrUnexpectedNilValue = errors.New("unexpected nil value")
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

func (b *Batch) validateID(data interface{}) error {
	id := GetId(data)
	if id == "" {
		return ErrInvalidID
	}

	return nil
}

func (b *Batch) appendInternalTransaction(action Action, data interface{}) error {
	if err := b.validateID(data); err != nil {
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

	payload, err := proto.Marshal(jsonPayload)
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

func (b *Batch) appendInternalJSONMapTransaction(action Action, data map[string]interface{}) error {
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

	payload, err := proto.Marshal(jsonPayload)
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

func (b *Batch) Clear() {
	b.Reset()
}

func (b *Batch) Merge(data interface{}) error {
	return b.appendInternalTransaction(Action_MERGE, data)
}

func (b *Batch) MergeJSONMap(data map[string]interface{}) error {
	return b.appendInternalJSONMapTransaction(Action_MERGE, data)
}

func (b *Batch) Insert(data interface{}) error {
	return b.appendInternalTransaction(Action_INSERT, data)
}

func (b *Batch) InsertJSONMap(data map[string]interface{}) error {
	return b.appendInternalJSONMapTransaction(Action_INSERT, data)
}

func (b *Batch) Update(data interface{}) error {
	return b.appendInternalTransaction(Action_UPDATE, data)
}

func (b *Batch) UpdateJSONMap(data map[string]interface{}) error {
	return b.appendInternalJSONMapTransaction(Action_UPDATE, data)
}

func (b *Batch) Upsert(data interface{}) error {
	return b.appendInternalTransaction(Action_UPSERT, data)
}

func (b *Batch) UpsertJSONMap(data map[string]interface{}) error {
	return b.appendInternalJSONMapTransaction(Action_UPSERT, data)
}

func (b *Batch) Delete(data interface{}) error {
	return b.appendInternalTransaction(Action_DELETE, data)
}

func (b *Batch) DeleteJSONMap(data map[string]interface{}) error {
	return b.appendInternalJSONMapTransaction(Action_DELETE, data)
}

func (b *Batch) AppendTransaction(txn *pb.Transaction) error {
	if txn == nil {
		return ErrUnexpectedNilValue
	}

	if b.mTransactions == nil {
		b.mTransactions = make([]*pb.Transaction, 0, b.mMaxBatchLength)
	}
	b.mTransactions = append(b.mTransactions, txn)

	return nil
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

func (c *Client) Get(id string, object interface{}) (interface{}, error) {
	data, err := c.mRW.Read(c.mClusterID, variant.LookupRequest{
		Query:         id,
		Context:       nil,
		FamilyName:    Name,
		FamilyVersion: Version,
	}, c.mTimeout)

	if err != nil {
		return nil, err
	}

	if v, ok := data.([]byte); ok {
		if err := json.Unmarshal(v, object); err != nil {
			return nil, err
		} else {
			return object, nil
		}
	}

	return nil, x.ErrUnknownValue
}

func (c *Client) GetBytes(id string) ([]byte, error) {
	data, err := c.mRW.Read(c.mClusterID, variant.LookupRequest{
		Query:         id,
		Context:       nil,
		FamilyName:    Name,
		FamilyVersion: Version,
	}, c.mTimeout)

	if err != nil {
		return nil, err
	}

	if v, ok := data.([]byte); ok {
		return v, nil
	}

	return nil, x.ErrUnknownValue
}

func (c *Client) GetJSONMap(id string) (map[string]interface{}, error) {
	data, err := c.mRW.Read(c.mClusterID, variant.LookupRequest{
		Query:         id,
		Context:       nil,
		FamilyName:    Name,
		FamilyVersion: Version,
	}, c.mTimeout)

	if err != nil {
		return nil, err
	}

	if v, ok := data.([]byte); ok {
		jsonMap := make(map[string]interface{})
		if err := json.Unmarshal(v, &jsonMap); err != nil {
			return nil, err
		} else {
			return jsonMap, nil
		}
	}

	return nil, x.ErrUnknownValue
}

func (c *Client) Merge(data interface{}) (*pb.ProposalResponse, error) {
	b := &Batch{
		mNamespace:      []byte(c.mNamespace),
		mMaxBatchLength: 1,
		mTransactions:   nil,
	}

	if err := b.Merge(data); err != nil {
		return nil, err
	}
	return c.ApplyBatch(b)
}

func (c *Client) MergeJSONMap(data map[string]interface{}) (*pb.ProposalResponse, error) {
	b := &Batch{
		mNamespace:      []byte(c.mNamespace),
		mMaxBatchLength: 1,
		mTransactions:   nil,
	}

	if err := b.MergeJSONMap(data); err != nil {
		return nil, err
	}
	return c.ApplyBatch(b)
}

func (c *Client) Insert(data interface{}) (*pb.ProposalResponse, error) {
	b := &Batch{
		mNamespace:      []byte(c.mNamespace),
		mMaxBatchLength: 1,
		mTransactions:   nil,
	}

	if err := b.Insert(data); err != nil {
		return nil, err
	}
	return c.ApplyBatch(b)
}

func (c *Client) InsertJSONMap(data map[string]interface{}) (*pb.ProposalResponse, error) {
	b := &Batch{
		mNamespace:      []byte(c.mNamespace),
		mMaxBatchLength: 1,
		mTransactions:   nil,
	}

	if err := b.InsertJSONMap(data); err != nil {
		return nil, err
	}
	return c.ApplyBatch(b)
}

func (c *Client) Upsert(data interface{}) (*pb.ProposalResponse, error) {
	b := &Batch{
		mNamespace:      []byte(c.mNamespace),
		mMaxBatchLength: 1,
		mTransactions:   nil,
	}

	if err := b.Upsert(data); err != nil {
		return nil, err
	}
	return c.ApplyBatch(b)
}

func (c *Client) UpsertJSONMap(data map[string]interface{}) (*pb.ProposalResponse, error) {
	b := &Batch{
		mNamespace:      []byte(c.mNamespace),
		mMaxBatchLength: 1,
		mTransactions:   nil,
	}

	if err := b.UpsertJSONMap(data); err != nil {
		return nil, err
	}
	return c.ApplyBatch(b)
}

func (c *Client) Update(data interface{}) (*pb.ProposalResponse, error) {
	b := &Batch{
		mNamespace:      []byte(c.mNamespace),
		mMaxBatchLength: 1,
		mTransactions:   nil,
	}

	if err := b.Update(data); err != nil {
		return nil, err
	}
	return c.ApplyBatch(b)
}

func (c *Client) UpdateJSONMap(data map[string]interface{}) (*pb.ProposalResponse, error) {
	b := &Batch{
		mNamespace:      []byte(c.mNamespace),
		mMaxBatchLength: 1,
		mTransactions:   nil,
	}

	if err := b.UpdateJSONMap(data); err != nil {
		return nil, err
	}
	return c.ApplyBatch(b)
}

func (c *Client) Delete(data interface{}) (*pb.ProposalResponse, error) {
	b := &Batch{
		mNamespace:      []byte(c.mNamespace),
		mMaxBatchLength: 1,
		mTransactions:   nil,
	}

	if err := b.Delete(data); err != nil {
		return nil, err
	}
	return c.ApplyBatch(b)
}

func (c *Client) DeleteJSONMap(data map[string]interface{}) (*pb.ProposalResponse, error) {
	b := &Batch{
		mNamespace:      []byte(c.mNamespace),
		mMaxBatchLength: 1,
		mTransactions:   nil,
	}

	if err := b.DeleteJSONMap(data); err != nil {
		return nil, err
	}
	return c.ApplyBatch(b)
}
