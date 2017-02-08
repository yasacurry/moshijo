package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	sc.Scan()
	c := strings.Split(sc.Text(), "/")
	if c[0] == "" {
		c = c[1:]
	}
	sc.Scan()
	p := strings.Split(sc.Text(), "/")

	for _, v := range p {
		if v == ".." {
			if len(c) == 0 {
				continue
			}
			c = c[:len(c)-1]
		} else if v != "." {
			c = append(c, v)
		}
	}

	fmt.Println("/" + strings.Join(c, "/"))
}
