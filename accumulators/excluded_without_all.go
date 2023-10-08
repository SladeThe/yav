package accumulators

import (
	"reflect"
	"time"

	"github.com/SladeThe/yav"
)

type ExcludedWithoutAll[T any] struct {
	provideFunc ProvideFunc[T]
	excluded    bool
}

func NewExcludedWithoutAll[T any](provideFunc ProvideFunc[T]) ExcludedWithoutAll[T] {
	return ExcludedWithoutAll[T]{
		provideFunc: provideFunc,
		excluded:    true,
	}
}

func (r ExcludedWithoutAll[T]) String(value string) ExcludedWithoutAll[T] {
	r.excluded = r.excluded && value == ""
	return r
}

func (r ExcludedWithoutAll[T]) Bytes(value []byte) ExcludedWithoutAll[T] {
	r.excluded = r.excluded && len(value) == 0
	return r
}

func (r ExcludedWithoutAll[T]) Slice(value any) ExcludedWithoutAll[T] {
	r.excluded = r.excluded && (value == nil || reflect.ValueOf(value).Len() == 0)
	return r
}

func (r ExcludedWithoutAll[T]) Map(value any) ExcludedWithoutAll[T] {
	r.excluded = r.excluded && (value == nil || reflect.ValueOf(value).Len() == 0)
	return r
}

func (r ExcludedWithoutAll[T]) Pointer(value any) ExcludedWithoutAll[T] {
	r.excluded = r.excluded && reflect.ValueOf(value).IsNil()
	return r
}

func (r ExcludedWithoutAll[T]) Time(value time.Time) ExcludedWithoutAll[T] {
	r.excluded = r.excluded && value.IsZero()
	return r
}

func (r ExcludedWithoutAll[T]) Zeroer(value yav.Zeroer) ExcludedWithoutAll[T] {
	r.excluded = r.excluded && value.IsZero()
	return r
}

func (r ExcludedWithoutAll[T]) Bool(value bool) ExcludedWithoutAll[T] {
	r.excluded = r.excluded && !value
	return r
}

func (r ExcludedWithoutAll[T]) Int(value int) ExcludedWithoutAll[T] {
	r.excluded = r.excluded && value == 0
	return r
}

func (r ExcludedWithoutAll[T]) Int8(value int8) ExcludedWithoutAll[T] {
	r.excluded = r.excluded && value == 0
	return r
}

func (r ExcludedWithoutAll[T]) Int16(value int16) ExcludedWithoutAll[T] {
	r.excluded = r.excluded && value == 0
	return r
}

func (r ExcludedWithoutAll[T]) Int32(value int32) ExcludedWithoutAll[T] {
	r.excluded = r.excluded && value == 0
	return r
}

func (r ExcludedWithoutAll[T]) Int64(value int64) ExcludedWithoutAll[T] {
	r.excluded = r.excluded && value == 0
	return r
}

func (r ExcludedWithoutAll[T]) Uint(value int) ExcludedWithoutAll[T] {
	r.excluded = r.excluded && value == 0
	return r
}

func (r ExcludedWithoutAll[T]) Uint8(value uint8) ExcludedWithoutAll[T] {
	r.excluded = r.excluded && value == 0
	return r
}

func (r ExcludedWithoutAll[T]) Uint16(value uint16) ExcludedWithoutAll[T] {
	r.excluded = r.excluded && value == 0
	return r
}

func (r ExcludedWithoutAll[T]) Uint32(value uint32) ExcludedWithoutAll[T] {
	r.excluded = r.excluded && value == 0
	return r
}

func (r ExcludedWithoutAll[T]) Uint64(value uint64) ExcludedWithoutAll[T] {
	r.excluded = r.excluded && value == 0
	return r
}

func (r ExcludedWithoutAll[T]) Names(names string) yav.ValidateFunc[T] {
	return r.provideFunc(names, r.excluded)
}
