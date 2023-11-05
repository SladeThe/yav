package main

import (
	"fmt"

	"github.com/SladeThe/yav"
	"github.com/SladeThe/yav/vmap"
	"github.com/SladeThe/yav/vnumber"
	"github.com/SladeThe/yav/vstring"
)

type Example struct {
	LengthByName map[string]int
}

func (e Example) Validate() error {
	return yav.Chain(
		"LengthByName", e.LengthByName,
		vmap.OmitEmpty[map[string]int],
		vmap.Keys[map[string]int](
			vstring.Required,
			vstring.Text,
			vstring.ExcludesWhitespace,
		),
		vmap.Values[map[string]int](
			vnumber.GreaterThanOrEqualInt(0),
		),
		vmap.Entries[map[string]int](
			func(name string, value vmap.Entry[string, int]) (stop bool, err error) {
				if len(value.Key) != value.Value {
					return true, yav.Error{
						CheckName: "key_match_value",
						ValueName: name,
						Value:     value,
					}
				}

				return false, nil
			},
		),
	)
}

func main() {
	valid := Example{
		LengthByName: map[string]int{
			"Jim":  3,
			"John": 4,
		},
	}

	invalid := Example{
		LengthByName: map[string]int{
			"":    0,
			"< >": 3,
			"Jim": 4,
		},
	}

	fmt.Println(valid.Validate())
	fmt.Println(invalid.Validate())
}
