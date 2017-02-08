package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	a := sc.Text()

	for {
		i := seekA(a)
		if i == -1 {
			break
		}
		a = roundA(a, i)
	}
	fmt.Println(a)
}

func seekA(a string) int {
	for i := range a {
		n, _ := strconv.Atoi(a[i : i+1])
		if n >= 5 {
			return i
		}
	}
	return -1
}

func roundA(a string, i int) string {
	aint, _ := strconv.Atoi(a)
	p := math.Pow10(len(a) - i)
	f := float64(aint) / p
	around := strconv.Itoa(int(math.Floor(f+.5) * p))
	return around
}
