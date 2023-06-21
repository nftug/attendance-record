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
	api          model.ITimeStatusApi
	receiver     *model.TimeStatusReceiver
	update       []func()
	Data         []dto.TimeStatusDto
	SelIdx       int
	Window       fyne.Window
	OnUnselected func()
}

func NewHistoryViewModel(a *model.AppContainer, w fyne.Window) *HistoryViewModel {
	vm := &HistoryViewModel{api: a.Api, receiver: a.Receiver, Window: w, SelIdx: -1}
	a.Receiver.AddUpdateOuterFunc(vm.InvokeUpdate)
	vm.InvokeUpdate()
	return vm
}

func (vm *HistoryViewModel) AddUpdateFunc(f ...func()) {
	vm.update = append(vm.update, f...)
}

func (vm *HistoryViewModel) InvokeUpdate() {
	vm.Data = vm.api.GetAll()
	for _, f := range vm.update {
		f()
	}
}

func (vm *HistoryViewModel) Delete(item dto.TimeStatusDto) error {
	if err := vm.api.Delete(item.Type, item.Id); err != nil {
		return err
	}

	// vm.InvokeUpdate()
	vm.receiver.SetCurrentStatus()

	vm.SelIdx = -1
	if vm.OnUnselected != nil {
		vm.OnUnselected()
	}
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

	// vm.InvokeUpdate()
	vm.receiver.SetCurrentStatus()
	return nil
}

func (vm *HistoryViewModel) GetSelected() (dto.TimeStatusDto, error) {
	if vm.SelIdx < 0 {
		return dto.TimeStatusDto{}, errors.New("no items selected")
	} else if vm.SelIdx > len(vm.Data)-1 {
		return dto.TimeStatusDto{}, errors.New("too large index")
	}
	return vm.Data[vm.SelIdx], nil
}
