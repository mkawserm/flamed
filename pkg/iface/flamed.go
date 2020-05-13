package iface

type IFlamedConfiguration interface {
	NodeConfiguration() INodeConfiguration
	StoragedConfiguration() IStoragedConfiguration
}
