package vtime

import (
	"time"

	"github.com/SladeThe/yav"
)

func Min(parameter time.Time) yav.ValidateFunc[time.Time] {
	return min(parameter).validate
}

func Max(parameter time.Time) yav.ValidateFunc[time.Time] {
	return max(parameter).validate
}

func Between(min, max time.Time) yav.ValidateFunc[time.Time] {
	return between{min: min, max: max}.validate
}

func LessThan(parameter time.Time) yav.ValidateFunc[time.Time] {
	return lessThan(parameter).validate
}

func LessThanOrEqual(parameter time.Time) yav.ValidateFunc[time.Time] {
	return lessThanOrEqual(parameter).validate
}

func GreaterThan(parameter time.Time) yav.ValidateFunc[time.Time] {
	return greaterThan(parameter).validate
}

func GreaterThanOrEqual(parameter time.Time) yav.ValidateFunc[time.Time] {
	return greaterThanOrEqual(parameter).validate
}

type min time.Time

func (m min) validate(name string, value time.Time) (stop bool, err error) {
	if value.Before(time.Time(m)) {
		return true, yav.Error{
			CheckName: yav.CheckNameMin,
			Parameter: time.Time(m).String(),
			ValueName: name,
			Value:     value,
		}
	}

	return false, nil
}

type max time.Time

func (m max) validate(name string, value time.Time) (stop bool, err error) {
	if value.After(time.Time(m)) {
		return true, yav.Error{
			CheckName: yav.CheckNameMax,
			Parameter: time.Time(m).String(),
			ValueName: name,
			Value:     value,
		}
	}

	return false, nil
}

type between struct {
	min, max time.Time
}

func (b between) validate(name string, value time.Time) (stop bool, err error) {
	if value.Before(b.min) {
		return true, yav.Error{
			CheckName: yav.CheckNameMin,
			Parameter: b.min.String(),
			ValueName: name,
			Value:     value,
		}
	}

	if value.After(b.max) {
		return true, yav.Error{
			CheckName: yav.CheckNameMax,
			Parameter: b.max.String(),
			ValueName: name,
			Value:     value,
		}
	}

	return false, nil
}

type lessThan time.Time

func (l lessThan) validate(name string, value time.Time) (stop bool, err error) {
	if !value.Before(time.Time(l)) {
		return true, yav.Error{
			CheckName: yav.CheckNameLessThan,
			Parameter: time.Time(l).String(),
			ValueName: name,
			Value:     value,
		}
	}

	return false, nil
}

type lessThanOrEqual time.Time

func (l lessThanOrEqual) validate(name string, value time.Time) (stop bool, err error) {
	if value.After(time.Time(l)) {
		return true, yav.Error{
			CheckName: yav.CheckNameLessThanOrEqual,
			Parameter: time.Time(l).String(),
			ValueName: name,
			Value:     value,
		}
	}

	return false, nil
}

type greaterThan time.Time

func (g greaterThan) validate(name string, value time.Time) (stop bool, err error) {
	if !value.After(time.Time(g)) {
		return true, yav.Error{
			CheckName: yav.CheckNameGreaterThan,
			Parameter: time.Time(g).String(),
			ValueName: name,
			Value:     value,
		}
	}

	return false, nil
}

type greaterThanOrEqual time.Time

func (g greaterThanOrEqual) validate(name string, value time.Time) (stop bool, err error) {
	if value.Before(time.Time(g)) {
		return true, yav.Error{
			CheckName: yav.CheckNameGreaterThanOrEqual,
			Parameter: time.Time(g).String(),
			ValueName: name,
			Value:     value,
		}
	}

	return false, nil
}
