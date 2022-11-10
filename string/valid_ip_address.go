package string

import (
	"strings"
	"unicode"
)

// ValidIPAddress
// @summary: LeetCode #468 (Medium); CodeTop MS
// @author : binqibang
// @created: 2022/11/10 15:21
func ValidIPAddress(queryIP string) string {
	if IsIPv4(queryIP) {
		return "IPv4"
	} else if IsIPv6(queryIP) {
		return "IPv6"
	} else {
		return "Neither"
	}
}

func IsIPv6(ip string) bool {
	const IPv6Len = 8
	segments := strings.Split(ip, ":")

	// A valid IPv6 address is an IP in the form "x1:x2:x3:x4:x5:x6:x7:x8" where:
	//  1) 1 <= xi.length <= 4
	//  2) xi is a hexadecimal string which may contain digits, lowercase English letter
	//     ('a' to 'f') and upper-case English letters ('A' to 'F').
	if len(segments) != IPv6Len {
		return false
	}
	for i := 0; i < IPv6Len; i++ {
		seg := segments[i]
		if len(seg) < 1 || len(seg) > 4 {
			return false
		}
		for j := 0; j < len(seg); j++ {
			ch := seg[j]
			if (ch < '0' || ch > '9') && (ch < 'a' || ch > 'f') &&
				(ch < 'A' || ch > 'F') {
				return false
			}
		}
	}
	return true
}

func IsIPv4(ip string) bool {
	const IPv4Len = 4
	segments := strings.Split(ip, ".")

	// A valid IPv4 address is an IP in the form "x1.x2.x3.x4"
	// where 0 <= xi <= 255 and xi cannot contain leading zeros.
	if len(segments) != IPv4Len {
		return false
	}
	for i := 0; i < IPv4Len; i++ {
		seg := segments[i]
		if len(seg) < 1 || len(seg) > 3 {
			return false
		}
		if len(seg) > 1 && seg[0] == '0' {
			return false
		}
		num := 0
		for _, ch := range seg {
			if !unicode.IsDigit(ch) {
				return false
			}
			num = num*10 + int(ch-'0')
		}
		if num > 255 {
			return false
		}
	}
	return true
}
