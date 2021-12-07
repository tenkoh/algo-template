package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)

func scanInts(sep string) ([]int, error) {
	sc.Scan()
	ss := strings.Split(sc.Text(), sep)
	ans := make([]int, len(ss))
	for i := 0; i < len(ss); i++ {
		n, err := strconv.Atoi(ss[i])
		if err != nil {
			return nil, err
		}
		ans[i] = n
	}
	return ans, nil
}

func scanStrings(sep string) []string {
	sc.Scan()
	return strings.Split(sc.Text(), sep)
}
