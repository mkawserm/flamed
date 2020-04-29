package main

import (
	"bufio"
	"fmt"
	"github.com/mkawserm/flamed/pkg/conf"
	"os"
	"strings"
)
import "github.com/mkawserm/flamed/pkg/node"

func main() {
	n := &node.Node{}
	defer n.StopNode()

	err := n.ConfigureNode(
		conf.SimpleNodeConfiguration(1, "/tmp/1", "/tmp/1", "localhost:63001"),
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

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Printf(">> ")
		text, _ := reader.ReadString('\n')
		t := strings.Trim(text, "\n")

		switch t {
		case "tcs":
			fmt.Println(n.TotalCluster())
		case "quit":
			break
		case "q":
			break
		case "":
		}
	}
}
