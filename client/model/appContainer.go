package model

import (
	"attendance-record/infrastructure/localpath"

	"fyne.io/fyne/v2"
)

type AppContainer struct {
	Api       ITimeStatusApi
	Receiver  *TimeStatusReceiver
	ConfigApi IConfigApi
	LocalPath *localpath.LocalPathService
	App       fyne.App
}

func NewAppContainer(api ITimeStatusApi, cfgApi IConfigApi, lp *localpath.LocalPathService, app fyne.App) *AppContainer {
	return &AppContainer{
		Api:       api,
		ConfigApi: cfgApi,
		Receiver:  NewTimeStatusReceiverSingleton(api),
		LocalPath: lp,
		App:       app,
	}
}
