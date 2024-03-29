package accumulators

import (
	"reflect"
	"time"

	"github.com/SladeThe/yav"
)

type RequiredWithAll[T any] struct {
	provideFunc ProvideFunc[T]
	required    bool
}

func NewRequiredWithAll[T any](provideFunc ProvideFunc[T]) RequiredWithAll[T] {
	return RequiredWithAll[T]{
		provideFunc: provideFunc,
		required:    true,
	}
}

func (r RequiredWithAll[T]) String(value string) RequiredWithAll[T] {
	r.required = r.required && value != ""
	return r
}

func (r RequiredWithAll[T]) Bytes(value []byte) RequiredWithAll[T] {
	r.required = r.required && len(value) != 0
	return r
}

func (r RequiredWithAll[T]) Slice(value any) RequiredWithAll[T] {
	r.required = r.required && value != nil && reflect.ValueOf(value).Len() != 0
	return r
}

func (r RequiredWithAll[T]) Map(value any) RequiredWithAll[T] {
	r.required = r.required && value != nil && reflect.ValueOf(value).Len() != 0
	return r
}

func (r RequiredWithAll[T]) Pointer(value any) RequiredWithAll[T] {
	r.required = r.required && !reflect.ValueOf(value).IsNil()
	return r
}

func (r RequiredWithAll[T]) Time(value time.Time) RequiredWithAll[T] {
	r.required = r.required && !value.IsZero()
	return r
}

func (r RequiredWithAll[T]) Zeroer(value yav.Zeroer) RequiredWithAll[T] {
	r.required = r.required && !value.IsZero()
	return r
}

func (r RequiredWithAll[T]) Bool(value bool) RequiredWithAll[T] {
	r.required = r.required && value
	return r
}

func (r RequiredWithAll[T]) Int(value int) RequiredWithAll[T] {
	r.required = r.required && value != 0
	return r
}

func (r RequiredWithAll[T]) Int8(value int8) RequiredWithAll[T] {
	r.required = r.required && value != 0
	return r
}

func (r RequiredWithAll[T]) Int16(value int16) RequiredWithAll[T] {
	r.required = r.required && value != 0
	return r
}

func (r RequiredWithAll[T]) Int32(value int32) RequiredWithAll[T] {
	r.required = r.required && value != 0
	return r
}

func (r RequiredWithAll[T]) Int64(value int64) RequiredWithAll[T] {
	r.required = r.required && value != 0
	return r
}

func (r RequiredWithAll[T]) Uint(value int) RequiredWithAll[T] {
	r.required = r.required && value != 0
	return r
}

func (r RequiredWithAll[T]) Uint8(value uint8) RequiredWithAll[T] {
	r.required = r.required && value != 0
	return r
}

func (r RequiredWithAll[T]) Uint16(value uint16) RequiredWithAll[T] {
	r.required = r.required && value != 0
	return r
}

func (r RequiredWithAll[T]) Uint32(value uint32) RequiredWithAll[T] {
	r.required = r.required && value != 0
	return r
}

func (r RequiredWithAll[T]) Uint64(value uint64) RequiredWithAll[T] {
	r.required = r.required && value != 0
	return r
}

func (r RequiredWithAll[T]) Names(names string) yav.ValidateFunc[T] {
	return r.provideFunc(names, r.required)
}
