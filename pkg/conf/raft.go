package conf

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
	RaftConfigurationInput RaftConfigurationInput
}

func (c *RaftConfiguration) NodeID() uint64 {
	return c.RaftConfigurationInput.NodeID
}

func (c *RaftConfiguration) CheckQuorum() bool {
	return c.RaftConfigurationInput.CheckQuorum
}

func (c *RaftConfiguration) ElectionRTT() uint64 {
	return c.RaftConfigurationInput.ElectionRTT
}

func (c *RaftConfiguration) HeartbeatRTT() uint64 {
	return c.RaftConfigurationInput.HeartbeatRTT
}

func (c *RaftConfiguration) SnapshotEntries() uint64 {
	return c.RaftConfigurationInput.SnapshotEntries
}

func (c *RaftConfiguration) CompactionOverhead() uint64 {
	return c.RaftConfigurationInput.CompactionOverhead
}

func (c *RaftConfiguration) OrderedConfigChange() bool {
	return c.RaftConfigurationInput.OrderedConfigChange
}

func (c *RaftConfiguration) MaxInMemLogSize() uint64 {
	return c.RaftConfigurationInput.MaxInMemLogSize
}

func (c *RaftConfiguration) DisableAutoCompactions() bool {
	return c.RaftConfigurationInput.DisableAutoCompactions
}

func (c *RaftConfiguration) IsObserver() bool {
	return c.RaftConfigurationInput.IsObserver
}

func (c *RaftConfiguration) IsWitness() bool {
	return c.RaftConfigurationInput.IsWitness
}

func (c *RaftConfiguration) Quiesce() bool {
	return c.RaftConfigurationInput.Quiesce
}
