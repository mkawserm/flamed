package main

import "fmt"
import "github.com/mkawserm/flamed"
import "github.com/mkawserm/flamed/pkg/node"

func main() {
	fmt.Println("Hello world")
	node := &node.Node{}
	defer node.StopNode()

	err := node.ConfigureNode(
		flamed.SimpleNodeConfiguration(1, "/tmp/1", "localhost:63001"),
		flamed.SimpleStoragedConfiguration("/tmp/1", nil),
	)

	if err != nil {
		panic(err)
	}

	err = node.StartCluster(
		flamed.SimpleClusterConfiguration(1,
			"example",
			map[uint64]string{1: "localhost:63001"}, false),
	)

	if err != nil {
		panic(err)
	}

}
