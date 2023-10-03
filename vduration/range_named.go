package vduration

import (
	"time"

	"github.com/SladeThe/yav"
)

func LessThanNamed(name string, parameter time.Duration) yav.ValidateFunc[time.Duration] {
	return lessThanNamed{name: name, parameter: parameter}.validate
}

func LessThanOrEqualNamed(name string, parameter time.Duration) yav.ValidateFunc[time.Duration] {
	return lessThanOrEqualNamed{name: name, parameter: parameter}.validate
}

func GreaterThanNamed(name string, parameter time.Duration) yav.ValidateFunc[time.Duration] {
	return greaterThanNamed{name: name, parameter: parameter}.validate
}

func GreaterThanOrEqualNamed(name string, parameter time.Duration) yav.ValidateFunc[time.Duration] {
	return greaterThanOrEqualNamed{name: name, parameter: parameter}.validate
}

type lessThanNamed struct {
	name      string
	parameter time.Duration
}

func (l lessThanNamed) validate(name string, value time.Duration) (stop bool, err error) {
	if value >= l.parameter {
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
	parameter time.Duration
}

func (l lessThanOrEqualNamed) validate(name string, value time.Duration) (stop bool, err error) {
	if value > l.parameter {
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
	parameter time.Duration
}

func (g greaterThanNamed) validate(name string, value time.Duration) (stop bool, err error) {
	if value <= g.parameter {
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
	parameter time.Duration
}

func (g greaterThanOrEqualNamed) validate(name string, value time.Duration) (stop bool, err error) {
	if value < g.parameter {
		return true, yav.Error{
			CheckName: yav.CheckNameGreaterThanOrEqualNamed,
			Parameter: g.name,
			ValueName: name,
			Value:     value,
		}
	}

	return false, nil
}
