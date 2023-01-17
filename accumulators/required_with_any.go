package accumulators

import (
	"github.com/SladeThe/yav"
)

type RequiredWithAny[T any] struct {
	provideFunc ProvideFunc[T]
	enabled     bool
}

func NewRequiredWithAny[T any](provideFunc ProvideFunc[T]) RequiredWithAny[T] {
	return RequiredWithAny[T]{
		provideFunc: provideFunc,
		enabled:     false,
	}
}

func (r RequiredWithAny[T]) String(value string) RequiredWithAny[T] {
	r.enabled = r.enabled || value != ""
	return r
}

func (r RequiredWithAny[T]) Bytes(value []byte) RequiredWithAny[T] {
	r.enabled = r.enabled || len(value) != 0
	return r
}

func (r RequiredWithAny[T]) Bool(value bool) RequiredWithAny[T] {
	r.enabled = r.enabled || value
	return r
}

func (r RequiredWithAny[T]) Int(value int) RequiredWithAny[T] {
	r.enabled = r.enabled || value != 0
	return r
}

func (r RequiredWithAny[T]) Int8(value int8) RequiredWithAny[T] {
	r.enabled = r.enabled || value != 0
	return r
}

func (r RequiredWithAny[T]) Int16(value int16) RequiredWithAny[T] {
	r.enabled = r.enabled || value != 0
	return r
}

func (r RequiredWithAny[T]) Int32(value int32) RequiredWithAny[T] {
	r.enabled = r.enabled || value != 0
	return r
}

func (r RequiredWithAny[T]) Int64(value int64) RequiredWithAny[T] {
	r.enabled = r.enabled || value != 0
	return r
}

func (r RequiredWithAny[T]) Uint(value int) RequiredWithAny[T] {
	r.enabled = r.enabled || value != 0
	return r
}

func (r RequiredWithAny[T]) Uint8(value uint8) RequiredWithAny[T] {
	r.enabled = r.enabled || value != 0
	return r
}

func (r RequiredWithAny[T]) Uint16(value uint16) RequiredWithAny[T] {
	r.enabled = r.enabled || value != 0
	return r
}

func (r RequiredWithAny[T]) Uint32(value uint32) RequiredWithAny[T] {
	r.enabled = r.enabled || value != 0
	return r
}

func (r RequiredWithAny[T]) Uint64(value uint64) RequiredWithAny[T] {
	r.enabled = r.enabled || value != 0
	return r
}

func (r RequiredWithAny[T]) Names(names string) yav.ValidateFunc[T] {
	return r.provideFunc(names, r.enabled)
}
