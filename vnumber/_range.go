package vnumber

import (
	"fmt"

	"github.com/cheekybits/genny/generic"

	"github.com/SladeThe/yav"
)

type Element generic.Type

func MinElement(parameter Element) yav.ValidateFunc[Element] {
	return minimumElement{parameter: parameter}.validate
}

func MaxElement(parameter Element) yav.ValidateFunc[Element] {
	return maximumElement{parameter: parameter}.validate
}

func BetweenElement(minimum, maximum Element) yav.ValidateFunc[Element] {
	return betweenElement{min: minimum, max: maximum}.validate
}

func LessThanElement(parameter Element) yav.ValidateFunc[Element] {
	return lessThanElement{parameter: parameter}.validate
}

func LessThanOrEqualElement(parameter Element) yav.ValidateFunc[Element] {
	return lessThanOrEqualElement{parameter: parameter}.validate
}

func GreaterThanElement(parameter Element) yav.ValidateFunc[Element] {
	return greaterThanElement{parameter: parameter}.validate
}

func GreaterThanOrEqualElement(parameter Element) yav.ValidateFunc[Element] {
	return greaterThanOrEqualElement{parameter: parameter}.validate
}

type minimumElement struct {
	parameter Element
}

func (m minimumElement) validate(name string, value Element) (stop bool, err error) {
	if !(value >= m.parameter) {
		return true, yav.Error{
			CheckName: yav.CheckNameMin,
			Parameter: fmt.Sprintf("%v", m.parameter),
			ValueName: name,
			Value:     value,
		}
	}

	return false, nil
}

type maximumElement struct {
	parameter Element
}

func (m maximumElement) validate(name string, value Element) (stop bool, err error) {
	if !(value <= m.parameter) {
		return true, yav.Error{
			CheckName: yav.CheckNameMax,
			Parameter: fmt.Sprintf("%v", m.parameter),
			ValueName: name,
			Value:     value,
		}
	}

	return false, nil
}

type betweenElement struct {
	min, max Element
}

func (b betweenElement) validate(name string, value Element) (stop bool, err error) {
	if !(value >= b.min) {
		return true, yav.Error{
			CheckName: yav.CheckNameMin,
			Parameter: fmt.Sprintf("%v", b.min),
			ValueName: name,
			Value:     value,
		}
	}

	if !(value <= b.max) {
		return true, yav.Error{
			CheckName: yav.CheckNameMax,
			Parameter: fmt.Sprintf("%v", b.max),
			ValueName: name,
			Value:     value,
		}
	}

	return false, nil
}

type lessThanElement struct {
	parameter Element
}

func (l lessThanElement) validate(name string, value Element) (stop bool, err error) {
	if !(value < l.parameter) {
		return true, yav.Error{
			CheckName: yav.CheckNameLessThan,
			Parameter: fmt.Sprintf("%v", l.parameter),
			ValueName: name,
			Value:     value,
		}
	}

	return false, nil
}

type lessThanOrEqualElement struct {
	parameter Element
}

func (l lessThanOrEqualElement) validate(name string, value Element) (stop bool, err error) {
	if !(value <= l.parameter) {
		return true, yav.Error{
			CheckName: yav.CheckNameLessThanOrEqual,
			Parameter: fmt.Sprintf("%v", l.parameter),
			ValueName: name,
			Value:     value,
		}
	}

	return false, nil
}

type greaterThanElement struct {
	parameter Element
}

func (g greaterThanElement) validate(name string, value Element) (stop bool, err error) {
	if !(value > g.parameter) {
		return true, yav.Error{
			CheckName: yav.CheckNameGreaterThan,
			Parameter: fmt.Sprintf("%v", g.parameter),
			ValueName: name,
			Value:     value,
		}
	}

	return false, nil
}

type greaterThanOrEqualElement struct {
	parameter Element
}

func (g greaterThanOrEqualElement) validate(name string, value Element) (stop bool, err error) {
	if !(value >= g.parameter) {
		return true, yav.Error{
			CheckName: yav.CheckNameGreaterThanOrEqual,
			Parameter: fmt.Sprintf("%v", g.parameter),
			ValueName: name,
			Value:     value,
		}
	}

	return false, nil
}
