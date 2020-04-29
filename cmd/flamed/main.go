package main

import "github.com/mkawserm/flamed/pkg/conf"
import "github.com/mkawserm/flamed/pkg/node"

func main() {
	n := &node.Node{}
	defer n.StopNode()

	err := n.ConfigureNode(
		conf.SimpleNodeConfiguration(1, "/tmp/1", "localhost:63001"),
		conf.SimpleStoragedConfiguration("/tmp/1", nil),
	)

	if err != nil {
		panic(err)
	}

	err = n.StartCluster(
		conf.SimpleClusterConfiguration(1,
			"example",
			map[uint64]string{1: "localhost:63001"}, false),
	)

	if err != nil {
		panic(err)
	}

}
