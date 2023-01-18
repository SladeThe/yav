package accumulators

import (
	"github.com/SladeThe/yav"
)

type ProvideFunc[T any] func(names string, required bool) yav.ValidateFunc[T]
