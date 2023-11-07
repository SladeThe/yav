package vpointer

import (
	"reflect"
	"runtime"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/SladeThe/yav"
)

func TestExcludedIf(t *testing.T) {
	type args struct {
		condition       bool
		conditionString string
		name            string
		value           *time.Time
	}

	type want struct {
		stop bool
		err  error
	}

	test := func(a args, w want) func(t *testing.T) {
		return func(t *testing.T) {
			stop, err := ExcludedIf[time.Time](a.conditionString, a.condition)(a.name, a.value)
			assert.Equalf(t, w.stop, stop, "invalid stop: want = %v, got = %v", w.stop, stop)
			assert.Truef(t, reflect.DeepEqual(w.err, err), "invalid error: want = %v, got = %v", w.err, err)
		}
	}

	tests := []struct {
		name string
		args args
		want want
	}{{
		name: "not nil excluded",
		args: args{
			condition:       true,
			conditionString: "a == a",
			name:            "v",
			value:           new(time.Time),
		},
		want: want{
			stop: true,
			err: yav.Error{
				CheckName: yav.CheckNameExcludedIf,
				Parameter: "a == a",
				ValueName: "v",
			},
		},
	}, {
		name: "not nil not excluded",
		args: args{
			condition:       false,
			conditionString: "a == b",
			name:            "v",
			value:           new(time.Time),
		},
		want: want{
			stop: false,
			err:  nil,
		},
	}, {
		name: "nil",
		args: args{
			condition:       true,
			conditionString: "a == a",
			name:            "v",
			value:           nil,
		},
		want: want{
			stop: false,
			err:  nil,
		},
	}}

	for _, tt := range tests {
		t.Run(tt.name, test(tt.args, tt.want))
	}
}

func TestExcludedUnless(t *testing.T) {
	type args struct {
		condition       bool
		conditionString string
		name            string
		value           *time.Time
	}

	type want struct {
		stop bool
		err  error
	}

	test := func(a args, w want) func(t *testing.T) {
		return func(t *testing.T) {
			stop, err := ExcludedUnless[time.Time](a.conditionString, a.condition)(a.name, a.value)
			assert.Equalf(t, w.stop, stop, "invalid stop: want = %v, got = %v", w.stop, stop)
			assert.Truef(t, reflect.DeepEqual(w.err, err), "invalid error: want = %v, got = %v", w.err, err)
		}
	}

	tests := []struct {
		name string
		args args
		want want
	}{{
		name: "not nil excluded",
		args: args{
			condition:       false,
			conditionString: "a == b",
			name:            "v",
			value:           new(time.Time),
		},
		want: want{
			stop: true,
			err: yav.Error{
				CheckName: yav.CheckNameExcludedUnless,
				Parameter: "a == b",
				ValueName: "v",
			},
		},
	}, {
		name: "not nil not excluded",
		args: args{
			condition:       true,
			conditionString: "a == a",
			name:            "v",
			value:           new(time.Time),
		},
		want: want{
			stop: false,
			err:  nil,
		},
	}, {
		name: "nil",
		args: args{
			condition:       false,
			conditionString: "a == b",
			name:            "v",
			value:           nil,
		},
		want: want{
			stop: false,
			err:  nil,
		},
	}}

	for _, tt := range tests {
		t.Run(tt.name, test(tt.args, tt.want))
	}
}

func TestExcludedWithAny(t *testing.T) {
	type args struct {
		parameters []int
		name       string
		value      *time.Time
	}

	type want struct {
		stop bool
		err  error
	}

	test := func(a args, w want) func(t *testing.T) {
		return func(t *testing.T) {
			accumulator := ExcludedWithAny[time.Time]()

			for _, parameter := range a.parameters {
				accumulator = accumulator.Int(parameter)
			}

			stop, err := accumulator.Names("pp")(a.name, a.value)
			assert.Equalf(t, w.stop, stop, "invalid stop: want = %v, got = %v", w.stop, stop)
			assert.Truef(t, reflect.DeepEqual(w.err, err), "invalid error: want = %v, got = %v", w.err, err)
		}
	}

	tests := []struct {
		name string
		args args
		want want
	}{{
		name: "not nil excluded",
		args: args{
			parameters: []int{-1, 0},
			name:       "v",
			value:      new(time.Time),
		},
		want: want{
			stop: true,
			err: yav.Error{
				CheckName: yav.CheckNameExcludedWithAny,
				Parameter: "pp",
				ValueName: "v",
			},
		},
	}, {
		name: "not nil not excluded",
		args: args{
			parameters: []int{0, 0},
			name:       "v",
			value:      new(time.Time),
		},
		want: want{
			stop: false,
			err:  nil,
		},
	}, {
		name: "nil",
		args: args{
			parameters: []int{1},
			name:       "v",
			value:      nil,
		},
		want: want{
			stop: false,
			err:  nil,
		},
	}}

	for _, tt := range tests {
		t.Run(tt.name, test(tt.args, tt.want))
	}
}

func TestExcludedWithoutAny(t *testing.T) {
	type args struct {
		parameters []int
		name       string
		value      *time.Time
	}

	type want struct {
		stop bool
		err  error
	}

	test := func(a args, w want) func(t *testing.T) {
		return func(t *testing.T) {
			accumulator := ExcludedWithoutAny[time.Time]()

			for _, parameter := range a.parameters {
				accumulator = accumulator.Int(parameter)
			}

			stop, err := accumulator.Names("pp")(a.name, a.value)
			assert.Equalf(t, w.stop, stop, "invalid stop: want = %v, got = %v", w.stop, stop)
			assert.Truef(t, reflect.DeepEqual(w.err, err), "invalid error: want = %v, got = %v", w.err, err)
		}
	}

	tests := []struct {
		name string
		args args
		want want
	}{{
		name: "not nil excluded",
		args: args{
			parameters: []int{-1, 0},
			name:       "v",
			value:      new(time.Time),
		},
		want: want{
			stop: true,
			err: yav.Error{
				CheckName: yav.CheckNameExcludedWithoutAny,
				Parameter: "pp",
				ValueName: "v",
			},
		},
	}, {
		name: "not nil not excluded",
		args: args{
			parameters: []int{-1, 1},
			name:       "v",
			value:      new(time.Time),
		},
		want: want{
			stop: false,
			err:  nil,
		},
	}, {
		name: "nil",
		args: args{
			parameters: []int{0},
			name:       "v",
			value:      nil,
		},
		want: want{
			stop: false,
			err:  nil,
		},
	}}

	for _, tt := range tests {
		t.Run(tt.name, test(tt.args, tt.want))
	}
}

func TestExcludedWithAll(t *testing.T) {
	type args struct {
		parameters []int
		name       string
		value      *time.Time
	}

	type want struct {
		stop bool
		err  error
	}

	test := func(a args, w want) func(t *testing.T) {
		return func(t *testing.T) {
			accumulator := ExcludedWithAll[time.Time]()

			for _, parameter := range a.parameters {
				accumulator = accumulator.Int(parameter)
			}

			stop, err := accumulator.Names("pp")(a.name, a.value)
			assert.Equalf(t, w.stop, stop, "invalid stop: want = %v, got = %v", w.stop, stop)
			assert.Truef(t, reflect.DeepEqual(w.err, err), "invalid error: want = %v, got = %v", w.err, err)
		}
	}

	tests := []struct {
		name string
		args args
		want want
	}{{
		name: "not nil excluded",
		args: args{
			parameters: []int{-1, 1},
			name:       "v",
			value:      new(time.Time),
		},
		want: want{
			stop: true,
			err: yav.Error{
				CheckName: yav.CheckNameExcludedWithAll,
				Parameter: "pp",
				ValueName: "v",
			},
		},
	}, {
		name: "not nil not excluded",
		args: args{
			parameters: []int{-1, 0},
			name:       "v",
			value:      new(time.Time),
		},
		want: want{
			stop: false,
			err:  nil,
		},
	}, {
		name: "nil",
		args: args{
			parameters: []int{1},
			name:       "v",
			value:      nil,
		},
		want: want{
			stop: false,
			err:  nil,
		},
	}}

	for _, tt := range tests {
		t.Run(tt.name, test(tt.args, tt.want))
	}
}

func TestExcludedWithoutAll(t *testing.T) {
	type args struct {
		parameters []int
		name       string
		value      *time.Time
	}

	type want struct {
		stop bool
		err  error
	}

	test := func(a args, w want) func(t *testing.T) {
		return func(t *testing.T) {
			accumulator := ExcludedWithoutAll[time.Time]()

			for _, parameter := range a.parameters {
				accumulator = accumulator.Int(parameter)
			}

			stop, err := accumulator.Names("pp")(a.name, a.value)
			assert.Equalf(t, w.stop, stop, "invalid stop: want = %v, got = %v", w.stop, stop)
			assert.Truef(t, reflect.DeepEqual(w.err, err), "invalid error: want = %v, got = %v", w.err, err)
		}
	}

	tests := []struct {
		name string
		args args
		want want
	}{{
		name: "not nil excluded",
		args: args{
			parameters: []int{0, 0},
			name:       "v",
			value:      new(time.Time),
		},
		want: want{
			stop: true,
			err: yav.Error{
				CheckName: yav.CheckNameExcludedWithoutAll,
				Parameter: "pp",
				ValueName: "v",
			},
		},
	}, {
		name: "not nil not excluded",
		args: args{
			parameters: []int{-1, 0},
			name:       "v",
			value:      new(time.Time),
		},
		want: want{
			stop: false,
			err:  nil,
		},
	}, {
		name: "nil",
		args: args{
			parameters: []int{0},
			name:       "v",
			value:      nil,
		},
		want: want{
			stop: false,
			err:  nil,
		},
	}}

	for _, tt := range tests {
		t.Run(tt.name, test(tt.args, tt.want))
	}
}

func BenchmarkExcludedAccumulators(b *testing.B) {
	var pointer *time.Time

	b.ReportAllocs()

	var err error

	for i := 0; i < b.N; i++ {
		err = yav.Join(
			yav.Chain(
				"pointer", pointer,
				ExcludedIf[time.Time]("", true),
				ExcludedUnless[time.Time]("", false),
				ExcludedWithAny[time.Time]().Bool(true).Names(""),
				ExcludedWithoutAny[time.Time]().Bool(false).Names(""),
				ExcludedWithAll[time.Time]().Bool(true).Names(""),
				ExcludedWithoutAll[time.Time]().Bool(false).Names(""),
			),
		)
	}

	runtime.KeepAlive(err)
}

func BenchmarkExcludedAccumulatorsParallel(b *testing.B) {
	var pointer *time.Time

	b.RunParallel(func(pb *testing.PB) {
		var err error

		for pb.Next() {
			err = yav.Join(
				yav.Chain(
					"pointer", pointer,
					ExcludedIf[time.Time]("", true),
					ExcludedUnless[time.Time]("", false),
					ExcludedWithAny[time.Time]().Bool(true).Names(""),
					ExcludedWithoutAny[time.Time]().Bool(false).Names(""),
					ExcludedWithAll[time.Time]().Bool(true).Names(""),
					ExcludedWithoutAll[time.Time]().Bool(false).Names(""),
				),
			)
		}

		runtime.KeepAlive(err)
	})
}
