package app

import (
	"context"
	"fmt"
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
	Short: "Run Flamed server",
	Run: func(cmd *cobra.Command, args []string) {
		logger.GetLoggerFactory().ChangeLogLevel(viper.GetString("LogLevel"))
		if len(viper.GetString(StoragePath)) == 0 {
			fmt.Println("StoragePath can not be empty")
			return
		}

		if len(viper.GetString(RaftAddress)) == 0 {
			fmt.Println("RaftAddress can not be empty")
			return
		}

		if len(viper.GetString(HTTPAddress)) == 0 {
			fmt.Println("HTTPAddress can not be empty")
			return
		}

		im := getInitialMembers(strings.Split(viper.GetString(InitialMembers), ";"))

		if len(im) == 0 {
			if !viper.GetBool("Join") {
				im[viper.GetUint64(NodeID)] = viper.GetString(RaftAddress)
			}
		}

		raftStoragePath := viper.GetString(StoragePath) + "/raft"

		nodeConfiguration := &conf.NodeConfiguration{
			NodeConfigurationInput: conf.NodeConfigurationInput{
				NodeHostDir:                   raftStoragePath,
				WALDir:                        raftStoragePath,
				DeploymentID:                  viper.GetUint64(DeploymentID),
				RTTMillisecond:                viper.GetUint64(RTTMillisecond),
				RaftAddress:                   viper.GetString(RaftAddress),
				ListenAddress:                 "",
				MutualTLS:                     viper.GetBool(MutualTLS),
				CAFile:                        viper.GetString(CAFile),
				CertFile:                      viper.GetString(CertFile),
				KeyFile:                       viper.GetString(KeyFile),
				MaxSendQueueSize:              viper.GetUint64(MaxSendQueueSize),
				MaxReceiveQueueSize:           viper.GetUint64(MaxReceiveQueueSize),
				EnableMetrics:                 viper.GetBool(EnableMetrics),
				MaxSnapshotSendBytesPerSecond: viper.GetUint64(MaxSnapshotSendBytesPerSecond),
				MaxSnapshotRecvBytesPerSecond: viper.GetUint64(MaxSnapshotRecvBytesPerSecond),
				NotifyCommit:                  viper.GetBool(NotifyCommit),

				SystemTickerPrecision: viper.GetDuration(SystemTickerPrecision),

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

		clusterStoragePath := viper.GetString(StoragePath) + "/cluster-1"

		storagedConfiguration := conf.SimpleStoragedConfiguration(clusterStoragePath, nil)
		for _, tp := range GetApp().mTransactionProcessor {
			storagedConfiguration.AddTransactionProcessor(tp)
		}

		clusterConfiguration := conf.SimpleOnDiskClusterConfiguration(
			1,
			"cluster-1",
			im,
			viper.GetBool(Join))

		raftConfiguration := &conf.RaftConfiguration{
			RaftConfigurationInput: conf.RaftConfigurationInput{
				NodeID:                 viper.GetUint64(NodeID),
				CheckQuorum:            viper.GetBool(CheckQuorum),
				ElectionRTT:            viper.GetUint64(ElectionRTT),
				HeartbeatRTT:           viper.GetUint64(HeartbeatRTT),
				SnapshotEntries:        viper.GetUint64(SnapshotEntries),
				CompactionOverhead:     viper.GetUint64(CompactionOverhead),
				OrderedConfigChange:    viper.GetBool(OrderedConfigChange),
				MaxInMemLogSize:        viper.GetUint64(MaxInMemLogSize),
				DisableAutoCompactions: viper.GetBool(DisableAutoCompactions),
				IsObserver:             viper.GetBool(IsObserver),
				IsWitness:              viper.GetBool(IsWitness),
				Quiesce:                viper.GetBool(Quiesce),
			},
		}

		err = GetApp().GetFlamed().StartOnDiskCluster(
			clusterConfiguration,
			storagedConfiguration,
			raftConfiguration)

		if err != nil {
			panic(err)
		}

		// initialize cluster defaults
		// like admin user and other things
		initializeClusterDefaults()

		// initialize views
		GetApp().initViews()

		// run server and wait for shutdown
		runServerAndWaitForShutDown()
	},
} // Command

func runServerAndWaitForShutDown() {
	// idleChan channel is dedicated for shutting down all active connections.
	// Once actual shutdown occurred by closing this channel, the main goroutine
	// is shutdown.
	idleChan := make(chan struct{})
	go func() {
		// signChan channel is used to transmit signal notifications.
		signChan := make(chan os.Signal, 1)
		// Catch and relay certain signal(s) to signChan channel.
		signal.Notify(signChan, os.Interrupt, syscall.SIGTERM)
		// Blocking until a signal is sent over signChan channel.
		sig := <-signChan

		logger.L("app").Info("shutdown signal received",
			zap.String("signal", sig.String()))

		// Create a new context with a timeout duration. It helps allowing
		// timeout duration to all active connections in order for them to
		// finish their job. Any connections that wont complete within the
		// allowed timeout duration gets halted.
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		if err := GetApp().getServer().Shutdown(ctx); err == context.DeadlineExceeded {
			logger.L("app").Info("shutdown: halted active connections")
		}

		// Actual shutdown trigger.
		close(idleChan)
	}()

	if err := runHTTPServer(); err == http.ErrServerClosed {
		logger.L("app").Info("flamed shutdown started")
	}

	// Blocking until the shutdown to complete then inform the main goroutine.
	<-idleChan
	logger.L("app").Info("flamed shutdown complete")
}

func runHTTPServer() error {
	logger.L("app").Info("Running HTTP Server")
	if viper.GetBool(HTTPServerTLS) {
		logger.L("app").Info("HTTP Server with TLS started")
		server := &http.Server{Addr: viper.GetString(HTTPAddress), Handler: appIns.getServerMux()}
		appIns.mHTTPServer = server

		err := server.ListenAndServeTLS(viper.GetString(HTTPServerCertFile),
			viper.GetString(HTTPServerKeyFile))

		return err
	} else {
		logger.L("app").Info("HTTP Server started")
		server := &http.Server{Addr: viper.GetString(HTTPAddress), Handler: appIns.getServerMux()}
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
	if viper.GetBool("Join") {
		return
	}

	clusterAdmin := GetApp().
		GetFlamed().
		NewClusterAdmin(1, viper.GetDuration(GlobalRequestTimeout))

	if clusterAdmin == nil {
		panic("Failed to create new cluster admin")
	}

	for {
		leaderID, leaderAvailable, _ := clusterAdmin.GetLeaderID()
		if leaderAvailable {
			logger.L("app").Info("leader found", zap.Uint64("leaderID", leaderID))
			break
		}

		time.Sleep(500 * time.Millisecond)
	}

	// Creating default super user
	lastAppliedIndex, err := clusterAdmin.GetAppliedIndex()
	if err != nil {
		return
	}

	logger.L("app").Info("last applied index", zap.Uint64("lastAppliedIndex", lastAppliedIndex))

	if lastAppliedIndex > 0 {
		return
	}

	admin := GetApp().
		GetFlamed().
		NewAdmin(1, viper.GetDuration(GlobalRequestTimeout))

	if admin == nil {
		logger.L("app").Error("failed to create new Admin")
		return
	}

	pha := GetApp().GetPasswordHashAlgorithmFactory()
	if !pha.IsAlgorithmAvailable(DefaultPasswordHashAlgorithm) {
		logger.L("app").Error(DefaultPasswordHashAlgorithm +
			" password hash algorithm is to available")
		return
	}

	encoded, err := pha.MakePassword("admin",
		crypto.GetRandomString(12),
		DefaultPasswordHashAlgorithm)

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
