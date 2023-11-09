package main

import (
	"fmt"

	"github.com/SladeThe/yav"
	"github.com/SladeThe/yav/vbool"
	"github.com/SladeThe/yav/vnumber"
	"github.com/SladeThe/yav/vstring"
)

type Example struct {
	RequiredWithAny bool
	RequiredWithAll bool
	OptionalString  string
	OptionalInt     int
}

func (e Example) Validate() error {
	return yav.Join(
		yav.Chain(
			"RequiredWithAny", e.RequiredWithAny,
			vbool.RequiredWithAny().String(e.OptionalString).Int(e.OptionalInt).Names("OptionalString OptionalInt"),
		),
		yav.Chain(
			"RequiredWithAll", e.RequiredWithAll,
			vbool.RequiredWithAll().String(e.OptionalString).Int(e.OptionalInt).Names("OptionalString OptionalInt"),
		),
		yav.Chain(
			"OptionalString", e.OptionalString,
			vstring.Max(1),
		),
		yav.Chain(
			"OptionalInt", e.OptionalInt,
			vnumber.GreaterThanOrEqualInt(0),
		),
	)
}

func main() {
	valid := Example{
		RequiredWithAny: false,
		RequiredWithAll: false,
		OptionalString:  "",
		OptionalInt:     0,
	}

	invalidAny := Example{
		RequiredWithAny: false,
		RequiredWithAll: false,
		OptionalString:  "a",
		OptionalInt:     0,
	}

	invalidAll := Example{
		RequiredWithAny: true,
		RequiredWithAll: false,
		OptionalString:  "a",
		OptionalInt:     1,
	}

	fmt.Println(valid.Validate())
	fmt.Println(invalidAny.Validate())
	fmt.Println(invalidAll.Validate())
}
