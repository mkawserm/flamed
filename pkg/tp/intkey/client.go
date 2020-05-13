package intkey

import (
	"fmt"
	"time"

	"github.com/golang/protobuf/proto"
	"github.com/google/uuid"
	"github.com/mkawserm/flamed/pkg/iface"
	"github.com/mkawserm/flamed/pkg/pb"
	"github.com/mkawserm/flamed/pkg/utility"
	"github.com/mkawserm/flamed/pkg/x"
)

type Client struct {
	mRW        iface.IRW
	mClusterID uint64
	mNamespace string
}

func (c *Client) GetIntKeyState(name string, timeout time.Duration) (*IntKeyState, error) {
	request := Request{
		Name:      name,
		Namespace: c.mNamespace,
	}

	r, err := c.mRW.Read(c.mClusterID, request, timeout)

	if err != nil {
		return nil, err
	}

	if v, ok := r.(*IntKeyState); ok {
		return v, nil
	}

	return nil, x.ErrUnknownValue
}

func (c *Client) applyPayload(payload *Payload, timeout time.Duration) error {
	payloadBytes, err := proto.Marshal(payload)

	if err != nil {
		return err
	}

	uuidValue := uuid.New()
	proposal := &pb.Proposal{
		Uuid:      uuidValue[:],
		CreatedAt: uint64(time.Now().UnixNano()),
		Transactions: []*pb.Transaction{
			{
				Payload:       payloadBytes,
				Namespace:     []byte(c.mNamespace),
				FamilyName:    Name,
				FamilyVersion: Version,
			},
		},
	}

	r, err := c.mRW.Write(c.mClusterID, proposal, timeout)

	if err != nil {
		return err
	}

	fmt.Println(r.Value)
	fmt.Println(r.Data)
	fmt.Println(err)

	return nil
}

func (c *Client) Insert(name string, value uint32, timeout time.Duration) error {
	payload := &Payload{
		Verb:  Verb_INSERT,
		Name:  name,
		Value: value,
	}

	return c.applyPayload(payload, timeout)
}

func (c *Client) Upsert(name string, value uint32, timeout time.Duration) error {
	payload := &Payload{
		Verb:  Verb_UPSERT,
		Name:  name,
		Value: value,
	}

	return c.applyPayload(payload, timeout)
}

func (c *Client) Delete(name string, timeout time.Duration) error {
	payload := &Payload{
		Verb: Verb_DELETE,
		Name: name,
	}

	return c.applyPayload(payload, timeout)
}

func (c *Client) Increment(name string, value uint32, timeout time.Duration) error {
	payload := &Payload{
		Verb:  Verb_INCREMENT,
		Name:  name,
		Value: value,
	}

	return c.applyPayload(payload, timeout)
}

func (c *Client) Decrement(name string, value uint32, timeout time.Duration) error {
	payload := &Payload{
		Verb:  Verb_DECREMENT,
		Name:  name,
		Value: value,
	}

	return c.applyPayload(payload, timeout)
}

func NewClient(namespace string, clusterID uint64, rw iface.IRW) (*Client, error) {
	if !utility.IsNamespaceValid([]byte(namespace)) {
		return nil, x.ErrInvalidNamespace
	}

	if rw == nil {
		return nil, x.ErrUnexpectedNilValue
	}

	return &Client{
		mRW:        rw,
		mClusterID: clusterID,
		mNamespace: namespace,
	}, nil
}
