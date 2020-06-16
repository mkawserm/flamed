package app

import (
	"fmt"
	"github.com/mkawserm/flamed/pkg/app/graphql"
	graphql2 "github.com/mkawserm/flamed/pkg/app/http/view/graphql"
	utility2 "github.com/mkawserm/flamed/pkg/app/utility"
	"github.com/mkawserm/flamed/pkg/constant"
	"github.com/mkawserm/flamed/pkg/context"
	"github.com/mkawserm/flamed/pkg/crypto"
	"github.com/mkawserm/flamed/pkg/iface"
	"github.com/mkawserm/flamed/pkg/logger"
	"github.com/mkawserm/flamed/pkg/pb"
	"github.com/mkawserm/flamed/pkg/tp/accesscontrol"
	"github.com/mkawserm/flamed/pkg/tp/indexmeta"
	"github.com/mkawserm/flamed/pkg/tp/intkey"
	"github.com/mkawserm/flamed/pkg/tp/json"
	"github.com/mkawserm/flamed/pkg/tp/user"
	"github.com/mkawserm/flamed/pkg/utility"
	"github.com/mkawserm/flamed/pkg/variable"
	"github.com/spf13/viper"
	"go.uber.org/zap"
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

	mRootCommand   *cobra.Command
	mFlamedContext *context.FlamedContext

	mProposalReceiver func(*pb.Proposal, pb.Status)

	mGraphQL             *graphql.GraphQL
	mGraphQLQuery        map[string]graphql.GQLHandler
	mGraphQLMutation     map[string]graphql.GQLHandler
	mGraphQLSubscription map[string]graphql.GQLHandler

	mMutex sync.Mutex
}

func (a *App) GetFlamedContext() *context.FlamedContext {
	a.mMutex.Lock()
	defer a.mMutex.Unlock()

	return a.mFlamedContext
}

func (a *App) GetTPMap() map[string]iface.ITransactionProcessor {
	a.mMutex.Lock()
	defer a.mMutex.Unlock()

	return a.mFlamedContext.TransactionProcessorMap()
}

func (a *App) GetTransactionProcessorMap() map[string]iface.ITransactionProcessor {
	a.mMutex.Lock()
	defer a.mMutex.Unlock()

	return a.mFlamedContext.TransactionProcessorMap()
}

func (a *App) SetProposalReceiver(pr func(*pb.Proposal, pb.Status)) {
	a.mMutex.Lock()
	defer a.mMutex.Unlock()

	a.mProposalReceiver = pr
}

func (a *App) GetProposalReceiver() func(*pb.Proposal, pb.Status) {
	a.mMutex.Lock()
	defer a.mMutex.Unlock()

	return a.mProposalReceiver
}

func (a *App) getHTTPServer() *http.Server {
	a.mMutex.Lock()
	defer a.mMutex.Unlock()

	return a.mHTTPServer
}

func (a *App) getServerMux() *http.ServeMux {
	a.mMutex.Lock()
	defer a.mMutex.Unlock()

	return a.mServerMux
}

func (a *App) SetPasswordHashAlgorithmFactory(f iface.IPasswordHashAlgorithmFactory) {
	a.mMutex.Lock()
	defer a.mMutex.Unlock()

	variable.DefaultPasswordHashAlgorithmFactory = f
}

func (a *App) GetPasswordHashAlgorithmFactory() iface.IPasswordHashAlgorithmFactory {
	a.mMutex.Lock()
	defer a.mMutex.Unlock()

	return variable.DefaultPasswordHashAlgorithmFactory
}

func (a *App) EnableDefaultTransactionProcessors() {
	a.mMutex.Lock()
	defer a.mMutex.Unlock()

	a.mDefaultTPFlag = true
}

func (a *App) DisableDefaultTransactionProcessors() {
	a.mMutex.Lock()
	defer a.mMutex.Unlock()

	a.mDefaultTPFlag = false
}

func (a *App) EnableDefaultViews() {
	a.mMutex.Lock()
	defer a.mMutex.Unlock()

	a.mDefaultViewFlag = true
}

func (a *App) DisableDefaultViews() {
	a.mMutex.Lock()
	defer a.mMutex.Unlock()

	a.mDefaultViewFlag = false
}

func (a *App) EnableDefaultCommands() {
	a.mMutex.Lock()
	defer a.mMutex.Unlock()

	a.mDefaultCommandFlag = true
}

func (a *App) DisableDefaultCommands() {
	a.mMutex.Lock()
	defer a.mMutex.Unlock()

	a.mDefaultCommandFlag = false
}

func (a *App) UpdateGlobalRequestTimeout(timeout time.Duration) {
	a.mMutex.Lock()
	defer a.mMutex.Unlock()

	viper.Set(constant.GlobalRequestTimeout, timeout)
}

func (a *App) AddGraphQLQuery(name string, handler graphql.GQLHandler) {
	a.mMutex.Lock()
	defer a.mMutex.Unlock()

	a.mGraphQLQuery[name] = handler
}

func (a *App) AddGraphQLMutation(name string, handler graphql.GQLHandler) {
	a.mMutex.Lock()
	defer a.mMutex.Unlock()

	a.mGraphQLMutation[name] = handler
}

func (a *App) AddGraphQLSubscription(name string, handler graphql.GQLHandler) {
	a.mMutex.Lock()
	defer a.mMutex.Unlock()

	a.mGraphQLSubscription[name] = handler
}

func (a *App) AddView(pattern string, handler func(http.ResponseWriter, *http.Request)) {
	a.mMutex.Lock()
	defer a.mMutex.Unlock()

	a.mServerMux.HandleFunc(pattern, handler)
}

func (a *App) AddTransactionProcessor(tp iface.ITransactionProcessor) {
	a.mMutex.Lock()
	defer a.mMutex.Unlock()

	a.mFlamedContext.AddTP(tp)
}

func (a *App) GetFlamed() *flamed.Flamed {
	a.mMutex.Lock()
	defer a.mMutex.Unlock()

	return a.mFlamedContext.Flamed()
}

func (a *App) AddCommand(commands ...*cobra.Command) {
	a.mMutex.Lock()
	defer a.mMutex.Unlock()

	a.mRootCommand.AddCommand(commands...)
}

func (a *App) setup() {
	//a.mMutex.Lock()
	//defer a.mMutex.Unlock()

	/*initialize all attributes*/
	a.mServerMux = &http.ServeMux{}
	a.mFlamedContext = context.NewFlamedContext()

	a.mCommandsInitialized = false
	a.mDefaultCommandFlag = true
	a.mViewsInitialized = false
	a.mDefaultViewFlag = true
	a.mTPInitialized = false
	a.mDefaultTPFlag = true
	a.mGraphQLQuery = make(map[string]graphql.GQLHandler)
	a.mGraphQLMutation = make(map[string]graphql.GQLHandler)
	a.mGraphQLSubscription = make(map[string]graphql.GQLHandler)

	a.mFlamedContext.SetFlamed(flamed.NewFlamed())
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

func (a *App) initGraphQL() {
	if !viper.GetBool(constant.EnableHTTPServer) &&
		!viper.GetBool(constant.EnableGRPCServer) {
		return
	}

	if !viper.GetBool(constant.EnableGraphQLOverHTTP) &&
		!viper.GetBool(constant.EnableGraphQLOverGRPC) {
		return
	}

	if viper.GetBool(constant.EnableHTTPServer) ||
		viper.GetBool(constant.EnableGRPCServer) {
		if !viper.GetBool(constant.EnableGraphQLOverHTTP) &&
			!viper.GetBool(constant.EnableGraphQLOverGRPC) {
			return
		}

		if viper.GetBool(constant.EnableHTTPServer) &&
			!viper.GetBool(constant.EnableGRPCServer) &&
			!viper.GetBool(constant.EnableGraphQLOverHTTP) {
			return
		}

		if !viper.GetBool(constant.EnableHTTPServer) &&
			viper.GetBool(constant.EnableGRPCServer) &&
			!viper.GetBool(constant.EnableGraphQLOverGRPC) {
			return
		}
	}

	a.mGraphQL = graphql.NewGraphQL(a.mFlamedContext)

	for k, v := range a.mGraphQLQuery {
		a.mGraphQL.AddQueryField(k, v)
	}

	for k, v := range a.mGraphQLMutation {
		a.mGraphQL.AddMutationField(k, v)
	}

	for k, v := range a.mGraphQLSubscription {
		a.mGraphQL.AddSubscriptionField(k, v)
	}

	_, _ = a.mGraphQL.BuildSchema()
}

func (a *App) initGraphQLView() {
	if !viper.GetBool(constant.EnableGraphQLOverHTTP) {
		return
	}

	if a.mGraphQL == nil {
		return
	}

	// graphql view
	schema, _ := a.mGraphQL.BuildSchema()
	a.AddView("/graphql", graphql2.NewGraphQLView(a.mFlamedContext, schema).GetHTTPHandler())
}

func (a *App) initViews() {
	if !viper.GetBool(constant.EnableHTTPServer) {
		return
	}

	if !a.mDefaultViewFlag {
		return
	}

	if !a.mViewsInitialized {
		/* initialize all views here */
		a.initGraphQLView()

		a.mViewsInitialized = true
	}
}

func (a *App) initTransactionProcessors() {
	if !a.mDefaultTPFlag {
		return
	}

	if !a.mTPInitialized {
		/* initialize all transaction processors here */
		a.AddTransactionProcessor(&user.User{})
		a.AddTransactionProcessor(&json.JSON{})
		a.AddTransactionProcessor(&intkey.IntKey{})
		a.AddTransactionProcessor(&indexmeta.IndexMeta{})
		a.AddTransactionProcessor(&accesscontrol.AccessControl{})
		a.mTPInitialized = true
	}
}

func (a *App) initBeforeCommandExecution() {
	/* setup defaults */

	// set log level
	logger.GetLoggerFactory().ChangeLogLevel(viper.GetString(constant.LogLevel))
	a.mFlamedContext.SetGlobalTimeout(viper.GetDuration(constant.GlobalRequestTimeout))

	// set cors options
	utility2.GetCORSOptions().AllowAllOrigins = viper.GetBool(constant.CORSAllowAllOrigins)
	utility2.GetCORSOptions().AllowOrigins = viper.GetStringSlice(constant.CORSAllowOrigins)
	utility2.GetCORSOptions().AllowCredentials = viper.GetBool(constant.CORSAllowCredentials)
	utility2.GetCORSOptions().AllowMethods = viper.GetStringSlice(constant.CORSAllowMethods)
	utility2.GetCORSOptions().AllowHeaders = viper.GetStringSlice(constant.CORSAllowHeaders)
	utility2.GetCORSOptions().ExposeHeaders = viper.GetStringSlice(constant.CORSExposeHeaders)
	utility2.GetCORSOptions().MaxAge = viper.GetDuration(constant.CORSMaxAge)

	if a.mProposalReceiver == nil {
		a.mProposalReceiver = func(proposal *pb.Proposal, status pb.Status) {

			logger.L("app").Debug("received proposal",
				zap.Int32("status", int32(status)),
				zap.String("uuid", utility.UUIDToString(proposal.Uuid)),
				zap.Uint64("createdAt", proposal.CreatedAt),
				zap.Int("transactionLength", len(proposal.Transactions)),
			)
		}
	}

	/* init graphql */
	a.initGraphQL()

	/* init views */
	a.initViews()
}

func (a *App) Execute() error {
	a.initTransactionProcessors()
	a.initCommands()

	if variable.DefaultPasswordHashAlgorithmFactory == nil {
		variable.DefaultPasswordHashAlgorithmFactory = crypto.DefaultPasswordHashAlgorithmFactory()
	}

	return a.mRootCommand.Execute()
}

func (a *App) setHTTPServer(server *http.Server) {
	a.mMutex.Lock()
	defer a.mMutex.Unlock()
	a.mHTTPServer = server
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
