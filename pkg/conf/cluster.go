package conf

import (
	"github.com/mkawserm/flamed/pkg/iface"
	"github.com/mkawserm/flamed/pkg/storaged"
)

type OnDiskClusterConfigurationInput struct {
	ClusterID      uint64                               `json:"clusterID"`
	ClusterName    string                               `json:"clusterName"`
	InitialMembers map[uint64]string                    `json:"initialMembers"`
	Join           bool                                 `json:"join"`
	StateMachine   func(uint64, uint64) iface.IStoraged `json:"-"`
}

type OnDiskClusterConfiguration struct {
	OnDiskClusterConfigurationInput OnDiskClusterConfigurationInput
}

func (c *OnDiskClusterConfiguration) ClusterID() uint64 {
	if c.OnDiskClusterConfigurationInput.ClusterID == 0 {
		return 1
	} else {
		return c.OnDiskClusterConfigurationInput.ClusterID
	}
}

func (c *OnDiskClusterConfiguration) ClusterName() string {
	return c.OnDiskClusterConfigurationInput.ClusterName
}

func (c *OnDiskClusterConfiguration) InitialMembers() map[uint64]string {
	if c.OnDiskClusterConfigurationInput.InitialMembers == nil {
		c.OnDiskClusterConfigurationInput.InitialMembers = map[uint64]string{}
	}

	return c.OnDiskClusterConfigurationInput.InitialMembers
}

func (c *OnDiskClusterConfiguration) Join() bool {
	return c.OnDiskClusterConfigurationInput.Join
}

func (c *OnDiskClusterConfiguration) StateMachine(sc iface.IStoragedConfiguration) func(uint64, uint64) iface.IStoraged {
	if c.OnDiskClusterConfigurationInput.StateMachine == nil {
		c.OnDiskClusterConfigurationInput.StateMachine = storaged.NewStoraged(sc)
	}

	return c.OnDiskClusterConfigurationInput.StateMachine
}
