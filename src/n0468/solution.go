package n0468

import (
	"strconv"
	"strings"
)

func validIPAddress(queryIP string) string {
	if checkIPv4(queryIP) {
		return "IPv4"
	} else if checkIPv6(queryIP) {
		return "IPv6"
	} else {
		return "Neither"
	}
}

func checkIPv4(query string) bool {
	ips := strings.Split(query, ".")
	if len(ips) != 4 {
		return false
	}
	for _, ip := range ips {
		intVal, err := strconv.Atoi(ip)
		if err != nil || intVal > 255 || intVal < 0 {
			return false
		}
		if ip == "0" {
			continue
		}
		count := 0
		for intVal > 0 {
			count++
			intVal /= 10
		}
		if count != len(ip) {
			return false
		}
	}
	return true
}

func checkIPv6(query string) bool {
	ips := strings.Split(query, ":")
	if len(ips) != 8 {
		return false
	}
	for _, ip := range ips {
		if len(ip) > 4 || len(ip) < 1 {
			return false
		}
		for i := range ip {
			if (ip[i] >= '0' && ip[i] <= '9') || (ip[i] >= 'a' && ip[i] <= 'f') || (ip[i] >= 'A' && ip[i] <= 'F') {
				continue
			} else {
				return false
			}
		}
	}
	return true
}