package vnumeric

import (
	"fmt"

	"golang.org/x/exp/constraints"

	"github.com/SladeThe/yav"
)

// TODO get rid of allocations

func Min[T constraints.Integer](parameter T) yav.ValidateFunc[T] {
	return min[T]{parameter: parameter}.validate
}

func Max[T constraints.Integer](parameter T) yav.ValidateFunc[T] {
	return max[T]{parameter: parameter}.validate
}

func LessThan[T constraints.Integer](parameter T) yav.ValidateFunc[T] {
	return lessThan[T]{parameter: parameter}.validate
}

func LessThanOrEqual[T constraints.Integer](parameter T) yav.ValidateFunc[T] {
	return lessThanOrEqual[T]{parameter: parameter}.validate
}

func GreaterThan[T constraints.Integer](parameter T) yav.ValidateFunc[T] {
	return greaterThan[T]{parameter: parameter}.validate
}

func GreaterThanOrEqual[T constraints.Integer](parameter T) yav.ValidateFunc[T] {
	return greaterThanOrEqual[T]{parameter: parameter}.validate
}

type min[T constraints.Integer] struct {
	parameter T
}

func (m min[T]) validate(name string, value T) (stop bool, err error) {
	if value < m.parameter {
		return false, yav.Error{
			CheckName: yav.CheckNameMin,
			Parameter: fmt.Sprintf("%d", m.parameter),
			ValueName: name,
			Value:     value,
		}
	}

	return false, nil
}

type max[T constraints.Integer] struct {
	parameter T
}

func (m max[T]) validate(name string, value T) (stop bool, err error) {
	if value > m.parameter {
		return false, yav.Error{
			CheckName: yav.CheckNameMax,
			Parameter: fmt.Sprintf("%d", m.parameter),
			ValueName: name,
			Value:     value,
		}
	}

	return false, nil
}

type lessThan[T constraints.Integer] struct {
	parameter T
}

func (l lessThan[T]) validate(name string, value T) (stop bool, err error) {
	if value >= l.parameter {
		return false, yav.Error{
			CheckName: yav.CheckNameLessThan,
			Parameter: fmt.Sprintf("%d", l.parameter),
			ValueName: name,
			Value:     value,
		}
	}

	return false, nil
}

type lessThanOrEqual[T constraints.Integer] struct {
	parameter T
}

func (l lessThanOrEqual[T]) validate(name string, value T) (stop bool, err error) {
	if value > l.parameter {
		return false, yav.Error{
			CheckName: yav.CheckNameLessThanOrEqual,
			Parameter: fmt.Sprintf("%d", l.parameter),
			ValueName: name,
			Value:     value,
		}
	}

	return false, nil
}

type greaterThan[T constraints.Integer] struct {
	parameter T
}

func (l greaterThan[T]) validate(name string, value T) (stop bool, err error) {
	if value <= l.parameter {
		return false, yav.Error{
			CheckName: yav.CheckNameGreaterThan,
			Parameter: fmt.Sprintf("%d", l.parameter),
			ValueName: name,
			Value:     value,
		}
	}

	return false, nil
}

type greaterThanOrEqual[T constraints.Integer] struct {
	parameter T
}

func (l greaterThanOrEqual[T]) validate(name string, value T) (stop bool, err error) {
	if value < l.parameter {
		return false, yav.Error{
			CheckName: yav.CheckNameGreaterThanOrEqual,
			Parameter: fmt.Sprintf("%d", l.parameter),
			ValueName: name,
			Value:     value,
		}
	}

	return false, nil
}
