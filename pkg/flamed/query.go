package flamed

import (
	"github.com/mkawserm/flamed/pkg/crypto"
	"github.com/mkawserm/flamed/pkg/iface"
	"github.com/mkawserm/flamed/pkg/pb"
	"github.com/mkawserm/flamed/pkg/utility"
	"github.com/mkawserm/flamed/pkg/x"
	"time"
)

type Query struct {
	mNamespace string
	mClusterID uint64
	mTimeout   time.Duration
	mRW        iface.IReaderWriter
}

func (q *Query) Namespace() string {
	return q.mNamespace
}

func (q *Query) ClusterID() uint64 {
	return q.mClusterID
}

func (q *Query) Timeout() time.Duration {
	return q.mTimeout
}

func (q *Query) UpdateTimeout(timeout time.Duration) {
	q.mTimeout = timeout
}

func (q *Query) Setup(clusterID uint64, namespace string, rw iface.IReaderWriter, timeout time.Duration) error {
	if !utility.IsNamespaceValid([]byte(namespace)) {
		return x.ErrInvalidNamespace
	}

	if rw == nil {
		return x.ErrUnexpectedNilValue
	}

	q.mClusterID = clusterID
	q.mNamespace = namespace
	q.mRW = rw
	q.mTimeout = timeout
	return nil
}

func (q *Query) Search(globalSearchInput *pb.GlobalSearchInput) (iface.ISearchResult, error) {
	output, err := q.mRW.Read(q.mClusterID, globalSearchInput, q.mTimeout)
	if err != nil {
		return nil, err
	}

	if v, ok := output.(iface.ISearchResult); ok {
		return v, nil
	} else {
		return nil, x.ErrUnknownValue
	}
}

func (q *Query) Retrieve(addresses []interface{}) (interface{}, error) {
	globalRetrieveInput := &pb.GlobalRetrieveInput{Namespace: []byte(q.mNamespace)}

	for _, addr := range addresses {
		addrString := addr.(string)
		addrBytes := crypto.GetStateAddressFromHexString(addrString)
		globalRetrieveInput.Addresses = append(globalRetrieveInput.Addresses, addrBytes)
	}
	return q.mRW.Read(q.mClusterID, globalRetrieveInput, q.mTimeout)
}
