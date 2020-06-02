package app

import (
	"fmt"
	"github.com/mitchellh/go-homedir"
	"github.com/mkawserm/flamed/pkg/constant"
	"github.com/mkawserm/flamed/pkg/variable"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
	"time"
)

var homePath string
var configFile string

func initConfig() {
	// Find home directory.
	home, err := homedir.Dir()
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	homePath = home + "/" + variable.Name

	if configFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(configFile)
	} else {
		// Search config in home directory with name ".Name" (default .flamed) (without extension).
		viper.AddConfigPath(homePath)
		viper.SetConfigName("." + variable.Name)
	}

	viper.SetEnvPrefix(variable.Name)
	viper.AutomaticEnv()

	// SET DEFAULTS
	initAllDefaults()

	if err := viper.ReadInConfig(); err == nil {
		//fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}

func initAllDefaults() {
	/*Log related settings*/
	viper.SetDefault(constant.LogLevel, "error")
	viper.SetDefault(constant.GlobalRequestTimeout, 30*time.Second)
	viper.SetDefault(constant.StoragePath, homePath)
	viper.SetDefault(constant.RaftAddress, "localhost:2281")
	viper.SetDefault(constant.Join, false)
	viper.SetDefault(constant.InitialMembers, "")

	viper.SetDefault(constant.EnableHTTPServer, true)
	viper.SetDefault(constant.HTTPAddress, "localhost:8081")
	viper.SetDefault(constant.HTTPServerTLS, false)
	viper.SetDefault(constant.HTTPServerCertFile, "")
	viper.SetDefault(constant.HTTPServerKeyFile, "")

	viper.SetDefault(constant.DeploymentID, 1)
	viper.SetDefault(constant.RTTMillisecond, 200)
	viper.SetDefault(constant.MutualTLS, false)
	viper.SetDefault(constant.CAFile, "")
	viper.SetDefault(constant.CertFile, "")
	viper.SetDefault(constant.KeyFile, "")
	viper.SetDefault(constant.MaxSendQueueSize, 0)
	viper.SetDefault(constant.MaxReceiveQueueSize, 0)
	viper.SetDefault(constant.EnableMetrics, false)
	viper.SetDefault(constant.MaxSnapshotSendBytesPerSecond, 0)
	viper.SetDefault(constant.MaxSnapshotRecvBytesPerSecond, 0)
	viper.SetDefault(constant.NotifyCommit, false)
	viper.SetDefault(constant.LogDBConfig, "tiny")
	viper.SetDefault(constant.SystemTickerPrecision, time.Duration(0))

	viper.SetDefault(constant.NodeID, 1)
	viper.SetDefault(constant.CheckQuorum, true)
	viper.SetDefault(constant.ElectionRTT, 5)
	viper.SetDefault(constant.HeartbeatRTT, 1)
	viper.SetDefault(constant.SnapshotEntries, 100)
	viper.SetDefault(constant.CompactionOverhead, 5)
	viper.SetDefault(constant.OrderedConfigChange, false)
	viper.SetDefault(constant.MaxInMemLogSize, 0)
	viper.SetDefault(constant.DisableAutoCompactions, false)
	viper.SetDefault(constant.IsObserver, false)
	viper.SetDefault(constant.IsWitness, false)
	viper.SetDefault(constant.Quiesce, false)
}

func initAllPersistentFlags(cmd *cobra.Command) {
	/*Log related settings*/
	cmd.PersistentFlags().
		String("log-level", "info", "Log level")
	_ = viper.BindPFlag(constant.LogLevel, cmd.PersistentFlags().Lookup("log-level"))

	cmd.PersistentFlags().
		Duration("global-request-timeout", 30*time.Second, "Global request timeout")
	_ = viper.BindPFlag(constant.GlobalRequestTimeout, cmd.PersistentFlags().Lookup("global-request-timeout"))

	cmd.PersistentFlags().
		String("storage-path", homePath, "Data storage path")
	_ = viper.BindPFlag(constant.StoragePath, cmd.PersistentFlags().Lookup("storage-path"))

	cmd.PersistentFlags().
		String("raft-address", "localhost:2281", "Raft rpc address")
	_ = viper.BindPFlag(constant.RaftAddress, cmd.PersistentFlags().Lookup("raft-address"))

	cmd.PersistentFlags().
		Bool("join", false, "Node join flag")
	_ = viper.BindPFlag(constant.Join, cmd.PersistentFlags().Lookup("join"))

	cmd.PersistentFlags().
		String("initial-members", "", "Initial raft members "+
			"format: [node_id,raft_address;] (ex: 1,localhost:6001; 2,localhost:6002;)")
	_ = viper.BindPFlag(constant.InitialMembers, cmd.PersistentFlags().Lookup("initial-members"))

	cmd.PersistentFlags().
		Bool("enable-http-server", true, "Enable HTTP server flag")
	_ = viper.BindPFlag(constant.EnableHTTPServer, cmd.PersistentFlags().Lookup("enable-http-server"))

	cmd.PersistentFlags().
		String("http-address", "localhost:8081", "HTTP server address")
	_ = viper.BindPFlag(constant.HTTPAddress, cmd.PersistentFlags().Lookup("http-address"))

	cmd.PersistentFlags().
		Bool("http-server-tls", false, "HTTP server TLS flag")
	_ = viper.BindPFlag(constant.HTTPServerTLS, cmd.PersistentFlags().Lookup("http-server-tls"))

	cmd.PersistentFlags().
		String("http-server-cert-file", "", "HTTP server cert file")
	_ = viper.BindPFlag(constant.HTTPServerCertFile, cmd.PersistentFlags().Lookup("http-server-cert-file"))

	cmd.PersistentFlags().
		String("http-server-key-file", "", "HTTP server cert file")
	_ = viper.BindPFlag(constant.HTTPServerKeyFile, cmd.PersistentFlags().Lookup("http-server-key-file"))

	cmd.PersistentFlags().
		Uint64("deployment-id", 1, "HTTP server cert file")
	_ = viper.BindPFlag(constant.DeploymentID, cmd.PersistentFlags().Lookup("deployment-id"))

	cmd.PersistentFlags().
		Uint64("rtt-millisecond", 200, "Round trip time in milli second")
	_ = viper.BindPFlag(constant.RTTMillisecond, cmd.PersistentFlags().Lookup("rtt-millisecond"))

	cmd.PersistentFlags().
		Bool("mutual-tls", false, "Raft mutual TLS flag")
	_ = viper.BindPFlag(constant.MutualTLS, cmd.PersistentFlags().Lookup("mutual-tls"))

	cmd.PersistentFlags().
		String("ca-file", "", "Raft TLS ca file")
	_ = viper.BindPFlag(constant.CAFile, cmd.PersistentFlags().Lookup("ca-file"))

	cmd.PersistentFlags().
		String("cert-file", "", "Raft TLS cert file")
	_ = viper.BindPFlag(constant.CertFile, cmd.PersistentFlags().Lookup("cert-file"))

	cmd.PersistentFlags().
		String("key-file", "", "Raft TLS key file")
	_ = viper.BindPFlag(constant.KeyFile, cmd.PersistentFlags().Lookup("key-file"))

	cmd.PersistentFlags().
		Uint64("max-send-queue-size", 0, "Raft max send queue size")
	_ = viper.BindPFlag(constant.MaxSendQueueSize, cmd.PersistentFlags().Lookup("max-send-queue-size"))

	cmd.PersistentFlags().
		Uint64("max-receive-queue-size", 0, "Raft max receive queue size")
	_ = viper.BindPFlag(constant.MaxReceiveQueueSize, cmd.PersistentFlags().Lookup("max-receive-queue-size"))

	cmd.PersistentFlags().
		Bool("enable-metrics", false, "Enable metrics")
	_ = viper.BindPFlag(constant.EnableMetrics, cmd.PersistentFlags().Lookup("enable-metrics"))

	cmd.PersistentFlags().
		Uint64(
			"max-snapshot-send-bytes-per-second",
			0,
			"Max snapshot send bytes per second")
	_ = viper.BindPFlag(constant.MaxSnapshotSendBytesPerSecond,
		cmd.PersistentFlags().Lookup("max-snapshot-send-bytes-per-second"))

	cmd.PersistentFlags().
		Uint64(
			"max-snapshot-recv-bytes-per-second",
			0,
			"Max snapshot recv bytes per second")
	_ = viper.BindPFlag(constant.MaxSnapshotRecvBytesPerSecond,
		cmd.PersistentFlags().Lookup("max-snapshot-recv-bytes-per-second"))

	cmd.PersistentFlags().
		Bool("notify-commit",
			false,
			"Notify commit")
	_ = viper.BindPFlag(constant.NotifyCommit, cmd.PersistentFlags().Lookup("notify-commit"))

	cmd.PersistentFlags().
		String("log-db-config",
			"tiny",
			"Log db config")
	_ = viper.BindPFlag(constant.LogDBConfig, cmd.PersistentFlags().Lookup("log-db-config"))

	cmd.PersistentFlags().
		Duration("system-ticker-precision",
			0,
			"System ticker precision")
	_ = viper.BindPFlag(constant.SystemTickerPrecision, cmd.PersistentFlags().Lookup("system-ticker-precision"))

	cmd.PersistentFlags().
		Uint64("node-id",
			1,
			"Node id")
	_ = viper.BindPFlag(constant.NodeID, cmd.PersistentFlags().Lookup("node-id"))

	cmd.PersistentFlags().
		Bool("check-quorum",
			true,
			"Check quorum")
	_ = viper.BindPFlag(constant.CheckQuorum, cmd.PersistentFlags().Lookup("check-quorum"))

	cmd.PersistentFlags().
		Uint64("election-rtt",
			5,
			"Election RTT")
	_ = viper.BindPFlag(constant.ElectionRTT, cmd.PersistentFlags().Lookup("election-rtt"))

	cmd.PersistentFlags().
		Uint64("heartbeat-rtt",
			1,
			"Heartbeat RTT")
	_ = viper.BindPFlag(constant.HeartbeatRTT, cmd.PersistentFlags().Lookup("heartbeat-rtt"))

	cmd.PersistentFlags().
		Uint64("snapshot-entries",
			100,
			"Snapshot entries")
	_ = viper.BindPFlag(constant.SnapshotEntries, cmd.PersistentFlags().Lookup("snapshot-entries"))

	cmd.PersistentFlags().
		Uint64("compaction-overhead",
			5,
			"Compaction overhead")
	_ = viper.BindPFlag(constant.CompactionOverhead, cmd.PersistentFlags().Lookup("compaction-overhead"))

	cmd.PersistentFlags().
		Bool("ordered-config-change",
			false,
			"Ordered config change")
	_ = viper.BindPFlag(constant.OrderedConfigChange, cmd.PersistentFlags().Lookup("ordered-config-change"))

	cmd.PersistentFlags().
		Uint64("max-in-mem-log-size",
			0,
			"Max in mem log size")
	_ = viper.BindPFlag(constant.MaxInMemLogSize, cmd.PersistentFlags().Lookup("max-in-mem-log-size"))

	cmd.PersistentFlags().
		Bool("disable-auto-compactions",
			false,
			"Disable auto compactions")
	_ = viper.BindPFlag(constant.DisableAutoCompactions,
		cmd.PersistentFlags().Lookup("disable-auto-compactions"))

	cmd.PersistentFlags().
		Bool("is-observer",
			false,
			"Is observer")
	_ = viper.BindPFlag(constant.IsObserver, cmd.PersistentFlags().Lookup("is-observer"))

	cmd.PersistentFlags().
		Bool("is-witness",
			false,
			"Is witness")
	_ = viper.BindPFlag(constant.IsWitness, cmd.PersistentFlags().Lookup("is-witness"))

	cmd.PersistentFlags().
		Bool("quiesce",
			false,
			"Quiesce")
	_ = viper.BindPFlag(constant.Quiesce, cmd.PersistentFlags().Lookup("quiesce"))
}
