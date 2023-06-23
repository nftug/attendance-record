package model

import "fyne.io/fyne/v2"

type AppContainer struct {
	Api       ITimeStatusApi
	Receiver  *TimeStatusReceiver
	ConfigApi IConfigApi
	App       fyne.App
}

func NewAppContainer(api ITimeStatusApi, cfgApi IConfigApi, app fyne.App) *AppContainer {
	return &AppContainer{
		Api:       api,
		ConfigApi: cfgApi,
		Receiver:  NewTimeStatusReceiverSingleton(api),
		App:       app,
	}
}
