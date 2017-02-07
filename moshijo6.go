package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var n, m int
	type psd struct {
		p, s, d int
	}

	var ngroup []psd

	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	n, _ = strconv.Atoi(sc.Text())

	for i := 0; i < n; i++ {
		sc.Scan()
		psdlist := strings.Split(sc.Text(), " ")
		pt, _ := strconv.Atoi(psdlist[0])
		st, _ := strconv.Atoi(psdlist[1])
		dt, _ := strconv.Atoi(psdlist[2])
		ngroup = append(ngroup, psd{pt, st, dt})
	}

	sc.Scan()
	m, _ = strconv.Atoi(sc.Text())

	for i := 0; i < m; i++ {
		sc.Scan()
		ca := strings.Split(sc.Text(), " ")
		c, _ := strconv.Atoi(ca[0])
		a, _ := strconv.Atoi(ca[1])
		fmt.Println((ngroup[c-1].p * a) - ((a / ngroup[c-1].s) * ngroup[c-1].d))
	}
}
