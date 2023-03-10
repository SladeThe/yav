package vstring

import (
	"net"
	"regexp"
	"strconv"

	"github.com/SladeThe/yav"
)

var (
	rfc952Regex  = regexp.MustCompile(`^[a-zA-Z]([a-zA-Z0-9\-]+[.]?)*[a-zA-Z0-9]$`)
	rfc1123Regex = regexp.MustCompile(`^([a-zA-Z0-9][a-zA-Z0-9-]{0,62})(\.[a-zA-Z0-9][a-zA-Z0-9-]{0,62})*?$`)
	fqdnRegex    = regexp.MustCompile(`^([a-zA-Z0-9][a-zA-Z0-9-]{0,62})(\.[a-zA-Z0-9][a-zA-Z0-9-]{0,62})*?(\.[a-zA-Z][a-zA-Z0-9]{0,62})\.?$`)
)

func Hostname(name string, value string) (stop bool, err error) {
	if !rfc952Regex.MatchString(value) {
		return true, yav.Error{
			CheckName: yav.CheckNameHostname,
			ValueName: name,
			Value:     value,
		}
	}

	return false, nil
}

func HostnameRFC1123(name string, value string) (stop bool, err error) {
	if !rfc1123Regex.MatchString(value) {
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

	if _, errPort := strconv.ParseUint(port, 10, 16); errPort != nil {
		return true, errHostnamePort(name, value)
	}

	if host != "" && !rfc1123Regex.MatchString(host) {
		return true, errHostnamePort(name, value)
	}

	return false, nil
}

func FQDN(name string, value string) (stop bool, err error) {
	if !fqdnRegex.MatchString(value) {
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
