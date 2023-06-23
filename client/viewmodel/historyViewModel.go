package viewmodel

import (
	"attendance-record/client/model"
	"attendance-record/domain/dto"
	"attendance-record/domain/enum"
	"errors"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/dialog"
)

type HistoryViewModel struct {
	api          model.ITimeStatusApi
	receiver     *model.TimeStatusReceiver
	update       []func()
	Data         []dto.TimeStatusDto
	SelIdx       int
	CurDt        time.Time
	CurDtData    binding.String
	Window       fyne.Window
	OnUnselected func()
}

func NewHistoryViewModel(a *model.AppContainer, w fyne.Window) *HistoryViewModel {
	now := time.Now()
	vm := &HistoryViewModel{
		api:       a.Api,
		receiver:  a.Receiver,
		Window:    w,
		SelIdx:    -1,
		CurDt:     time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, time.Local),
		CurDtData: binding.NewString(),
	}
	a.Receiver.AddUpdateOuterFunc(vm.InvokeUpdate)

	vm.InvokeUpdate()
	return vm
}

func (vm *HistoryViewModel) AddUpdateFunc(f ...func()) {
	vm.update = append(vm.update, f...)
}

func (vm *HistoryViewModel) InvokeUpdate() {
	vm.CurDtData.Set(vm.CurDt.Format("2006年 01月"))
	d, err := vm.api.FindByMonth(vm.CurDt.Year(), vm.CurDt.Month())
	if err != nil {
		dialog.ShowError(err, vm.Window)
		return
	}
	vm.Data = d

	for _, f := range vm.update {
		f()
	}
}

func (vm *HistoryViewModel) NextMonth() {
	vm.CurDt = vm.CurDt.AddDate(0, 1, 0)
	vm.InvokeUpdate()
}

func (vm *HistoryViewModel) PrevMonth() {
	vm.CurDt = vm.CurDt.AddDate(0, -1, 0)
	vm.InvokeUpdate()
}

func (vm *HistoryViewModel) Delete(item dto.TimeStatusDto) error {
	if err := vm.api.Delete(item.Type, item.Id); err != nil {
		dialog.ShowError(err, vm.Window)
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
	date := item.StartedOn.Format(dto.DateFormat)
	cmd, err := dto.NewTimeStatusCommandDto(date, start, end)
	if err != nil {
		dialog.ShowError(err, vm.Window)
		return err
	}

	if err = vm.api.Update(item.Type, item.Id, cmd); err != nil {
		dialog.ShowError(err, vm.Window)
		return err
	}

	// vm.InvokeUpdate()
	vm.receiver.SetCurrentStatus()
	return nil
}

func (vm *HistoryViewModel) Create(t enum.TimeStatusType, date string, start string, end string) error {
	cmd, err := dto.NewTimeStatusCommandDto(date, start, end)
	if err != nil {
		dialog.ShowError(err, vm.Window)
		return err
	}

	if err = vm.api.Create(t, cmd); err != nil {
		dialog.ShowError(err, vm.Window)
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
