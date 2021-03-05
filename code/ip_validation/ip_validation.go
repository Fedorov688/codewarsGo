package ip_validation

import (
	"strconv"
	"strings"
)

func Is_valid_ip(ip string) bool {
	var octet = strings.Split(ip, ".")
	if len(octet) != 4 {
		return false
	}
	for _, values := range octet {
		if s, err := strconv.Atoi(values); err == nil {
			if s > 255 {
				return false
			}
		}
		for key, value := range values {
			if value < 48 || value > 57 {
				return false
			} else if key == 0 && value == 48 && len(values) > 1 {
				return false
			}
		}
	}
	return true
}
