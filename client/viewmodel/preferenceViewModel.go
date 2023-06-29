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
	api         model.IConfigApi
	receiver    *model.TimeStatusReceiver
	window      fyne.Window
	config      config.Config
	workHrsData binding.Int

	WorkHrsLabelData   binding.String
	WorkAlarmEnabled   binding.Bool
	WorkAlarmBeforeMin binding.Int
	LocalPathData      binding.String
}

func NewPreferenceViewModel(a *model.AppContainer, w fyne.Window) *PreferenceViewModel {
	config, _ := a.ConfigApi.LoadConfig()
	workHrsData := binding.NewInt()
	workHrsLabelData := binding.NewSprintf("%d時間", workHrsData)
	localPathData := binding.NewString()
	localPathData.Set(a.LocalPath.GetLocalPath())

	workAlarmEnabled := binding.NewBool()
	workAlarmEnabled.Set(config.WorkAlarm.IsEnabled)
	workAlarmBeforeMin := binding.NewInt()
	workAlarmBeforeMin.Set(config.WorkAlarm.BeforeMinutes)

	vm := PreferenceViewModel{
		api:                a.ConfigApi,
		receiver:           a.Receiver,
		window:             w,
		config:             *config,
		workHrsData:        workHrsData,
		WorkHrsLabelData:   workHrsLabelData,
		LocalPathData:      localPathData,
		WorkAlarmEnabled:   workAlarmEnabled,
		WorkAlarmBeforeMin: workAlarmBeforeMin,
	}

	workAlarmEnabled.AddListener(binding.NewDataListener(func() {
		if v, err := workAlarmEnabled.Get(); err == nil {
			vm.config.WorkAlarm.IsEnabled = v
		}
	}))
	workAlarmBeforeMin.AddListener(binding.NewDataListener(func() {
		if v, err := workAlarmBeforeMin.Get(); err == nil {
			vm.config.WorkAlarm.BeforeMinutes = v
		}
	}))

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
