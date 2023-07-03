package viewmodel

import (
	"attendance-record/client/model"
	"attendance-record/domain/config"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/dialog"
	"github.com/skratchdot/open-golang/open"
)

type PreferenceViewModel struct {
	api      model.IConfigApi
	receiver *model.TimeStatusReceiver
	window   fyne.Window
	config   config.Config

	workHrsData      binding.Int
	WorkHrsLabelData binding.String

	WorkAlarmEnabled   binding.Bool
	WorkAlarmBeforeMin binding.Int
	WorkAlarmSnoozeMin binding.Int
	RestAlarmEnabled   binding.Bool
	RestAlarmHrs       binding.Int
	RestAlarmMin       binding.Int
	RestAlarmSnoozeMin binding.Int

	LocalPathData binding.String
}

func NewPreferenceViewModel(a *model.AppContainer, w fyne.Window) *PreferenceViewModel {
	config, _ := a.ConfigApi.LoadConfig()
	workHrsData := binding.NewInt()
	workHrsLabelData := binding.NewSprintf("%d時間", workHrsData)

	vm := PreferenceViewModel{
		api:                a.ConfigApi,
		receiver:           a.Receiver,
		window:             w,
		config:             *config,
		workHrsData:        workHrsData,
		WorkHrsLabelData:   workHrsLabelData,
		WorkAlarmEnabled:   binding.NewBool(),
		WorkAlarmBeforeMin: binding.NewInt(),
		WorkAlarmSnoozeMin: binding.NewInt(),
		RestAlarmEnabled:   binding.NewBool(),
		RestAlarmHrs:       binding.NewInt(),
		RestAlarmMin:       binding.NewInt(),
		RestAlarmSnoozeMin: binding.NewInt(),
		LocalPathData:      binding.NewString(),
	}

	vm.OnChangeWorkHrsData(float64(config.WorkHours))

	vm.WorkAlarmEnabled.Set(config.WorkAlarm.IsEnabled)
	vm.WorkAlarmBeforeMin.Set(config.WorkAlarm.BeforeMinutes)
	vm.WorkAlarmSnoozeMin.Set(config.WorkAlarm.SnoozeMinutes)

	vm.WorkAlarmEnabled.AddListener(binding.NewDataListener(func() {
		v, _ := vm.WorkAlarmEnabled.Get()
		vm.config.WorkAlarm.IsEnabled = v
	}))
	vm.WorkAlarmBeforeMin.AddListener(binding.NewDataListener(func() {
		v, _ := vm.WorkAlarmBeforeMin.Get()
		vm.config.WorkAlarm.BeforeMinutes = v
	}))
	vm.WorkAlarmSnoozeMin.AddListener(binding.NewDataListener(func() {
		v, _ := vm.WorkAlarmSnoozeMin.Get()
		vm.config.WorkAlarm.SnoozeMinutes = v
	}))

	vm.RestAlarmEnabled.Set(config.RestAlarm.IsEnabled)
	vm.RestAlarmHrs.Set(config.RestAlarm.Hours)
	vm.RestAlarmMin.Set(config.RestAlarm.Minutes)
	vm.RestAlarmSnoozeMin.Set(config.RestAlarm.SnoozeMinutes)

	vm.RestAlarmEnabled.AddListener(binding.NewDataListener(func() {
		v, _ := vm.RestAlarmEnabled.Get()
		vm.config.RestAlarm.IsEnabled = v
	}))
	vm.RestAlarmHrs.AddListener(binding.NewDataListener(func() {
		v, _ := vm.RestAlarmHrs.Get()
		vm.config.RestAlarm.Hours = v
	}))
	vm.RestAlarmMin.AddListener(binding.NewDataListener(func() {
		v, _ := vm.RestAlarmMin.Get()
		vm.config.RestAlarm.Minutes = v
	}))
	vm.RestAlarmSnoozeMin.AddListener(binding.NewDataListener(func() {
		v, _ := vm.RestAlarmSnoozeMin.Get()
		vm.config.RestAlarm.SnoozeMinutes = v
	}))

	vm.LocalPathData.Set(a.LocalPath.GetLocalPath())

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

func (vm *PreferenceViewModel) OpenLocalPath() {
	path, _ := vm.LocalPathData.Get()
	open.Start(path)
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
