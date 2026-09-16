package internal

//nolint:gocritic // It is necessary to have a pointer here.
func RegisterMapEntry[K comparable, V any](m *map[K]V, key K, value V) V {
	// NaN keys cannot be found by a map lookup.
	// noinspection GoBoolExpressions,GoDfaConstantCondition
	if key != key {
		return value
	}

	src := *m

	if _, ok := src[key]; ok {
		return value
	}

	dst := make(map[K]V, len(src)+1)

	for k, v := range src {
		dst[k] = v
	}

	dst[key] = value
	*m = dst
	return value
}
