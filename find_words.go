package bench

// my entry for the gobench "Find words" competition:
// https://github.com/gobench/competitions/tree/master/00000001
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

	lines, err := readLines(path)
	if err != nil {
		return "", err
	}

	T := kmpBuildTable(s)

	var matchesByLine [][]int
	matchesByLine = make([][]int, len(lines))

	// find the matches as int offsets in each line
	for row, line := range lines {
		matchesByLine[row] = kmpSearch(T, s, line)
	}

	// join the matches together into a comma-separated string
	result := ""
	sep    := ""
	for row, matches := range matchesByLine {
		for _, col := range matches {
			result = result + fmt.Sprintf("%s%d:%d", sep, row + 1, col)
			sep = ","
		}
	}

	return result, nil
}

// Knuth-Morris-Pratt algorithm, modified slightly to return all occurrences
// via: http://en.wikipedia.org/wiki/Knuth–Morris–Pratt_algorithm
func kmpSearch(T []int, word, line string) []int {
	var result []int

	m := 0
	i := 0

	for ; m + i < len(line); {
		if word[i] == line[m + i] {
			if i == len(word) - 1 {
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

	for ; pos < len(word); {
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


// read the file at path and return as array of lines
func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var result []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		result = append(result, scanner.Text())
	}
	return result, scanner.Err()
}
