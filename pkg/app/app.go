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
	appOnce    sync.Once
	appIns     *App
	configFile string
)

type App struct {
	mDataStoragePath      string
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
	flamedHomePath := home + "/flamed"

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
