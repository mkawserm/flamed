package app

import (
	"fmt"
	"github.com/lni/dragonboat/v3/config"
	"github.com/mkawserm/flamed/pkg/conf"
	"github.com/mkawserm/flamed/pkg/logger"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"net/http"
	"strconv"
	"strings"
)

var RunCMD = &cobra.Command{
	Use:   "run",
	Short: "Run Flamed server",
	Run: func(cmd *cobra.Command, args []string) {
		logger.GetLoggerFactory().ChangeLogLevel(viper.GetString("LogLevel"))
		if len(viper.GetString("StoragePath")) == 0 {
			fmt.Println("StoragePath can not be empty")
			return
		}

		if len(viper.GetString("RaftAddress")) == 0 {
			fmt.Println("RaftAddress can not be empty")
			return
		}

		if len(viper.GetString("HTTPAddress")) == 0 {
			fmt.Println("HTTPAddress can not be empty")
			return
		}

		im := getInitialMembers(strings.Split(viper.GetString("InitialMembers"), ";"))

		if len(im) == 0 {
			if !viper.GetBool("Join") {
				im[viper.GetUint64("NodeID")] = viper.GetString("RaftAddress")
			}
		}

		raftStoragePath := viper.GetString("StoragePath") + "/raft"

		nodeConfiguration := &conf.NodeConfiguration{
			NodeConfigurationInput: conf.NodeConfigurationInput{
				NodeHostDir:                   raftStoragePath,
				WALDir:                        raftStoragePath,
				DeploymentID:                  viper.GetUint64("DeploymentID"),
				RTTMillisecond:                viper.GetUint64("RTTMillisecond"),
				RaftAddress:                   viper.GetString("RaftAddress"),
				ListenAddress:                 "",
				MutualTLS:                     viper.GetBool("MutualTLS"),
				CAFile:                        viper.GetString("CAFile"),
				CertFile:                      viper.GetString("CertFile"),
				KeyFile:                       viper.GetString("KeyFile"),
				MaxSendQueueSize:              viper.GetUint64("MaxSendQueueSize"),
				MaxReceiveQueueSize:           viper.GetUint64("MaxReceiveQueueSize"),
				EnableMetrics:                 viper.GetBool("EnableMetrics"),
				MaxSnapshotSendBytesPerSecond: viper.GetUint64("MaxSnapshotSendBytesPerSecond"),
				MaxSnapshotRecvBytesPerSecond: viper.GetUint64("MaxSnapshotRecvBytesPerSecond"),
				NotifyCommit:                  viper.GetBool("NotifyCommit"),

				SystemTickerPrecision: viper.GetDuration("SystemTickerPrecision"),

				LogDBConfig:         config.GetTinyMemLogDBConfig(),
				LogDBFactory:        nil,
				RaftRPCFactory:      nil,
				RaftEventListener:   nil,
				SystemEventListener: nil,
			},
		}

		if viper.GetString("LogDBConfig") == "tiny" {
			nodeConfiguration.NodeConfigurationInput.LogDBConfig = config.GetTinyMemLogDBConfig()
		} else {
			nodeConfiguration.NodeConfigurationInput.LogDBConfig = config.GetTinyMemLogDBConfig()
		}

		err := GetApp().GetFlamed().ConfigureNode(nodeConfiguration)

		if err != nil {
			panic(err)
		}

		clusterStoragePath := viper.GetString("StoragePath") + "/cluster-1"

		storagedConfiguration := conf.SimpleStoragedConfiguration(clusterStoragePath, nil)
		for _, tp := range GetApp().mTransactionProcessor {
			storagedConfiguration.AddTransactionProcessor(tp)
		}

		clusterConfiguration := conf.SimpleOnDiskClusterConfiguration(
			1,
			"cluster-1",
			im,
			viper.GetBool("Join"))

		raftConfiguration := &conf.RaftConfiguration{
			RaftConfigurationInput: conf.RaftConfigurationInput{
				NodeID:                 viper.GetUint64("NodeID"),
				CheckQuorum:            viper.GetBool("CheckQuorum"),
				ElectionRTT:            viper.GetUint64("ElectionRTT"),
				HeartbeatRTT:           viper.GetUint64("HeartbeatRTT"),
				SnapshotEntries:        viper.GetUint64("SnapshotEntries"),
				CompactionOverhead:     viper.GetUint64("CompactionOverhead"),
				OrderedConfigChange:    viper.GetBool("OrderedConfigChange"),
				MaxInMemLogSize:        viper.GetUint64("MaxInMemLogSize"),
				DisableAutoCompactions: viper.GetBool("DisableAutoCompactions"),
				IsObserver:             viper.GetBool("IsObserver"),
				IsWitness:              viper.GetBool("IsWitness"),
				Quiesce:                viper.GetBool("Quiesce"),
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

		err = http.ListenAndServe(viper.GetString("HTTPAddress"), nil)
		if err != nil {
			panic(err)
		}

	},
}

func getInitialMembers(stringList []string) map[uint64]string {
	var im = make(map[uint64]string)
	for _, value := range stringList {
		v := strings.TrimSpace(value)
		if v != "" {
			idAndAddress := strings.Split(v, ",")
			if len(idAndAddress) != 2 {
				continue
			}

			idString := strings.TrimSpace(idAndAddress[0])
			address := strings.TrimSpace(idAndAddress[1])
			id, err := strconv.Atoi(idString)
			if err != nil {
				panic(err)
			}

			if address != "" {
				im[uint64(id)] = address
			}
		}
	}

	return im
}
