package vmap

import (
	"errors"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/SladeThe/yav"
)

func TestKeys(t *testing.T) {
	type args struct {
		validateFuncs []yav.ValidateFunc[int]
		name          string
		value         map[int]int
	}

	type want struct {
		stop bool
		err  error
	}

	test := func(a args, w want) func(t *testing.T) {
		return func(t *testing.T) {
			stop, err := Keys[map[int]int](a.validateFuncs...)(a.name, a.value)
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
			value: map[int]int{1: 2},
		},
		want: want{
			stop: false,
			err: yav.Errors{
				Validation: []yav.Error{{
					CheckName: "continues",
					ValueName: "v[1]",
					Value:     1,
				}, {
					CheckName: "stops",
					ValueName: "v[1]",
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
				func(string, int) (bool, error) {
					return true, reusedErr
				},
			},
			name:  "v",
			value: map[int]int{1: 2},
		},
		want: want{
			stop: false,
			err: yav.Errors{
				Validation: []yav.Error{
					yav.ErrRequired("v[1]"),
					yav.ErrRequired("v[1]"),
				},
			},
		},
	}}

	for _, tt := range tests {
		t.Run(tt.name, test(tt.args, tt.want))
	}
}

func TestValues(t *testing.T) {
	type args struct {
		validateFuncs []yav.ValidateFunc[int]
		name          string
		value         map[int]int
	}

	type want struct {
		stop bool
		err  error
	}

	test := func(a args, w want) func(t *testing.T) {
		return func(t *testing.T) {
			stop, err := Values[map[int]int](a.validateFuncs...)(a.name, a.value)
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
			value: map[int]int{1: 2},
		},
		want: want{
			stop: false,
			err: yav.Errors{
				Validation: []yav.Error{{
					CheckName: "continues",
					ValueName: "v[1]",
					Value:     2,
				}, {
					CheckName: "stops",
					ValueName: "v[1]",
					Value:     2,
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
				func(string, int) (bool, error) {
					return true, reusedErr
				},
			},
			name:  "v",
			value: map[int]int{1: 2},
		},
		want: want{
			stop: false,
			err: yav.Errors{
				Validation: []yav.Error{
					yav.ErrRequired("v[1]"),
					yav.ErrRequired("v[1]"),
				},
			},
		},
	}}

	for _, tt := range tests {
		t.Run(tt.name, test(tt.args, tt.want))
	}
}

func TestEntries(t *testing.T) {
	type args struct {
		validateFuncs []yav.ValidateFunc[Entry[int, int]]
		name          string
		value         map[int]int
	}

	type want struct {
		stop bool
		err  error
	}

	test := func(a args, w want) func(t *testing.T) {
		return func(t *testing.T) {
			stop, err := Entries[map[int]int](a.validateFuncs...)(a.name, a.value)
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
			validateFuncs: []yav.ValidateFunc[Entry[int, int]]{
				func(name string, value Entry[int, int]) (stop bool, err error) {
					return false, yav.Error{
						CheckName: "continues",
						ValueName: name,
						Value:     value,
					}
				},
				func(name string, value Entry[int, int]) (stop bool, err error) {
					return true, yav.Error{
						CheckName: "stops",
						ValueName: name,
						Value:     value,
					}
				},
			},
			name:  "v",
			value: map[int]int{1: 2},
		},
		want: want{
			stop: false,
			err: yav.Errors{
				Validation: []yav.Error{{
					CheckName: "continues",
					ValueName: "v[1]",
					Value:     Entry[int, int]{Key: 1, Value: 2},
				}, {
					CheckName: "stops",
					ValueName: "v[1]",
					Value:     Entry[int, int]{Key: 1, Value: 2},
				}},
			},
		},
	}, {
		name: "reused aggregate",
		args: args{
			validateFuncs: []yav.ValidateFunc[Entry[int, int]]{
				func(string, Entry[int, int]) (bool, error) {
					return false, reusedErr
				},
				func(string, Entry[int, int]) (bool, error) {
					return true, reusedErr
				},
			},
			name:  "v",
			value: map[int]int{1: 2},
		},
		want: want{
			stop: false,
			err: yav.Errors{
				Validation: []yav.Error{
					yav.ErrRequired("v[1]"),
					yav.ErrRequired("v[1]"),
				},
			},
		},
	}}

	for _, tt := range tests {
		t.Run(tt.name, test(tt.args, tt.want))
	}
}

func TestWithKeyName(t *testing.T) {
	type args struct {
		err error
	}

	type want struct {
		isNil bool
	}

	test := func(a args, w want) func(t *testing.T) {
		return func(t *testing.T) {
			convertFunc := withKeyName("v", 1, a.err)
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
