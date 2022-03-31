package main

import (
	"strconv"
	"strings"
)

// 468.验证IP地址
func validIPAddress(IP string) string {
	v4Count := strings.Count(IP, ".")
	v6Count := strings.Count(IP, ":")
	if v4Count > 0 && v6Count > 0 {
		return "Neither"
	}

	if v4Count > 0 {
		return validIP4(IP)
	}

	if v6Count > 0 {
		return validIP6(IP)
	}

	return "Neither"
}

func validIP4(ip string) string {
	strs := strings.Split(ip, ".")
	if len(strs) != 4 {
		return "Neither"
	}

	for _, str := range strs {
		if len(str) == 0 {
			return "Neither"
		}

		if len(str) > 1 && str[0] == '0' {
			return "Neither"
		}

		n, err := strconv.Atoi(str)
		if err != nil {
			return "Neither"
		}

		if n >= 256 || n < 0 {
			return "Neither"
		}
	}

	return "IPv4"
}

func validIP6(ip string) string {
	strs := strings.Split(ip, ":")
	if len(strs) != 8 {
		return "Neither"
	}

	for _, str := range strs {
		if len(str) == 0 || len(str) > 4 {
			return "Neither"
		}
		for j := 0; j < len(str); j++ {
			if str[j] >= 'g' && str[j] <= 'z' || str[j] >= 'G' && str[j] <= 'Z' {
				return "Neither"
			}
		}
	}

	return "IPv6"
}
