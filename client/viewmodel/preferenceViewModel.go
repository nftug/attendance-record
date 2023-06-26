package viewmodel

import (
	"attendance-record/client/model"
	"attendance-record/domain/config"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/dialog"
)

type PreferenceViewModel struct {
	api              model.IConfigApi
	receiver         *model.TimeStatusReceiver
	window           fyne.Window
	config           config.Config
	workHrsData      binding.Int
	WorkHrsLabelData binding.String
}

func NewPreferenceViewModel(a *model.AppContainer, w fyne.Window) *PreferenceViewModel {
	config, _ := a.ConfigApi.LoadConfig()
	workHrsData := binding.NewInt()
	workHrsLabelData := binding.NewSprintf("%d時間", workHrsData)

	vm := PreferenceViewModel{
		api:              a.ConfigApi,
		receiver:         a.Receiver,
		window:           w,
		config:           *config,
		workHrsData:      workHrsData,
		WorkHrsLabelData: workHrsLabelData,
	}

	vm.OnChangeWorkHrsData(float64(config.WorkHours))
	return &vm
}

func (vm *PreferenceViewModel) GetWorkHour() float64 {
	v, err := vm.workHrsData.Get()
	if err != nil {
		return float64(vm.config.WorkHours)
	}
	return float64(v)
}

func (vm *PreferenceViewModel) OnChangeWorkHrsData(val float64) {
	vm.config.WorkHours = int(val)
	vm.workHrsData.Set(int(val))
}

func (vm *PreferenceViewModel) OnClickApply() {
	if err := vm.api.SaveConfig(vm.config); err != nil {
		dialog.ShowError(err, vm.window)
	}
	vm.receiver.InvokeUpdate()
}

func (vm *PreferenceViewModel) OnClickSave() {
	if err := vm.api.SaveConfig(vm.config); err != nil {
		dialog.ShowError(err, vm.window)
	} else {
		vm.receiver.InvokeUpdate()
		vm.window.Hide()
	}
}

func (vm *PreferenceViewModel) OnClickCancel() {
	vm.window.Hide()
}