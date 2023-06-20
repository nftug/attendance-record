package viewmodel

import (
	"time"

	"github.com/multiplay/go-cticker"
)

type RenderHandler func(v string)

func updateClock(f RenderHandler) {
	f(time.Now().Format("15:04:05"))
}

func UpdateByTick(f RenderHandler) {
	updateClock(f)

	go func() {
		for range cticker.New(time.Second, 100*time.Millisecond).C {
			updateClock(f)
		}
	}()
}
