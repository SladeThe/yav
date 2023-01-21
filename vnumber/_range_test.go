package vnumber

import (
	"reflect"
	"testing"

	"github.com/cheekybits/genny/generic"
	"github.com/stretchr/testify/assert"

	"github.com/SladeThe/yav"
)

type Element generic.Type

func TestMinElement(t *testing.T) {
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
			value:     2,
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
			value:     1,
		},
		want: want{
			stop: true,
			err: yav.Error{
				CheckName: "min",
				Parameter: "2",
				ValueName: "v",
				Value:     Element(1),
			},
		},
	}}

	for _, tt := range tests {
		t.Run(tt.name, test(tt.args, tt.want))
	}
}

func TestMaxElement(t *testing.T) {
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
			value:     2,
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
			value:     3,
		},
		want: want{
			stop: true,
			err: yav.Error{
				CheckName: "max",
				Parameter: "2",
				ValueName: "v",
				Value:     Element(3),
			},
		},
	}}

	for _, tt := range tests {
		t.Run(tt.name, test(tt.args, tt.want))
	}
}

func TestBetweenElement(t *testing.T) {
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
			value: 2,
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
			value: 0,
		},
		want: want{
			stop: true,
			err: yav.Error{
				CheckName: "min",
				Parameter: "1",
				ValueName: "v",
				Value:     Element(0),
			},
		},
	}, {
		name: "invalid higher",
		args: args{
			min:   1,
			max:   2,
			name:  "v",
			value: 3,
		},
		want: want{
			stop: true,
			err: yav.Error{
				CheckName: "max",
				Parameter: "2",
				ValueName: "v",
				Value:     Element(3),
			},
		},
	}}

	for _, tt := range tests {
		t.Run(tt.name, test(tt.args, tt.want))
	}
}

func TestLessThanElement(t *testing.T) {
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
			value:     1,
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
			value:     2,
		},
		want: want{
			stop: true,
			err: yav.Error{
				CheckName: "lt",
				Parameter: "2",
				ValueName: "v",
				Value:     Element(2),
			},
		},
	}}

	for _, tt := range tests {
		t.Run(tt.name, test(tt.args, tt.want))
	}
}

func TestLessThanOrEqualElement(t *testing.T) {
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
			value:     2,
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
			value:     3,
		},
		want: want{
			stop: true,
			err: yav.Error{
				CheckName: "lte",
				Parameter: "2",
				ValueName: "v",
				Value:     Element(3),
			},
		},
	}}

	for _, tt := range tests {
		t.Run(tt.name, test(tt.args, tt.want))
	}
}

func TestGreaterThanElement(t *testing.T) {
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
			parameter: 1,
			name:      "",
			value:     2,
		},
		want: want{
			stop: false,
			err:  nil,
		},
	}, {
		name: "invalid",
		args: args{
			parameter: 1,
			name:      "v",
			value:     1,
		},
		want: want{
			stop: true,
			err: yav.Error{
				CheckName: "gt",
				Parameter: "1",
				ValueName: "v",
				Value:     Element(1),
			},
		},
	}}

	for _, tt := range tests {
		t.Run(tt.name, test(tt.args, tt.want))
	}
}

func TestGreaterThanOrEqualElement(t *testing.T) {
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
			value:     2,
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
			value:     1,
		},
		want: want{
			stop: true,
			err: yav.Error{
				CheckName: "gte",
				Parameter: "2",
				ValueName: "v",
				Value:     Element(1),
			},
		},
	}}

	for _, tt := range tests {
		t.Run(tt.name, test(tt.args, tt.want))
	}
}
