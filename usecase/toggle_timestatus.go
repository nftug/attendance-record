package usecase

import (
	"domain/entity"
	"fmt"
	"shared"
)

func ToggleWork(tss *entity.TimeStatusSet) {
	tss.ToggleWork()
	shared.PrintAsJson(tss.Work)
	if !tss.Work.IsActive {
		fmt.Println(tss.Work.TotalTime)
	}
}

func ToggleRest(tss *entity.TimeStatusSet) {
	tss.ToggleRest()
	shared.PrintAsJson(tss.Rest)
}
