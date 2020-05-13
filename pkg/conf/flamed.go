package conf

import "github.com/mkawserm/flamed/pkg/iface"

type FlamedConfiguration struct {
	NodeConfigurationInput     iface.INodeConfiguration
	StoragedConfigurationInput iface.IStoragedConfiguration
}

func (n *FlamedConfiguration) NodeConfiguration() iface.INodeConfiguration {
	return n.NodeConfigurationInput
}

func (n *FlamedConfiguration) StoragedConfiguration() iface.IStoragedConfiguration {
	return n.StoragedConfigurationInput
}
