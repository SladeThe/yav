package vduration

import (
	"reflect"
	"runtime"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/SladeThe/yav"
)

func TestMin(t *testing.T) {
	type args struct {
		parameter time.Duration
		name      string
		value     time.Duration
	}

	type want struct {
		stop bool
		err  error
	}

	test := func(a args, w want) func(t *testing.T) {
		return func(t *testing.T) {
			stop, err := Min(a.parameter)(a.name, a.value)
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
			parameter: time.Second,
			name:      "",
			value:     time.Second,
		},
		want: want{
			stop: false,
			err:  nil,
		},
	}, {
		name: "invalid",
		args: args{
			parameter: time.Second,
			name:      "t",
			value:     time.Second - 1,
		},
		want: want{
			stop: true,
			err: yav.Error{
				CheckName: "min",
				Parameter: time.Second.String(),
				ValueName: "t",
				Value:     time.Second - 1,
			},
		},
	}}

	for _, tt := range tests {
		t.Run(tt.name, test(tt.args, tt.want))
	}
}

func TestMax(t *testing.T) {
	type args struct {
		parameter time.Duration
		name      string
		value     time.Duration
	}

	type want struct {
		stop bool
		err  error
	}

	test := func(a args, w want) func(t *testing.T) {
		return func(t *testing.T) {
			stop, err := Max(a.parameter)(a.name, a.value)
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
			parameter: time.Second,
			name:      "",
			value:     time.Second,
		},
		want: want{
			stop: false,
			err:  nil,
		},
	}, {
		name: "invalid",
		args: args{
			parameter: time.Second,
			name:      "t",
			value:     time.Second + 1,
		},
		want: want{
			stop: true,
			err: yav.Error{
				CheckName: "max",
				Parameter: time.Second.String(),
				ValueName: "t",
				Value:     time.Second + 1,
			},
		},
	}}

	for _, tt := range tests {
		t.Run(tt.name, test(tt.args, tt.want))
	}
}

func TestBetween(t *testing.T) {
	type args struct {
		min, max time.Duration
		name     string
		value    time.Duration
	}

	type want struct {
		stop bool
		err  error
	}

	test := func(a args, w want) func(t *testing.T) {
		return func(t *testing.T) {
			stop, err := Between(a.min, a.max)(a.name, a.value)
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
			min:   time.Second,
			max:   time.Second + 1,
			name:  "",
			value: time.Second + 1,
		},
		want: want{
			stop: false,
			err:  nil,
		},
	}, {
		name: "invalid lower",
		args: args{
			min:   time.Second,
			max:   time.Second + 1,
			name:  "t",
			value: time.Second - 1,
		},
		want: want{
			stop: true,
			err: yav.Error{
				CheckName: "min",
				Parameter: time.Second.String(),
				ValueName: "t",
				Value:     time.Second - 1,
			},
		},
	}, {
		name: "invalid higher",
		args: args{
			min:   time.Second,
			max:   time.Second + 1,
			name:  "t",
			value: time.Second + 2,
		},
		want: want{
			stop: true,
			err: yav.Error{
				CheckName: "max",
				Parameter: (time.Second + 1).String(),
				ValueName: "t",
				Value:     time.Second + 2,
			},
		},
	}, {
		name: "invalid with range shift",
		args: args{
			min:   time.Second + 3,
			max:   time.Second + 3,
			name:  "t",
			value: time.Second + 1,
		},
		want: want{
			stop: true,
			err: yav.Error{
				CheckName: "min",
				Parameter: (time.Second + 3).String(),
				ValueName: "t",
				Value:     time.Second + 1,
			},
		},
	}}

	for _, tt := range tests {
		t.Run(tt.name, test(tt.args, tt.want))
	}
}

func TestLessThan(t *testing.T) {
	type args struct {
		parameter time.Duration
		name      string
		value     time.Duration
	}

	type want struct {
		stop bool
		err  error
	}

	test := func(a args, w want) func(t *testing.T) {
		return func(t *testing.T) {
			stop, err := LessThan(a.parameter)(a.name, a.value)
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
			parameter: time.Second,
			name:      "",
			value:     time.Second - 1,
		},
		want: want{
			stop: false,
			err:  nil,
		},
	}, {
		name: "invalid",
		args: args{
			parameter: time.Second,
			name:      "t",
			value:     time.Second,
		},
		want: want{
			stop: true,
			err: yav.Error{
				CheckName: "lt",
				Parameter: time.Second.String(),
				ValueName: "t",
				Value:     time.Second,
			},
		},
	}}

	for _, tt := range tests {
		t.Run(tt.name, test(tt.args, tt.want))
	}
}

func TestLessThanOrEqual(t *testing.T) {
	type args struct {
		parameter time.Duration
		name      string
		value     time.Duration
	}

	type want struct {
		stop bool
		err  error
	}

	test := func(a args, w want) func(t *testing.T) {
		return func(t *testing.T) {
			stop, err := LessThanOrEqual(a.parameter)(a.name, a.value)
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
			parameter: time.Second,
			name:      "",
			value:     time.Second,
		},
		want: want{
			stop: false,
			err:  nil,
		},
	}, {
		name: "invalid",
		args: args{
			parameter: time.Second,
			name:      "t",
			value:     time.Second + 1,
		},
		want: want{
			stop: true,
			err: yav.Error{
				CheckName: "lte",
				Parameter: time.Second.String(),
				ValueName: "t",
				Value:     time.Second + 1,
			},
		},
	}}

	for _, tt := range tests {
		t.Run(tt.name, test(tt.args, tt.want))
	}
}

func TestGreaterThan(t *testing.T) {
	type args struct {
		parameter time.Duration
		name      string
		value     time.Duration
	}

	type want struct {
		stop bool
		err  error
	}

	test := func(a args, w want) func(t *testing.T) {
		return func(t *testing.T) {
			stop, err := GreaterThan(a.parameter)(a.name, a.value)
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
			parameter: time.Second,
			name:      "",
			value:     time.Second + 1,
		},
		want: want{
			stop: false,
			err:  nil,
		},
	}, {
		name: "invalid",
		args: args{
			parameter: time.Second,
			name:      "t",
			value:     time.Second,
		},
		want: want{
			stop: true,
			err: yav.Error{
				CheckName: "gt",
				Parameter: time.Second.String(),
				ValueName: "t",
				Value:     time.Second,
			},
		},
	}}

	for _, tt := range tests {
		t.Run(tt.name, test(tt.args, tt.want))
	}
}

func TestGreaterThanOrEqual(t *testing.T) {
	type args struct {
		parameter time.Duration
		name      string
		value     time.Duration
	}

	type want struct {
		stop bool
		err  error
	}

	test := func(a args, w want) func(t *testing.T) {
		return func(t *testing.T) {
			stop, err := GreaterThanOrEqual(a.parameter)(a.name, a.value)
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
			parameter: time.Second,
			name:      "",
			value:     time.Second,
		},
		want: want{
			stop: false,
			err:  nil,
		},
	}, {
		name: "invalid",
		args: args{
			parameter: time.Second,
			name:      "t",
			value:     time.Second - 1,
		},
		want: want{
			stop: true,
			err: yav.Error{
				CheckName: "gte",
				Parameter: time.Second.String(),
				ValueName: "t",
				Value:     time.Second - 1,
			},
		},
	}}

	for _, tt := range tests {
		t.Run(tt.name, test(tt.args, tt.want))
	}
}

func BenchmarkRange(b *testing.B) {
	b.ReportAllocs()

	var err error

	for i := 0; i < b.N; i++ {
		err = yav.Join(
			yav.Chain(
				"second", time.Second,
				Min(time.Second),
				Max(time.Second),
				Between(time.Second, time.Second),
				LessThan(time.Second+1),
				LessThanOrEqual(time.Second),
				GreaterThan(time.Second-1),
				GreaterThanOrEqual(time.Second),
			),
		)
	}

	runtime.KeepAlive(err)
}
