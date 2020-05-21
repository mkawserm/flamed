package app

import (
	"fmt"
	"github.com/mkawserm/flamed/pkg/iface"
	"github.com/mkawserm/flamed/pkg/tp/accesscontrol"
	"github.com/mkawserm/flamed/pkg/tp/indexmeta"
	"github.com/mkawserm/flamed/pkg/tp/intkey"
	"github.com/mkawserm/flamed/pkg/tp/json"
	"github.com/mkawserm/flamed/pkg/tp/user"
	"os"
	"sync"
	"time"

	"github.com/mitchellh/go-homedir"
	"github.com/mkawserm/flamed/pkg/flamed"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	appOnce        sync.Once
	appIns         *App
	configFile     string
	flamedHomePath string
)

type App struct {
	mStoragePath          string
	mGlobalRequestTimeout time.Duration

	mFlamed      *flamed.Flamed
	mRootCommand *cobra.Command

	mTransactionProcessor []iface.ITransactionProcessor
}

func (a *App) UpdateGlobalRequestTimeout(timeout time.Duration) {
	a.mGlobalRequestTimeout = timeout
}

func (a *App) AddTransactionProcessor(tp iface.ITransactionProcessor) {
	a.mTransactionProcessor = append(a.mTransactionProcessor, tp)
}

func (a *App) GetFlamed() *flamed.Flamed {
	return a.mFlamed
}

func (a *App) setup() {
	a.mGlobalRequestTimeout = 30 * time.Second
	a.mFlamed = flamed.NewFlamed()

	a.mTransactionProcessor = append(a.mTransactionProcessor, &user.User{})
	a.mTransactionProcessor = append(a.mTransactionProcessor, &json.JSON{})
	a.mTransactionProcessor = append(a.mTransactionProcessor, &intkey.IntKey{})
	a.mTransactionProcessor = append(a.mTransactionProcessor, &indexmeta.IndexMeta{})
	a.mTransactionProcessor = append(a.mTransactionProcessor, &accesscontrol.AccessControl{})

	a.mRootCommand = &cobra.Command{
		Use:   "flamed",
		Short: "Flamed is an open-source distributed embeddable NoSQL database",
		Long:  "Flamed is an open-source distributed embeddable NoSQL database",
		Run: func(cmd *cobra.Command, _ []string) {
			fmt.Println(cmd.UsageString())
		},
	}

	a.mRootCommand.
		PersistentFlags().
		StringVar(
			&configFile,
			"config",
			"",
			"config file (default is $HOME/flamed/.flamed.yaml)")

	a.mRootCommand.PersistentFlags().
		String("storage-path", flamedHomePath, "Data storage path")
	_ = viper.BindPFlag("StoragePath", a.mRootCommand.PersistentFlags().Lookup("storage-path"))

	a.mRootCommand.PersistentFlags().
		String("raft-address", "", "Raft rpc address")
	_ = viper.BindPFlag("RaftAddress", a.mRootCommand.PersistentFlags().Lookup("raft-address"))

	a.mRootCommand.PersistentFlags().
		String("http-address", "", "HTTP server address")
	_ = viper.BindPFlag("HTTPAddress", a.mRootCommand.PersistentFlags().Lookup("http-address"))

	a.mRootCommand.PersistentFlags().
		Bool("join", false, "Node join flag")
	_ = viper.BindPFlag("Join", a.mRootCommand.PersistentFlags().Lookup("join"))

	a.mRootCommand.PersistentFlags().
		String("initial-members", "", "Initial raft members")
	_ = viper.BindPFlag("InitialMembers", a.mRootCommand.PersistentFlags().Lookup("initial-members"))

	a.mRootCommand.PersistentFlags().
		String("http-server-tls", "", "HTTP server TLS flag")
	_ = viper.BindPFlag("HTTPServerTLS", a.mRootCommand.PersistentFlags().Lookup("http-server-tls"))

	a.mRootCommand.PersistentFlags().
		String("http-server-cert-file", "", "HTTP server cert file")
	_ = viper.BindPFlag("HTTPServerCertFile", a.mRootCommand.PersistentFlags().Lookup("http-server-cert-file"))

	a.mRootCommand.PersistentFlags().
		String("http-server-key-file", "", "HTTP server cert file")
	_ = viper.BindPFlag("HTTPServerKeyFile", a.mRootCommand.PersistentFlags().Lookup("http-server-key-file"))

	a.mRootCommand.PersistentFlags().
		Uint64("deployment-id", 1, "HTTP server cert file")
	_ = viper.BindPFlag("DeploymentID", a.mRootCommand.PersistentFlags().Lookup("deployment-id"))

	a.mRootCommand.PersistentFlags().
		Uint64("rtt-millisecond", 200, "Round trip time in milli second")
	_ = viper.BindPFlag("RTTMillisecond", a.mRootCommand.PersistentFlags().Lookup("rtt-millisecond"))

	a.mRootCommand.PersistentFlags().
		Bool("mutual-tls", false, "Raft mutual TLS flag")
	_ = viper.BindPFlag("MutualTLS", a.mRootCommand.PersistentFlags().Lookup("mutual-tls"))

	a.mRootCommand.PersistentFlags().
		String("ca-file", "", "Raft TLS ca file")
	_ = viper.BindPFlag("CAFile", a.mRootCommand.PersistentFlags().Lookup("ca-file"))

	a.mRootCommand.PersistentFlags().
		String("cert-file", "", "Raft TLS cert file")
	_ = viper.BindPFlag("CertFile", a.mRootCommand.PersistentFlags().Lookup("cert-file"))

	a.mRootCommand.PersistentFlags().
		String("key-file", "", "Raft TLS key file")
	_ = viper.BindPFlag("KeyFile", a.mRootCommand.PersistentFlags().Lookup("key-file"))

	a.mRootCommand.PersistentFlags().
		Uint64("max-send-queue-size", 0, "Raft max send queue size")
	_ = viper.BindPFlag("MaxSendQueueSize", a.mRootCommand.PersistentFlags().Lookup("max-send-queue-size"))

	a.mRootCommand.PersistentFlags().
		Uint64("max-receive-queue-size", 0, "Raft max receive queue size")
	_ = viper.BindPFlag("MaxReceiveQueueSize", a.mRootCommand.PersistentFlags().Lookup("max-receive-queue-size"))

	a.mRootCommand.PersistentFlags().
		Bool("enable-metrics", false, "Enable metrics")
	_ = viper.BindPFlag("EnableMetrics", a.mRootCommand.PersistentFlags().Lookup("enable-metrics"))

	a.mRootCommand.PersistentFlags().
		Uint64(
			"max-snapshot-send-bytes-per-second",
			0,
			"Max snapshot send bytes per second")
	_ = viper.BindPFlag("MaxSnapshotSendBytesPerSecond",
		a.mRootCommand.PersistentFlags().Lookup("max-snapshot-send-bytes-per-second"))

	a.mRootCommand.PersistentFlags().
		Uint64(
			"max-snapshot-recv-bytes-per-second",
			0,
			"Max snapshot recv bytes per second")
	_ = viper.BindPFlag("MaxSnapshotRecvBytesPerSecond",
		a.mRootCommand.PersistentFlags().Lookup("max-snapshot-recv-bytes-per-second"))

	a.mRootCommand.PersistentFlags().
		Bool("notify-commit",
			false,
			"Notify commit")
	_ = viper.BindPFlag("NotifyCommit", a.mRootCommand.PersistentFlags().Lookup("notify-commit"))

	a.mRootCommand.PersistentFlags().
		String("log-db-config",
			"tiny",
			"Log db config")
	_ = viper.BindPFlag("LogDBConfig", a.mRootCommand.PersistentFlags().Lookup("log-db-config"))

	a.mRootCommand.PersistentFlags().
		Uint64("node-id",
			1,
			"Node id")
	_ = viper.BindPFlag("NodeID", a.mRootCommand.PersistentFlags().Lookup("node-id"))

	a.mRootCommand.PersistentFlags().
		Bool("check-quorum",
			true,
			"Check quorum")
	_ = viper.BindPFlag("CheckQuorum", a.mRootCommand.PersistentFlags().Lookup("check-quorum"))

	a.mRootCommand.PersistentFlags().
		Uint64("election-rtt",
			5,
			"Election RTT")
	_ = viper.BindPFlag("ElectionRTT", a.mRootCommand.PersistentFlags().Lookup("election-rtt"))

	a.mRootCommand.PersistentFlags().
		Uint64("heartbeat-rtt",
			1,
			"Heartbeat RTT")
	_ = viper.BindPFlag("HeartbeatRTT", a.mRootCommand.PersistentFlags().Lookup("heartbeat-rtt"))

	a.mRootCommand.PersistentFlags().
		Uint64("snapshot-entries",
			100,
			"Snapshot entries")
	_ = viper.BindPFlag("SnapshotEntries", a.mRootCommand.PersistentFlags().Lookup("snapshot-entries"))

	a.mRootCommand.PersistentFlags().
		Uint64("compaction-overhead",
			5,
			"Compaction overhead")
	_ = viper.BindPFlag("CompactionOverhead", a.mRootCommand.PersistentFlags().Lookup("compaction-overhead"))

	a.mRootCommand.PersistentFlags().
		Bool("ordered-config-change",
			false,
			"Ordered config change")
	_ = viper.BindPFlag("OrderedConfigChange", a.mRootCommand.PersistentFlags().Lookup("ordered-config-change"))

	a.mRootCommand.PersistentFlags().
		Uint64("max-in-mem-log-size",
			0,
			"Max in mem log size")
	_ = viper.BindPFlag("MaxInMemLogSize", a.mRootCommand.PersistentFlags().Lookup("max-in-mem-log-size"))

	a.mRootCommand.PersistentFlags().
		Bool("disable-auto-compactions",
			false,
			"Disable auto compactions")
	_ = viper.BindPFlag("DisableAutoCompactions",
		a.mRootCommand.PersistentFlags().Lookup("disable-auto-compactions"))

	a.mRootCommand.PersistentFlags().
		Bool("is-observer",
			false,
			"Is observer")
	_ = viper.BindPFlag("IsObserver", a.mRootCommand.PersistentFlags().Lookup("is-observer"))

	a.mRootCommand.PersistentFlags().
		Bool("is-witness",
			false,
			"Is witness")
	_ = viper.BindPFlag("IsWitness", a.mRootCommand.PersistentFlags().Lookup("is-witness"))

	a.mRootCommand.PersistentFlags().
		Bool("quiesce",
			false,
			"Quiesce")
	_ = viper.BindPFlag("Quiesce", a.mRootCommand.PersistentFlags().Lookup("quiesce"))

	a.mRootCommand.AddCommand(runCMD)
	a.mRootCommand.AddCommand(authorCMD)
	a.mRootCommand.AddCommand(versionCMD)
}

func (a *App) Execute() error {
	return a.mRootCommand.Execute()
}

func GetApp() *App {
	return appIns
}

func initConfig() {
	// Find home directory.
	home, err := homedir.Dir()
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	flamedHomePath = home + "/flamed"

	if configFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(configFile)
	} else {
		// Search config in home directory with name ".cobra" (without extension).
		viper.AddConfigPath(flamedHomePath)
		viper.SetConfigName(".flamed")
	}

	viper.SetEnvPrefix("flamed")
	viper.AutomaticEnv()

	// SET DEFAULTS
	viper.SetDefault("StoragePath", flamedHomePath)
	viper.SetDefault("RaftAddress", "")
	viper.SetDefault("HTTPAddress", "")
	viper.SetDefault("Join", false)
	viper.SetDefault("InitialMembers", "")

	viper.SetDefault("HTTPServerTLS", false)
	viper.SetDefault("HTTPServerCertFile", "")
	viper.SetDefault("HTTPServerKeyFile", "")

	viper.SetDefault("DeploymentID", 1)
	viper.SetDefault("RTTMillisecond", 200)
	viper.SetDefault("MutualTLS", false)
	viper.SetDefault("CAFile", "")
	viper.SetDefault("CertFile", "")
	viper.SetDefault("KeyFile", "")
	viper.SetDefault("MaxSendQueueSize", 0)
	viper.SetDefault("MaxReceiveQueueSize", 0)
	viper.SetDefault("EnableMetrics", false)
	viper.SetDefault("MaxSnapshotSendBytesPerSecond", 0)
	viper.SetDefault("MaxSnapshotRecvBytesPerSecond", 0)
	viper.SetDefault("NotifyCommit", false)
	viper.SetDefault("LogDBConfig", "tiny")

	viper.SetDefault("NodeID", 1)
	viper.SetDefault("CheckQuorum", true)
	viper.SetDefault("ElectionRTT", 5)
	viper.SetDefault("HeartbeatRTT", 1)
	viper.SetDefault("SnapshotEntries", 100)
	viper.SetDefault("CompactionOverhead", 5)
	viper.SetDefault("OrderedConfigChange", false)
	viper.SetDefault("MaxInMemLogSize", 0)
	viper.SetDefault("DisableAutoCompactions", false)
	viper.SetDefault("IsObserver", false)
	viper.SetDefault("IsWitness", false)
	viper.SetDefault("Quiesce", false)

	if err := viper.ReadInConfig(); err == nil {
		//fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	appOnce.Do(func() {
		appIns = &App{}
		appIns.setup()
	})
}
