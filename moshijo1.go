package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var l1, l2, l3 int
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	l1, _ = strconv.Atoi(sc.Text())
	sc.Scan()
	l2, _ = strconv.Atoi(sc.Text())
	sc.Scan()
	l3, _ = strconv.Atoi(sc.Text())
	fmt.Println(l1 + l2 + l3)

}
