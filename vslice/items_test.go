package vslice

import (
	"errors"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/SladeThe/yav"
)

func TestItems(t *testing.T) {
	type args struct {
		validateFuncs []yav.ValidateFunc[int]
		name          string
		value         []int
	}

	type want struct {
		stop bool
		err  error
	}

	test := func(a args, w want) func(t *testing.T) {
		return func(t *testing.T) {
			stop, err := Items[[]int](a.validateFuncs...)(a.name, a.value)
			assert.Equalf(t, w.stop, stop, "invalid stop: want = %v, got = %v", w.stop, stop)
			assert.Truef(t, reflect.DeepEqual(w.err, err), "invalid error: want = %v, got = %v", w.err, err)
		}
	}

	reusedErr := yav.Errors{Validation: []yav.Error{yav.ErrRequired("v")}}

	tests := []struct {
		name string
		args args
		want want
	}{{
		name: "error without stop",
		args: args{
			validateFuncs: []yav.ValidateFunc[int]{
				func(name string, value int) (stop bool, err error) {
					return false, yav.Error{
						CheckName: "continues",
						ValueName: name,
						Value:     value,
					}
				},
				func(name string, value int) (stop bool, err error) {
					return true, yav.Error{
						CheckName: "stops",
						ValueName: name,
						Value:     value,
					}
				},
			},
			name:  "v",
			value: []int{1},
		},
		want: want{
			stop: false,
			err: yav.Errors{
				Validation: []yav.Error{{
					CheckName: "continues",
					ValueName: "v[0]",
					Value:     1,
				}, {
					CheckName: "stops",
					ValueName: "v[0]",
					Value:     1,
				}},
			},
		},
	}, {
		name: "reused aggregate",
		args: args{
			validateFuncs: []yav.ValidateFunc[int]{
				func(string, int) (bool, error) {
					return false, reusedErr
				},
			},
			name:  "v",
			value: []int{1, 2},
		},
		want: want{
			stop: false,
			err: yav.Errors{
				Validation: []yav.Error{
					yav.ErrRequired("v[0]"),
					yav.ErrRequired("v[1]"),
				},
			},
		},
	}}

	for _, tt := range tests {
		t.Run(tt.name, test(tt.args, tt.want))
	}
}

func TestWithIndex(t *testing.T) {
	type args struct {
		err error
	}

	type want struct {
		isNil bool
	}

	test := func(a args, w want) func(t *testing.T) {
		return func(t *testing.T) {
			convertFunc := withIndex("v", 1, a.err)
			assert.Equalf(t, w.isNil, convertFunc == nil, "invalid nil convertFunc: want = %v, got = %v", w.isNil, convertFunc == nil)
		}
	}

	tests := []struct {
		name string
		args args
		want want
	}{{
		name: "nil error",
		want: want{isNil: true},
	}, {
		name: "unknown error",
		args: args{err: errors.New("unknown")},
		want: want{isNil: true},
	}, {
		name: "empty aggregate",
		args: args{err: yav.Errors{}},
		want: want{isNil: true},
	}, {
		name: "unknown aggregate",
		args: args{err: yav.Errors{Unknown: []error{errors.New("unknown")}}},
		want: want{isNil: true},
	}, {
		name: "validation error",
		args: args{err: yav.ErrRequired("v")},
		want: want{isNil: false},
	}, {
		name: "validation aggregate",
		args: args{err: yav.Errors{Validation: []yav.Error{yav.ErrRequired("v")}}},
		want: want{isNil: false},
	}}

	for _, tt := range tests {
		t.Run(tt.name, test(tt.args, tt.want))
	}
}

func TestWithIndexYAV(t *testing.T) {
	type args struct {
		name  string
		index int
		err   yav.Error
	}

	type want struct {
		err yav.Error
	}

	test := func(a args, w want) func(t *testing.T) {
		return func(t *testing.T) {
			err := withIndexYAV(a.name, a.index)(a.err)
			assert.Truef(t, reflect.DeepEqual(w.err, err), "invalid error: want = %v, got = %v", w.err, err)
		}
	}

	tests := []struct {
		name string
		args args
		want want
	}{{
		name: "exact name",
		args: args{
			name:  "items",
			index: 0,
			err:   yav.ErrRequired("items"),
		},
		want: want{
			err: yav.ErrRequired("items[0]"),
		},
	}, {
		name: "child field",
		args: args{
			name:  "items",
			index: 0,
			err:   yav.ErrRequired("items.child"),
		},
		want: want{
			err: yav.ErrRequired("items[0].child"),
		},
	}, {
		name: "nested index",
		args: args{
			name:  "items",
			index: 0,
			err:   yav.ErrRequired("items[1]"),
		},
		want: want{
			err: yav.ErrRequired("items[0][1]"),
		},
	}, {
		name: "unrelated prefix",
		args: args{
			name:  "a",
			index: 0,
			err:   yav.ErrRequired("age"),
		},
		want: want{
			err: yav.ErrRequired("age"),
		},
	}, {
		name: "unrelated name",
		args: args{
			name:  "items",
			index: 0,
			err:   yav.ErrRequired("other"),
		},
		want: want{
			err: yav.ErrRequired("other"),
		},
	}, {
		name: "unnamed collection field",
		args: args{
			name:  "",
			index: 0,
			err:   yav.ErrRequired("child"),
		},
		want: want{
			err: yav.ErrRequired("[0].child"),
		},
	}, {
		name: "unnamed collection index",
		args: args{
			name:  "",
			index: 0,
			err:   yav.ErrRequired("[1]"),
		},
		want: want{
			err: yav.ErrRequired("[0][1]"),
		},
	}, {
		name: "unnamed collection value",
		args: args{
			name:  "",
			index: 0,
			err:   yav.ErrRequired(""),
		},
		want: want{
			err: yav.ErrRequired("[0]"),
		},
	}}

	for _, tt := range tests {
		t.Run(tt.name, test(tt.args, tt.want))
	}
}
