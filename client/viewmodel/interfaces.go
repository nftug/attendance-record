package viewmodel

type WidgetWithText interface {
	SetText(v string)
}

type Binding[T any] interface {
	Get() (T, error)
	Set(T) error
}
