package app

import (
	"context"
	"fmt"
	"github.com/mkawserm/flamed/pkg/constant"
	"github.com/mkawserm/flamed/pkg/iface"
	"github.com/mkawserm/flamed/pkg/variable"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"
	"time"

	"github.com/lni/dragonboat/v3/config"
	"github.com/mkawserm/flamed/pkg/conf"
	"github.com/mkawserm/flamed/pkg/crypto"
	"github.com/mkawserm/flamed/pkg/logger"
	"github.com/mkawserm/flamed/pkg/pb"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var RunCMD = &cobra.Command{
	Use:   "run",
	Short: "Run server",
	Run: func(cmd *cobra.Command, args []string) {
		if len(viper.GetString(constant.StoragePath)) == 0 {
			fmt.Println("StoragePath can not be empty")
			return
		}

		if len(viper.GetString(constant.RaftAddress)) == 0 {
			fmt.Println("RaftAddress can not be empty")
			return
		}

		if len(viper.GetString(constant.HTTPServerAddress)) == 0 {
			fmt.Println("HTTPServerAddress can not be empty")
			return
		}

		raftStoragePath := viper.GetString(constant.StoragePath) + "/raft"
		// node configuration
		nodeConfiguration := getNodeConfiguration(raftStoragePath)
		if viper.GetString("LogDBConfig") == "tiny" {
			nodeConfiguration.NodeConfigurationInput.LogDBConfig = config.GetTinyMemLogDBConfig()
		} else {
			nodeConfiguration.NodeConfigurationInput.LogDBConfig = config.GetTinyMemLogDBConfig()
		}
		err := GetApp().GetFlamed().ConfigureNode(nodeConfiguration)
		if err != nil {
			panic(err)
		}

		// start cluster
		startCluster()

		// initialize cluster defaults
		// like admin user and other things
		initializeClusterDefaults()

		// run server and wait for shutdown
		runServerAndWaitForShutdown()
	},
} // Command

func runServerAndWaitForShutdown() {
	idleChan := make(chan struct{})
	go func() {
		signChan := make(chan os.Signal, 1)
		signal.Notify(signChan, os.Interrupt, syscall.SIGTERM)
		sig := <-signChan

		logger.L("app").Info("shutdown signal received",
			zap.String("signal", sig.String()))
		logger.L("app").Info("preparing for shutdown")

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		if err := GetApp().getServer().Shutdown(ctx); err == context.DeadlineExceeded {
			logger.L("app").Info("shutdown: halted active connections")
		}

		// Stop node
		GetApp().GetFlamed().StopNode()

		// Actual shutdown trigger.
		close(idleChan)
	}()

	// HTTP Server
	if viper.GetBool(constant.EnableHTTPServer) {
		go func() {
			if err := runHTTPServer(); err == http.ErrServerClosed {
				logger.L("app").Info("http server closed")
			}
		}()
	}

	// Blocking until the shutdown is complete
	<-idleChan
	logger.L("app").Info("shutdown complete")
}

func runHTTPServer() error {
	logger.L("app").Info("preparing http server")
	if viper.GetBool(constant.HTTPServerTLS) {
		logger.L("app").Info("http server with tls started @ " +
			viper.GetString(constant.HTTPServerAddress))

		server := &http.Server{Addr: viper.GetString(constant.HTTPServerAddress), Handler: appIns.getServerMux()}
		appIns.mHTTPServer = server

		err := server.ListenAndServeTLS(viper.GetString(constant.HTTPServerCertFile),
			viper.GetString(constant.HTTPServerKeyFile))

		return err
	} else {
		logger.L("app").Info("http server started @ " +
			viper.GetString(constant.HTTPServerAddress))

		server := &http.Server{Addr: viper.GetString(constant.HTTPServerAddress), Handler: appIns.getServerMux()}
		appIns.mHTTPServer = server

		err := server.ListenAndServe()
		return err
	}
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

func initializeClusterDefaults() {
	/*Initialize cluster defaults*/
	if viper.GetBool(constant.Join) {
		return
	}

	nodeAdmin := GetApp().
		GetFlamed().
		NewNodeAdmin(1, viper.GetDuration(constant.GlobalRequestTimeout))

	if nodeAdmin == nil {
		panic("Failed to create new node admin")
	}

	for {
		leaderID, leaderAvailable, _ := nodeAdmin.GetLeaderID()
		if leaderAvailable {
			logger.L("app").Info("leader found", zap.Uint64("leaderID", leaderID))
			break
		}

		time.Sleep(500 * time.Millisecond)
	}

	// Creating default super user
	lastAppliedIndex, err := nodeAdmin.GetAppliedIndex()
	if err != nil {
		return
	}

	logger.L("app").Info("last applied index", zap.Uint64("lastAppliedIndex", lastAppliedIndex))

	if lastAppliedIndex > 0 {
		return
	}

	admin := GetApp().
		GetFlamed().
		NewAdmin(1, viper.GetDuration(constant.GlobalRequestTimeout))

	if admin == nil {
		logger.L("app").Error("failed to create new Admin")
		return
	}

	pha := GetApp().GetPasswordHashAlgorithmFactory()
	if !pha.IsAlgorithmAvailable(variable.DefaultPasswordHashAlgorithm) {
		logger.L("app").Error(variable.DefaultPasswordHashAlgorithm +
			" password hash algorithm is to available")
		return
	}

	encoded, err := pha.MakePassword("admin",
		crypto.GetRandomString(12),
		variable.DefaultPasswordHashAlgorithm)

	if err != nil {
		logger.L("app").Error("make password returned error", zap.Error(err))
		return
	}

	superUser := &pb.User{
		UserType:  pb.UserType_SUPER_USER,
		Roles:     "*",
		Username:  "admin",
		Password:  encoded,
		CreatedAt: uint64(time.Now().UnixNano()),
		UpdatedAt: uint64(time.Now().UnixNano()),
	}

	pr, err := admin.UpsertUser(superUser)

	if err != nil {
		logger.L("app").Error("upsert user error", zap.Error(err))
		return
	}

	if pr.Status == pb.Status_ACCEPTED {
		logger.L("app").Info("admin user created")
	} else {
		logger.L("app").Error("failed to create admin user")
	}
}

func getNodeConfiguration(raftStoragePath string) *conf.NodeConfiguration {
	return &conf.NodeConfiguration{
		NodeConfigurationInput: conf.NodeConfigurationInput{
			NodeHostDir:                   raftStoragePath,
			WALDir:                        raftStoragePath,
			DeploymentID:                  viper.GetUint64(constant.DeploymentID),
			RTTMillisecond:                viper.GetUint64(constant.RTTMillisecond),
			RaftAddress:                   viper.GetString(constant.RaftAddress),
			ListenAddress:                 "",
			MutualTLS:                     viper.GetBool(constant.MutualTLS),
			CAFile:                        viper.GetString(constant.CAFile),
			CertFile:                      viper.GetString(constant.CertFile),
			KeyFile:                       viper.GetString(constant.KeyFile),
			MaxSendQueueSize:              viper.GetUint64(constant.MaxSendQueueSize),
			MaxReceiveQueueSize:           viper.GetUint64(constant.MaxReceiveQueueSize),
			EnableMetrics:                 viper.GetBool(constant.EnableMetrics),
			MaxSnapshotSendBytesPerSecond: viper.GetUint64(constant.MaxSnapshotSendBytesPerSecond),
			MaxSnapshotRecvBytesPerSecond: viper.GetUint64(constant.MaxSnapshotRecvBytesPerSecond),
			NotifyCommit:                  viper.GetBool(constant.NotifyCommit),

			SystemTickerPrecision: viper.GetDuration(constant.SystemTickerPrecision),

			LogDBConfig:         config.GetTinyMemLogDBConfig(),
			LogDBFactory:        variable.DefaultLogDbFactory,
			RaftRPCFactory:      variable.DefaultRaftRPCFactory,
			RaftEventListener:   variable.DefaultRaftEventListener,
			SystemEventListener: variable.DefaultSystemEventListener,
		},
	}
}

func getRaftConfiguration() *conf.RaftConfiguration {
	return &conf.RaftConfiguration{
		RaftConfigurationInput: conf.RaftConfigurationInput{
			NodeID:                 viper.GetUint64(constant.NodeID),
			CheckQuorum:            viper.GetBool(constant.CheckQuorum),
			ElectionRTT:            viper.GetUint64(constant.ElectionRTT),
			HeartbeatRTT:           viper.GetUint64(constant.HeartbeatRTT),
			SnapshotEntries:        viper.GetUint64(constant.SnapshotEntries),
			CompactionOverhead:     viper.GetUint64(constant.CompactionOverhead),
			OrderedConfigChange:    viper.GetBool(constant.OrderedConfigChange),
			MaxInMemLogSize:        viper.GetUint64(constant.MaxInMemLogSize),
			DisableAutoCompactions: viper.GetBool(constant.DisableAutoCompactions),
			IsObserver:             viper.GetBool(constant.IsObserver),
			IsWitness:              viper.GetBool(constant.IsWitness),
			Quiesce:                viper.GetBool(constant.Quiesce),
		},
	}
}

func startCluster() {
	clusterID := uint64(1)
	clusterName := "cluster-1"

	im := getInitialMembers(strings.Split(viper.GetString(constant.InitialMembers), ";"))

	if len(im) == 0 {
		if !viper.GetBool(constant.Join) {
			im[viper.GetUint64(constant.NodeID)] = viper.GetString(constant.RaftAddress)
		}
	}

	clusterStoragePath := viper.GetString(constant.StoragePath) + "/" + clusterName

	storagedConfiguration := &conf.StoragedConfiguration{
		StoragedConfigurationInput: conf.StoragedConfigurationInput{
			AutoIndexMeta:         true,
			IndexEnable:           true,
			StateStoragePath:      clusterStoragePath + "/state",
			StateStorageSecretKey: nil,
			IndexStoragePath:      clusterStoragePath + "/index",
			IndexStorageSecretKey: nil,

			AutoBuildIndex: true,

			ProposalReceiver: GetApp().GetProposalReceiver(),
		},
		TransactionProcessorMap: GetApp().GetTPMap(),
	}

	//for _, tp := range GetApp().mTransactionProcessorMap {
	//	storagedConfiguration.AddTransactionProcessor(tp)
	//}

	var clusterConfiguration iface.IOnDiskClusterConfiguration

	if viper.GetBool(constant.Join) {
		clusterConfiguration = conf.SimpleOnDiskClusterConfiguration(
			clusterID,
			clusterName,
			nil,
			true)
	} else {
		clusterConfiguration = conf.SimpleOnDiskClusterConfiguration(
			clusterID,
			clusterName,
			im,
			false)
	}

	raftConfiguration := getRaftConfiguration()

	err := GetApp().GetFlamed().StartOnDiskCluster(
		clusterConfiguration,
		storagedConfiguration,
		raftConfiguration)

	if err != nil {
		panic(err)
	}
}
