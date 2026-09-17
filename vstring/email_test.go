package vstring

import (
	"reflect"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/SladeThe/yav"
)

func TestEmail(t *testing.T) {
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
			stop, err := Email(a.name, a.value)
			assert.Equalf(t, w.stop, stop, "invalid stop: want = %v, got = %v", w.stop, stop)
			assert.Truef(t, reflect.DeepEqual(w.err, err), "invalid error: want = %v, got = %v", w.err, err)
		}
	}

	tests := []struct {
		name string
		args args
		want want
	}{{
		name: "address",
		args: args{
			name:  "v",
			value: "test@example.com",
		},
		want: want{
			stop: false,
			err:  nil,
		},
	}, {
		name: "single-character parts",
		args: args{
			name:  "v",
			value: "a@b.c",
		},
		want: want{
			stop: false,
			err:  nil,
		},
	}, {
		name: "dots and alias",
		args: args{
			name:  "v",
			value: "First.Last+tag@Sub.Example.COM",
		},
		want: want{
			stop: false,
			err:  nil,
		},
	}, {
		name: "local punctuation",
		args: args{
			name:  "v",
			value: "!#$%&'*+-/=?^_`{|}~@example.com",
		},
		want: want{
			stop: false,
			err:  nil,
		},
	}, {
		name: "Unicode",
		args: args{
			name:  "v",
			value: "Dörte@Sörensen.example.com",
		},
		want: want{
			stop: false,
			err:  nil,
		},
	}, {
		name: "multibyte Unicode",
		args: args{
			name:  "v",
			value: "用户@例子.广告",
		},
		want: want{
			stop: false,
			err:  nil,
		},
	}, {
		name: "quoted space",
		args: args{
			name:  "v",
			value: `"test test"@example.com`,
		},
		want: want{
			stop: false,
			err:  nil,
		},
	}, {
		name: "quoted at sign",
		args: args{
			name:  "v",
			value: `"user@tag"@example.com`,
		},
		want: want{
			stop: false,
			err:  nil,
		},
	}, {
		name: "quoted escapes",
		args: args{
			name:  "v",
			value: `"a\"b\\c"@example.com`,
		},
		want: want{
			stop: false,
			err:  nil,
		},
	}, {
		name: "quoted tab",
		args: args{
			name:  "v",
			value: "\"a\tb\"@example.com",
		},
		want: want{
			stop: false,
			err:  nil,
		},
	}, {
		name: "quoted Unicode",
		args: args{
			name:  "v",
			value: `"用户"@例子.广告`,
		},
		want: want{
			stop: false,
			err:  nil,
		},
	}, {
		name: "punycode TLD",
		args: args{
			name:  "v",
			value: "test@sub-domain.xn--p1ai",
		},
		want: want{
			stop: false,
			err:  nil,
		},
	}, {
		name: "TLD ending with digit",
		args: args{
			name:  "v",
			value: "test@example.a1",
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
				CheckName: "email",
				ValueName: "v",
				Value:     "",
			},
		},
	}, {
		name: "missing at sign",
		args: args{
			name:  "v",
			value: "test.example.com",
		},
		want: want{
			stop: true,
			err: yav.Error{
				CheckName: "email",
				ValueName: "v",
				Value:     "test.example.com",
			},
		},
	}, {
		name: "empty local part",
		args: args{
			name:  "v",
			value: "@example.com",
		},
		want: want{
			stop: true,
			err: yav.Error{
				CheckName: "email",
				ValueName: "v",
				Value:     "@example.com",
			},
		},
	}, {
		name: "empty domain",
		args: args{
			name:  "v",
			value: "test@",
		},
		want: want{
			stop: true,
			err: yav.Error{
				CheckName: "email",
				ValueName: "v",
				Value:     "test@",
			},
		},
	}, {
		name: "multiple at signs",
		args: args{
			name:  "v",
			value: "test@test@example.com",
		},
		want: want{
			stop: true,
			err: yav.Error{
				CheckName: "email",
				ValueName: "v",
				Value:     "test@test@example.com",
			},
		},
	}, {
		name: "leading local dot",
		args: args{
			name:  "v",
			value: ".test@example.com",
		},
		want: want{
			stop: true,
			err: yav.Error{
				CheckName: "email",
				ValueName: "v",
				Value:     ".test@example.com",
			},
		},
	}, {
		name: "trailing local dot",
		args: args{
			name:  "v",
			value: "test.@example.com",
		},
		want: want{
			stop: true,
			err: yav.Error{
				CheckName: "email",
				ValueName: "v",
				Value:     "test.@example.com",
			},
		},
	}, {
		name: "consecutive local dots",
		args: args{
			name:  "v",
			value: "test..name@example.com",
		},
		want: want{
			stop: true,
			err: yav.Error{
				CheckName: "email",
				ValueName: "v",
				Value:     "test..name@example.com",
			},
		},
	}, {
		name: "unquoted space",
		args: args{
			name:  "v",
			value: "test name@example.com",
		},
		want: want{
			stop: true,
			err: yav.Error{
				CheckName: "email",
				ValueName: "v",
				Value:     "test name@example.com",
			},
		},
	}, {
		name: "unquoted punctuation",
		args: args{
			name:  "v",
			value: "test,name@example.com",
		},
		want: want{
			stop: true,
			err: yav.Error{
				CheckName: "email",
				ValueName: "v",
				Value:     "test,name@example.com",
			},
		},
	}, {
		name: "display name",
		args: args{
			name:  "v",
			value: "Test Name <test@example.com>",
		},
		want: want{
			stop: true,
			err: yav.Error{
				CheckName: "email",
				ValueName: "v",
				Value:     "Test Name <test@example.com>",
			},
		},
	}, {
		name: "empty domain label",
		args: args{
			name:  "v",
			value: "a@foo..bar.com",
		},
		want: want{
			stop: true,
			err: yav.Error{
				CheckName: "email",
				ValueName: "v",
				Value:     "a@foo..bar.com",
			},
		},
	}, {
		name: "domain tilde",
		args: args{
			name:  "v",
			value: "a@a~b.com",
		},
		want: want{
			stop: true,
			err: yav.Error{
				CheckName: "email",
				ValueName: "v",
				Value:     "a@a~b.com",
			},
		},
	}, {
		name: "domain underscore",
		args: args{
			name:  "v",
			value: "a@a_b.com",
		},
		want: want{
			stop: true,
			err: yav.Error{
				CheckName: "email",
				ValueName: "v",
				Value:     "a@a_b.com",
			},
		},
	}, {
		name: "leading domain hyphen",
		args: args{
			name:  "v",
			value: "a@-example.com",
		},
		want: want{
			stop: true,
			err: yav.Error{
				CheckName: "email",
				ValueName: "v",
				Value:     "a@-example.com",
			},
		},
	}, {
		name: "trailing domain hyphen",
		args: args{
			name:  "v",
			value: "a@example-.com",
		},
		want: want{
			stop: true,
			err: yav.Error{
				CheckName: "email",
				ValueName: "v",
				Value:     "a@example-.com",
			},
		},
	}, {
		name: "trailing TLD hyphen",
		args: args{
			name:  "v",
			value: "a@example.com-",
		},
		want: want{
			stop: true,
			err: yav.Error{
				CheckName: "email",
				ValueName: "v",
				Value:     "a@example.com-",
			},
		},
	}, {
		name: "leading domain dot",
		args: args{
			name:  "v",
			value: "a@.example.com",
		},
		want: want{
			stop: true,
			err: yav.Error{
				CheckName: "email",
				ValueName: "v",
				Value:     "a@.example.com",
			},
		},
	}, {
		name: "trailing domain dot",
		args: args{
			name:  "v",
			value: "a@example.com.",
		},
		want: want{
			stop: true,
			err: yav.Error{
				CheckName: "email",
				ValueName: "v",
				Value:     "a@example.com.",
			},
		},
	}, {
		name: "single domain label",
		args: args{
			name:  "v",
			value: "a@localhost",
		},
		want: want{
			stop: true,
			err: yav.Error{
				CheckName: "email",
				ValueName: "v",
				Value:     "a@localhost",
			},
		},
	}, {
		name: "numeric TLD",
		args: args{
			name:  "v",
			value: "a@example.123",
		},
		want: want{
			stop: true,
			err: yav.Error{
				CheckName: "email",
				ValueName: "v",
				Value:     "a@example.123",
			},
		},
	}, {
		name: "empty quoted local part",
		args: args{
			name:  "v",
			value: `""@example.com`,
		},
		want: want{
			stop: true,
			err: yav.Error{
				CheckName: "email",
				ValueName: "v",
				Value:     `""@example.com`,
			},
		},
	}, {
		name: "unclosed quote",
		args: args{
			name:  "v",
			value: `"test@example.com`,
		},
		want: want{
			stop: true,
			err: yav.Error{
				CheckName: "email",
				ValueName: "v",
				Value:     `"test@example.com`,
			},
		},
	}, {
		name: "unescaped quote",
		args: args{
			name:  "v",
			value: `"test"name"@example.com`,
		},
		want: want{
			stop: true,
			err: yav.Error{
				CheckName: "email",
				ValueName: "v",
				Value:     `"test"name"@example.com`,
			},
		},
	}, {
		name: "escaped closing quote",
		args: args{
			name:  "v",
			value: `"test\"@example.com`,
		},
		want: want{
			stop: true,
			err: yav.Error{
				CheckName: "email",
				ValueName: "v",
				Value:     `"test\"@example.com`,
			},
		},
	}, {
		name: "quoted newline",
		args: args{
			name:  "v",
			value: "\"test\nname\"@example.com",
		},
		want: want{
			stop: true,
			err: yav.Error{
				CheckName: "email",
				ValueName: "v",
				Value:     "\"test\nname\"@example.com",
			},
		},
	}, {
		name: "quoted control",
		args: args{
			name:  "v",
			value: "\"test\x7fname\"@example.com",
		},
		want: want{
			stop: true,
			err: yav.Error{
				CheckName: "email",
				ValueName: "v",
				Value:     "\"test\x7fname\"@example.com",
			},
		},
	}, {
		name: "excluded Unicode",
		args: args{
			name:  "v",
			value: "\uE000@example.com",
		},
		want: want{
			stop: true,
			err: yav.Error{
				CheckName: "email",
				ValueName: "v",
				Value:     "\uE000@example.com",
			},
		},
	}, {
		name: "invalid local UTF-8",
		args: args{
			name:  "v",
			value: "\xff@example.com",
		},
		want: want{
			stop: true,
			err: yav.Error{
				CheckName: "email",
				ValueName: "v",
				Value:     "\xff@example.com",
			},
		},
	}, {
		name: "invalid domain UTF-8",
		args: args{
			name:  "v",
			value: "a@exa\xffmple.com",
		},
		want: want{
			stop: true,
			err: yav.Error{
				CheckName: "email",
				ValueName: "v",
				Value:     "a@exa\xffmple.com",
			},
		},
	}, {
		name: "invalid quoted UTF-8",
		args: args{
			name:  "v",
			value: "\"a\xff\"@example.com",
		},
		want: want{
			stop: true,
			err: yav.Error{
				CheckName: "email",
				ValueName: "v",
				Value:     "\"a\xff\"@example.com",
			},
		},
	}}

	for _, tt := range tests {
		t.Run(tt.name, test(tt.args, tt.want))
	}
}

func BenchmarkEmail(b *testing.B) {
	tests := []struct {
		name  string
		value string
	}{
		{"ASCII", "test@example.com"},
		{"alias", "first.last+tag@example.co.uk"},
		{"Unicode", "用户@例子.广告"},
		{"quoted", `"test test"@example.com`},
	}

	for _, tt := range tests {
		b.Run(tt.name, func(b *testing.B) {
			b.ReportAllocs()

			var stop bool
			var err error

			for i := 0; i < b.N; i++ {
				stop, err = Email("email", tt.value)
			}

			runtime.KeepAlive(stop)
			runtime.KeepAlive(err)
		})
	}
}
