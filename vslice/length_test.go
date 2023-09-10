package vslice

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/SladeThe/yav"
)

func TestMin(t *testing.T) {
	type args struct {
		parameter int
		name      string
		value     []int
	}

	type want struct {
		stop bool
		err  error
	}

	test := func(a args, w want) func(t *testing.T) {
		return func(t *testing.T) {
			stop, err := Min[[]int](a.parameter)(a.name, a.value)
			assert.Equalf(t, w.stop, stop, "invalid stop: want = %v, got = %v", w.stop, stop)
			assert.Truef(t, reflect.DeepEqual(w.err, err), "invalid error: want = %v, got = %v", w.err, err)
		}
	}

	tests := []struct {
		name string
		args args
		want want
	}{{
		name: "valid",
		args: args{
			parameter: 2,
			name:      "",
			value:     []int{1, 2},
		},
		want: want{
			stop: false,
			err:  nil,
		},
	}, {
		name: "invalid",
		args: args{
			parameter: 2,
			name:      "v",
			value:     []int{1},
		},
		want: want{
			stop: true,
			err: yav.Error{
				CheckName: "min",
				Parameter: "2",
				ValueName: "v",
				Value:     []int{1},
			},
		},
	}}

	for _, tt := range tests {
		t.Run(tt.name, test(tt.args, tt.want))
	}
}

func TestMax(t *testing.T) {
	type args struct {
		parameter int
		name      string
		value     []int
	}

	type want struct {
		stop bool
		err  error
	}

	test := func(a args, w want) func(t *testing.T) {
		return func(t *testing.T) {
			stop, err := Max[[]int](a.parameter)(a.name, a.value)
			assert.Equalf(t, w.stop, stop, "invalid stop: want = %v, got = %v", w.stop, stop)
			assert.Truef(t, reflect.DeepEqual(w.err, err), "invalid error: want = %v, got = %v", w.err, err)
		}
	}

	tests := []struct {
		name string
		args args
		want want
	}{{
		name: "valid",
		args: args{
			parameter: 2,
			name:      "",
			value:     []int{1, 2},
		},
		want: want{
			stop: false,
			err:  nil,
		},
	}, {
		name: "invalid",
		args: args{
			parameter: 2,
			name:      "v",
			value:     []int{1, 2, 3},
		},
		want: want{
			stop: true,
			err: yav.Error{
				CheckName: "max",
				Parameter: "2",
				ValueName: "v",
				Value:     []int{1, 2, 3},
			},
		},
	}}

	for _, tt := range tests {
		t.Run(tt.name, test(tt.args, tt.want))
	}
}

func TestBetween(t *testing.T) {
	type args struct {
		min, max int
		name     string
		value    []int
	}

	type want struct {
		stop bool
		err  error
	}

	test := func(a args, w want) func(t *testing.T) {
		return func(t *testing.T) {
			stop, err := Between[[]int](a.min, a.max)(a.name, a.value)
			assert.Equalf(t, w.stop, stop, "invalid stop: want = %v, got = %v", w.stop, stop)
			assert.Truef(t, reflect.DeepEqual(w.err, err), "invalid error: want = %v, got = %v", w.err, err)
		}
	}

	tests := []struct {
		name string
		args args
		want want
	}{{
		name: "valid",
		args: args{
			min:   1,
			max:   2,
			name:  "",
			value: []int{1, 2},
		},
		want: want{
			stop: false,
			err:  nil,
		},
	}, {
		name: "invalid lower",
		args: args{
			min:   1,
			max:   2,
			name:  "v",
			value: nil,
		},
		want: want{
			stop: true,
			err: yav.Error{
				CheckName: "min",
				Parameter: "1",
				ValueName: "v",
				Value:     []int(nil),
			},
		},
	}, {
		name: "invalid higher",
		args: args{
			min:   1,
			max:   2,
			name:  "v",
			value: []int{1, 2, 3},
		},
		want: want{
			stop: true,
			err: yav.Error{
				CheckName: "max",
				Parameter: "2",
				ValueName: "v",
				Value:     []int{1, 2, 3},
			},
		},
	}, {
		name: "invalid with range shift",
		args: args{
			min:   3,
			max:   3,
			name:  "v",
			value: []int{1, 2},
		},
		want: want{
			stop: true,
			err: yav.Error{
				CheckName: "min",
				Parameter: "3",
				ValueName: "v",
				Value:     []int{1, 2},
			},
		},
	}}

	for _, tt := range tests {
		t.Run(tt.name, test(tt.args, tt.want))
	}
}
