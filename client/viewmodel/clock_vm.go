package viewmodel

import (
	"time"
)

type RenderHandler func(v string)

func updateClock(f RenderHandler) {
	f(time.Now().Format("03:04:05"))
}

func UpdateByTick(f RenderHandler) {
	updateClock(f)

	go func() {
		for range time.Tick(time.Second) {
			updateClock(f)
		}
	}()
}
