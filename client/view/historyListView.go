package view

import (
	"attendance-record/client/viewmodel"
	"attendance-record/domain/enum"
	"attendance-record/shared/util"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

func NewHistoryListView(vm *viewmodel.HistoryViewModel) fyne.CanvasObject {
	colName := []string{"種類", "開始時刻", "終了時刻", "時間数"}

	list := widget.NewTable(
		func() (int, int) { return len(vm.Data) + 1, 4 },
		func() fyne.CanvasObject {
			return widget.NewLabel("wide content")
		},
		func(i widget.TableCellID, o fyne.CanvasObject) {
			l := o.(*widget.Label)
			if i.Row == 0 {
				l.SetText(colName[i.Col])
			} else {
				item := vm.Data[i.Row-1]

				switch i.Col {
				case 0:
					if item.Type == enum.Work {
						l.SetText("勤務")
					} else {
						l.SetText("休憩")
					}
				case 1:
					l.SetText(util.FormatDateTime(item.StartedOn))
				case 2:
					l.SetText(util.FormatDateTime(item.EndedOn))
				case 3:
					l.SetText(item.TotalTime.String())
				}
			}
		},
	)

	list.OnSelected = func(i widget.TableCellID) {
		if i.Row == 0 {
			return
		}
		vm.Selected = &vm.Data[i.Row-1]
	}

	vm.AddUpdateFunc(list.Refresh)

	return list
}
