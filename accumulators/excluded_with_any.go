package accumulators

import (
	"reflect"
	"time"

	"github.com/SladeThe/yav"
)

type ExcludedWithAny[T any] struct {
	provideFunc ProvideFunc[T]
	excluded    bool
}

func NewExcludedWithAny[T any](provideFunc ProvideFunc[T]) ExcludedWithAny[T] {
	return ExcludedWithAny[T]{
		provideFunc: provideFunc,
		excluded:    false,
	}
}

func (r ExcludedWithAny[T]) String(value string) ExcludedWithAny[T] {
	r.excluded = r.excluded || value != ""
	return r
}

func (r ExcludedWithAny[T]) Bytes(value []byte) ExcludedWithAny[T] {
	r.excluded = r.excluded || len(value) != 0
	return r
}

func (r ExcludedWithAny[T]) Slice(value any) ExcludedWithAny[T] {
	r.excluded = r.excluded || value != nil && reflect.ValueOf(value).Len() != 0
	return r
}

func (r ExcludedWithAny[T]) Map(value any) ExcludedWithAny[T] {
	r.excluded = r.excluded || value != nil && reflect.ValueOf(value).Len() != 0
	return r
}

func (r ExcludedWithAny[T]) Time(value time.Time) ExcludedWithAny[T] {
	r.excluded = r.excluded || !value.IsZero()
	return r
}

func (r ExcludedWithAny[T]) Zeroer(value yav.Zeroer) ExcludedWithAny[T] {
	r.excluded = r.excluded || !value.IsZero()
	return r
}

func (r ExcludedWithAny[T]) Bool(value bool) ExcludedWithAny[T] {
	r.excluded = r.excluded || value
	return r
}

func (r ExcludedWithAny[T]) Int(value int) ExcludedWithAny[T] {
	r.excluded = r.excluded || value != 0
	return r
}

func (r ExcludedWithAny[T]) Int8(value int8) ExcludedWithAny[T] {
	r.excluded = r.excluded || value != 0
	return r
}

func (r ExcludedWithAny[T]) Int16(value int16) ExcludedWithAny[T] {
	r.excluded = r.excluded || value != 0
	return r
}

func (r ExcludedWithAny[T]) Int32(value int32) ExcludedWithAny[T] {
	r.excluded = r.excluded || value != 0
	return r
}

func (r ExcludedWithAny[T]) Int64(value int64) ExcludedWithAny[T] {
	r.excluded = r.excluded || value != 0
	return r
}

func (r ExcludedWithAny[T]) Uint(value int) ExcludedWithAny[T] {
	r.excluded = r.excluded || value != 0
	return r
}

func (r ExcludedWithAny[T]) Uint8(value uint8) ExcludedWithAny[T] {
	r.excluded = r.excluded || value != 0
	return r
}

func (r ExcludedWithAny[T]) Uint16(value uint16) ExcludedWithAny[T] {
	r.excluded = r.excluded || value != 0
	return r
}

func (r ExcludedWithAny[T]) Uint32(value uint32) ExcludedWithAny[T] {
	r.excluded = r.excluded || value != 0
	return r
}

func (r ExcludedWithAny[T]) Uint64(value uint64) ExcludedWithAny[T] {
	r.excluded = r.excluded || value != 0
	return r
}

func (r ExcludedWithAny[T]) Names(names string) yav.ValidateFunc[T] {
	return r.provideFunc(names, r.excluded)
}
