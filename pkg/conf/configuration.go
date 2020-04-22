package conf

type Configuration struct {
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
