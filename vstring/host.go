package vstring

import (
	"net"
	"strconv"
	"strings"

	"github.com/SladeThe/yav"
	"github.com/SladeThe/yav/common"
)

const (
	maxHostnameLength      = 253
	maxHostnameLabelLength = 63
)

func Hostname(name string, value string) (stop bool, err error) {
	if !isHostname(value) {
		return true, yav.Error{
			CheckName: yav.CheckNameHostname,
			ValueName: name,
			Value:     value,
		}
	}

	return false, nil
}

func HostnameRFC1123(name string, value string) (stop bool, err error) {
	if !isHostnameRFC1123(value) {
		return true, yav.Error{
			CheckName: yav.CheckNameHostnameRFC1123,
			ValueName: name,
			Value:     value,
		}
	}

	return false, nil
}

func HostnamePort(name string, value string) (stop bool, err error) {
	host, port, errSplit := net.SplitHostPort(value)
	if errSplit != nil {
		return true, errHostnamePort(name, value)
	}

	if portNumber, errPort := strconv.ParseUint(port, 10, 16); errPort != nil || portNumber == 0 {
		return true, errHostnamePort(name, value)
	}

	if host != "" && !isHostnameRFC1123(host) {
		return true, errHostnamePort(name, value)
	}

	return false, nil
}

func FQDN(name string, value string) (stop bool, err error) {
	if !isFQDN(value) {
		return true, yav.Error{
			CheckName: yav.CheckNameFQDN,
			ValueName: name,
			Value:     value,
		}
	}

	return false, nil
}

func errHostnamePort(name string, value string) yav.Error {
	return yav.Error{
		CheckName: yav.CheckNameHostnamePort,
		ValueName: name,
		Value:     value,
	}
}

// Equivalent regexp with len(value) <= maxHostnameLength:
// `^[a-zA-Z]([a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(\.[a-zA-Z0-9]([a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$`
func isHostname(value string) bool {
	return value != "" && common.IsAlpha(rune(value[0])) && isHostnameRFC1123(value)
}

// Equivalent regexp with len(value) <= maxHostnameLength:
// `^[a-zA-Z0-9]([a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(\.[a-zA-Z0-9]([a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$`
func isHostnameRFC1123(value string) bool {
	if value == "" || len(value) > maxHostnameLength {
		return false
	}

	labelLength := 0
	for i, c := range []byte(value) {
		if c == '.' {
			if labelLength == 0 || value[i-1] == '-' {
				return false
			}

			labelLength = 0
			continue
		}

		if c == '-' {
			if labelLength == 0 {
				return false
			}
		} else if !common.IsAlphanumeric(rune(c)) {
			return false
		}

		labelLength++
		if labelLength > maxHostnameLabelLength {
			return false
		}
	}

	return labelLength > 0 && value[len(value)-1] != '-'
}

// Equivalent regexp with len(strings.TrimSuffix(value, ".")) <= maxHostnameLength:
// `^([a-zA-Z0-9]([a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)(\.[a-zA-Z0-9]([a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*?(\.[a-zA-Z]([a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)\.?$`
func isFQDN(value string) bool {
	value = strings.TrimSuffix(value, ".")
	if !isHostnameRFC1123(value) {
		return false
	}

	tldIndex := strings.LastIndexByte(value, '.') + 1
	return tldIndex > 0 && common.IsAlpha(rune(value[tldIndex]))
}
