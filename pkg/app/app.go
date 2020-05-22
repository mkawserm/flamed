package app

import (
	"fmt"
	"github.com/mkawserm/flamed/pkg/iface"
	"github.com/mkawserm/flamed/pkg/logger"
	"github.com/mkawserm/flamed/pkg/tp/accesscontrol"
	"github.com/mkawserm/flamed/pkg/tp/indexmeta"
	"github.com/mkawserm/flamed/pkg/tp/intkey"
	"github.com/mkawserm/flamed/pkg/tp/json"
	"github.com/mkawserm/flamed/pkg/tp/user"
	"github.com/spf13/viper"
	"sync"
	"time"

	"github.com/mkawserm/flamed/pkg/flamed"
	"github.com/spf13/cobra"
)

var (
	appOnce sync.Once
	appIns  *App
)

type App struct {
	mGlobalRequestTimeout time.Duration
	mFlamed               *flamed.Flamed
	mRootCommand          *cobra.Command
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
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			a.initBeforeExecution()
		},
	}

	a.mRootCommand.
		PersistentFlags().
		StringVar(
			&configFile,
			"config",
			"",
			"config file (default is $HOME/flamed/.flamed.yaml)")

	initAllPersistentFlags(a.mRootCommand)

	a.mRootCommand.AddCommand(runCMD)
	a.mRootCommand.AddCommand(configCMD)
	a.mRootCommand.AddCommand(authorCMD)
	a.mRootCommand.AddCommand(versionCMD)
}

func (a *App) initBeforeExecution() {
	/* setup defaults */
	a.mGlobalRequestTimeout = viper.GetDuration("GlobalRequestTimeout")
	logger.GetLoggerFactory().ChangeLogLevel(viper.GetString("LogLevel"))
}

func (a *App) Execute() error {
	return a.mRootCommand.Execute()
}

func GetApp() *App {
	return appIns
}

func init() {
	cobra.OnInitialize(initConfig)
	appOnce.Do(func() {
		appIns = &App{}
		appIns.setup()
	})
}
