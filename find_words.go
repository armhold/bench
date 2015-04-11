package bench

// my entry for the gobench "Find words" competition:
// https://github.com/gobench/competitions/tree/master/00000001
//
// benchmark with:  $ go test -bench . -benchmem
//
// George Armhold, March 2015

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

func Find(path, s string) (string, error) {
	if s == "" {
		return "", errors.New("s cannot be empty")
	}

	T := kmpBuildTable(s)
	sBytes := []byte(s)

	file, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer file.Close()

	var matchesByLine [][]int

	searchResultBuffer := make([]int, 0, 5000000)
	result := ""
	sep := ""
	row := 1

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Bytes()
		lineMatchCount := kmpSearch(T, sBytes, line, searchResultBuffer)
		matches := searchResultBuffer[0:lineMatchCount]

		for _, col := range matches {
			result = result + fmt.Sprintf("%s%d:%d", sep, row, col)
			sep = ","
		}

		matchesByLine = append(matchesByLine, matches)
		row++
	}

	return result, nil
}

// Knuth-Morris-Pratt algorithm, modified slightly to return all occurrences
// via: http://en.wikipedia.org/wiki/Knuth–Morris–Pratt_algorithm
func kmpSearch(T []int, word, line []byte, result []int) int {
	m := 0
	i := 0

	matchCount := 0

	for m+i < len(line) {
		if word[i] == line[m+i] {
			if i == len(word)-1 {
				// got a match
				result = append(result, m)
				matchCount++
				m = m + i
				i = 0
			} else {
				i++
			}

		} else {
			if T[i] > -1 {
				m = m + i - T[i]
				i = T[i]
			} else {
				i = 0
				m++
			}
		}
	}

	return matchCount
}

// builds the table "T" for Knuth-Morris-Pratt string search
// via: http://en.wikipedia.org/wiki/Knuth–Morris–Pratt_algorithm
func kmpBuildTable(word string) []int {
	T := make([]int, len(word))

	pos := 2
	cnd := 0

	T[0] = -1
	T[1] = 0

	for pos < len(word) {
		if word[pos-1] == word[cnd] {
			cnd++
			T[pos] = cnd
			pos++
		} else if cnd > 0 {
			cnd = T[cnd]
		} else {
			T[pos] = 0
			pos++
		}
	}

	return T
}
