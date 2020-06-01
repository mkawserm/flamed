package flamed

import (
	"github.com/golang/protobuf/proto"
	"github.com/mkawserm/flamed/pkg/crypto"
	"github.com/mkawserm/flamed/pkg/iface"
	"github.com/mkawserm/flamed/pkg/pb"
	"github.com/mkawserm/flamed/pkg/utility"
	"github.com/mkawserm/flamed/pkg/x"
	"time"
)

type GlobalOperation struct {
	mNamespace string
	mClusterID uint64
	mTimeout   time.Duration
	mRW        iface.IReaderWriter
}

func (g *GlobalOperation) Namespace() string {
	return g.mNamespace
}

func (g *GlobalOperation) ClusterID() uint64 {
	return g.mClusterID
}

func (g *GlobalOperation) Timeout() time.Duration {
	return g.mTimeout
}

func (g *GlobalOperation) UpdateTimeout(timeout time.Duration) {
	g.mTimeout = timeout
}

func (g *GlobalOperation) Setup(clusterID uint64, namespace string, rw iface.IReaderWriter, timeout time.Duration) error {
	if !utility.IsNamespaceValid([]byte(namespace)) {
		return x.ErrInvalidNamespace
	}

	if rw == nil {
		return x.ErrUnexpectedNilValue
	}

	g.mClusterID = clusterID
	g.mNamespace = namespace
	g.mRW = rw
	g.mTimeout = timeout
	return nil
}

func (g *GlobalOperation) Search(globalSearchInput *pb.GlobalSearchInput) (iface.ISearchResult, error) {
	output, err := g.mRW.Read(g.mClusterID, globalSearchInput, g.mTimeout)
	if err != nil {
		return nil, err
	}

	if v, ok := output.(iface.ISearchResult); ok {
		return v, nil
	} else {
		return nil, x.ErrUnknownValue
	}
}

func (g *GlobalOperation) Retrieve(addresses []interface{}) (interface{}, error) {
	globalRetrieveInput := &pb.GlobalRetrieveInput{Namespace: []byte(g.mNamespace)}

	for _, addr := range addresses {
		addrString := addr.(string)
		addrBytes := crypto.GetStateAddressFromHexString(addrString)
		globalRetrieveInput.Addresses = append(globalRetrieveInput.Addresses, addrBytes)
	}
	return g.mRW.Read(g.mClusterID, globalRetrieveInput, g.mTimeout)
}

func (g *GlobalOperation) Iterate(from string, prefix string, limit uint64) (interface{}, error) {
	input := &pb.GlobalIterateInput{
		Namespace: []byte(g.mNamespace),
		Prefix:    []byte(g.mNamespace),
		Limit:     limit,
	}

	if from == "" {
		input.From = nil
	} else {
		input.From = crypto.GetStateAddressFromHexString(from)
	}

	if prefix != "" {
		input.Prefix = crypto.GetStateAddressFromHexString(prefix)
	}

	return g.mRW.Read(g.mClusterID, input, g.mTimeout)
}

func (g *GlobalOperation) Propose(proposal *pb.Proposal) (*pb.ProposalResponse, error) {
	r, err := g.mRW.Write(g.mClusterID, proposal, g.mTimeout)

	if err != nil {
		return nil, err
	}

	pr := &pb.ProposalResponse{}

	if err := proto.Unmarshal(r.Data, pr); err != nil {
		return nil, err
	}

	return pr, nil
}
