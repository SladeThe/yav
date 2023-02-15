package vstring

import (
	"regexp"

	"github.com/SladeThe/yav"
)

// Regexp validates string value using regexp.Regexp.MatchString.
// That means the value must contain any match of the regular expression re.
// Wrap your expression with ^...$ to test the value for a full match.
func Regexp(re *regexp.Regexp) yav.ValidateFunc[string] {
	return _regexp{Regexp: re}.validate
}

type _regexp struct {
	*regexp.Regexp
}

func (re _regexp) validate(name string, value string) (stop bool, err error) {
	if !re.MatchString(value) {
		return true, yav.Error{
			CheckName: yav.CheckNameRegexp,
			Parameter: re.String(),
			ValueName: name,
			Value:     value,
		}
	}

	return false, nil
}
