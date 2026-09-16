package vnumber

import (
	"math"
	"testing"

	"github.com/cheekybits/genny/generic"
	"github.com/stretchr/testify/assert"

	"github.com/SladeThe/yav"
)

type Element generic.Type

// noinspection GoRedundantConversion
func TestMinSpecialElement(t *testing.T) {
	type args struct {
		parameter Element
		name      string
		value     Element
	}

	type want struct {
		stop bool
		err  error
	}

	test := func(a args, w want) func(t *testing.T) {
		return func(t *testing.T) {
			stop, err := MinElement(a.parameter)(a.name, a.value)
			assert.Equalf(t, w.stop, stop, "invalid stop: want = %v, got = %v", w.stop, stop)
			assert.ErrorIs(t, err, w.err)
			if wantErr, ok := w.err.(yav.Error); ok && assert.IsType(t, w.err, err) {
				gotErr := err.(yav.Error)
				assert.IsType(t, wantErr.Value, gotErr.Value)
				assert.InDelta(t, wantErr.Value, gotErr.Value, 0)
			}
		}
	}

	tests := []struct {
		name string
		args args
		want want
	}{{
		name: "NaN value",
		args: args{
			parameter: 0,
			name:      "v",
			value:     Element(math.NaN()),
		},
		want: want{
			stop: true,
			err: yav.Error{
				CheckName: "min",
				Parameter: "0",
				ValueName: "v",
				Value:     Element(math.NaN()),
			},
		},
	}, {
		name: "NaN bound",
		args: args{
			parameter: Element(math.NaN()),
			name:      "v",
			value:     Element(1),
		},
		want: want{
			stop: true,
			err: yav.Error{
				CheckName: "min",
				Parameter: "NaN",
				ValueName: "v",
				Value:     Element(1),
			},
		},
	}, {
		name: "positive infinity",
		args: args{
			parameter: 0,
			name:      "v",
			value:     Element(math.Inf(1)),
		},
		want: want{
			stop: false,
			err:  nil,
		},
	}, {
		name: "negative infinity",
		args: args{
			parameter: 0,
			name:      "v",
			value:     Element(math.Inf(-1)),
		},
		want: want{
			stop: true,
			err: yav.Error{
				CheckName: "min",
				Parameter: "0",
				ValueName: "v",
				Value:     Element(math.Inf(-1)),
			},
		},
	}, {
		name: "infinite bound",
		args: args{
			parameter: Element(math.Inf(1)),
			name:      "v",
			value:     Element(math.Inf(1)),
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

// noinspection GoRedundantConversion
func TestMaxSpecialElement(t *testing.T) {
	type args struct {
		parameter Element
		name      string
		value     Element
	}

	type want struct {
		stop bool
		err  error
	}

	test := func(a args, w want) func(t *testing.T) {
		return func(t *testing.T) {
			stop, err := MaxElement(a.parameter)(a.name, a.value)
			assert.Equalf(t, w.stop, stop, "invalid stop: want = %v, got = %v", w.stop, stop)
			assert.ErrorIs(t, err, w.err)
			if wantErr, ok := w.err.(yav.Error); ok && assert.IsType(t, w.err, err) {
				gotErr := err.(yav.Error)
				assert.IsType(t, wantErr.Value, gotErr.Value)
				assert.InDelta(t, wantErr.Value, gotErr.Value, 0)
			}
		}
	}

	tests := []struct {
		name string
		args args
		want want
	}{{
		name: "NaN value",
		args: args{
			parameter: 1,
			name:      "v",
			value:     Element(math.NaN()),
		},
		want: want{
			stop: true,
			err: yav.Error{
				CheckName: "max",
				Parameter: "1",
				ValueName: "v",
				Value:     Element(math.NaN()),
			},
		},
	}, {
		name: "NaN bound",
		args: args{
			parameter: Element(math.NaN()),
			name:      "v",
			value:     Element(1),
		},
		want: want{
			stop: true,
			err: yav.Error{
				CheckName: "max",
				Parameter: "NaN",
				ValueName: "v",
				Value:     Element(1),
			},
		},
	}, {
		name: "negative infinity",
		args: args{
			parameter: 1,
			name:      "v",
			value:     Element(math.Inf(-1)),
		},
		want: want{
			stop: false,
			err:  nil,
		},
	}, {
		name: "positive infinity",
		args: args{
			parameter: 1,
			name:      "v",
			value:     Element(math.Inf(1)),
		},
		want: want{
			stop: true,
			err: yav.Error{
				CheckName: "max",
				Parameter: "1",
				ValueName: "v",
				Value:     Element(math.Inf(1)),
			},
		},
	}, {
		name: "infinite bound",
		args: args{
			parameter: Element(math.Inf(-1)),
			name:      "v",
			value:     Element(math.Inf(-1)),
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

// noinspection GoRedundantConversion
func TestBetweenSpecialElement(t *testing.T) {
	type args struct {
		min, max Element
		name     string
		value    Element
	}

	type want struct {
		stop bool
		err  error
	}

	test := func(a args, w want) func(t *testing.T) {
		return func(t *testing.T) {
			stop, err := BetweenElement(a.min, a.max)(a.name, a.value)
			assert.Equalf(t, w.stop, stop, "invalid stop: want = %v, got = %v", w.stop, stop)
			assert.ErrorIs(t, err, w.err)
			if wantErr, ok := w.err.(yav.Error); ok && assert.IsType(t, w.err, err) {
				gotErr := err.(yav.Error)
				assert.IsType(t, wantErr.Value, gotErr.Value)
				assert.InDelta(t, wantErr.Value, gotErr.Value, 0)
			}
		}
	}

	tests := []struct {
		name string
		args args
		want want
	}{{
		name: "NaN value",
		args: args{
			min:   0,
			max:   1,
			name:  "v",
			value: Element(math.NaN()),
		},
		want: want{
			stop: true,
			err: yav.Error{
				CheckName: "min",
				Parameter: "0",
				ValueName: "v",
				Value:     Element(math.NaN()),
			},
		},
	}, {
		name: "NaN lower bound",
		args: args{
			min:   Element(math.NaN()),
			max:   1,
			name:  "v",
			value: Element(1),
		},
		want: want{
			stop: true,
			err: yav.Error{
				CheckName: "min",
				Parameter: "NaN",
				ValueName: "v",
				Value:     Element(1),
			},
		},
	}, {
		name: "NaN upper bound",
		args: args{
			min:   0,
			max:   Element(math.NaN()),
			name:  "v",
			value: Element(1),
		},
		want: want{
			stop: true,
			err: yav.Error{
				CheckName: "max",
				Parameter: "NaN",
				ValueName: "v",
				Value:     Element(1),
			},
		},
	}, {
		name: "negative infinity",
		args: args{
			min:   0,
			max:   1,
			name:  "v",
			value: Element(math.Inf(-1)),
		},
		want: want{
			stop: true,
			err: yav.Error{
				CheckName: "min",
				Parameter: "0",
				ValueName: "v",
				Value:     Element(math.Inf(-1)),
			},
		},
	}, {
		name: "positive infinity",
		args: args{
			min:   0,
			max:   1,
			name:  "v",
			value: Element(math.Inf(1)),
		},
		want: want{
			stop: true,
			err: yav.Error{
				CheckName: "max",
				Parameter: "1",
				ValueName: "v",
				Value:     Element(math.Inf(1)),
			},
		},
	}, {
		name: "unbounded negative infinity",
		args: args{
			min:   Element(math.Inf(-1)),
			max:   Element(math.Inf(1)),
			name:  "v",
			value: Element(math.Inf(-1)),
		},
		want: want{
			stop: false,
			err:  nil,
		},
	}, {
		name: "unbounded positive infinity",
		args: args{
			min:   Element(math.Inf(-1)),
			max:   Element(math.Inf(1)),
			name:  "v",
			value: Element(math.Inf(1)),
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

// noinspection GoRedundantConversion
func TestLessThanSpecialElement(t *testing.T) {
	type args struct {
		parameter Element
		name      string
		value     Element
	}

	type want struct {
		stop bool
		err  error
	}

	test := func(a args, w want) func(t *testing.T) {
		return func(t *testing.T) {
			stop, err := LessThanElement(a.parameter)(a.name, a.value)
			assert.Equalf(t, w.stop, stop, "invalid stop: want = %v, got = %v", w.stop, stop)
			assert.ErrorIs(t, err, w.err)
			if wantErr, ok := w.err.(yav.Error); ok && assert.IsType(t, w.err, err) {
				gotErr := err.(yav.Error)
				assert.IsType(t, wantErr.Value, gotErr.Value)
				assert.InDelta(t, wantErr.Value, gotErr.Value, 0)
			}
		}
	}

	tests := []struct {
		name string
		args args
		want want
	}{{
		name: "NaN value",
		args: args{
			parameter: 1,
			name:      "v",
			value:     Element(math.NaN()),
		},
		want: want{
			stop: true,
			err: yav.Error{
				CheckName: "lt",
				Parameter: "1",
				ValueName: "v",
				Value:     Element(math.NaN()),
			},
		},
	}, {
		name: "NaN bound",
		args: args{
			parameter: Element(math.NaN()),
			name:      "v",
			value:     Element(1),
		},
		want: want{
			stop: true,
			err: yav.Error{
				CheckName: "lt",
				Parameter: "NaN",
				ValueName: "v",
				Value:     Element(1),
			},
		},
	}, {
		name: "infinite bound",
		args: args{
			parameter: Element(math.Inf(1)),
			name:      "v",
			value:     Element(math.Inf(1)),
		},
		want: want{
			stop: true,
			err: yav.Error{
				CheckName: "lt",
				Parameter: "+Inf",
				ValueName: "v",
				Value:     Element(math.Inf(1)),
			},
		},
	}}

	for _, tt := range tests {
		t.Run(tt.name, test(tt.args, tt.want))
	}
}

// noinspection GoRedundantConversion
func TestLessThanOrEqualSpecialElement(t *testing.T) {
	type args struct {
		parameter Element
		name      string
		value     Element
	}

	type want struct {
		stop bool
		err  error
	}

	test := func(a args, w want) func(t *testing.T) {
		return func(t *testing.T) {
			stop, err := LessThanOrEqualElement(a.parameter)(a.name, a.value)
			assert.Equalf(t, w.stop, stop, "invalid stop: want = %v, got = %v", w.stop, stop)
			assert.ErrorIs(t, err, w.err)
			if wantErr, ok := w.err.(yav.Error); ok && assert.IsType(t, w.err, err) {
				gotErr := err.(yav.Error)
				assert.IsType(t, wantErr.Value, gotErr.Value)
				assert.InDelta(t, wantErr.Value, gotErr.Value, 0)
			}
		}
	}

	tests := []struct {
		name string
		args args
		want want
	}{{
		name: "NaN value",
		args: args{
			parameter: 1,
			name:      "v",
			value:     Element(math.NaN()),
		},
		want: want{
			stop: true,
			err: yav.Error{
				CheckName: "lte",
				Parameter: "1",
				ValueName: "v",
				Value:     Element(math.NaN()),
			},
		},
	}, {
		name: "NaN bound",
		args: args{
			parameter: Element(math.NaN()),
			name:      "v",
			value:     Element(1),
		},
		want: want{
			stop: true,
			err: yav.Error{
				CheckName: "lte",
				Parameter: "NaN",
				ValueName: "v",
				Value:     Element(1),
			},
		},
	}, {
		name: "infinite bound",
		args: args{
			parameter: Element(math.Inf(1)),
			name:      "v",
			value:     Element(math.Inf(1)),
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

// noinspection GoRedundantConversion
func TestGreaterThanSpecialElement(t *testing.T) {
	type args struct {
		parameter Element
		name      string
		value     Element
	}

	type want struct {
		stop bool
		err  error
	}

	test := func(a args, w want) func(t *testing.T) {
		return func(t *testing.T) {
			stop, err := GreaterThanElement(a.parameter)(a.name, a.value)
			assert.Equalf(t, w.stop, stop, "invalid stop: want = %v, got = %v", w.stop, stop)
			assert.ErrorIs(t, err, w.err)
			if wantErr, ok := w.err.(yav.Error); ok && assert.IsType(t, w.err, err) {
				gotErr := err.(yav.Error)
				assert.IsType(t, wantErr.Value, gotErr.Value)
				assert.InDelta(t, wantErr.Value, gotErr.Value, 0)
			}
		}
	}

	tests := []struct {
		name string
		args args
		want want
	}{{
		name: "NaN value",
		args: args{
			parameter: 0,
			name:      "v",
			value:     Element(math.NaN()),
		},
		want: want{
			stop: true,
			err: yav.Error{
				CheckName: "gt",
				Parameter: "0",
				ValueName: "v",
				Value:     Element(math.NaN()),
			},
		},
	}, {
		name: "NaN bound",
		args: args{
			parameter: Element(math.NaN()),
			name:      "v",
			value:     Element(1),
		},
		want: want{
			stop: true,
			err: yav.Error{
				CheckName: "gt",
				Parameter: "NaN",
				ValueName: "v",
				Value:     Element(1),
			},
		},
	}, {
		name: "infinite bound",
		args: args{
			parameter: Element(math.Inf(-1)),
			name:      "v",
			value:     Element(math.Inf(-1)),
		},
		want: want{
			stop: true,
			err: yav.Error{
				CheckName: "gt",
				Parameter: "-Inf",
				ValueName: "v",
				Value:     Element(math.Inf(-1)),
			},
		},
	}}

	for _, tt := range tests {
		t.Run(tt.name, test(tt.args, tt.want))
	}
}

// noinspection GoRedundantConversion
func TestGreaterThanOrEqualSpecialElement(t *testing.T) {
	type args struct {
		parameter Element
		name      string
		value     Element
	}

	type want struct {
		stop bool
		err  error
	}

	test := func(a args, w want) func(t *testing.T) {
		return func(t *testing.T) {
			stop, err := GreaterThanOrEqualElement(a.parameter)(a.name, a.value)
			assert.Equalf(t, w.stop, stop, "invalid stop: want = %v, got = %v", w.stop, stop)
			assert.ErrorIs(t, err, w.err)
			if wantErr, ok := w.err.(yav.Error); ok && assert.IsType(t, w.err, err) {
				gotErr := err.(yav.Error)
				assert.IsType(t, wantErr.Value, gotErr.Value)
				assert.InDelta(t, wantErr.Value, gotErr.Value, 0)
			}
		}
	}

	tests := []struct {
		name string
		args args
		want want
	}{{
		name: "NaN value",
		args: args{
			parameter: 0,
			name:      "v",
			value:     Element(math.NaN()),
		},
		want: want{
			stop: true,
			err: yav.Error{
				CheckName: "gte",
				Parameter: "0",
				ValueName: "v",
				Value:     Element(math.NaN()),
			},
		},
	}, {
		name: "NaN bound",
		args: args{
			parameter: Element(math.NaN()),
			name:      "v",
			value:     Element(1),
		},
		want: want{
			stop: true,
			err: yav.Error{
				CheckName: "gte",
				Parameter: "NaN",
				ValueName: "v",
				Value:     Element(1),
			},
		},
	}, {
		name: "infinite bound",
		args: args{
			parameter: Element(math.Inf(-1)),
			name:      "v",
			value:     Element(math.Inf(-1)),
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
