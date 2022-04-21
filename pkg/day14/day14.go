package day14

import (
	"crypto/md5"
	"encoding/hex"
	"strconv"
	"strings"

	"github.com/alexjercan/aoc-2016/pkg/util"
)

type HashData struct {
	three []byte
	five  []byte
}

func hash(salt string, i int) string {
	hash := md5.Sum([]byte(salt + strconv.Itoa(i)))
	return hex.EncodeToString(hash[:])
}

func hashN(salt string, i int, n int) string {
	hash := md5.Sum([]byte(salt + strconv.Itoa(i)))

	for j := 0; j < n-1; j++ {
		hash = md5.Sum([]byte(hex.EncodeToString(hash[:])))
	}

	return hex.EncodeToString(hash[:])
}

func getThreeConsecutive(s string) []byte {
	var result []byte

	for i := 0; i < len(s)-2; i++ {
		if s[i] == s[i+1] && s[i] == s[i+2] {
			result = append(result, s[i])
			break
		}
	}

	return result
}

func getFiveConsecutive(s string) []byte {
	var result []byte

	for i := 0; i < len(s)-4; i++ {
		if s[i] == s[i+1] && s[i] == s[i+2] && s[i] == s[i+3] && s[i] == s[i+4] {
			result = append(result, s[i])
		}
	}

	return result
}

func solve1(salt string) int {
	hashes := make([]HashData, 0)
	indices := make([]int, 0)

	for i := 0; i < 1001; i++ {
		hash := hash(salt, i)
		hashes = append(hashes, HashData{getThreeConsecutive(hash), getFiveConsecutive(hash)})
	}

	j := 0
	i := 1001

	for {
		for h := j + 1; h < j+1001; h++ {
			if util.IntersectNotNull(hashes[j].three, hashes[h].five) {
				indices = append(indices, j)
			}
		}

		if len(indices) > 100 {
			break
		}

		hash := hash(salt, i)
		hashes = append(hashes, HashData{getThreeConsecutive(hash), getFiveConsecutive(hash)})

		i++
		j++
	}

	return indices[63]
}

func solve2(salt string) int {
	hashes := make([]HashData, 0)
	indices := make([]int, 0)

	for i := 0; i < 1001; i++ {
		hash := hashN(salt, i, 2017)
		hashes = append(hashes, HashData{getThreeConsecutive(hash), getFiveConsecutive(hash)})
	}

	j := 0
	i := 1001

	for {
		for h := j + 1; h < j+1001; h++ {
			if util.IntersectNotNull(hashes[j].three, hashes[h].five) {
				indices = append(indices, j)
			}
		}

		if len(indices) > 100 {
			break
		}

		hash := hashN(salt, i, 2017)
		hashes = append(hashes, HashData{getThreeConsecutive(hash), getFiveConsecutive(hash)})

		i++
		j++
	}

	return indices[63]
}

func Solve(input string) string {
	salt := strings.TrimSpace(input)

	return "Day14\nPart1: " + strconv.Itoa(solve1(salt)) + "\nPart2: " + strconv.Itoa(solve2(salt))
}
