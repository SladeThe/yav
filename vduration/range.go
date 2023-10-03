package vduration

import (
	"time"

	"github.com/SladeThe/yav"
)

func Min(parameter time.Duration) yav.ValidateFunc[time.Duration] {
	return min(parameter).validate
}

func Max(parameter time.Duration) yav.ValidateFunc[time.Duration] {
	return max(parameter).validate
}

func Between(min, max time.Duration) yav.ValidateFunc[time.Duration] {
	return between{min: min, max: max}.validate
}

func LessThan(parameter time.Duration) yav.ValidateFunc[time.Duration] {
	return lessThan(parameter).validate
}

func LessThanOrEqual(parameter time.Duration) yav.ValidateFunc[time.Duration] {
	return lessThanOrEqual(parameter).validate
}

func GreaterThan(parameter time.Duration) yav.ValidateFunc[time.Duration] {
	return greaterThan(parameter).validate
}

func GreaterThanOrEqual(parameter time.Duration) yav.ValidateFunc[time.Duration] {
	return greaterThanOrEqual(parameter).validate
}

type min time.Duration

func (m min) validate(name string, value time.Duration) (stop bool, err error) {
	if value < time.Duration(m) {
		return true, yav.Error{
			CheckName: yav.CheckNameMin,
			Parameter: time.Duration(m).String(),
			ValueName: name,
			Value:     value,
		}
	}

	return false, nil
}

type max time.Duration

func (m max) validate(name string, value time.Duration) (stop bool, err error) {
	if value > time.Duration(m) {
		return true, yav.Error{
			CheckName: yav.CheckNameMax,
			Parameter: time.Duration(m).String(),
			ValueName: name,
			Value:     value,
		}
	}

	return false, nil
}

type between struct {
	min, max time.Duration
}

func (b between) validate(name string, value time.Duration) (stop bool, err error) {
	if value < b.min {
		return true, yav.Error{
			CheckName: yav.CheckNameMin,
			Parameter: b.min.String(),
			ValueName: name,
			Value:     value,
		}
	}

	if value > b.max {
		return true, yav.Error{
			CheckName: yav.CheckNameMax,
			Parameter: b.max.String(),
			ValueName: name,
			Value:     value,
		}
	}

	return false, nil
}

type lessThan time.Duration

func (l lessThan) validate(name string, value time.Duration) (stop bool, err error) {
	if value >= time.Duration(l) {
		return true, yav.Error{
			CheckName: yav.CheckNameLessThan,
			Parameter: time.Duration(l).String(),
			ValueName: name,
			Value:     value,
		}
	}

	return false, nil
}

type lessThanOrEqual time.Duration

func (l lessThanOrEqual) validate(name string, value time.Duration) (stop bool, err error) {
	if value > time.Duration(l) {
		return true, yav.Error{
			CheckName: yav.CheckNameLessThanOrEqual,
			Parameter: time.Duration(l).String(),
			ValueName: name,
			Value:     value,
		}
	}

	return false, nil
}

type greaterThan time.Duration

func (g greaterThan) validate(name string, value time.Duration) (stop bool, err error) {
	if value <= time.Duration(g) {
		return true, yav.Error{
			CheckName: yav.CheckNameGreaterThan,
			Parameter: time.Duration(g).String(),
			ValueName: name,
			Value:     value,
		}
	}

	return false, nil
}

type greaterThanOrEqual time.Duration

func (g greaterThanOrEqual) validate(name string, value time.Duration) (stop bool, err error) {
	if value < time.Duration(g) {
		return true, yav.Error{
			CheckName: yav.CheckNameGreaterThanOrEqual,
			Parameter: time.Duration(g).String(),
			ValueName: name,
			Value:     value,
		}
	}

	return false, nil
}
