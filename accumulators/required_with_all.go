package accumulators

import (
	"github.com/SladeThe/yav"
)

type RequiredWithAll[T any] struct {
	provideFunc ProvideFunc[T]
	enabled     bool
}

func NewRequiredWithAll[T any](provideFunc ProvideFunc[T]) RequiredWithAll[T] {
	return RequiredWithAll[T]{
		provideFunc: provideFunc,
		enabled:     true,
	}
}

func (r RequiredWithAll[T]) String(value string) RequiredWithAll[T] {
	r.enabled = r.enabled && value != ""
	return r
}

func (r RequiredWithAll[T]) Bytes(value []byte) RequiredWithAll[T] {
	r.enabled = r.enabled && len(value) != 0
	return r
}

func (r RequiredWithAll[T]) Bool(value bool) RequiredWithAll[T] {
	r.enabled = r.enabled && value
	return r
}

func (r RequiredWithAll[T]) Int(value int) RequiredWithAll[T] {
	r.enabled = r.enabled && value != 0
	return r
}

func (r RequiredWithAll[T]) Int8(value int8) RequiredWithAll[T] {
	r.enabled = r.enabled && value != 0
	return r
}

func (r RequiredWithAll[T]) Int16(value int16) RequiredWithAll[T] {
	r.enabled = r.enabled && value != 0
	return r
}

func (r RequiredWithAll[T]) Int32(value int32) RequiredWithAll[T] {
	r.enabled = r.enabled && value != 0
	return r
}

func (r RequiredWithAll[T]) Int64(value int64) RequiredWithAll[T] {
	r.enabled = r.enabled && value != 0
	return r
}

func (r RequiredWithAll[T]) Uint(value int) RequiredWithAll[T] {
	r.enabled = r.enabled && value != 0
	return r
}

func (r RequiredWithAll[T]) Uint8(value uint8) RequiredWithAll[T] {
	r.enabled = r.enabled && value != 0
	return r
}

func (r RequiredWithAll[T]) Uint16(value uint16) RequiredWithAll[T] {
	r.enabled = r.enabled && value != 0
	return r
}

func (r RequiredWithAll[T]) Uint32(value uint32) RequiredWithAll[T] {
	r.enabled = r.enabled && value != 0
	return r
}

func (r RequiredWithAll[T]) Uint64(value uint64) RequiredWithAll[T] {
	r.enabled = r.enabled && value != 0
	return r
}

func (r RequiredWithAll[T]) Fields(fields string) yav.ValidateFunc[T] {
	return r.provideFunc(fields, r.enabled)
}
