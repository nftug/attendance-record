package client

import (
	"attendance-record/client/model"
	"attendance-record/client/resource"
	"attendance-record/client/view"
	"attendance-record/shared/appinfo"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"github.com/mitchellh/go-ps"
	"github.com/sqweek/dialog"
)

type Client struct {
	appContainer *model.AppContainer
}

func NewClient(appContainer *model.AppContainer) *Client {
	return &Client{appContainer: appContainer}
}

func (c *Client) Run() {
	a := app.NewWithID(appinfo.AppName)
	a.Settings().SetTheme(&resource.MyTheme{})
	w := a.NewWindow(appinfo.AppTitle)
	w.SetIcon(resource.ResourceIconPng)

	view.SetSystemTrayMenu(c.appContainer, w)

	w.SetContent(view.NewTimeStatusView(c.appContainer, w))
	w.Resize(fyne.NewSize(360, 330))
	w.SetFixedSize(true)

	w.ShowAndRun()
}

func CheckIfAppRunning() {
	pid := os.Getpid()
	curProc, _ := ps.FindProcess(pid)

	procs, _ := ps.Processes()
	for _, p := range procs {
		if p.Executable() == curProc.Executable() && p.Pid() != pid {
			dialog.Message("既にアプリが起動しています。\n二重起動はできません。").Title(appinfo.AppTitle).Error()
			os.Exit(1)
		}
	}
}
