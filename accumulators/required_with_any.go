package accumulators

import (
	"reflect"
	"time"

	"github.com/SladeThe/yav"
)

type RequiredWithAny[T any] struct {
	provideFunc ProvideFunc[T]
	required    bool
}

func NewRequiredWithAny[T any](provideFunc ProvideFunc[T]) RequiredWithAny[T] {
	return RequiredWithAny[T]{
		provideFunc: provideFunc,
		required:    false,
	}
}

func (r RequiredWithAny[T]) String(value string) RequiredWithAny[T] {
	r.required = r.required || value != ""
	return r
}

func (r RequiredWithAny[T]) Bytes(value []byte) RequiredWithAny[T] {
	r.required = r.required || len(value) != 0
	return r
}

func (r RequiredWithAny[T]) Slice(value any) RequiredWithAny[T] {
	r.required = r.required || value != nil && reflect.ValueOf(value).Len() != 0
	return r
}

func (r RequiredWithAny[T]) Map(value any) RequiredWithAny[T] {
	r.required = r.required || value != nil && reflect.ValueOf(value).Len() != 0
	return r
}

func (r RequiredWithAny[T]) Time(value time.Time) RequiredWithAny[T] {
	r.required = r.required || !value.IsZero()
	return r
}

func (r RequiredWithAny[T]) Zeroer(value yav.Zeroer) RequiredWithAny[T] {
	r.required = r.required || !value.IsZero()
	return r
}

func (r RequiredWithAny[T]) Bool(value bool) RequiredWithAny[T] {
	r.required = r.required || value
	return r
}

func (r RequiredWithAny[T]) Int(value int) RequiredWithAny[T] {
	r.required = r.required || value != 0
	return r
}

func (r RequiredWithAny[T]) Int8(value int8) RequiredWithAny[T] {
	r.required = r.required || value != 0
	return r
}

func (r RequiredWithAny[T]) Int16(value int16) RequiredWithAny[T] {
	r.required = r.required || value != 0
	return r
}

func (r RequiredWithAny[T]) Int32(value int32) RequiredWithAny[T] {
	r.required = r.required || value != 0
	return r
}

func (r RequiredWithAny[T]) Int64(value int64) RequiredWithAny[T] {
	r.required = r.required || value != 0
	return r
}

func (r RequiredWithAny[T]) Uint(value int) RequiredWithAny[T] {
	r.required = r.required || value != 0
	return r
}

func (r RequiredWithAny[T]) Uint8(value uint8) RequiredWithAny[T] {
	r.required = r.required || value != 0
	return r
}

func (r RequiredWithAny[T]) Uint16(value uint16) RequiredWithAny[T] {
	r.required = r.required || value != 0
	return r
}

func (r RequiredWithAny[T]) Uint32(value uint32) RequiredWithAny[T] {
	r.required = r.required || value != 0
	return r
}

func (r RequiredWithAny[T]) Uint64(value uint64) RequiredWithAny[T] {
	r.required = r.required || value != 0
	return r
}

func (r RequiredWithAny[T]) Names(names string) yav.ValidateFunc[T] {
	return r.provideFunc(names, r.required)
}
