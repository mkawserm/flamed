package conf

import "sync"

type RaftConfigurationInput struct {
	NodeID       uint64 `json:"nodeId"`
	CheckQuorum  bool   `json:"checkQuorum"`
	ElectionRTT  uint64 `json:"electionRTT"`
	HeartbeatRTT uint64 `json:"heartbeatRTT"`

	SnapshotEntries     uint64 `json:"snapshotEntries"`
	CompactionOverhead  uint64 `json:"compactionOverhead"`
	OrderedConfigChange bool   `json:"orderedConfigChange"`
	MaxInMemLogSize     uint64 `json:"maxInMemLogSize"`

	DisableAutoCompactions bool `json:"disableAutoCompactions"`

	IsObserver bool `json:"isObserver"`
	IsWitness  bool `json:"isWitness"`
	Quiesce    bool `json:"quiesce"`
}

type RaftConfiguration struct {
	mMutex                 sync.Mutex
	RaftConfigurationInput RaftConfigurationInput
}

func (c *RaftConfiguration) NodeID() uint64 {
	c.mMutex.Lock()
	defer c.mMutex.Unlock()
	return c.RaftConfigurationInput.NodeID
}

func (c *RaftConfiguration) CheckQuorum() bool {
	c.mMutex.Lock()
	defer c.mMutex.Unlock()
	return c.RaftConfigurationInput.CheckQuorum
}

func (c *RaftConfiguration) ElectionRTT() uint64 {
	c.mMutex.Lock()
	defer c.mMutex.Unlock()
	return c.RaftConfigurationInput.ElectionRTT
}

func (c *RaftConfiguration) HeartbeatRTT() uint64 {
	c.mMutex.Lock()
	defer c.mMutex.Unlock()
	return c.RaftConfigurationInput.HeartbeatRTT
}

func (c *RaftConfiguration) SnapshotEntries() uint64 {
	c.mMutex.Lock()
	defer c.mMutex.Unlock()
	return c.RaftConfigurationInput.SnapshotEntries
}

func (c *RaftConfiguration) CompactionOverhead() uint64 {
	c.mMutex.Lock()
	defer c.mMutex.Unlock()
	return c.RaftConfigurationInput.CompactionOverhead
}

func (c *RaftConfiguration) OrderedConfigChange() bool {
	c.mMutex.Lock()
	defer c.mMutex.Unlock()
	return c.RaftConfigurationInput.OrderedConfigChange
}

func (c *RaftConfiguration) MaxInMemLogSize() uint64 {
	c.mMutex.Lock()
	defer c.mMutex.Unlock()
	return c.RaftConfigurationInput.MaxInMemLogSize
}

func (c *RaftConfiguration) DisableAutoCompactions() bool {
	c.mMutex.Lock()
	defer c.mMutex.Unlock()
	return c.RaftConfigurationInput.DisableAutoCompactions
}

func (c *RaftConfiguration) IsObserver() bool {
	c.mMutex.Lock()
	defer c.mMutex.Unlock()
	return c.RaftConfigurationInput.IsObserver
}

func (c *RaftConfiguration) IsWitness() bool {
	c.mMutex.Lock()
	defer c.mMutex.Unlock()
	return c.RaftConfigurationInput.IsWitness
}

func (c *RaftConfiguration) Quiesce() bool {
	c.mMutex.Lock()
	defer c.mMutex.Unlock()
	return c.RaftConfigurationInput.Quiesce
}
