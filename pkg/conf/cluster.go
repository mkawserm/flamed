package conf

import (
	sm "github.com/mkawserm/dragonboat/v3/statemachine"
	"github.com/mkawserm/flamed/pkg/iface"
	"github.com/mkawserm/flamed/pkg/storaged"
)

type ClusterConfigurationInput struct {
	ClusterID      uint64                                      `json:"clusterID"`
	ClusterName    string                                      `json:"clusterName"`
	InitialMembers map[uint64]string                           `json:"initialMembers"`
	Join           bool                                        `json:"join"`
	StateMachine   func(uint64, uint64) sm.IOnDiskStateMachine `json:"-"`
}

type ClusterConfiguration struct {
	ClusterConfigurationInput ClusterConfigurationInput
}

func (c *ClusterConfiguration) ClusterID() uint64 {
	if c.ClusterConfigurationInput.ClusterID == 0 {
		return 1
	} else {
		return c.ClusterConfigurationInput.ClusterID
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

func (c *ClusterConfiguration) StateMachine(sc iface.IStoragedConfiguration) func(uint64, uint64) sm.IOnDiskStateMachine {
	if c.ClusterConfigurationInput.StateMachine == nil {
		return storaged.NewStoraged(sc)
	} else {
		return c.ClusterConfigurationInput.StateMachine
	}
}
