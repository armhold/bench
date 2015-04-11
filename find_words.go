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

type Job struct {
	Line     string // the line of text to search
	RowIndex int    // the (zero-based) index into the matchesByLine array where the job should write the result
}

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

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Bytes()
		matches := kmpSearch(T, sBytes, line)
		matchesByLine = append(matchesByLine, matches)

	}

	// join the matches together into a comma-separated string
	result := ""
	sep := ""
	for row, matches := range matchesByLine {
		for _, col := range matches {
			result = result + fmt.Sprintf("%s%d:%d", sep, row+1, col)
			sep = ","
		}
	}

	return result, nil
}

// Knuth-Morris-Pratt algorithm, modified slightly to return all occurrences
// via: http://en.wikipedia.org/wiki/Knuth–Morris–Pratt_algorithm
func kmpSearch(T []int, word, line []byte) []int {
	result := make([]int, 0, 0)

	m := 0
	i := 0

	for m+i < len(line) {
		if word[i] == line[m+i] {
			if i == len(word)-1 {
				// got a match
				result = append(result, m)
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

	return result
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
