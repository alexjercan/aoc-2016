package day05

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strings"
)

func GetMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

func Solve1(doorID string) string {
	password := ""
	cnt := 0
	i := 0

	for {
		hash := GetMD5Hash(doorID + fmt.Sprint(i))
		if strings.HasPrefix(hash, "00000") {
			password += string(hash[5])
			cnt++
		}
		i++

		if cnt == 8 {
			break
		}
	}

	return password
}

func Solve2(doorID string) string {
	password := make(map[byte]byte)

	i := 0
	for {
		hash := GetMD5Hash(doorID + fmt.Sprint(i))

		if strings.HasPrefix(hash, "00000") {
			pos := hash[5]
			if pos >= '0' && pos <= '7' {
				if _, ok := password[pos]; !ok {
					password[pos] = hash[6]
				}
			}
		}

		i++

		if len(password) == 8 {
			break
		}
	}

	result := ""
	for i := '0'; i <= '8'; i++ {
		result += string(password[byte(i)])
	}

	return result
}

func Solve(input string) string {
	text := strings.TrimSpace(input)

	return "Day05\nPart1: " + Solve1(text) + "\nPart2: " + Solve2(text)
}
