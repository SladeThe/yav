package accumulators

import (
	"reflect"
	"time"

	"github.com/SladeThe/yav"
)

type ExcludedWithoutAny[T any] struct {
	provideFunc ProvideFunc[T]
	excluded    bool
}

func NewExcludedWithoutAny[T any](provideFunc ProvideFunc[T]) ExcludedWithoutAny[T] {
	return ExcludedWithoutAny[T]{
		provideFunc: provideFunc,
		excluded:    false,
	}
}

func (r ExcludedWithoutAny[T]) String(value string) ExcludedWithoutAny[T] {
	r.excluded = r.excluded || value == ""
	return r
}

func (r ExcludedWithoutAny[T]) Bytes(value []byte) ExcludedWithoutAny[T] {
	r.excluded = r.excluded || len(value) == 0
	return r
}

func (r ExcludedWithoutAny[T]) Slice(value any) ExcludedWithoutAny[T] {
	r.excluded = r.excluded || value == nil || reflect.ValueOf(value).Len() == 0
	return r
}

func (r ExcludedWithoutAny[T]) Map(value any) ExcludedWithoutAny[T] {
	r.excluded = r.excluded || value == nil || reflect.ValueOf(value).Len() == 0
	return r
}

func (r ExcludedWithoutAny[T]) Pointer(value any) ExcludedWithoutAny[T] {
	r.excluded = r.excluded || reflect.ValueOf(value).IsNil()
	return r
}

func (r ExcludedWithoutAny[T]) Time(value time.Time) ExcludedWithoutAny[T] {
	r.excluded = r.excluded || value.IsZero()
	return r
}

func (r ExcludedWithoutAny[T]) Zeroer(value yav.Zeroer) ExcludedWithoutAny[T] {
	r.excluded = r.excluded || value.IsZero()
	return r
}

func (r ExcludedWithoutAny[T]) Bool(value bool) ExcludedWithoutAny[T] {
	r.excluded = r.excluded || !value
	return r
}

func (r ExcludedWithoutAny[T]) Int(value int) ExcludedWithoutAny[T] {
	r.excluded = r.excluded || value == 0
	return r
}

func (r ExcludedWithoutAny[T]) Int8(value int8) ExcludedWithoutAny[T] {
	r.excluded = r.excluded || value == 0
	return r
}

func (r ExcludedWithoutAny[T]) Int16(value int16) ExcludedWithoutAny[T] {
	r.excluded = r.excluded || value == 0
	return r
}

func (r ExcludedWithoutAny[T]) Int32(value int32) ExcludedWithoutAny[T] {
	r.excluded = r.excluded || value == 0
	return r
}

func (r ExcludedWithoutAny[T]) Int64(value int64) ExcludedWithoutAny[T] {
	r.excluded = r.excluded || value == 0
	return r
}

func (r ExcludedWithoutAny[T]) Uint(value int) ExcludedWithoutAny[T] {
	r.excluded = r.excluded || value == 0
	return r
}

func (r ExcludedWithoutAny[T]) Uint8(value uint8) ExcludedWithoutAny[T] {
	r.excluded = r.excluded || value == 0
	return r
}

func (r ExcludedWithoutAny[T]) Uint16(value uint16) ExcludedWithoutAny[T] {
	r.excluded = r.excluded || value == 0
	return r
}

func (r ExcludedWithoutAny[T]) Uint32(value uint32) ExcludedWithoutAny[T] {
	r.excluded = r.excluded || value == 0
	return r
}

func (r ExcludedWithoutAny[T]) Uint64(value uint64) ExcludedWithoutAny[T] {
	r.excluded = r.excluded || value == 0
	return r
}

func (r ExcludedWithoutAny[T]) Names(names string) yav.ValidateFunc[T] {
	return r.provideFunc(names, r.excluded)
}
