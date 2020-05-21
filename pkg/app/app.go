package app

import (
	"sync"

	"github.com/mkawserm/flamed/pkg/flamed"
	"github.com/spf13/cobra"
)

var (
	appOnce sync.Once
	appIns  *App
)

type App struct {
	mFlamed      *flamed.Flamed
	mRootCommand *cobra.Command
}

func (a *App) GetFlamed() *flamed.Flamed {
	return a.mFlamed
}

func (a *App) setup() {
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
		appIns = &App{
			mFlamed: flamed.NewFlamed(),
		}
		appIns.setup()
	})
}
