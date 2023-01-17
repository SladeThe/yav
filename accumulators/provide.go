package accumulators

import (
	"github.com/SladeThe/yav"
)

type ProvideFunc[T any] func(names string, enabled bool) yav.ValidateFunc[T]
