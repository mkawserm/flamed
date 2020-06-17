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
	viper.SetDefault(constant.HTTPServerAddress, "localhost:8081")
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

	// CORS defaults
	viper.SetDefault(constant.CORSAllowAllOrigins, false)
	viper.SetDefault(constant.CORSAllowOrigins, []string{})
	viper.SetDefault(constant.CORSAllowCredentials, false)
	viper.SetDefault(constant.CORSAllowMethods, []string{"GET"})
	viper.SetDefault(constant.CORSAllowHeaders, []string{"Accept",
		"Content-Type",
		"Content-Length",
		"Accept-Encoding",
		"Authorization"})
	viper.SetDefault(constant.CORSExposeHeaders, []string{})
	viper.SetDefault(constant.CORSMaxAge, time.Minute*5)

	viper.SetDefault(constant.EnableGraphQLOverHTTP, true)
	viper.SetDefault(constant.EnableGraphQLOverGRPC, false)

	viper.SetDefault(constant.EnableGRPCServer, false)
	viper.SetDefault(constant.GRPCServerAddress, "localhost:9091")
	viper.SetDefault(constant.GRPCServerTLS, false)
	viper.SetDefault(constant.GRPCServerCertFile, "")
	viper.SetDefault(constant.GRPCServerKeyFile, "")
}

func initAllPersistentFlags(cmd *cobra.Command) {
	/*Log related settings*/
	cmd.PersistentFlags().
		String(constant.LogLevelFlag, "error", "Log level")
	_ = viper.BindPFlag(constant.LogLevel, cmd.PersistentFlags().Lookup(constant.LogLevelFlag))

	cmd.PersistentFlags().
		Duration(constant.GlobalRequestTimeoutFlag, 30*time.Second, "Global request timeout")
	_ = viper.BindPFlag(constant.GlobalRequestTimeout, cmd.PersistentFlags().Lookup(constant.GlobalRequestTimeoutFlag))

	cmd.PersistentFlags().
		String(constant.StoragePathFlag, "~/"+variable.Name, "Data storage path")
	_ = viper.BindPFlag(constant.StoragePath, cmd.PersistentFlags().Lookup(constant.StoragePathFlag))

	cmd.PersistentFlags().
		String(constant.RaftAddressFlag, "localhost:2281", "Raft rpc address")
	_ = viper.BindPFlag(constant.RaftAddress, cmd.PersistentFlags().Lookup(constant.RaftAddressFlag))

	cmd.PersistentFlags().
		Bool(constant.JoinFlag, false, "Node join flag")
	_ = viper.BindPFlag(constant.Join, cmd.PersistentFlags().Lookup(constant.JoinFlag))

	cmd.PersistentFlags().
		String(constant.InitialMembersFlag, "", "Initial raft members "+
			"format: [node_id,raft_address;] (ex: 1,localhost:6001; 2,localhost:6002;)")
	_ = viper.BindPFlag(constant.InitialMembers, cmd.PersistentFlags().Lookup(constant.InitialMembersFlag))

	cmd.PersistentFlags().
		Bool(constant.EnableHTTPServerFlag, true, "Enable HTTP server flag")
	_ = viper.BindPFlag(constant.EnableHTTPServer, cmd.PersistentFlags().Lookup(constant.EnableHTTPServerFlag))

	cmd.PersistentFlags().
		String(constant.HTTPServerAddressFlag, "localhost:8081", "HTTP server address")
	_ = viper.BindPFlag(constant.HTTPServerAddress, cmd.PersistentFlags().Lookup(constant.HTTPServerAddressFlag))

	cmd.PersistentFlags().
		Bool(constant.HTTPServerTLSFlag, false, "HTTP server TLS flag")
	_ = viper.BindPFlag(constant.HTTPServerTLS, cmd.PersistentFlags().Lookup(constant.HTTPServerTLSFlag))

	cmd.PersistentFlags().
		String(constant.HTTPServerCertFileFlag, "", "HTTP server cert file")
	_ = viper.BindPFlag(constant.HTTPServerCertFile, cmd.PersistentFlags().Lookup(constant.HTTPServerCertFileFlag))

	cmd.PersistentFlags().
		String(constant.HTTPServerKeyFileFlag, "", "HTTP server key file")
	_ = viper.BindPFlag(constant.HTTPServerKeyFile, cmd.PersistentFlags().Lookup(constant.HTTPServerKeyFileFlag))

	cmd.PersistentFlags().
		Uint64(constant.DeploymentIDFlag, 1, "HTTP server cert file")
	_ = viper.BindPFlag(constant.DeploymentID, cmd.PersistentFlags().Lookup(constant.DeploymentIDFlag))

	cmd.PersistentFlags().
		Uint64(constant.RTTMillisecondFlag, 200, "Round trip time in milli second")
	_ = viper.BindPFlag(constant.RTTMillisecond, cmd.PersistentFlags().Lookup(constant.RTTMillisecondFlag))

	cmd.PersistentFlags().
		Bool(constant.MutualTLSFlag, false, "Raft mutual TLS flag")
	_ = viper.BindPFlag(constant.MutualTLS, cmd.PersistentFlags().Lookup(constant.MutualTLSFlag))

	cmd.PersistentFlags().
		String(constant.CAFileFlag, "", "Raft TLS ca file")
	_ = viper.BindPFlag(constant.CAFile, cmd.PersistentFlags().Lookup(constant.CAFileFlag))

	cmd.PersistentFlags().
		String(constant.CertFileFlag, "", "Raft TLS cert file")
	_ = viper.BindPFlag(constant.CertFile, cmd.PersistentFlags().Lookup(constant.CertFileFlag))

	cmd.PersistentFlags().
		String(constant.KeyFileFlag, "", "Raft TLS key file")
	_ = viper.BindPFlag(constant.KeyFile, cmd.PersistentFlags().Lookup(constant.KeyFileFlag))

	cmd.PersistentFlags().
		Uint64(constant.MaxSendQueueSizeFlag, 0, "Raft max send queue size")
	_ = viper.BindPFlag(constant.MaxSendQueueSize, cmd.PersistentFlags().Lookup(constant.MaxSendQueueSizeFlag))

	cmd.PersistentFlags().
		Uint64(constant.MaxReceiveQueueSizeFlag, 0, "Raft max receive queue size")
	_ = viper.BindPFlag(constant.MaxReceiveQueueSize, cmd.PersistentFlags().Lookup(constant.MaxReceiveQueueSizeFlag))

	cmd.PersistentFlags().
		Bool(constant.EnableMetricsFlag, false, "Enable metrics")
	_ = viper.BindPFlag(constant.EnableMetrics, cmd.PersistentFlags().Lookup(constant.EnableMetricsFlag))

	cmd.PersistentFlags().
		Uint64(
			constant.MaxSnapshotSendBytesPerSecondFlag,
			0,
			"Max snapshot send bytes per second")
	_ = viper.BindPFlag(constant.MaxSnapshotSendBytesPerSecond,
		cmd.PersistentFlags().Lookup(constant.MaxSnapshotSendBytesPerSecondFlag))

	cmd.PersistentFlags().
		Uint64(
			constant.MaxSnapshotRecvBytesPerSecondFlag,
			0,
			"Max snapshot recv bytes per second")
	_ = viper.BindPFlag(constant.MaxSnapshotRecvBytesPerSecond,
		cmd.PersistentFlags().Lookup(constant.MaxSnapshotRecvBytesPerSecondFlag))

	cmd.PersistentFlags().
		Bool(constant.NotifyCommitFlag,
			false,
			"Notify commit")
	_ = viper.BindPFlag(constant.NotifyCommit, cmd.PersistentFlags().Lookup(constant.NotifyCommitFlag))

	cmd.PersistentFlags().
		String(constant.LogDBConfigFlag,
			"tiny",
			"Log db config")
	_ = viper.BindPFlag(constant.LogDBConfig, cmd.PersistentFlags().Lookup(constant.LogDBConfigFlag))

	cmd.PersistentFlags().
		Duration(constant.SystemTickerPrecisionFlag,
			0,
			"System ticker precision")
	_ = viper.BindPFlag(constant.SystemTickerPrecision, cmd.PersistentFlags().
		Lookup(constant.SystemTickerPrecisionFlag))

	cmd.PersistentFlags().
		Uint64(constant.NodeIDFlag,
			1,
			"Node id")
	_ = viper.BindPFlag(constant.NodeID, cmd.PersistentFlags().Lookup(constant.NodeIDFlag))

	cmd.PersistentFlags().
		Bool(constant.CheckQuorumFlag,
			true,
			"Check quorum")
	_ = viper.BindPFlag(constant.CheckQuorum, cmd.PersistentFlags().Lookup(constant.CheckQuorumFlag))

	cmd.PersistentFlags().
		Uint64(constant.ElectionRTTFlag,
			5,
			"Election RTT")
	_ = viper.BindPFlag(constant.ElectionRTT, cmd.PersistentFlags().Lookup(constant.ElectionRTTFlag))

	cmd.PersistentFlags().
		Uint64(constant.HeartbeatRTTFlag,
			1,
			"Heartbeat RTT")
	_ = viper.BindPFlag(constant.HeartbeatRTT, cmd.PersistentFlags().Lookup(constant.HeartbeatRTTFlag))

	cmd.PersistentFlags().
		Uint64(constant.SnapshotEntriesFlag,
			1000,
			"Snapshot entries")
	_ = viper.BindPFlag(constant.SnapshotEntries, cmd.PersistentFlags().Lookup(constant.SnapshotEntriesFlag))

	cmd.PersistentFlags().
		Uint64(constant.CompactionOverheadFlag,
			5,
			"Compaction overhead")
	_ = viper.BindPFlag(constant.CompactionOverhead, cmd.PersistentFlags().Lookup(constant.CompactionOverheadFlag))

	cmd.PersistentFlags().
		Bool(constant.OrderedConfigChangeFlag,
			false,
			"Ordered config change")
	_ = viper.BindPFlag(constant.OrderedConfigChange, cmd.PersistentFlags().Lookup(constant.OrderedConfigChangeFlag))

	cmd.PersistentFlags().
		Uint64(constant.MaxInMemLogSizeFlag,
			0,
			"Max in mem log size")
	_ = viper.BindPFlag(constant.MaxInMemLogSize, cmd.PersistentFlags().Lookup(constant.MaxInMemLogSizeFlag))

	cmd.PersistentFlags().
		Bool(constant.DisableAutoCompactionsFlag,
			false,
			"Disable auto compactions")
	_ = viper.BindPFlag(constant.DisableAutoCompactions,
		cmd.PersistentFlags().Lookup(constant.DisableAutoCompactionsFlag))

	cmd.PersistentFlags().
		Bool(constant.IsObserverFlag,
			false,
			"Is observer")
	_ = viper.BindPFlag(constant.IsObserver, cmd.PersistentFlags().Lookup(constant.IsObserverFlag))

	cmd.PersistentFlags().
		Bool(constant.IsWitnessFlag,
			false,
			"Is witness")
	_ = viper.BindPFlag(constant.IsWitness, cmd.PersistentFlags().Lookup(constant.IsWitnessFlag))

	cmd.PersistentFlags().
		Bool(constant.QuiesceFlag,
			false,
			"Quiesce")
	_ = viper.BindPFlag(constant.Quiesce, cmd.PersistentFlags().Lookup(constant.QuiesceFlag))

	// CORS related flags
	cmd.PersistentFlags().
		Bool(constant.CORSAllowAllOriginsFlag,
			false,
			"CORS allow all origins")
	_ = viper.BindPFlag(constant.CORSAllowAllOrigins,
		cmd.PersistentFlags().Lookup(constant.CORSAllowAllOriginsFlag))

	cmd.PersistentFlags().
		StringSlice(constant.CORSAllowOriginsFlag,
			[]string{},
			"CORS allow origins")
	_ = viper.BindPFlag(constant.CORSAllowOrigins,
		cmd.PersistentFlags().Lookup(constant.CORSAllowOriginsFlag))

	cmd.PersistentFlags().
		Bool(constant.CORSAllowCredentialsFlag,
			false,
			"CORS allow credentials")
	_ = viper.BindPFlag(constant.CORSAllowCredentials,
		cmd.PersistentFlags().Lookup(constant.CORSAllowCredentialsFlag))

	cmd.PersistentFlags().
		StringSlice(constant.CORSAllowMethodsFlag,
			[]string{"GET"},
			"CORS allow methods")
	_ = viper.BindPFlag(constant.CORSAllowMethods,
		cmd.PersistentFlags().Lookup(constant.CORSAllowMethodsFlag))

	cmd.PersistentFlags().
		StringSlice(constant.CORSAllowHeadersFlag,
			[]string{"Accept", "Content-Type", "Content-Length", "Accept-Encoding", "Authorization"},
			"CORS allow headers")
	_ = viper.BindPFlag(constant.CORSAllowHeaders,
		cmd.PersistentFlags().Lookup(constant.CORSAllowHeadersFlag))

	cmd.PersistentFlags().
		StringSlice(constant.CORSExposeHeadersFlag,
			[]string{},
			"CORS expose headers")
	_ = viper.BindPFlag(constant.CORSExposeHeaders,
		cmd.PersistentFlags().Lookup(constant.CORSExposeHeadersFlag))

	cmd.PersistentFlags().
		Duration(constant.CORSMaxAgeFlag,
			5*time.Minute,
			"CORS max age")
	_ = viper.BindPFlag(constant.CORSMaxAge,
		cmd.PersistentFlags().Lookup(constant.CORSMaxAgeFlag))

	cmd.PersistentFlags().
		Bool(constant.EnableGraphQLOverHTTPFlag,
			true,
			"Enable GraphQL Over HTTP")
	_ = viper.BindPFlag(constant.EnableGraphQLOverHTTP,
		cmd.PersistentFlags().Lookup(constant.EnableGraphQLOverHTTPFlag))

	cmd.PersistentFlags().
		Bool(constant.EnableGraphQLOverGRPCFlag,
			false,
			"Enable GraphQL Over GRPC")
	_ = viper.BindPFlag(constant.EnableGraphQLOverGRPC,
		cmd.PersistentFlags().Lookup(constant.EnableGraphQLOverGRPCFlag))

	cmd.PersistentFlags().
		Bool(constant.EnableGRPCServerFlag,
			false,
			"Enable GRPC Server")
	_ = viper.BindPFlag(constant.EnableGRPCServer,
		cmd.PersistentFlags().Lookup(constant.EnableGRPCServerFlag))

	cmd.PersistentFlags().
		String(constant.GRPCServerAddressFlag, "localhost:9091", "GRPC server address")
	_ = viper.BindPFlag(constant.GRPCServerAddress, cmd.PersistentFlags().Lookup(constant.GRPCServerAddressFlag))

	cmd.PersistentFlags().
		Bool(constant.GRPCServerTLSFlag, false, "GRPC server TLS flag")
	_ = viper.BindPFlag(constant.GRPCServerTLS, cmd.PersistentFlags().Lookup(constant.GRPCServerTLSFlag))

	cmd.PersistentFlags().
		String(constant.GRPCServerCertFileFlag, "", "GRPC server cert file")
	_ = viper.BindPFlag(constant.GRPCServerCertFile, cmd.PersistentFlags().Lookup(constant.GRPCServerCertFileFlag))

	cmd.PersistentFlags().
		String(constant.GRPCServerKeyFileFlag, "", "GRPC server key file")
	_ = viper.BindPFlag(constant.GRPCServerKeyFile, cmd.PersistentFlags().Lookup(constant.GRPCServerKeyFileFlag))
}
