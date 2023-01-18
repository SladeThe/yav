package accumulators

import (
	"reflect"
	"time"

	"github.com/SladeThe/yav"
)

type RequiredWithoutAny[T any] struct {
	provideFunc ProvideFunc[T]
	required    bool
}

func NewRequiredWithoutAny[T any](provideFunc ProvideFunc[T]) RequiredWithoutAny[T] {
	return RequiredWithoutAny[T]{
		provideFunc: provideFunc,
		required:    false,
	}
}

func (r RequiredWithoutAny[T]) String(value string) RequiredWithoutAny[T] {
	r.required = r.required || value == ""
	return r
}

func (r RequiredWithoutAny[T]) Bytes(value []byte) RequiredWithoutAny[T] {
	r.required = r.required || len(value) == 0
	return r
}

func (r RequiredWithoutAny[T]) Slice(value any) RequiredWithoutAny[T] {
	r.required = r.required || reflect.ValueOf(value).Len() == 0
	return r
}

func (r RequiredWithoutAny[T]) Map(value any) RequiredWithoutAny[T] {
	r.required = r.required || reflect.ValueOf(value).Len() == 0
	return r
}

func (r RequiredWithoutAny[T]) Time(value time.Time) RequiredWithoutAny[T] {
	r.required = r.required || value.IsZero()
	return r
}

func (r RequiredWithoutAny[T]) Bool(value bool) RequiredWithoutAny[T] {
	r.required = r.required || !value
	return r
}

func (r RequiredWithoutAny[T]) Int(value int) RequiredWithoutAny[T] {
	r.required = r.required || value == 0
	return r
}

func (r RequiredWithoutAny[T]) Int8(value int8) RequiredWithoutAny[T] {
	r.required = r.required || value == 0
	return r
}

func (r RequiredWithoutAny[T]) Int16(value int16) RequiredWithoutAny[T] {
	r.required = r.required || value == 0
	return r
}

func (r RequiredWithoutAny[T]) Int32(value int32) RequiredWithoutAny[T] {
	r.required = r.required || value == 0
	return r
}

func (r RequiredWithoutAny[T]) Int64(value int64) RequiredWithoutAny[T] {
	r.required = r.required || value == 0
	return r
}

func (r RequiredWithoutAny[T]) Uint(value int) RequiredWithoutAny[T] {
	r.required = r.required || value == 0
	return r
}

func (r RequiredWithoutAny[T]) Uint8(value uint8) RequiredWithoutAny[T] {
	r.required = r.required || value == 0
	return r
}

func (r RequiredWithoutAny[T]) Uint16(value uint16) RequiredWithoutAny[T] {
	r.required = r.required || value == 0
	return r
}

func (r RequiredWithoutAny[T]) Uint32(value uint32) RequiredWithoutAny[T] {
	r.required = r.required || value == 0
	return r
}

func (r RequiredWithoutAny[T]) Uint64(value uint64) RequiredWithoutAny[T] {
	r.required = r.required || value == 0
	return r
}

func (r RequiredWithoutAny[T]) Names(names string) yav.ValidateFunc[T] {
	return r.provideFunc(names, r.required)
}
