package view

import (
	"attendance-record/client/viewmodel"
	"attendance-record/domain/enum"
	"attendance-record/shared/util"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

func NewHistoryListView(vm *viewmodel.HistoryViewModel) fyne.CanvasObject {
	colName := []string{"日付", "種類", "開始時刻", "終了時刻", "時間数"}

	list := widget.NewTable(
		func() (int, int) { return len(vm.Data) + 1, 5 },
		func() fyne.CanvasObject {
			return widget.NewLabel("wide content")
		},
		func(i widget.TableCellID, o fyne.CanvasObject) {
			l := o.(*widget.Label)
			if i.Row == 0 {
				l.SetText(colName[i.Col])
				l.TextStyle = fyne.TextStyle{Bold: true}
			} else {
				item := vm.Data[i.Row-1]

				switch i.Col {
				case 0:
					l.SetText(util.GetDate(item.StartedOn).Format("2006-01-02"))
				case 1:
					if item.Type == enum.Work {
						l.SetText("勤務")
					} else {
						l.SetText("休憩")
					}
				case 2:
					l.SetText(util.FormatDateTime(item.StartedOn))
				case 3:
					l.SetText(util.FormatDateTime(item.EndedOn))
				case 4:
					l.SetText(item.TotalTime.String())
				}
			}
		},
	)

	list.OnSelected = func(i widget.TableCellID) {
		list.Select(widget.TableCellID{Row: i.Row, Col: 0})
		if i.Row == 0 {
			list.Unselect(i)
		}
		vm.SelIdx = i.Row - 1
	}
	vm.OnUnselected = list.UnselectAll

	vm.AddUpdateFunc(list.Refresh)

	return list
}
