package bench

import (
	"os"
	"bufio"
	"strings"
	"fmt"
)



func Find(path, s string) (string, error) {

	lines, err := readLines(path)
	if err != nil {
		return "", err
	}

	var result []string

	for i, line := range lines {
		result = append(result, findInLine(s, line, i)...)
//		fmt.Printf("read line %s\n", line)
	}


	return strings.Join(result, ","), nil
}


func findInLine(s, line string, lineNum int) []string {
	var result []string

	row := lineNum + 1

	for col := 0; col < len(line); col++ {
		if matchAtOffset(s, line, col) {
		    result = append(result, fmt.Sprintf("%d:%d", row, col))
		}
	}

	return result
}

func matchAtOffset(s, line string, offset int) bool {

	// past the end
	if offset + len(s) > len(line) {
		return false
	}

	for i:= 0; i < len(s); i++ {
		if s[i] != line[offset + i] {
			return false
		}
	}

	return true
}




// readLines reads a whole file into memory
// and returns a slice of its lines.
func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}
