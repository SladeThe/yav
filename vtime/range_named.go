package vtime

import (
	"time"

	"github.com/SladeThe/yav"
)

func LessThanNamed(name string, parameter time.Time) yav.ValidateFunc[time.Time] {
	return lessThanNamed{name: name, parameter: parameter}.validate
}

func LessThanOrEqualNamed(name string, parameter time.Time) yav.ValidateFunc[time.Time] {
	return lessThanOrEqualNamed{name: name, parameter: parameter}.validate
}

func GreaterThanNamed(name string, parameter time.Time) yav.ValidateFunc[time.Time] {
	return greaterThanNamed{name: name, parameter: parameter}.validate
}

func GreaterThanOrEqualNamed(name string, parameter time.Time) yav.ValidateFunc[time.Time] {
	return greaterThanOrEqualNamed{name: name, parameter: parameter}.validate
}

type lessThanNamed struct {
	name      string
	parameter time.Time
}

func (l lessThanNamed) validate(name string, value time.Time) (stop bool, err error) {
	if !value.Before(l.parameter) {
		return true, yav.Error{
			CheckName: yav.CheckNameLessThanNamed,
			Parameter: l.name,
			ValueName: name,
			Value:     value,
		}
	}

	return false, nil
}

type lessThanOrEqualNamed struct {
	name      string
	parameter time.Time
}

func (l lessThanOrEqualNamed) validate(name string, value time.Time) (stop bool, err error) {
	if value.After(l.parameter) {
		return true, yav.Error{
			CheckName: yav.CheckNameLessThanOrEqualNamed,
			Parameter: l.name,
			ValueName: name,
			Value:     value,
		}
	}

	return false, nil
}

type greaterThanNamed struct {
	name      string
	parameter time.Time
}

func (g greaterThanNamed) validate(name string, value time.Time) (stop bool, err error) {
	if !value.After(g.parameter) {
		return true, yav.Error{
			CheckName: yav.CheckNameGreaterThanNamed,
			Parameter: g.name,
			ValueName: name,
			Value:     value,
		}
	}

	return false, nil
}

type greaterThanOrEqualNamed struct {
	name      string
	parameter time.Time
}

func (g greaterThanOrEqualNamed) validate(name string, value time.Time) (stop bool, err error) {
	if value.Before(g.parameter) {
		return true, yav.Error{
			CheckName: yav.CheckNameGreaterThanOrEqualNamed,
			Parameter: g.name,
			ValueName: name,
			Value:     value,
		}
	}

	return false, nil
}
