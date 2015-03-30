package bench

// my "naive" entry for the gobench "Find words" competition:
// https://github.com/gobench/competitions/tree/master/00000001

import (
	"errors"
	"fmt"
	"strings"
)

func FindNaive(path, s string) (string, error) {
	if s == "" {
		return "", errors.New("s cannot be empty")
	}

	lines, err := readLines(path)
	if err != nil {
		return "", err
	}

	var result []string

	for i, line := range lines {
		result = append(result, naiveFindInLine(s, line, i)...)
	}

	return strings.Join(result, ","), nil
}

// naive version that just iterates the string with no backtrack-prevention
func naiveFindInLine(s, line string, lineNum int) []string {
	var result []string

	row := lineNum + 1

	for col := 0; col < len(line); col++ {
		if matchAtOffset(s, line, col) {
			result = append(result, fmt.Sprintf("%d:%d", row, col))
		}
	}

	return result
}

// return true if the pattern s exists in line at the given offset
func matchAtOffset(s, line string, offset int) bool {
	// past the end
	if offset+len(s) > len(line) {
		return false
	}

	for i := 0; i < len(s); i++ {
		if s[i] != line[offset+i] {
			return false
		}
	}

	return true
}
