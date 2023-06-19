package viewmodel

type WidgetWithText interface {
	SetText(v string)
}

type Button interface {
	WidgetWithText
	Enable()
	Disable()
}

type Binding[T any] interface {
	Get() (T, error)
	Set(T) error
}

type Window interface {
	SetTitle(v string)
}
