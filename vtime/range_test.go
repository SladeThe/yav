package vtime

import (
	"reflect"
	"runtime"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"go.uber.org/multierr"

	"github.com/SladeThe/yav"
)

func TestMin(t *testing.T) {
	type args struct {
		parameter time.Time
		name      string
		value     time.Time
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

	now := time.Now()

	tests := []struct {
		name string
		args args
		want want
	}{{
		name: "valid",
		args: args{
			parameter: now,
			name:      "",
			value:     now,
		},
		want: want{
			stop: false,
			err:  nil,
		},
	}, {
		name: "invalid",
		args: args{
			parameter: now,
			name:      "t",
			value:     now.Add(-1),
		},
		want: want{
			stop: true,
			err: yav.Error{
				CheckName: "min",
				Parameter: now.String(),
				ValueName: "t",
				Value:     now.Add(-1),
			},
		},
	}}

	for _, tt := range tests {
		t.Run(tt.name, test(tt.args, tt.want))
	}
}

func TestMax(t *testing.T) {
	type args struct {
		parameter time.Time
		name      string
		value     time.Time
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

	now := time.Now()

	tests := []struct {
		name string
		args args
		want want
	}{{
		name: "valid",
		args: args{
			parameter: now,
			name:      "",
			value:     now,
		},
		want: want{
			stop: false,
			err:  nil,
		},
	}, {
		name: "invalid",
		args: args{
			parameter: now,
			name:      "t",
			value:     now.Add(1),
		},
		want: want{
			stop: true,
			err: yav.Error{
				CheckName: "max",
				Parameter: now.String(),
				ValueName: "t",
				Value:     now.Add(1),
			},
		},
	}}

	for _, tt := range tests {
		t.Run(tt.name, test(tt.args, tt.want))
	}
}

func TestBetween(t *testing.T) {
	type args struct {
		min, max time.Time
		name     string
		value    time.Time
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

	now := time.Now()

	tests := []struct {
		name string
		args args
		want want
	}{{
		name: "valid",
		args: args{
			min:   now,
			max:   now.Add(1),
			name:  "",
			value: now.Add(1),
		},
		want: want{
			stop: false,
			err:  nil,
		},
	}, {
		name: "invalid lower",
		args: args{
			min:   now,
			max:   now.Add(1),
			name:  "t",
			value: now.Add(-1),
		},
		want: want{
			stop: true,
			err: yav.Error{
				CheckName: "min",
				Parameter: now.String(),
				ValueName: "t",
				Value:     now.Add(-1),
			},
		},
	}, {
		name: "invalid higher",
		args: args{
			min:   now,
			max:   now.Add(1),
			name:  "t",
			value: now.Add(2),
		},
		want: want{
			stop: true,
			err: yav.Error{
				CheckName: "max",
				Parameter: now.Add(1).String(),
				ValueName: "t",
				Value:     now.Add(2),
			},
		},
	}}

	for _, tt := range tests {
		t.Run(tt.name, test(tt.args, tt.want))
	}
}

func TestLessThan(t *testing.T) {
	type args struct {
		parameter time.Time
		name      string
		value     time.Time
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

	now := time.Now()

	tests := []struct {
		name string
		args args
		want want
	}{{
		name: "valid",
		args: args{
			parameter: now,
			name:      "",
			value:     now.Add(-1),
		},
		want: want{
			stop: false,
			err:  nil,
		},
	}, {
		name: "invalid",
		args: args{
			parameter: now,
			name:      "t",
			value:     now,
		},
		want: want{
			stop: true,
			err: yav.Error{
				CheckName: "lt",
				Parameter: now.String(),
				ValueName: "t",
				Value:     now,
			},
		},
	}}

	for _, tt := range tests {
		t.Run(tt.name, test(tt.args, tt.want))
	}
}

func TestLessThanOrEqual(t *testing.T) {
	type args struct {
		parameter time.Time
		name      string
		value     time.Time
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

	now := time.Now()

	tests := []struct {
		name string
		args args
		want want
	}{{
		name: "valid",
		args: args{
			parameter: now,
			name:      "",
			value:     now,
		},
		want: want{
			stop: false,
			err:  nil,
		},
	}, {
		name: "invalid",
		args: args{
			parameter: now,
			name:      "t",
			value:     now.Add(1),
		},
		want: want{
			stop: true,
			err: yav.Error{
				CheckName: "lte",
				Parameter: now.String(),
				ValueName: "t",
				Value:     now.Add(1),
			},
		},
	}}

	for _, tt := range tests {
		t.Run(tt.name, test(tt.args, tt.want))
	}
}

func TestGreaterThan(t *testing.T) {
	type args struct {
		parameter time.Time
		name      string
		value     time.Time
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

	now := time.Now()

	tests := []struct {
		name string
		args args
		want want
	}{{
		name: "valid",
		args: args{
			parameter: now,
			name:      "",
			value:     now.Add(1),
		},
		want: want{
			stop: false,
			err:  nil,
		},
	}, {
		name: "invalid",
		args: args{
			parameter: now,
			name:      "t",
			value:     now,
		},
		want: want{
			stop: true,
			err: yav.Error{
				CheckName: "gt",
				Parameter: now.String(),
				ValueName: "t",
				Value:     now,
			},
		},
	}}

	for _, tt := range tests {
		t.Run(tt.name, test(tt.args, tt.want))
	}
}

func TestGreaterThanOrEqual(t *testing.T) {
	type args struct {
		parameter time.Time
		name      string
		value     time.Time
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

	now := time.Now()

	tests := []struct {
		name string
		args args
		want want
	}{{
		name: "valid",
		args: args{
			parameter: now,
			name:      "",
			value:     now,
		},
		want: want{
			stop: false,
			err:  nil,
		},
	}, {
		name: "invalid",
		args: args{
			parameter: now,
			name:      "t",
			value:     now.Add(-1),
		},
		want: want{
			stop: true,
			err: yav.Error{
				CheckName: "gte",
				Parameter: now.String(),
				ValueName: "t",
				Value:     now.Add(-1),
			},
		},
	}}

	for _, tt := range tests {
		t.Run(tt.name, test(tt.args, tt.want))
	}
}

func BenchmarkRange(b *testing.B) {
	now := time.Now()

	b.ReportAllocs()

	var err error

	for i := 0; i < b.N; i++ {
		err = multierr.Combine(
			yav.Chain(
				"now", now,
				Min(now),
				Max(now),
				Between(now, now),
				LessThan(now.Add(1)),
				LessThanOrEqual(now),
				GreaterThan(now.Add(-1)),
				GreaterThanOrEqual(now),
			),
		)
	}

	runtime.KeepAlive(err)
}
