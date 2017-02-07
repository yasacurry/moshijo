package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var c, r []int
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()

	sc.Scan()
	cs := strings.Split(sc.Text(), " ")
	for _, v := range cs {
		n, _ := strconv.Atoi(v)
		c = append(c, n)
	}

	sc.Scan()
	rs := strings.Split(sc.Text(), " ")
	for _, v := range rs {
		n, _ := strconv.Atoi(v)
		r = append(r, n)
	}

	for _, rv := range r {
		var a []string
		for _, cv := range c {
			n := strconv.Itoa(cv + rv)
			a = append(a, n)
		}
		fmt.Println(strings.Join(a, " "))
	}
}
