package app

import (
	"fmt"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
	"time"
)

var flamedHomePath string
var configFile string

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
		// Search config in home directory with name ".flamed" (without extension).
		viper.AddConfigPath(flamedHomePath)
		viper.SetConfigName(".flamed")
	}

	viper.SetEnvPrefix("flamed")
	viper.AutomaticEnv()

	// SET DEFAULTS
	InitAllDefaults()

	if err := viper.ReadInConfig(); err == nil {
		//fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}

func InitAllDefaults() {
	viper.SetDefault("GlobalRequestTimeout", 30*time.Second)

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
	viper.SetDefault("SystemTickerPrecision", time.Duration(0))

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
}

func InitAllPersistentFlags(cmd *cobra.Command) {
	cmd.PersistentFlags().
		Duration("global-request-timeout", 30*time.Second, "Data storage path")
	_ = viper.BindPFlag("GlobalRequestTimeout", cmd.PersistentFlags().Lookup("global-request-timeout"))

	cmd.PersistentFlags().
		String("storage-path", flamedHomePath, "Data storage path")
	_ = viper.BindPFlag("StoragePath", cmd.PersistentFlags().Lookup("storage-path"))

	cmd.PersistentFlags().
		String("raft-address", "", "Raft rpc address")
	_ = viper.BindPFlag("RaftAddress", cmd.PersistentFlags().Lookup("raft-address"))

	cmd.PersistentFlags().
		String("http-address", "", "HTTP server address")
	_ = viper.BindPFlag("HTTPAddress", cmd.PersistentFlags().Lookup("http-address"))

	cmd.PersistentFlags().
		Bool("join", false, "Node join flag")
	_ = viper.BindPFlag("Join", cmd.PersistentFlags().Lookup("join"))

	cmd.PersistentFlags().
		String("initial-members", "", "Initial raft members "+
			"format: [node_id,raft_address;] (ex: 1,localhost:6001; 2,localhost:6002;)")
	_ = viper.BindPFlag("InitialMembers", cmd.PersistentFlags().Lookup("initial-members"))

	cmd.PersistentFlags().
		String("http-server-tls", "", "HTTP server TLS flag")
	_ = viper.BindPFlag("HTTPServerTLS", cmd.PersistentFlags().Lookup("http-server-tls"))

	cmd.PersistentFlags().
		String("http-server-cert-file", "", "HTTP server cert file")
	_ = viper.BindPFlag("HTTPServerCertFile", cmd.PersistentFlags().Lookup("http-server-cert-file"))

	cmd.PersistentFlags().
		String("http-server-key-file", "", "HTTP server cert file")
	_ = viper.BindPFlag("HTTPServerKeyFile", cmd.PersistentFlags().Lookup("http-server-key-file"))

	cmd.PersistentFlags().
		Uint64("deployment-id", 1, "HTTP server cert file")
	_ = viper.BindPFlag("DeploymentID", cmd.PersistentFlags().Lookup("deployment-id"))

	cmd.PersistentFlags().
		Uint64("rtt-millisecond", 200, "Round trip time in milli second")
	_ = viper.BindPFlag("RTTMillisecond", cmd.PersistentFlags().Lookup("rtt-millisecond"))

	cmd.PersistentFlags().
		Bool("mutual-tls", false, "Raft mutual TLS flag")
	_ = viper.BindPFlag("MutualTLS", cmd.PersistentFlags().Lookup("mutual-tls"))

	cmd.PersistentFlags().
		String("ca-file", "", "Raft TLS ca file")
	_ = viper.BindPFlag("CAFile", cmd.PersistentFlags().Lookup("ca-file"))

	cmd.PersistentFlags().
		String("cert-file", "", "Raft TLS cert file")
	_ = viper.BindPFlag("CertFile", cmd.PersistentFlags().Lookup("cert-file"))

	cmd.PersistentFlags().
		String("key-file", "", "Raft TLS key file")
	_ = viper.BindPFlag("KeyFile", cmd.PersistentFlags().Lookup("key-file"))

	cmd.PersistentFlags().
		Uint64("max-send-queue-size", 0, "Raft max send queue size")
	_ = viper.BindPFlag("MaxSendQueueSize", cmd.PersistentFlags().Lookup("max-send-queue-size"))

	cmd.PersistentFlags().
		Uint64("max-receive-queue-size", 0, "Raft max receive queue size")
	_ = viper.BindPFlag("MaxReceiveQueueSize", cmd.PersistentFlags().Lookup("max-receive-queue-size"))

	cmd.PersistentFlags().
		Bool("enable-metrics", false, "Enable metrics")
	_ = viper.BindPFlag("EnableMetrics", cmd.PersistentFlags().Lookup("enable-metrics"))

	cmd.PersistentFlags().
		Uint64(
			"max-snapshot-send-bytes-per-second",
			0,
			"Max snapshot send bytes per second")
	_ = viper.BindPFlag("MaxSnapshotSendBytesPerSecond",
		cmd.PersistentFlags().Lookup("max-snapshot-send-bytes-per-second"))

	cmd.PersistentFlags().
		Uint64(
			"max-snapshot-recv-bytes-per-second",
			0,
			"Max snapshot recv bytes per second")
	_ = viper.BindPFlag("MaxSnapshotRecvBytesPerSecond",
		cmd.PersistentFlags().Lookup("max-snapshot-recv-bytes-per-second"))

	cmd.PersistentFlags().
		Bool("notify-commit",
			false,
			"Notify commit")
	_ = viper.BindPFlag("NotifyCommit", cmd.PersistentFlags().Lookup("notify-commit"))

	cmd.PersistentFlags().
		String("log-db-config",
			"tiny",
			"Log db config")
	_ = viper.BindPFlag("LogDBConfig", cmd.PersistentFlags().Lookup("log-db-config"))

	cmd.PersistentFlags().
		Duration("system-ticker-precision",
			0,
			"System ticker precision")
	_ = viper.BindPFlag("SystemTickerPrecision", cmd.PersistentFlags().Lookup("system-ticker-precision"))

	cmd.PersistentFlags().
		Uint64("node-id",
			1,
			"Node id")
	_ = viper.BindPFlag("NodeID", cmd.PersistentFlags().Lookup("node-id"))

	cmd.PersistentFlags().
		Bool("check-quorum",
			true,
			"Check quorum")
	_ = viper.BindPFlag("CheckQuorum", cmd.PersistentFlags().Lookup("check-quorum"))

	cmd.PersistentFlags().
		Uint64("election-rtt",
			5,
			"Election RTT")
	_ = viper.BindPFlag("ElectionRTT", cmd.PersistentFlags().Lookup("election-rtt"))

	cmd.PersistentFlags().
		Uint64("heartbeat-rtt",
			1,
			"Heartbeat RTT")
	_ = viper.BindPFlag("HeartbeatRTT", cmd.PersistentFlags().Lookup("heartbeat-rtt"))

	cmd.PersistentFlags().
		Uint64("snapshot-entries",
			100,
			"Snapshot entries")
	_ = viper.BindPFlag("SnapshotEntries", cmd.PersistentFlags().Lookup("snapshot-entries"))

	cmd.PersistentFlags().
		Uint64("compaction-overhead",
			5,
			"Compaction overhead")
	_ = viper.BindPFlag("CompactionOverhead", cmd.PersistentFlags().Lookup("compaction-overhead"))

	cmd.PersistentFlags().
		Bool("ordered-config-change",
			false,
			"Ordered config change")
	_ = viper.BindPFlag("OrderedConfigChange", cmd.PersistentFlags().Lookup("ordered-config-change"))

	cmd.PersistentFlags().
		Uint64("max-in-mem-log-size",
			0,
			"Max in mem log size")
	_ = viper.BindPFlag("MaxInMemLogSize", cmd.PersistentFlags().Lookup("max-in-mem-log-size"))

	cmd.PersistentFlags().
		Bool("disable-auto-compactions",
			false,
			"Disable auto compactions")
	_ = viper.BindPFlag("DisableAutoCompactions",
		cmd.PersistentFlags().Lookup("disable-auto-compactions"))

	cmd.PersistentFlags().
		Bool("is-observer",
			false,
			"Is observer")
	_ = viper.BindPFlag("IsObserver", cmd.PersistentFlags().Lookup("is-observer"))

	cmd.PersistentFlags().
		Bool("is-witness",
			false,
			"Is witness")
	_ = viper.BindPFlag("IsWitness", cmd.PersistentFlags().Lookup("is-witness"))

	cmd.PersistentFlags().
		Bool("quiesce",
			false,
			"Quiesce")
	_ = viper.BindPFlag("Quiesce", cmd.PersistentFlags().Lookup("quiesce"))
}
