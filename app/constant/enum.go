package constant

type Enum[T any] interface {
	GetName() string
	GetValue() T
	FindByName(name string) (Enum[T], bool)
	FindByValue(value T) (Enum[T], bool)
}
