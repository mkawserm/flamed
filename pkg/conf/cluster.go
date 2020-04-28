package conf

import (
	"github.com/mkawserm/flamed/pkg/iface"
	"github.com/mkawserm/flamed/pkg/storaged"
)

type ClusterConfigurationInput struct {
	ClusterId      uint64                    `json:"clusterId"`
	ClusterName    string                    `json:"clusterName"`
	InitialMembers map[uint64]string         `json:"initialMembers"`
	Join           bool                      `json:"join"`
	StateMachine   iface.NewStateMachineFunc `json:"-"`
}

type ClusterConfiguration struct {
	ClusterConfigurationInput ClusterConfigurationInput
}

func (c *ClusterConfiguration) ClusterId() uint64 {
	if c.ClusterConfigurationInput.ClusterId == 0 {
		return 1
	} else {
		return c.ClusterConfigurationInput.ClusterId
	}
}

func (c *ClusterConfiguration) ClusterName() string {
	return c.ClusterConfigurationInput.ClusterName
}

func (c *ClusterConfiguration) InitialMembers() map[uint64]string {
	if c.ClusterConfigurationInput.InitialMembers == nil {
		return map[uint64]string{}
	} else {
		return c.ClusterConfigurationInput.InitialMembers
	}
}

func (c *ClusterConfiguration) Join() bool {
	return c.ClusterConfigurationInput.Join
}

func (c *ClusterConfiguration) StateMachine() iface.NewStateMachineFunc {
	if c.ClusterConfigurationInput.StateMachine == nil {
		return c.ClusterConfigurationInput.StateMachine
	} else {
		return storaged.NewStoraged
	}
}
