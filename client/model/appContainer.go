package model

import (
	"attendance-record/infrastructure/localpath"
)

type AppContainer struct {
	Api       ITimeStatusApi
	Receiver  *TimeStatusReceiver
	ConfigApi IConfigApi
	LocalPath *localpath.LocalPathService
}

func NewAppContainer(api ITimeStatusApi, cfgApi IConfigApi, lp *localpath.LocalPathService) *AppContainer {
	return &AppContainer{
		Api:       api,
		ConfigApi: cfgApi,
		Receiver:  NewTimeStatusReceiverSingleton(api),
		LocalPath: lp,
	}
}
