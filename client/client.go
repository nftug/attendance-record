package client

import (
	"attendance-record/client/model"
	"attendance-record/client/resource"
	"attendance-record/client/view"

	"fyne.io/fyne/v2/app"
)

type Client struct {
	appContainer *model.AppContainer
}

func NewClient(appContainer *model.AppContainer) *Client {
	return &Client{appContainer: appContainer}
}

func (c *Client) Run() {
	a := app.NewWithID("attendance-record")
	a.Settings().SetTheme(&resource.MyTheme{})
	w := a.NewWindow("勤怠記録")

	view.SetSystemTrayMenu(c.appContainer, w)

	w.SetContent(view.NewTimeStatusView(c.appContainer, w))
	w.ShowAndRun()
}
