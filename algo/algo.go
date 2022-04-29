package algo

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func read_file(f *os.File) []string {

	file := bufio.NewReader(f)
	var ret []string
	var err error
	i := 0
	ret[i], err = file.ReadString('\n')
	i++
	for {
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println(err)
			os.Exit(1)
		}
		ret[i], err = file.ReadString('\n')
		i++
	}

	return ret

}

func Longest_common_sub(f1 *os.File, f2 *os.File) []string {

	var f1_array []string = read_file(f1)
	var f2_array []string = read_file(f2)

}
