package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/mkawserm/flamed/pkg/conf"
	"github.com/mkawserm/flamed/pkg/pb"
	"github.com/mkawserm/flamed/pkg/utility"
	"os"
	"strings"
	"time"
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
	//members := map[uint64]string{1: "localhost:63001", 2: "localhost:63002", 3: "localhost:63003"}
	members := map[uint64]string{1: "localhost:63001"}
	var clusterId uint64 = 1

	// flame 1
	flame1 := flamed.NewFlamed()
	defer flame1.StopNode()

	err := flame1.ConfigureNode(
		conf.SimpleNodeHostConfiguration(1, "/tmp/1/nh", "/tmp/1/wal", "localhost:63001"),
		conf.SimpleStoragedConfiguration("/tmp/1/storage", nil),
	)

	if err != nil {
		panic(err)
	}

	err = flame1.StartCluster(
		conf.SimpleClusterConfiguration(clusterId, "example", members, false),
	)

	if err != nil {
		panic(err)
	}

	manager1 := flame1.NewStorageManager(clusterId)
	admin1 := flame1.NewAdmin(clusterId)

	//// flame 2
	//flame2 := flamed.NewFlamed()
	//defer flame2.StopNode()
	//
	//err = flame2.ConfigureNode(
	//	conf.SimpleNodeHostConfiguration(2, "/tmp/2/nh", "/tmp/2/wal", "localhost:63002"),
	//	conf.SimpleStoragedConfiguration("/tmp/2/storage", nil),
	//)
	//
	//if err != nil {
	//	panic(err)
	//}
	//
	//err = flame2.StartCluster(
	//	conf.SimpleClusterConfiguration(clusterId, "example", nil, true),
	//)
	//
	//if err != nil {
	//	panic(err)
	//}
	//
	//manager2 := flame2.NewStorageManager(clusterId)
	//admin2 := flame2.NewAdmin(clusterId)
	//
	//// flame 3
	//flame3 := flamed.NewFlamed()
	//defer flame3.StopNode()
	//
	//err = flame3.ConfigureNode(
	//	conf.SimpleNodeHostConfiguration(3, "/tmp/3/nh", "/tmp/3/wal", "localhost:63003"),
	//	conf.SimpleStoragedConfiguration("/tmp/3/storage", nil),
	//)
	//
	//if err != nil {
	//	panic(err)
	//}
	//
	//err = flame3.StartCluster(
	//	conf.SimpleClusterConfiguration(clusterId, "example", nil, true),
	//)
	//
	//if err != nil {
	//	panic(err)
	//}
	//
	//manager3 := flame3.NewStorageManager(clusterId)
	//admin3 := flame3.NewAdmin(clusterId)

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
		//	r, err := flame1.ManagedSyncRead(1, data, 3*time.Minute)
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
		case "ai1":
			index := admin1.QueryAppliedIndex(3 * time.Minute)
			fmt.Println(index)
		//case "ai2":
		//	index := admin2.QueryAppliedIndex(3 * time.Minute)
		//	fmt.Println(index)
		//case "ai3":
		//	index := admin3.QueryAppliedIndex(3 * time.Minute)
		//	fmt.Println(index)

		case "p1":
			counter = counter + 1
			key := fmt.Sprintf("counter-%d", counter)
			co := &CounterObject{}
			fmt.Println("counter to propose:", counter)
			b := manager1.NewBatch("test")

			co.Counter = counter
			b.Create([]byte(key), getJson(co))

			//co.Counter = counter + 1
			//b.Create([]byte("counter2"), getJson(co))

			if err := manager1.ApplyBatch(b, 3*time.Minute); err != nil {
				fmt.Println(err)
			}

		//case "p2":
		//	counter = counter + 1
		//	key := fmt.Sprintf("counter-%d", counter)
		//	co := &CounterObject{}
		//	fmt.Println("counter to propose:", counter)
		//	b := manager2.NewBatch("test")
		//
		//	co.Counter = counter
		//	b.Create([]byte(key), getJson(co))
		//
		//	//co.Counter = counter + 1
		//	//b.Create([]byte("counter2"), getJson(co))
		//
		//	if err := manager2.ApplyBatch(b, 3*time.Minute); err != nil {
		//		fmt.Println(err)
		//	}
		//case "p3":
		//	counter = counter + 1
		//	key := fmt.Sprintf("counter-%d", counter)
		//	co := &CounterObject{}
		//	fmt.Println("counter to propose:", counter)
		//	b := manager3.NewBatch("test")
		//
		//	co.Counter = counter
		//	b.Create([]byte(key), getJson(co))
		//
		//	//co.Counter = counter + 1
		//	//b.Create([]byte("counter2"), getJson(co))
		//
		//	if err := manager3.ApplyBatch(b, 3*time.Minute); err != nil {
		//		fmt.Println(err)
		//	}
		case "r1":
			e := &pb.FlameEntry{
				Namespace: []byte("test"),
				Key:       []byte("counter-1"),
			}

			if err := manager1.Read(e, 3*time.Minute); err != nil {
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
		//case "r2":
		//	e := &pb.FlameEntry{
		//		Namespace: []byte("test"),
		//		Key:       []byte("counter-1"),
		//	}
		//
		//	if err := manager2.Read(e, 3*time.Minute); err != nil {
		//		fmt.Println(err)
		//	} else {
		//		co := &CounterObject{}
		//		err := json.Unmarshal(e.Value, co)
		//		if err != nil {
		//			fmt.Println(err)
		//		}
		//
		//		fmt.Println("Counter:", co.Counter)
		//		counter = co.Counter
		//	}
		//case "r3":
		//	e := &pb.FlameEntry{
		//		Namespace: []byte("test"),
		//		Key:       []byte("counter-1"),
		//	}
		//
		//	if err := manager3.Read(e, 3*time.Minute); err != nil {
		//		fmt.Println(err)
		//	} else {
		//		co := &CounterObject{}
		//		err := json.Unmarshal(e.Value, co)
		//		if err != nil {
		//			fmt.Println(err)
		//		}
		//
		//		fmt.Println("Counter:", co.Counter)
		//		counter = co.Counter
		//	}
		case "i1":
			data, err := manager1.Iterate(&pb.FlameEntry{
				Namespace: []byte("test"),
			}, 0, 3*time.Minute)

			if err != nil {
				fmt.Println(err)
			}

			fmt.Println(data)

			co := &CounterObject{}
			err = json.Unmarshal(data[len(data)-1].Value, co)
			if err != nil {
				fmt.Println(err)
			}

			fmt.Println("Counter:", co.Counter)
			counter = co.Counter

			//data, err = manager1.Iterate(data[0], 1, 3*time.Minute)
			//
			//if err != nil {
			//	fmt.Println(err)
			//}
			//fmt.Println(data)
		//case "i2":
		//	data, err := manager2.Iterate(&pb.FlameEntry{
		//		Namespace: []byte("test"),
		//	}, 0, 3*time.Minute)
		//
		//	if err != nil {
		//		fmt.Println(err)
		//	}
		//
		//	fmt.Println(data)
		//case "i3":
		//	data, err := manager3.Iterate(&pb.FlameEntry{
		//		Namespace: []byte("test"),
		//	}, 0, 3*time.Minute)
		//
		//	if err != nil {
		//		fmt.Println(err)
		//	}
		//
		//	fmt.Println(data)
		case "rs":
			//ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
			//r, e := flame1.GetDragonboatNodeHost().SyncRequestSnapshot(ctx, 1, dragonboat.SnapshotOption{
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
		case "ci1":
			nodeHostInfo := flame1.GetNodeHostInfo()
			if b, err := json.Marshal(utility.LowerCamelCaseMarshaller{Value: nodeHostInfo}); err == nil {
				fmt.Println(string(b))
			}
		//case "ci2":
		//	nodeHostInfo := flame2.GetNodeHostInfo()
		//	if b, err := json.Marshal(utility.LowerCamelCaseMarshaller{Value: nodeHostInfo}); err == nil {
		//		fmt.Println(string(b))
		//	}
		//case "ci3":
		//	nodeHostInfo := flame3.GetNodeHostInfo()
		//	if b, err := json.Marshal(utility.LowerCamelCaseMarshaller{Value: nodeHostInfo}); err == nil {
		//		fmt.Println(string(b))
		//	}
		case "mi":
			//ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
			//m, err := flame1.GetDragonboatNodeHost().SyncGetClusterMembership(ctx, 1)
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
			fmt.Println(flame1.ClusterIDList())
		case "tcs":
			fmt.Println(flame1.TotalCluster())
		case "quit":
			l = false
			break
		case "q":
			l = false
			break
		}
	}
}
