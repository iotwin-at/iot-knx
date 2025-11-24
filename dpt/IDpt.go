package dpt

type IDpt[T any] interface {
	Name() string
	Unit() string
	String() string
	Pack() []byte
	Value() T
}
