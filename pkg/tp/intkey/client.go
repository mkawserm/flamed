package intkey

import (
	"github.com/mkawserm/flamed/pkg/variant"
	"time"

	"github.com/golang/protobuf/proto"
	"github.com/mkawserm/flamed/pkg/iface"
	"github.com/mkawserm/flamed/pkg/pb"
	"github.com/mkawserm/flamed/pkg/utility"
	"github.com/mkawserm/flamed/pkg/x"
)

type Client struct {
	mClusterID uint64
	mNamespace string
	mTimeout   time.Duration
	mRW        iface.IReaderWriter
}

func (c *Client) UpdateTimeout(timeout time.Duration) {
	c.mTimeout = timeout
}

func (c *Client) GetIntKeyState(name string) (*IntKeyState, error) {
	request := Request{
		Name:      name,
		Namespace: c.mNamespace,
	}

	lookupRequest := variant.LookupRequest{
		Query:         request,
		Context:       nil,
		FamilyName:    Name,
		FamilyVersion: Version,
	}

	r, err := c.mRW.Read(c.mClusterID, lookupRequest, c.mTimeout)

	if err != nil {
		return nil, err
	}

	if v, ok := r.(*IntKeyState); ok {
		return v, nil
	}

	return nil, x.ErrUnknownValue
}

func (c *Client) sendProposal(payload *IntKeyPayload) (*pb.ProposalResponse, error) {
	payloadBytes, err := proto.Marshal(payload)

	if err != nil {
		return nil, err
	}

	proposal := pb.NewProposal()
	proposal.AddTransaction([]byte(c.mNamespace), Name, Version, payloadBytes)

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

func (c *Client) Insert(name string, value uint32) (*pb.ProposalResponse, error) {
	intKeyPayload := &IntKeyPayload{
		Verb:  Verb_INSERT,
		Name:  name,
		Value: value,
	}

	return c.sendProposal(intKeyPayload)
}

func (c *Client) Upsert(name string, value uint32) (*pb.ProposalResponse, error) {
	intKeyPayload := &IntKeyPayload{
		Verb:  Verb_UPSERT,
		Name:  name,
		Value: value,
	}

	return c.sendProposal(intKeyPayload)
}

func (c *Client) Delete(name string) (*pb.ProposalResponse, error) {
	intKeyPayload := &IntKeyPayload{
		Verb: Verb_DELETE,
		Name: name,
	}

	return c.sendProposal(intKeyPayload)
}

func (c *Client) Increment(name string, value uint32) (*pb.ProposalResponse, error) {
	intKeyPayload := &IntKeyPayload{
		Verb:  Verb_INCREMENT,
		Name:  name,
		Value: value,
	}

	return c.sendProposal(intKeyPayload)
}

func (c *Client) Decrement(name string, value uint32) (*pb.ProposalResponse, error) {
	intKeyPayload := &IntKeyPayload{
		Verb:  Verb_DECREMENT,
		Name:  name,
		Value: value,
	}

	return c.sendProposal(intKeyPayload)
}

func (c *Client) Setup(clusterID uint64, namespace string, rw iface.IReaderWriter, timeout time.Duration) error {
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
	return nil
}
