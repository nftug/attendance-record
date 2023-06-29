package component

import (
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/driver/mobile"
	"fyne.io/fyne/v2/widget"
)

type numericalEntry struct {
	widget.Entry
	intData binding.Int
}

func NewNumericalEntryWithData(data binding.Int) *numericalEntry {
	entry := &numericalEntry{}
	entry.ExtendBaseWidget(entry)
	entry.intData = data
	v, _ := data.Get()
	entry.Text = strconv.Itoa(v)
	return entry
}

func (e *numericalEntry) TypedRune(r rune) {
	if r >= '0' && r <= '9' {
		e.Entry.TypedRune(r)

		if v, err := strconv.Atoi(e.Text); err == nil {
			e.intData.Set(v)
		}
	}
}

func (e *numericalEntry) TypedShortcut(shortcut fyne.Shortcut) {
	paste, ok := shortcut.(*fyne.ShortcutPaste)
	if !ok {
		e.Entry.TypedShortcut(shortcut)
		return
	}

	content := paste.Clipboard.Content()
	if v, err := strconv.Atoi(content); err == nil {
		e.Entry.TypedShortcut(shortcut)
		e.intData.Set(v)
	}
}

func (e *numericalEntry) Keyboard() mobile.KeyboardType {
	return mobile.NumberKeyboard
}
