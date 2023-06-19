package client

import (
	"attendance-record/client/resource"
	"attendance-record/shared"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/driver/desktop"
)

type Client struct {
	app *shared.App
}

func NewClient(app *shared.App) *Client {
	return &Client{app: app}
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

	r := initTimeStatusReceiver(c.app)
	v := initTimeStatusView(w, r)
	w.SetContent(v.Container)
	w.ShowAndRun()
}
