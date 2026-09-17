package vstring

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/SladeThe/yav"
)

func TestIsTitle(t *testing.T) {
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
			stop, err := Title(a.name, a.value)
			assert.Equalf(t, w.stop, stop, "invalid stop: want = %v, got = %v", w.stop, stop)
			assert.Truef(t, reflect.DeepEqual(w.err, err), "invalid error: want = %v, got = %v", w.err, err)
		}
	}

	tests := []struct {
		name string
		args args
		want want
	}{{
		name: "empty",
		args: args{
			name:  "v",
			value: "",
		},
		want: want{
			stop: true,
			err: yav.Error{
				CheckName: yav.CheckNameTitle,
				ValueName: "v",
				Value:     "",
			},
		},
	}, {
		name: "leading space",
		args: args{
			name:  "v",
			value: " Title",
		},
		want: want{
			stop: true,
			err: yav.Error{
				CheckName: yav.CheckNameTitle,
				ValueName: "v",
				Value:     " Title",
			},
		},
	}, {
		name: "leading tab",
		args: args{
			name:  "v",
			value: "\tTitle",
		},
		want: want{
			stop: true,
			err: yav.Error{
				CheckName: yav.CheckNameTitle,
				ValueName: "v",
				Value:     "\tTitle",
			},
		},
	}, {
		name: "trailing space",
		args: args{
			name:  "v",
			value: "Title ",
		},
		want: want{
			stop: true,
			err: yav.Error{
				CheckName: yav.CheckNameTitle,
				ValueName: "v",
				Value:     "Title ",
			},
		},
	}, {
		name: "trailing tab",
		args: args{
			name:  "v",
			value: "Title\t",
		},
		want: want{
			stop: true,
			err: yav.Error{
				CheckName: yav.CheckNameTitle,
				ValueName: "v",
				Value:     "Title\t",
			},
		},
	}, {
		name: "repeating spaces",
		args: args{
			name:  "v",
			value: "Two  words",
		},
		want: want{
			stop: true,
			err: yav.Error{
				CheckName: yav.CheckNameTitle,
				ValueName: "v",
				Value:     "Two  words",
			},
		},
	}, {
		name: "tab",
		args: args{
			name:  "v",
			value: "Two\twords",
		},
		want: want{
			stop: true,
			err: yav.Error{
				CheckName: yav.CheckNameTitle,
				ValueName: "v",
				Value:     "Two\twords",
			},
		},
	}, {
		name: "nbsp",
		args: args{
			name:  "v",
			value: "Two\u00A0words",
		},
		want: want{
			stop: true,
			err: yav.Error{
				CheckName: yav.CheckNameTitle,
				ValueName: "v",
				Value:     "Two\u00A0words",
			},
		},
	}, {
		name: "valid one word",
		args: args{
			name:  "v",
			value: "Title",
		},
		want: want{
			stop: false,
			err:  nil,
		},
	}, {
		name: "valid two words",
		args: args{
			name:  "v",
			value: "Two words",
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

func TestLowercase(t *testing.T) {
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
			stop, err := Lowercase(a.name, a.value)
			assert.Equalf(t, w.stop, stop, "invalid stop: want = %v, got = %v", w.stop, stop)
			assert.Truef(t, reflect.DeepEqual(w.err, err), "invalid error: want = %v, got = %v", w.err, err)
		}
	}

	tests := []struct {
		name string
		args args
		want want
	}{{
		name: "empty",
		args: args{
			name:  "v",
			value: "",
		},
		want: want{
			stop: false,
			err:  nil,
		},
	}, {
		name: "valid",
		args: args{
			name:  "v",
			value: "abc éǆⅰ中文123\uFFFD🙂",
		},
		want: want{
			stop: false,
			err:  nil,
		},
	}, {
		name: "unchanged letter",
		args: args{
			name:  "v",
			value: "𝔄",
		},
		want: want{
			stop: false,
			err:  nil,
		},
	}, {
		name: "opposite ASCII case",
		args: args{
			name:  "v",
			value: "A",
		},
		want: want{
			stop: true,
			err: yav.Error{
				CheckName: yav.CheckNameLowercase,
				ValueName: "v",
				Value:     "A",
			},
		},
	}, {
		name: "opposite Latin-1 case",
		args: args{
			name:  "v",
			value: "É",
		},
		want: want{
			stop: true,
			err: yav.Error{
				CheckName: yav.CheckNameLowercase,
				ValueName: "v",
				Value:     "É",
			},
		},
	}, {
		name: "titlecase",
		args: args{
			name:  "v",
			value: "ǅ",
		},
		want: want{
			stop: true,
			err: yav.Error{
				CheckName: yav.CheckNameLowercase,
				ValueName: "v",
				Value:     "ǅ",
			},
		},
	}, {
		name: "opposite Roman numeral",
		args: args{
			name:  "v",
			value: "Ⅰ",
		},
		want: want{
			stop: true,
			err: yav.Error{
				CheckName: yav.CheckNameLowercase,
				ValueName: "v",
				Value:     "Ⅰ",
			},
		},
	}, {
		name: "invalid UTF-8",
		args: args{
			name:  "v",
			value: "\xff",
		},
		want: want{
			stop: true,
			err: yav.Error{
				CheckName: yav.CheckNameLowercase,
				ValueName: "v",
				Value:     "\xff",
			},
		},
	}}

	for _, tt := range tests {
		t.Run(tt.name, test(tt.args, tt.want))
	}
}

func TestUppercase(t *testing.T) {
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
			stop, err := Uppercase(a.name, a.value)
			assert.Equalf(t, w.stop, stop, "invalid stop: want = %v, got = %v", w.stop, stop)
			assert.Truef(t, reflect.DeepEqual(w.err, err), "invalid error: want = %v, got = %v", w.err, err)
		}
	}

	tests := []struct {
		name string
		args args
		want want
	}{{
		name: "empty",
		args: args{
			name:  "v",
			value: "",
		},
		want: want{
			stop: false,
			err:  nil,
		},
	}, {
		name: "valid",
		args: args{
			name:  "v",
			value: "ABC ÉǄⅠ中文123\uFFFD🙂",
		},
		want: want{
			stop: false,
			err:  nil,
		},
	}, {
		name: "unchanged letter",
		args: args{
			name:  "v",
			value: "ß",
		},
		want: want{
			stop: false,
			err:  nil,
		},
	}, {
		name: "opposite ASCII case",
		args: args{
			name:  "v",
			value: "a",
		},
		want: want{
			stop: true,
			err: yav.Error{
				CheckName: yav.CheckNameUppercase,
				ValueName: "v",
				Value:     "a",
			},
		},
	}, {
		name: "opposite Latin-1 case",
		args: args{
			name:  "v",
			value: "ÿ",
		},
		want: want{
			stop: true,
			err: yav.Error{
				CheckName: yav.CheckNameUppercase,
				ValueName: "v",
				Value:     "ÿ",
			},
		},
	}, {
		name: "titlecase",
		args: args{
			name:  "v",
			value: "ǅ",
		},
		want: want{
			stop: true,
			err: yav.Error{
				CheckName: yav.CheckNameUppercase,
				ValueName: "v",
				Value:     "ǅ",
			},
		},
	}, {
		name: "opposite Roman numeral",
		args: args{
			name:  "v",
			value: "ⅰ",
		},
		want: want{
			stop: true,
			err: yav.Error{
				CheckName: yav.CheckNameUppercase,
				ValueName: "v",
				Value:     "ⅰ",
			},
		},
	}, {
		name: "invalid UTF-8",
		args: args{
			name:  "v",
			value: "\xff",
		},
		want: want{
			stop: true,
			err: yav.Error{
				CheckName: yav.CheckNameUppercase,
				ValueName: "v",
				Value:     "\xff",
			},
		},
	}}

	for _, tt := range tests {
		t.Run(tt.name, test(tt.args, tt.want))
	}
}

type characterTest struct {
	name  string
	value string
	valid bool
}

func testCharacterCheck(t *testing.T, validate yav.ValidateFunc[string], checkName string, tests []characterTest) {
	t.Helper()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var wantErr error

			if !tt.valid {
				wantErr = yav.Error{
					CheckName: checkName,
					ValueName: "v",
					Value:     tt.value,
				}
			}

			stop, err := validate("v", tt.value)
			assert.Equalf(t, !tt.valid, stop, "invalid stop: want = %v, got = %v", !tt.valid, stop)
			assert.Truef(t, reflect.DeepEqual(wantErr, err), "invalid error: want = %v, got = %v", wantErr, err)
		})
	}
}

func TestAlpha(t *testing.T) {
	tests := []characterTest{
		{name: "empty", value: "", valid: true},
		{name: "bounds", value: "aAzZ", valid: true},
		{name: "word", value: "Hello", valid: true},
		{name: "digit", value: "abc0", valid: false},
		{name: "Unicode", value: "é", valid: false},
	}

	testCharacterCheck(t, Alpha, yav.CheckNameAlpha, tests)
}

func TestAlphanumeric(t *testing.T) {
	tests := []characterTest{
		{name: "empty", value: "", valid: true},
		{name: "bounds", value: "aAzZ09", valid: true},
		{name: "word", value: "value123", valid: true},
		{name: "punctuation", value: "x_y", valid: false},
		{name: "Unicode", value: "é", valid: false},
	}

	testCharacterCheck(t, Alphanumeric, yav.CheckNameAlphanumeric, tests)
}

func TestNumeric(t *testing.T) {
	tests := []characterTest{
		{name: "empty", value: "", valid: true},
		{name: "bounds", value: "09", valid: true},
		{name: "number", value: "1234567890", valid: true},
		{name: "signed fraction", value: "-1.5", valid: false},
		{name: "Unicode digit", value: "١", valid: false},
	}

	testCharacterCheck(t, Numeric, yav.CheckNameNumeric, tests)
}

func TestContainsAlpha(t *testing.T) {
	tests := []characterTest{
		{name: "empty", value: "", valid: false},
		{name: "no letters", value: "09_", valid: false},
		{name: "lowercase", value: "1a2", valid: true},
		{name: "uppercase", value: "0Z9", valid: true},
		{name: "Unicode", value: "é", valid: false},
	}

	testCharacterCheck(t, ContainsAlpha, yav.CheckNameContainsAlpha, tests)
}

func TestContainsLowerAlpha(t *testing.T) {
	tests := []characterTest{
		{name: "empty", value: "", valid: false},
		{name: "no lowercase", value: "AZ09_", valid: false},
		{name: "mixed", value: "1aZ", valid: true},
		{name: "bounds", value: "az", valid: true},
		{name: "Unicode", value: "é", valid: false},
	}

	testCharacterCheck(t, ContainsLowerAlpha, yav.CheckNameContainsLowerAlpha, tests)
}

func TestContainsUpperAlpha(t *testing.T) {
	tests := []characterTest{
		{name: "empty", value: "", valid: false},
		{name: "no uppercase", value: "az09_", valid: false},
		{name: "mixed", value: "1Za", valid: true},
		{name: "bounds", value: "AZ", valid: true},
		{name: "Unicode", value: "É", valid: false},
	}

	testCharacterCheck(t, ContainsUpperAlpha, yav.CheckNameContainsUpperAlpha, tests)
}

func TestContainsDigit(t *testing.T) {
	tests := []characterTest{
		{name: "empty", value: "", valid: false},
		{name: "no digits", value: "abc_", valid: false},
		{name: "mixed", value: "a0z", valid: true},
		{name: "number", value: "1239", valid: true},
		{name: "Unicode digit", value: "١", valid: false},
	}

	testCharacterCheck(t, ContainsDigit, yav.CheckNameContainsDigit, tests)
}

func TestContainsSpecialCharacter(t *testing.T) {
	tests := []characterTest{
		{name: "empty", value: "", valid: false},
		{name: "alphanumeric", value: "abc123", valid: false},
		{name: "exclamation", value: "name!", valid: true},
		{name: "underscore", value: "a_b", valid: true},
		{name: "unlisted punctuation", value: "?/", valid: false},
	}

	testCharacterCheck(t, ContainsSpecialCharacter, yav.CheckNameContainsSpecialCharacter, tests)
}

func TestExcludesWhitespace(t *testing.T) {
	tests := []characterTest{
		{name: "empty", value: "", valid: true},
		{name: "word", value: "hello世界", valid: true},
		{name: "space", value: "hello world", valid: false},
		{name: "tab", value: "a\tb", valid: false},
		{name: "Unicode space", value: "a\u00A0b", valid: false},
	}

	testCharacterCheck(t, ExcludesWhitespace, yav.CheckNameExcludesWhitespace, tests)
}

func TestStartsWithAlpha(t *testing.T) {
	tests := []characterTest{
		{name: "empty", value: "", valid: false},
		{name: "single character", value: "a", valid: true},
		{name: "sample", value: "Z123", valid: true},
		{name: "wrong end", value: "0abc", valid: false},
		{name: "Unicode", value: "éabc", valid: false},
	}

	testCharacterCheck(t, StartsWithAlpha, yav.CheckNameStartsWithAlpha, tests)
}

func TestStartsWithLowerAlpha(t *testing.T) {
	tests := []characterTest{
		{name: "empty", value: "", valid: false},
		{name: "single character", value: "z", valid: true},
		{name: "sample", value: "a123", valid: true},
		{name: "uppercase", value: "Az", valid: false},
		{name: "Unicode", value: "éabc", valid: false},
	}

	testCharacterCheck(t, StartsWithLowerAlpha, yav.CheckNameStartsWithLowerAlpha, tests)
}

func TestStartsWithUpperAlpha(t *testing.T) {
	tests := []characterTest{
		{name: "empty", value: "", valid: false},
		{name: "single character", value: "Z", valid: true},
		{name: "sample", value: "A123", valid: true},
		{name: "lowercase", value: "aZ", valid: false},
		{name: "Unicode", value: "Éabc", valid: false},
	}

	testCharacterCheck(t, StartsWithUpperAlpha, yav.CheckNameStartsWithUpperAlpha, tests)
}

func TestStartsWithDigit(t *testing.T) {
	tests := []characterTest{
		{name: "empty", value: "", valid: false},
		{name: "single character", value: "0", valid: true},
		{name: "sample", value: "9abc", valid: true},
		{name: "wrong end", value: "a1", valid: false},
		{name: "Unicode digit", value: "١abc", valid: false},
	}

	testCharacterCheck(t, StartsWithDigit, yav.CheckNameStartsWithDigit, tests)
}

func TestStartsWithSpecialCharacter(t *testing.T) {
	tests := []characterTest{
		{name: "empty", value: "", valid: false},
		{name: "single character", value: "!", valid: true},
		{name: "sample", value: "_abc", valid: true},
		{name: "wrong end", value: "abc!", valid: false},
		{name: "unlisted punctuation", value: "?abc", valid: false},
	}

	testCharacterCheck(t, StartsWithSpecialCharacter, yav.CheckNameStartsWithSpecialCharacter, tests)
}

func TestEndsWithAlpha(t *testing.T) {
	tests := []characterTest{
		{name: "empty", value: "", valid: false},
		{name: "single character", value: "z", valid: true},
		{name: "sample", value: "123A", valid: true},
		{name: "wrong end", value: "abc0", valid: false},
		{name: "Unicode", value: "abcé", valid: false},
	}

	testCharacterCheck(t, EndsWithAlpha, yav.CheckNameEndsWithAlpha, tests)
}

func TestEndsWithLowerAlpha(t *testing.T) {
	tests := []characterTest{
		{name: "empty", value: "", valid: false},
		{name: "single character", value: "a", valid: true},
		{name: "sample", value: "123z", valid: true},
		{name: "uppercase", value: "zA", valid: false},
		{name: "Unicode", value: "abcé", valid: false},
	}

	testCharacterCheck(t, EndsWithLowerAlpha, yav.CheckNameEndsWithLowerAlpha, tests)
}

func TestEndsWithUpperAlpha(t *testing.T) {
	tests := []characterTest{
		{name: "empty", value: "", valid: false},
		{name: "single character", value: "A", valid: true},
		{name: "sample", value: "123Z", valid: true},
		{name: "lowercase", value: "Za", valid: false},
		{name: "Unicode", value: "abcÉ", valid: false},
	}

	testCharacterCheck(t, EndsWithUpperAlpha, yav.CheckNameEndsWithUpperAlpha, tests)
}

func TestEndsWithDigit(t *testing.T) {
	tests := []characterTest{
		{name: "empty", value: "", valid: false},
		{name: "single character", value: "9", valid: true},
		{name: "sample", value: "abc0", valid: true},
		{name: "wrong end", value: "1a", valid: false},
		{name: "Unicode digit", value: "abc١", valid: false},
	}

	testCharacterCheck(t, EndsWithDigit, yav.CheckNameEndsWithDigit, tests)
}

func TestEndsWithSpecialCharacter(t *testing.T) {
	tests := []characterTest{
		{name: "empty", value: "", valid: false},
		{name: "single character", value: "-", valid: true},
		{name: "sample", value: "abc_", valid: true},
		{name: "wrong end", value: "!abc", valid: false},
		{name: "unlisted punctuation", value: "abc?", valid: false},
	}

	testCharacterCheck(t, EndsWithSpecialCharacter, yav.CheckNameEndsWithSpecialCharacter, tests)
}

func TestText(t *testing.T) {
	tests := []characterTest{
		{name: "empty", value: "", valid: true},
		{name: "multiline Unicode", value: "Hello, 世界!\r\nNext line.", valid: true},
		{name: "Unicode space", value: "a\u00A0b", valid: true},
		{name: "control", value: "\u0007", valid: false},
		{name: "tab", value: "\t", valid: false},
	}

	testCharacterCheck(t, Text, yav.CheckNameText, tests)
}
