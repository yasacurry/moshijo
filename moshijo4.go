package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	a := sc.Text()
	sc.Scan()
	s := sc.Text()
	sc.Scan()
	t := sc.Text()

	var buf bytes.Buffer

	if s == "encode" {
		for i := range t {
			n, _ := strconv.Atoi(t[i : i+1])
			buf.WriteString(a[2*n : 2*n+1])
		}
		fmt.Println(buf.String())
		return
	}

	for i := range t {
		n := strings.Index(a, t[i:i+1])
		buf.WriteString(strconv.Itoa(n / 2))
	}

	fmt.Println(buf.String())

}
