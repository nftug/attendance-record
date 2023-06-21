package model

import "fyne.io/fyne/v2"

type AppContainer struct {
	Api      ITimeStatusApi
	Receiver *TimeStatusReceiver
	App      fyne.App
}

func NewAppContainer(api ITimeStatusApi, app fyne.App) *AppContainer {
	return &AppContainer{
		Api:      api,
		Receiver: NewTimeStatusReceiverSingleton(api),
		App:      app,
	}
}
