package json

import "github.com/mkawserm/flamed/pkg/pb"

type Context struct {
	Client        *Client
	AccessControl *pb.AccessControl
}
