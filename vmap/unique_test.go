package vmap

import (
	"math/rand"
	"reflect"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/SladeThe/yav"
)

func TestUnique(t *testing.T) {
	type args struct {
		name  string
		value map[int]uint
	}

	type want struct {
		stop bool
		err  error
	}

	test := func(a args, w want) func(t *testing.T) {
		return func(t *testing.T) {
			stop, err := Unique(a.name, a.value)
			assert.Equalf(t, w.stop, stop, "invalid stop: want = %v, got = %v", w.stop, stop)
			assert.Truef(t, reflect.DeepEqual(w.err, err), "invalid error: want = %v, got = %v", w.err, err)
		}
	}

	tests := []struct {
		name string
		args args
		want want
	}{{
		name: "unique nil",
		args: args{
			name:  "",
			value: nil,
		},
		want: want{
			stop: false,
			err:  nil,
		},
	}, {
		name: "unique 0",
		args: args{
			name:  "",
			value: map[int]uint{},
		},
		want: want{
			stop: false,
			err:  nil,
		},
	}, {
		name: "unique 1",
		args: args{
			name:  "",
			value: map[int]uint{-1: 1},
		},
		want: want{
			stop: false,
			err:  nil,
		},
	}, {
		name: "unique 2",
		args: args{
			name:  "",
			value: map[int]uint{-1: 1, -2: 2},
		},
		want: want{
			stop: false,
			err:  nil,
		},
	}, {
		name: "duplicate 2",
		args: args{
			name:  "m",
			value: map[int]uint{-1: 1, 0: 1},
		},
		want: want{
			stop: true,
			err: yav.Error{
				CheckName: "unique",
				ValueName: "m",
				Value:     map[int]uint{-1: 1, 0: 1},
			},
		},
	}, {
		name: "duplicate 3",
		args: args{
			name:  "m",
			value: map[int]uint{-2: 2, -1: 1, 0: 1},
		},
		want: want{
			stop: true,
			err: yav.Error{
				CheckName: "unique",
				ValueName: "m",
				Value:     map[int]uint{-2: 2, -1: 1, 0: 1},
			},
		},
	}, {
		name: "duplicate 4",
		args: args{
			name:  "m",
			value: map[int]uint{-3: 3, -2: 2, -1: 1, 0: 1},
		},
		want: want{
			stop: true,
			err: yav.Error{
				CheckName: "unique",
				ValueName: "m",
				Value:     map[int]uint{-3: 3, -2: 2, -1: 1, 0: 1},
			},
		},
	}, {
		name: "duplicate 5",
		args: args{
			name:  "m",
			value: map[int]uint{-4: 4, -3: 3, -2: 2, -1: 1, 0: 1},
		},
		want: want{
			stop: true,
			err: yav.Error{
				CheckName: "unique",
				ValueName: "m",
				Value:     map[int]uint{-4: 4, -3: 3, -2: 2, -1: 1, 0: 1},
			},
		},
	}, {
		name: "duplicate 6",
		args: args{
			name:  "m",
			value: map[int]uint{-5: 5, -4: 4, -3: 3, -2: 2, -1: 1, 0: 1},
		},
		want: want{
			stop: true,
			err: yav.Error{
				CheckName: "unique",
				ValueName: "m",
				Value:     map[int]uint{-5: 5, -4: 4, -3: 3, -2: 2, -1: 1, 0: 1},
			},
		},
	}}

	for _, tt := range tests {
		t.Run(tt.name, test(tt.args, tt.want))
	}
}

func BenchmarkUnique(b *testing.B) {
	const maxLength = 16

	rnd := rand.New(rand.NewSource(1432164634643))

	var mm []map[int64]uint64

	for length := 0; length <= maxLength; length++ {
		m := make(map[int64]uint64, length)

		for len(m) < length {
			v := rnd.Int63()

			if _, ok := m[-v]; ok {
				continue
			}

			m[-v] = uint64(v)
		}

		mm = append(mm, m)
	}

	b.ReportAllocs()

	var stop bool
	var err error

	for i := 0; i < b.N; i++ {
		for _, m := range mm {
			stop, err = Unique("name", m)
		}
	}

	runtime.KeepAlive(stop)
	runtime.KeepAlive(err)
}
