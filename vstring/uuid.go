package vstring

import (
	"github.com/SladeThe/yav"
)

const (
	uuidLength = 36
)

var (
	uuidPartLengths = [5]int{8, 4, 4, 4, 12}
)

func UUID(name string, value string) (stop bool, err error) {
	if !isUUID(value) {
		return true, yav.Error{
			CheckName: yav.CheckNameUUID,
			ValueName: name,
			Value:     value,
		}
	}

	return false, nil
}

func isUUID(value string) bool {
	if len(value) != uuidLength {
		return false
	}

	partIndex := 0
	partLength := uuidPartLengths[0]

	for _, r := range []byte(value) {
		if partLength > 0 {
			if (r < '0' || r > '9') && (r < 'a' || r > 'f') {
				return false
			}

			partLength--
			continue
		}

		if r != '-' {
			return false
		}

		partIndex++
		partLength = uuidPartLengths[partIndex]
	}

	return true
}
