package vnumber

import (
	"reflect"
	"testing"

	"github.com/cheekybits/genny/generic"
	"github.com/stretchr/testify/assert"

	"github.com/SladeThe/yav"
)

type Element generic.Type

func TestRequiredWithAllElement(t *testing.T) {
	type args struct {
		parameters []int
		name       string
		value      Element
	}

	type want struct {
		stop bool
		err  error
	}

	test := func(a args, w want) func(t *testing.T) {
		return func(t *testing.T) {
			accumulator := RequiredWithAllElement()

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
		name: "empty required",
		args: args{
			parameters: []int{-1, 1},
			name:       "v",
			value:      0,
		},
		want: want{
			stop: true,
			err: yav.Error{
				CheckName: yav.CheckNameRequiredWithAll,
				Parameter: "pp",
				ValueName: "v",
			},
		},
	}, {
		name: "empty non required",
		args: args{
			parameters: []int{-1, 0},
			name:       "v",
			value:      0,
		},
		want: want{
			stop: true,
			err:  nil,
		},
	}, {
		name: "not empty",
		args: args{
			parameters: []int{1},
			name:       "v",
			value:      1,
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
