package conf

import (
	"github.com/mkawserm/flamed/pkg/iface"
	"github.com/mkawserm/flamed/pkg/storaged"
	"sync"
)

type OnDiskClusterConfigurationInput struct {
	ClusterID      uint64                               `json:"clusterID"`
	ClusterName    string                               `json:"clusterName"`
	InitialMembers map[uint64]string                    `json:"initialMembers"`
	Join           bool                                 `json:"join"`
	StateMachine   func(uint64, uint64) iface.IStoraged `json:"-"`
}

type OnDiskClusterConfiguration struct {
	mMutex                          sync.Mutex
	OnDiskClusterConfigurationInput OnDiskClusterConfigurationInput
}

func (c *OnDiskClusterConfiguration) ClusterID() uint64 {
	c.mMutex.Lock()
	defer c.mMutex.Unlock()
	if c.OnDiskClusterConfigurationInput.ClusterID == 0 {
		return 1
	} else {
		return c.OnDiskClusterConfigurationInput.ClusterID
	}
}

func (c *OnDiskClusterConfiguration) ClusterName() string {
	c.mMutex.Lock()
	defer c.mMutex.Unlock()
	return c.OnDiskClusterConfigurationInput.ClusterName
}

func (c *OnDiskClusterConfiguration) InitialMembers() map[uint64]string {
	c.mMutex.Lock()
	defer c.mMutex.Unlock()
	if c.OnDiskClusterConfigurationInput.InitialMembers == nil {
		c.OnDiskClusterConfigurationInput.InitialMembers = map[uint64]string{}
	}

	return c.OnDiskClusterConfigurationInput.InitialMembers
}

func (c *OnDiskClusterConfiguration) Join() bool {
	c.mMutex.Lock()
	defer c.mMutex.Unlock()
	return c.OnDiskClusterConfigurationInput.Join
}

func (c *OnDiskClusterConfiguration) StateMachine(sc iface.IStoragedConfiguration) func(uint64, uint64) iface.IStoraged {
	c.mMutex.Lock()
	defer c.mMutex.Unlock()
	if c.OnDiskClusterConfigurationInput.StateMachine == nil {
		c.OnDiskClusterConfigurationInput.StateMachine = storaged.NewStoraged(sc)
	}

	return c.OnDiskClusterConfigurationInput.StateMachine
}
