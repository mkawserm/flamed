package flamed

import (
	"github.com/mkawserm/flamed/pkg/iface"
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

func (a *Query) UpdateTimeout(timeout time.Duration) {
	a.mTimeout = timeout
}

func (c *Query) Setup(clusterID uint64, namespace string, rw iface.IReaderWriter, timeout time.Duration) error {
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
