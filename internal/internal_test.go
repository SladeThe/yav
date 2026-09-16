package internal

import (
	"math"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRegisterMapEntry(t *testing.T) {
	type args struct {
		cache map[float64]string
		keys  []float64
		value string
	}

	type want struct {
		cache map[float64]string
		value string
	}

	test := func(a args, w want) func(t *testing.T) {
		return func(t *testing.T) {
			for _, key := range a.keys {
				value := RegisterMapEntry(&a.cache, key, a.value)
				assert.Equalf(t, w.value, value, "invalid value: want = %v, got = %v", w.value, value)
			}

			assert.Truef(t, reflect.DeepEqual(w.cache, a.cache), "invalid cache: want = %v, got = %v", w.cache, a.cache)
		}
	}

	tests := []struct {
		name string
		args args
		want want
	}{{
		name: "ordinary keys",
		args: args{
			keys:  []float64{1, 2},
			value: "new",
		},
		want: want{
			cache: map[float64]string{1: "new", 2: "new"},
			value: "new",
		},
	}, {
		name: "NaN with nil cache",
		args: args{
			keys:  []float64{math.NaN(), math.NaN()},
			value: "new",
		},
		want: want{
			value: "new",
		},
	}, {
		name: "NaN with existing cache",
		args: args{
			cache: map[float64]string{1: "existing"},
			keys:  []float64{math.NaN(), math.NaN()},
			value: "new",
		},
		want: want{
			cache: map[float64]string{1: "existing"},
			value: "new",
		},
	}}

	for _, tt := range tests {
		t.Run(tt.name, test(tt.args, tt.want))
	}
}
