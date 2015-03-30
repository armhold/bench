package main

import (
	"github.com/armhold/find_words/bench"
	"fmt"
)

func main() {

	result, err := bench.Find("assets/data.txt", "aa")

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("result: %s\n", result)
}

