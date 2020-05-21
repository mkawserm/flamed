package app

import (
	"github.com/mkawserm/flamed/pkg/iface"
	"github.com/mkawserm/flamed/pkg/tp/accesscontrol"
	"github.com/mkawserm/flamed/pkg/tp/indexmeta"
	"github.com/mkawserm/flamed/pkg/tp/intkey"
	"github.com/mkawserm/flamed/pkg/tp/json"
	"github.com/mkawserm/flamed/pkg/tp/user"
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
		Run: func(_ *cobra.Command, _ []string) {
		},
	}

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

func init() {
	appOnce.Do(func() {
		appIns = &App{}
		appIns.setup()
	})
}
