package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/mkawserm/flamed/pkg/conf"
	"github.com/mkawserm/flamed/pkg/pb"
	"os"
	"strings"
	"time"
	//"time"
)
import "github.com/mkawserm/flamed/pkg/flamed"

type CounterObject struct {
	Counter uint64 `json:"counter"`
}

func getJson(object *CounterObject) []byte {
	if data, err := json.Marshal(object); err == nil {
		return data
	} else {
		return nil
	}
}

func main() {
	n := &flamed.NodeHost{}
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

	manager := n.NewStorageManager(1)

	l := true
	reader := bufio.NewReader(os.Stdin)

	var counter uint64 = 0

	for l {
		fmt.Printf(">> ")
		text, _ := reader.ReadString('\n')
		t := strings.Trim(text, "\n")

		//e := &pb.FlameEntry{
		//	Namespace:            []byte("test"),
		//	Key:                  []byte("counter"),
		//}
		//
		//if data, err := proto.Marshal(e); err == nil {
		//	r, err := n.ManagedSyncRead(1, data, 3*time.Minute)
		//
		//	if err != nil {
		//		fmt.Println(err)
		//	}
		//
		//	counter = uidutil.ByteSliceToUint64(r.([]byte))
		//} else {
		//	fmt.Println(err)
		//}

		switch t {
		case "p":
			counter = counter + 1
			co := &CounterObject{}
			co.Counter = counter
			fmt.Println("counter to propose:", counter)
			b := manager.NewBatch("test")
			b.Create([]byte("counter"), getJson(co))

			if err := manager.ApplyBatch(b, 3*time.Minute); err != nil {
				fmt.Println(err)
			}
		case "r":
			e := &pb.FlameEntry{
				Namespace: []byte("test"),
				Key:       []byte("counter"),
			}

			if err := manager.Read(e, 3*time.Minute); err != nil {
				fmt.Println(err)
			} else {
				co := &CounterObject{}
				err := json.Unmarshal(e.Value, co)
				if err != nil {
					fmt.Println(err)
				}

				fmt.Println("Counter:", co.Counter)
				counter = co.Counter
			}

		case "rs":
			//ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
			//r, e := n.GetDragonboatNodeHost().SyncRequestSnapshot(ctx, 1, dragonboat.SnapshotOption{
			//	CompactionOverhead:         0,
			//	ExportPath:                 "/tmp",
			//	Exported:                   false,
			//	OverrideCompactionOverhead: false,
			//})
			//
			//if e != nil {
			//	fmt.Println(e)
			//}
			//fmt.Println(r)
			//
			//cancel()
		case "ci":
			//nodeHostInfo := n.GetDragonboatNodeHost().GetNodeHostInfo(dragonboat.NodeHostInfoOption{SkipLogInfo: false})
			//if b, err := json.Marshal(utility.LowerCamelCaseMarshaller{Value: nodeHostInfo}); err == nil {
			//	fmt.Println(string(b))
			//}

		case "mi":
			//ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
			//m, err := n.GetDragonboatNodeHost().SyncGetClusterMembership(ctx, 1)
			//if err != nil {
			//	panic(err)
			//}
			//
			//if b, err := json.Marshal(m); err == nil {
			//	fmt.Println(string(b))
			//}
			//
			//cancel()
		case "csid":
			fmt.Println(n.ClusterIDList())
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
