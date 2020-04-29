package main

import "fmt"
import "github.com/mkawserm/flamed/pkg/conf"
import "github.com/mkawserm/flamed/pkg/node"

func main() {
	fmt.Println("Hello world")
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
		fmt.Println("panic")
		panic(err)
	}

}
