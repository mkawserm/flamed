package app

import (
	"fmt"
	"github.com/lni/dragonboat/v3/config"
	"github.com/mkawserm/flamed/pkg/conf"
	"github.com/spf13/cobra"
	"net/http"
	"strings"
)

var join bool
var dataStoragePath string
var raftAddress string
var httpAddress string
var initialMembers string

var runCMD = &cobra.Command{
	Use:   "run",
	Short: "Run Flamed server",
	Run: func(cmd *cobra.Command, args []string) {
		if len(dataStoragePath) == 0 {
			fmt.Println("data-storage-path can not be empty")
			return
		}

		if len(raftAddress) == 0 {
			fmt.Println("raft-address can not be empty")
			return
		}

		if len(httpAddress) == 0 {
			fmt.Println("http-address can not be empty")
			return
		}

		im := make(map[uint64]string)
		stringList := strings.Split(initialMembers, ",")

		for idx, value := range stringList {
			var index = uint64(1 + idx)
			im[index] = strings.TrimSpace(value)
		}

		if !join {
			im[1] = raftAddress
		}

		raftStoragePath := dataStoragePath + "/raft"

		nodeConfiguration := &conf.NodeConfiguration{
			NodeConfigurationInput: conf.NodeConfigurationInput{
				NodeHostDir:                   raftStoragePath,
				WALDir:                        raftStoragePath,
				DeploymentID:                  1,
				RTTMillisecond:                200,
				RaftAddress:                   raftAddress,
				ListenAddress:                 "",
				MutualTLS:                     false,
				CAFile:                        "",
				CertFile:                      "",
				KeyFile:                       "",
				MaxSendQueueSize:              0,
				MaxReceiveQueueSize:           0,
				LogDBFactory:                  nil,
				RaftRPCFactory:                nil,
				EnableMetrics:                 false,
				MaxSnapshotSendBytesPerSecond: 0,
				MaxSnapshotRecvBytesPerSecond: 0,
				LogDBConfig:                   config.GetTinyMemLogDBConfig(),
				NotifyCommit:                  false,
			},
		}

		err := GetApp().GetFlamed().ConfigureNode(nodeConfiguration)

		if err != nil {
			panic(err)
		}

		clusterStoragePath := dataStoragePath + "/cluster-1"

		storagedConfiguration := conf.SimpleStoragedConfiguration(clusterStoragePath, nil)
		for _, tp := range GetApp().mTransactionProcessor {
			storagedConfiguration.AddTransactionProcessor(tp)
		}

		clusterConfiguration := conf.SimpleOnDiskClusterConfiguration(
			1,
			"cluster-1",
			im,
			join)

		raftConfiguration := &conf.RaftConfiguration{
			RaftConfigurationInput: conf.RaftConfigurationInput{
				NodeID:                 1,
				CheckQuorum:            true,
				ElectionRTT:            5,
				HeartbeatRTT:           1,
				SnapshotEntries:        100,
				CompactionOverhead:     5,
				OrderedConfigChange:    false,
				MaxInMemLogSize:        0,
				DisableAutoCompactions: false,
				IsObserver:             false,
				IsWitness:              false,
				Quiesce:                false,
			},
		}

		err = GetApp().GetFlamed().StartOnDiskCluster(
			clusterConfiguration,
			storagedConfiguration,
			raftConfiguration)

		if err != nil {
			panic(err)
		}

		http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
			_, _ = fmt.Fprintf(writer, "Hello, %s!", request.URL.Path[1:])
		})

		err = http.ListenAndServe(":8080", nil)
		if err != nil {
			panic(err)
		}

	},
}

func init() {
	runCMD.Flags().StringVar(&dataStoragePath,
		"data-storage-path",
		"",
		"Data storage path")

	_ = runCMD.MarkFlagRequired("data-storage-path")

	runCMD.Flags().StringVar(&raftAddress,
		"raft-address",
		"",
		"Raft address")

	_ = runCMD.MarkFlagRequired("raft-address")

	runCMD.Flags().StringVar(&httpAddress,
		"http-address",
		"",
		"HTTP address")

	_ = runCMD.MarkFlagRequired("http-address")

	runCMD.Flags().BoolVar(&join,
		"join",
		false,
		"If true node will join the cluster")

	runCMD.Flags().StringVar(&initialMembers,
		"initial-members",
		"",
		"Initial raft nodes separated by comma")
}
