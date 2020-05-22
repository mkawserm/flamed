package app

import (
	"fmt"
	"github.com/mkawserm/flamed/pkg/app/view/graphql"
	"github.com/mkawserm/flamed/pkg/crypto"
	"github.com/mkawserm/flamed/pkg/iface"
	"github.com/mkawserm/flamed/pkg/logger"
	"github.com/mkawserm/flamed/pkg/pb"
	"github.com/mkawserm/flamed/pkg/tp/accesscontrol"
	"github.com/mkawserm/flamed/pkg/tp/indexmeta"
	"github.com/mkawserm/flamed/pkg/tp/intkey"
	"github.com/mkawserm/flamed/pkg/tp/json"
	"github.com/mkawserm/flamed/pkg/tp/user"
	"github.com/mkawserm/flamed/pkg/variable"
	"github.com/spf13/viper"
	"net/http"
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
	mServerMux  *http.ServeMux
	mHTTPServer *http.Server

	mViewsInitialized bool
	mDefaultViewFlag  bool

	mCommandsInitialized bool
	mDefaultCommandFlag  bool

	mTPInitialized bool
	mDefaultTPFlag bool

	mFlamed                       *flamed.Flamed
	mRootCommand                  *cobra.Command
	mTransactionProcessorList     []iface.ITransactionProcessor
	mPasswordHashAlgorithmFactory iface.IPasswordHashAlgorithmFactory

	mProposalReceiver func(*pb.Proposal, pb.Status)
}

func (a *App) SetProposalReceiver(pr func(*pb.Proposal, pb.Status)) {
	a.mProposalReceiver = pr
}

func (a *App) GetProposalReceiver() func(*pb.Proposal, pb.Status) {
	return a.mProposalReceiver
}

func (a *App) getServer() *http.Server {
	return a.mHTTPServer
}

func (a *App) getServerMux() *http.ServeMux {
	return a.mServerMux
}

func (a *App) SetPasswordHashAlgorithmFactory(f iface.IPasswordHashAlgorithmFactory) {
	a.mPasswordHashAlgorithmFactory = f
}

func (a *App) GetPasswordHashAlgorithmFactory() iface.IPasswordHashAlgorithmFactory {
	return a.mPasswordHashAlgorithmFactory
}

func (a *App) EnableDefaultTransactionProcessors() {
	a.mDefaultTPFlag = true
}

func (a *App) DisableDefaultTransactionProcessors() {
	a.mDefaultTPFlag = false
}

func (a *App) EnableDefaultViews() {
	a.mDefaultViewFlag = true
}

func (a *App) DisableDefaultViews() {
	a.mDefaultViewFlag = false
}

func (a *App) EnableDefaultCommands() {
	a.mDefaultCommandFlag = true
}

func (a *App) DisableDefaultCommands() {
	a.mDefaultCommandFlag = false
}

func (a *App) UpdateGlobalRequestTimeout(timeout time.Duration) {
	viper.Set("GlobalRequestTimeout", timeout)
}

func (a *App) AddView(pattern string, handler func(http.ResponseWriter, *http.Request)) {
	a.mServerMux.HandleFunc(pattern, handler)
}

func (a *App) AddTransactionProcessor(tp iface.ITransactionProcessor) {
	a.mTransactionProcessorList = append(a.mTransactionProcessorList, tp)
}

func (a *App) GetFlamed() *flamed.Flamed {
	return a.mFlamed
}

func (a *App) AddCommand(commands ...*cobra.Command) {
	a.mRootCommand.AddCommand(commands...)
}

func (a *App) setup() {
	a.mFlamed = flamed.NewFlamed()

	a.mRootCommand = &cobra.Command{
		Use:   variable.Name,
		Short: variable.ShortDescription,
		Long:  variable.LongDescription,
		Run: func(cmd *cobra.Command, _ []string) {
			fmt.Println(cmd.UsageString())
		},
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			a.initBeforeCommandExecution()
		},
	}

	a.mRootCommand.
		PersistentFlags().
		StringVar(
			&configFile,
			"config",
			"",
			"Config file")

	initAllPersistentFlags(a.mRootCommand)
}

func (a *App) initCommands() {
	if !a.mDefaultCommandFlag {
		return
	}

	if !a.mCommandsInitialized {
		/* initialize all commands here */

		a.mRootCommand.AddCommand(RunCMD)
		a.mRootCommand.AddCommand(ConfigCMD)
		a.mRootCommand.AddCommand(AuthorCMD)
		a.mRootCommand.AddCommand(VersionCMD)
		a.mCommandsInitialized = true
	}
}

func (a *App) initViews() {
	if !a.mDefaultViewFlag {
		return
	}

	if !a.mViewsInitialized {
		/* initialize all views here */

		// graphql view
		a.AddView("/graphql",
			graphql.NewView(a.mFlamed,
				a.mTransactionProcessorList,
				a.mPasswordHashAlgorithmFactory).GetHTTPHandler())

		a.mViewsInitialized = true
	}
}

func (a *App) initTransactionProcessors() {
	if !a.mDefaultTPFlag {
		return
	}

	if !a.mTPInitialized {
		/* initialize all transaction processors here */
		a.mTransactionProcessorList = append(a.mTransactionProcessorList, &user.User{})
		a.mTransactionProcessorList = append(a.mTransactionProcessorList, &json.JSON{})
		a.mTransactionProcessorList = append(a.mTransactionProcessorList, &intkey.IntKey{})
		a.mTransactionProcessorList = append(a.mTransactionProcessorList, &indexmeta.IndexMeta{})
		a.mTransactionProcessorList = append(a.mTransactionProcessorList, &accesscontrol.AccessControl{})

		a.mTPInitialized = true
	}
}

func (a *App) initBeforeCommandExecution() {
	/* setup defaults */
	logger.GetLoggerFactory().ChangeLogLevel(viper.GetString("LogLevel"))
}

func (a *App) Execute() error {
	a.initTransactionProcessors()
	a.initCommands()

	if a.mPasswordHashAlgorithmFactory == nil {
		a.mPasswordHashAlgorithmFactory = crypto.DefaultPasswordHashAlgorithmFactory()
	}

	return a.mRootCommand.Execute()
}

func GetApp() *App {
	return appIns
}

func init() {
	cobra.OnInitialize(initConfig)
	appOnce.Do(func() {
		appIns = &App{
			mServerMux:           &http.ServeMux{},
			mCommandsInitialized: false,
			mDefaultCommandFlag:  true,

			mViewsInitialized: false,
			mDefaultViewFlag:  true,

			mTPInitialized: false,
			mDefaultTPFlag: true,
		}

		appIns.setup()
	})
}
