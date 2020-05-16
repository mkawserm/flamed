package json

import (
	"github.com/mkawserm/flamed/pkg/iface"
	"github.com/mkawserm/flamed/pkg/utility"
	"github.com/mkawserm/flamed/pkg/x"
	"time"
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
