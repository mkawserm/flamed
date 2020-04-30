package main

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"github.com/lni/dragonboat/v3"
	"github.com/mkawserm/flamed/pkg/conf"
	"github.com/mkawserm/flamed/pkg/utility"
	"os"
	"strings"
	"time"
)
import "github.com/mkawserm/flamed/pkg/nodehost"

func main() {
	n := &nodehost.NodeHost{}
	defer n.StopNode()

	err := n.ConfigureNode(
		conf.SimpleNodeHostConfiguration(1, "/tmp/1", "/tmp/1", "localhost:63001"),
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

	l := true
	reader := bufio.NewReader(os.Stdin)

	for l {
		fmt.Printf(">> ")
		text, _ := reader.ReadString('\n')
		t := strings.Trim(text, "\n")

		switch t {
		case "rs":
			ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
			r, e := n.GetDragonboatNodeHost().SyncRequestSnapshot(ctx, 1, dragonboat.SnapshotOption{
				CompactionOverhead:         0,
				ExportPath:                 "/tmp",
				Exported:                   false,
				OverrideCompactionOverhead: false,
			})

			if e != nil {
				fmt.Println(e)
			}
			fmt.Println(r)

			cancel()
		case "ci":
			nodeHostInfo := n.GetDragonboatNodeHost().GetNodeHostInfo(dragonboat.NodeHostInfoOption{SkipLogInfo: false})
			if b, err := json.Marshal(utility.LowerCamelCaseMarshaller{Value: nodeHostInfo}); err == nil {
				fmt.Println(string(b))
			}

		case "mi":
			ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
			m, err := n.GetDragonboatNodeHost().SyncGetClusterMembership(ctx, 1)
			if err != nil {
				panic(err)
			}

			if b, err := json.Marshal(m); err == nil {
				fmt.Println(string(b))
			}

			cancel()
		case "csid":
			fmt.Println(n.ClusterIdList())
		case "tcs":
			fmt.Println(n.TotalCluster())
		case "quit":
			l = false
			break
		case "q":
			l = false
			break
		}
	}
}
