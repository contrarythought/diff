package algo

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

var (
	debug = false
)

func read_file(f *os.File) []string {

	file := bufio.NewReader(f)
	var ret []string = make([]string, 1)
	var err error
	var line string

	line, err = file.ReadString('\n')
	for ok := true; ok; line, err = file.ReadString('\n') {
		if err != nil {
			if err == io.EOF {
				ret = append(ret, line)
				ok = false
				continue
			} else {
				fmt.Println(err)
				os.Exit(1)
			}
		}

		ret = append(ret, line)
	}

	return ret
}

func Longest_common_sub(f1 *os.File, f2 *os.File) {

	var f1_array []string = read_file(f1)
	var f2_array []string = read_file(f2)

	if debug {
		fmt.Println("file 1")
		for _, line := range f1_array {
			fmt.Print(line)
		}

		fmt.Println("\nfile 2")
		for _, line := range f2_array {
			fmt.Print(line)
		}
	}

}
