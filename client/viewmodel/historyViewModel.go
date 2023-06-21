package viewmodel

import (
	"attendance-record/client/model"
	"attendance-record/domain/dto"
	"attendance-record/shared/util"
	"errors"
	"fmt"
	"time"

	"fyne.io/fyne/v2"
)

type HistoryViewModel struct {
	api      model.ITimeStatusApi
	receiver *model.TimeStatusReceiver
	update   []func()
	Data     []dto.TimeStatusDto
	Selected *dto.TimeStatusDto
	Window   fyne.Window
}

func NewHistoryViewModel(a *model.AppContainer, w fyne.Window) *HistoryViewModel {
	data := a.Api.GetAll()
	return &HistoryViewModel{api: a.Api, receiver: a.Receiver, Data: data, Window: w}
}

func (vm *HistoryViewModel) AddUpdateFunc(f ...func()) {
	vm.update = append(vm.update, f...)
}

func (vm *HistoryViewModel) InvokeUpdate() {
	vm.Data = vm.api.GetAll()
	vm.receiver.SetCurrentStatus()
	for _, f := range vm.update {
		f()
	}
}

func (vm *HistoryViewModel) DeleteSelected() error {
	if vm.Selected == nil {
		return errors.New("no items selected")
	}

	if err := vm.api.Delete((*vm.Selected).Type, (*vm.Selected).Id); err != nil {
		return err
	}
	vm.InvokeUpdate()
	return nil
}

func (vm *HistoryViewModel) Edit(item dto.TimeStatusDto, start string, end string) error {
	layout := "15:04"
	cmd := dto.TimeStatusCommandDto{}

	startedOn, err := time.Parse(layout, start)
	if err != nil {
		return err
	}
	cmd.StartedOn = util.SetHourAndMinute(item.StartedOn, startedOn)

	if end != "" {
		endedOn, err := time.Parse(layout, end)
		if err != nil {
			return err
		}
		cmd.EndedOn = util.SetHourAndMinute(item.StartedOn, endedOn)
	} else {
		cmd.EndedOn = *new(time.Time)
	}

	if err = vm.api.Update(item.Type, item.Id, cmd); err != nil {
		fmt.Println(err)
		return err
	}
	vm.InvokeUpdate()
	return nil
}
