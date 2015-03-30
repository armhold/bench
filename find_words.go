package bench

// my "naive" entry for the gobench "Find words" competition:
// https://github.com/gobench/competitions/tree/master/00000001

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
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

	var result []string

	for i, line := range lines {
		result = append(result, kmpSearch(T, s, line, i)...)
	}

	return strings.Join(result, ","), nil
}

// Knuth-Morris-Pratt algorithm, modified slightly to return all occurrences
// via: http://en.wikipedia.org/wiki/Knuth–Morris–Pratt_algorithm
func kmpSearch(T []int, word, line string, lineNum int) []string {
	var result []string

	row := lineNum + 1

	m := 0
	i := 0

	for ; m + i < len(line); {
		if word[i] == line[m + i] {
			if i == len(word) - 1 {
				result = append(result, fmt.Sprintf("%d:%d", row, m))
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
