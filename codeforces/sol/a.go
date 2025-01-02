package sol

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func A() int {
	sc := bufio.NewScanner(os.Stdin)
	sc.Text()
	nm := strings.Split(sc.Text(), " ")
	n, _ := strconv.Atoi(nm[0])
	m, _ := strconv.Atoi(nm[1])
	x := 0

	for i := 0; i < n; i++ {
		l := len(sc.Text())
		if i > 0 && x == 0 {
			return x
		}
		if x+l <= m {
			x += l
		}
	}
	return x
}

func Aft(n int, m int, ss []string) int {
	x := 0
	totalLen := 0
	for i, s := range ss {
		l := len(s)
		if i > 0 && x == 0 {
			return x
		}
		if totalLen + l <= m {
			totalLen += l
			x++
		}
	}
	return x
}
