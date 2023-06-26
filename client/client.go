package client

import (
	"attendance-record/client/model"
	"attendance-record/client/resource"
	"attendance-record/client/view"

	"fyne.io/fyne/v2/app"
)

type Client struct {
	api    model.ITimeStatusApi
	cfgApi model.IConfigApi
}

func NewClient(api model.ITimeStatusApi, cfgApi model.IConfigApi) *Client {
	return &Client{api: api, cfgApi: cfgApi}
}

func (c *Client) Run() {
	a := app.New()
	a.Settings().SetTheme(&resource.MyTheme{})
	w := a.NewWindow("勤怠記録")

	ac := model.NewAppContainer(c.api, c.cfgApi, a)
	view.SetSystemTrayMenu(ac, w)

	w.SetContent(view.NewTimeStatusView(ac, w))
	w.SetFixedSize(true)
	w.ShowAndRun()
}
