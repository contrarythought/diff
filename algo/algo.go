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

	for line, err = file.ReadString('\n'); err == nil; line, err = file.ReadString('\n') {
		ret = append(ret, line)
	}
	if err == io.EOF {
		ret = append(ret, line)
	} else {
		fmt.Println(err)
		os.Exit(1)
	}

	return ret
}

func max(a uint, b uint) uint {
	if a > b {
		return a
	}
	return b
}

func Longest_common_sub(f1 *os.File, f2 *os.File) []string {

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

	// TODO - understand this
	m, n := len(f1_array), len(f2_array)
	P := make([]uint, (m+1)*(n+1))
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if f1_array[i-1] == f2_array[j-1] {
				P[i*(n+1)+j] = 1 + P[(i-1)*(n+1)+(j-1)]
			} else {
				P[i*(n+1)+j] = max(P[i*(n+1)+(j-1)], P[(i-1)*(n+1)+j])
			}
		}
	}

	longest := P[m*(n+1)+n]
	i, j := m, n
	subsequence := make([]string, longest)
	for k := longest; k > 0; {
		if P[i*(n+1)+j] == P[i*(n+1)+(j-1)] {
			j-- // the two strings end with the same char
		} else if P[i*(n+1)+j] == P[(i-1)*(n+1)+j] {
			i--
		} else if P[i*(n+1)+j] == 1+P[(i-1)*(n+1)+(j-1)] {
			subsequence[k-1] = f1_array[i-1]
			k--
			i--
			j--
		}
	}

	return subsequence
}
