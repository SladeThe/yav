package vstring

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/SladeThe/yav"
)

func TestURI(t *testing.T) {
	type args struct {
		name  string
		value string
	}

	type want struct {
		stop bool
		err  error
	}

	test := func(a args, w want) func(t *testing.T) {
		return func(t *testing.T) {
			stop, err := URI(a.name, a.value)
			assert.Equalf(t, w.stop, stop, "invalid stop: want = %v, got = %v", w.stop, stop)
			assert.Truef(t, reflect.DeepEqual(w.err, err), "invalid error: want = %v, got = %v", w.err, err)
		}
	}

	tests := []struct {
		name string
		args args
		want want
	}{{
		name: "no fragment",
		args: args{
			name:  "v",
			value: "https://example.com/path",
		},
		want: want{
			stop: false,
			err:  nil,
		},
	}, {
		name: "plain fragment",
		args: args{
			name:  "v",
			value: "https://example.com/path#section",
		},
		want: want{
			stop: false,
			err:  nil,
		},
	}, {
		name: "empty fragment",
		args: args{
			name:  "v",
			value: "https://example.com/path#",
		},
		want: want{
			stop: false,
			err:  nil,
		},
	}, {
		name: "escaped fragment",
		args: args{
			name:  "v",
			value: "https://example.com/path#a%20b%2Fc%3fd%25zz%00%7f",
		},
		want: want{
			stop: false,
			err:  nil,
		},
	}, {
		name: "Unicode fragment",
		args: args{
			name:  "v",
			value: "https://example.com/path#раздел",
		},
		want: want{
			stop: false,
			err:  nil,
		},
	}, {
		name: "empty",
		args: args{
			name:  "v",
			value: "",
		},
		want: want{
			stop: true,
			err: yav.Error{
				CheckName: "uri",
				ValueName: "v",
				Value:     "",
			},
		},
	}, {
		name: "fragment only",
		args: args{
			name:  "v",
			value: "#section",
		},
		want: want{
			stop: true,
			err: yav.Error{
				CheckName: "uri",
				ValueName: "v",
				Value:     "#section",
			},
		},
	}, {
		name: "invalid prefix",
		args: args{
			name:  "v",
			value: "https://exa\nmple.com/#section",
		},
		want: want{
			stop: true,
			err: yav.Error{
				CheckName: "uri",
				ValueName: "v",
				Value:     "https://exa\nmple.com/#section",
			},
		},
	}, {
		name: "missing escape digits",
		args: args{
			name:  "v",
			value: "https://example.com/path#%",
		},
		want: want{
			stop: true,
			err: yav.Error{
				CheckName: "uri",
				ValueName: "v",
				Value:     "https://example.com/path#%",
			},
		},
	}, {
		name: "short escape",
		args: args{
			name:  "v",
			value: "https://example.com/path#%a",
		},
		want: want{
			stop: true,
			err: yav.Error{
				CheckName: "uri",
				ValueName: "v",
				Value:     "https://example.com/path#%a",
			},
		},
	}, {
		name: "invalid first escape digit",
		args: args{
			name:  "v",
			value: "https://example.com/path#%g0",
		},
		want: want{
			stop: true,
			err: yav.Error{
				CheckName: "uri",
				ValueName: "v",
				Value:     "https://example.com/path#%g0",
			},
		},
	}, {
		name: "invalid second escape digit",
		args: args{
			name:  "v",
			value: "https://example.com/path#%0g",
		},
		want: want{
			stop: true,
			err: yav.Error{
				CheckName: "uri",
				ValueName: "v",
				Value:     "https://example.com/path#%0g",
			},
		},
	}, {
		name: "fragment newline",
		args: args{
			name:  "v",
			value: "https://example.com/path#first\nsecond",
		},
		want: want{
			stop: true,
			err: yav.Error{
				CheckName: "uri",
				ValueName: "v",
				Value:     "https://example.com/path#first\nsecond",
			},
		},
	}, {
		name: "fragment NUL",
		args: args{
			name:  "v",
			value: "https://example.com/path#first\u0000second",
		},
		want: want{
			stop: true,
			err: yav.Error{
				CheckName: "uri",
				ValueName: "v",
				Value:     "https://example.com/path#first\u0000second",
			},
		},
	}, {
		name: "fragment DEL",
		args: args{
			name:  "v",
			value: "https://example.com/path#first\x7fsecond",
		},
		want: want{
			stop: true,
			err: yav.Error{
				CheckName: "uri",
				ValueName: "v",
				Value:     "https://example.com/path#first\x7fsecond",
			},
		},
	}, {
		name: "absolute path",
		args: args{
			name:  "v",
			value: "/path#section",
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

func TestURL(t *testing.T) {
	type args struct {
		name  string
		value string
	}

	type want struct {
		stop bool
		err  error
	}

	test := func(a args, w want) func(t *testing.T) {
		return func(t *testing.T) {
			stop, err := URL(a.name, a.value)
			assert.Equalf(t, w.stop, stop, "invalid stop: want = %v, got = %v", w.stop, stop)
			assert.Truef(t, reflect.DeepEqual(w.err, err), "invalid error: want = %v, got = %v", w.err, err)
		}
	}

	tests := []struct {
		name string
		args args
		want want
	}{{
		name: "no fragment",
		args: args{
			name:  "v",
			value: "https://example.com/path",
		},
		want: want{
			stop: false,
			err:  nil,
		},
	}, {
		name: "plain fragment",
		args: args{
			name:  "v",
			value: "https://example.com/path#section",
		},
		want: want{
			stop: false,
			err:  nil,
		},
	}, {
		name: "empty fragment",
		args: args{
			name:  "v",
			value: "https://example.com/path#",
		},
		want: want{
			stop: false,
			err:  nil,
		},
	}, {
		name: "escaped fragment",
		args: args{
			name:  "v",
			value: "https://example.com/path#a%20b%2Fc%3fd%25zz%00%7f",
		},
		want: want{
			stop: false,
			err:  nil,
		},
	}, {
		name: "Unicode fragment",
		args: args{
			name:  "v",
			value: "https://example.com/path#раздел",
		},
		want: want{
			stop: false,
			err:  nil,
		},
	}, {
		name: "empty",
		args: args{
			name:  "v",
			value: "",
		},
		want: want{
			stop: true,
			err: yav.Error{
				CheckName: "url",
				ValueName: "v",
				Value:     "",
			},
		},
	}, {
		name: "fragment only",
		args: args{
			name:  "v",
			value: "#section",
		},
		want: want{
			stop: true,
			err: yav.Error{
				CheckName: "url",
				ValueName: "v",
				Value:     "#section",
			},
		},
	}, {
		name: "invalid prefix",
		args: args{
			name:  "v",
			value: "https://exa\nmple.com/#section",
		},
		want: want{
			stop: true,
			err: yav.Error{
				CheckName: "url",
				ValueName: "v",
				Value:     "https://exa\nmple.com/#section",
			},
		},
	}, {
		name: "missing escape digits",
		args: args{
			name:  "v",
			value: "https://example.com/path#%",
		},
		want: want{
			stop: true,
			err: yav.Error{
				CheckName: "url",
				ValueName: "v",
				Value:     "https://example.com/path#%",
			},
		},
	}, {
		name: "short escape",
		args: args{
			name:  "v",
			value: "https://example.com/path#%a",
		},
		want: want{
			stop: true,
			err: yav.Error{
				CheckName: "url",
				ValueName: "v",
				Value:     "https://example.com/path#%a",
			},
		},
	}, {
		name: "invalid first escape digit",
		args: args{
			name:  "v",
			value: "https://example.com/path#%g0",
		},
		want: want{
			stop: true,
			err: yav.Error{
				CheckName: "url",
				ValueName: "v",
				Value:     "https://example.com/path#%g0",
			},
		},
	}, {
		name: "invalid second escape digit",
		args: args{
			name:  "v",
			value: "https://example.com/path#%0g",
		},
		want: want{
			stop: true,
			err: yav.Error{
				CheckName: "url",
				ValueName: "v",
				Value:     "https://example.com/path#%0g",
			},
		},
	}, {
		name: "fragment newline",
		args: args{
			name:  "v",
			value: "https://example.com/path#first\nsecond",
		},
		want: want{
			stop: true,
			err: yav.Error{
				CheckName: "url",
				ValueName: "v",
				Value:     "https://example.com/path#first\nsecond",
			},
		},
	}, {
		name: "fragment NUL",
		args: args{
			name:  "v",
			value: "https://example.com/path#first\u0000second",
		},
		want: want{
			stop: true,
			err: yav.Error{
				CheckName: "url",
				ValueName: "v",
				Value:     "https://example.com/path#first\u0000second",
			},
		},
	}, {
		name: "fragment DEL",
		args: args{
			name:  "v",
			value: "https://example.com/path#first\x7fsecond",
		},
		want: want{
			stop: true,
			err: yav.Error{
				CheckName: "url",
				ValueName: "v",
				Value:     "https://example.com/path#first\x7fsecond",
			},
		},
	}, {
		name: "absolute path",
		args: args{
			name:  "v",
			value: "/path#section",
		},
		want: want{
			stop: true,
			err: yav.Error{
				CheckName: "url",
				ValueName: "v",
				Value:     "/path#section",
			},
		},
	}}

	for _, tt := range tests {
		t.Run(tt.name, test(tt.args, tt.want))
	}
}
