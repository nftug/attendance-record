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
	receiver *model.TimeStatusReceiver
}

func NewClient(receiver *model.TimeStatusReceiver) *Client {
	return &Client{receiver: receiver}
}

func (c *Client) Run() {
	a := app.New()
	a.Settings().SetTheme(&resource.MyTheme{})
	w := a.NewWindow("勤怠記録")

	if desk, ok := a.(desktop.App); ok {
		m := fyne.NewMenu(
			"勤怠記録",
			fyne.NewMenuItem("表示", func() { w.Show() }),
		)
		desk.SetSystemTrayMenu(m)
	}
	w.SetCloseIntercept(func() { w.Hide() })

	w.SetContent(view.NewTimeStatusView(w, c.receiver))
	w.ShowAndRun()
}
