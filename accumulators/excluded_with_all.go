package accumulators

import (
	"reflect"
	"time"

	"github.com/SladeThe/yav"
)

type ExcludedWithAll[T any] struct {
	provideFunc ProvideFunc[T]
	excluded    bool
}

func NewExcludedWithAll[T any](provideFunc ProvideFunc[T]) ExcludedWithAll[T] {
	return ExcludedWithAll[T]{
		provideFunc: provideFunc,
		excluded:    true,
	}
}

func (r ExcludedWithAll[T]) String(value string) ExcludedWithAll[T] {
	r.excluded = r.excluded && value != ""
	return r
}

func (r ExcludedWithAll[T]) Bytes(value []byte) ExcludedWithAll[T] {
	r.excluded = r.excluded && len(value) != 0
	return r
}

func (r ExcludedWithAll[T]) Slice(value any) ExcludedWithAll[T] {
	r.excluded = r.excluded && value != nil && reflect.ValueOf(value).Len() != 0
	return r
}

func (r ExcludedWithAll[T]) Map(value any) ExcludedWithAll[T] {
	r.excluded = r.excluded && value != nil && reflect.ValueOf(value).Len() != 0
	return r
}

func (r ExcludedWithAll[T]) Pointer(value any) ExcludedWithAll[T] {
	r.excluded = r.excluded && !reflect.ValueOf(value).IsNil()
	return r
}

func (r ExcludedWithAll[T]) Time(value time.Time) ExcludedWithAll[T] {
	r.excluded = r.excluded && !value.IsZero()
	return r
}

func (r ExcludedWithAll[T]) Zeroer(value yav.Zeroer) ExcludedWithAll[T] {
	r.excluded = r.excluded && !value.IsZero()
	return r
}

func (r ExcludedWithAll[T]) Bool(value bool) ExcludedWithAll[T] {
	r.excluded = r.excluded && value
	return r
}

func (r ExcludedWithAll[T]) Int(value int) ExcludedWithAll[T] {
	r.excluded = r.excluded && value != 0
	return r
}

func (r ExcludedWithAll[T]) Int8(value int8) ExcludedWithAll[T] {
	r.excluded = r.excluded && value != 0
	return r
}

func (r ExcludedWithAll[T]) Int16(value int16) ExcludedWithAll[T] {
	r.excluded = r.excluded && value != 0
	return r
}

func (r ExcludedWithAll[T]) Int32(value int32) ExcludedWithAll[T] {
	r.excluded = r.excluded && value != 0
	return r
}

func (r ExcludedWithAll[T]) Int64(value int64) ExcludedWithAll[T] {
	r.excluded = r.excluded && value != 0
	return r
}

func (r ExcludedWithAll[T]) Uint(value int) ExcludedWithAll[T] {
	r.excluded = r.excluded && value != 0
	return r
}

func (r ExcludedWithAll[T]) Uint8(value uint8) ExcludedWithAll[T] {
	r.excluded = r.excluded && value != 0
	return r
}

func (r ExcludedWithAll[T]) Uint16(value uint16) ExcludedWithAll[T] {
	r.excluded = r.excluded && value != 0
	return r
}

func (r ExcludedWithAll[T]) Uint32(value uint32) ExcludedWithAll[T] {
	r.excluded = r.excluded && value != 0
	return r
}

func (r ExcludedWithAll[T]) Uint64(value uint64) ExcludedWithAll[T] {
	r.excluded = r.excluded && value != 0
	return r
}

func (r ExcludedWithAll[T]) Names(names string) yav.ValidateFunc[T] {
	return r.provideFunc(names, r.excluded)
}
