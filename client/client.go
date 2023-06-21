package client

import (
	"attendance-record/client/model"
	"attendance-record/client/resource"
	"attendance-record/client/view"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/driver/desktop"
)

type Client struct {
	Api model.ITimeStatusApi
}

func NewClient(api model.ITimeStatusApi) *Client {
	return &Client{Api: api}
}

func (c *Client) Run() {
	a := app.New()
	a.Settings().SetTheme(&resource.MyTheme{})
	w := a.NewWindow("勤怠記録")

	appContainer := model.NewAppContainer(c.Api, a)

	if desk, ok := a.(desktop.App); ok {
		m := fyne.NewMenu(
			"勤怠記録",
			fyne.NewMenuItem("表示", func() { w.Show() }),
		)
		desk.SetSystemTrayMenu(m)
	}
	w.SetCloseIntercept(func() { w.Hide() })

	w.SetContent(view.NewTimeStatusView(appContainer, w))
	w.ShowAndRun()
}
