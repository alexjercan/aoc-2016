package day07

import (
	"strconv"
	"strings"
)

func SupportTLSIPv7(ipv7 string) bool {
	hasABBA := false
	insideSquareBrackets := false

	for i := 0; i < len(ipv7)-3; i++ {
		if insideSquareBrackets {
			if ipv7[i] == ']' {
				insideSquareBrackets = false
			} else {
				if ipv7[i] == ipv7[i+3] && ipv7[i+1] == ipv7[i+2] && ipv7[i] != ipv7[i+1] {
					return false
				}
			}
		} else {
			if ipv7[i] == '[' {
				insideSquareBrackets = true
			} else {
				if ipv7[i] == ipv7[i+3] && ipv7[i+1] == ipv7[i+2] && ipv7[i] != ipv7[i+1] {
					hasABBA = true
				}
			}
		}
	}

	return hasABBA
}

func SupportSSLIPv7(ipv7 string) bool {
	abas := make([]string, 0)
	babs := make([]string, 0)

	insideSquareBrackets := false

	for i := 0; i < len(ipv7)-2; i++ {
		if insideSquareBrackets {
			if ipv7[i] == ']' {
				insideSquareBrackets = false
			} else {
				if ipv7[i] == ipv7[i+2] && ipv7[i] != ipv7[i+1] {
					babs = append(babs, ipv7[i:i+3])
				}
			}
		} else {
			if ipv7[i] == '[' {
				insideSquareBrackets = true
			} else {
				if ipv7[i] == ipv7[i+2] && ipv7[i] != ipv7[i+1] {
					abas = append(abas, ipv7[i:i+3])
				}
			}
		}
	}

	for _, aba := range abas {
		for _, bab := range babs {
			if aba[0] == bab[1] && aba[1] == bab[0] {
				return true
			}
		}
	}

	return false
}

func Solve1(ips []string) int {
	cnt := 0

	for _, ip := range ips {
		if SupportTLSIPv7(ip) {
			cnt++
		}
	}

	return cnt
}

func Solve2(ips []string) int {
	cnt := 0

	for _, ip := range ips {
		if SupportSSLIPv7(ip) {
			cnt++
		}
	}

	return cnt
}

func Solve(input string) string {
	ips := strings.Split(strings.TrimSpace(input), "\n")

	return "Day07\nPart1: " + strconv.Itoa(Solve1(ips)) + "\nPart2: " + strconv.Itoa(Solve2(ips))
}
