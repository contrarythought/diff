package main

import (
	"diff/algo"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Printf("%s <file1.txt> <file2.txt>\n", os.Args[0])
		os.Exit(1)
	}

	file1, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	file2, err := os.Open(os.Args[2])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var res []string = algo.Longest_common_sub(file1, file2)

	for _, line := range res {
		fmt.Print(line)
	}

	file2.Close()
	file1.Close()
}
