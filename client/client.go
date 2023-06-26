package client

import (
	"attendance-record/client/model"
	"attendance-record/client/resource"
	"attendance-record/client/view"
	"attendance-record/infrastructure/localpath"

	"fyne.io/fyne/v2/app"
)

type Client struct {
	api       model.ITimeStatusApi
	cfgApi    model.IConfigApi
	localpath *localpath.LocalPathService
}

func NewClient(api model.ITimeStatusApi, cfgApi model.IConfigApi, lp *localpath.LocalPathService) *Client {
	return &Client{api: api, cfgApi: cfgApi, localpath: lp}
}

func (c *Client) Run() {
	a := app.New()
	a.Settings().SetTheme(&resource.MyTheme{})
	w := a.NewWindow("勤怠記録")

	ac := model.NewAppContainer(c.api, c.cfgApi, c.localpath, a)
	view.SetSystemTrayMenu(ac, w)

	w.SetContent(view.NewTimeStatusView(ac, w))
	w.ShowAndRun()
}
