package conf

type Configuration struct {
	Path string
}

func (c *Configuration) FlamedPath() string {
	return c.Path
}

func (c *Configuration) StoragePluginKV() string {
	return "default"
}

func (c *Configuration) StoragePluginIndex() string {
	return "default"
}

func (c *Configuration) StoragePluginRaftLog() string {
	return "default"
}

func (c *Configuration) KVStorageCustomConfiguration() interface{} {
	return nil
}

func (c *Configuration) IndexStorageCustomConfiguration() interface{} {
	return nil
}

func (c *Configuration) RaftLogStorageCustomConfiguration() interface{} {
	return nil
}
