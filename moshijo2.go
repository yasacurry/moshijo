package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	if strings.Contains(sc.Text(), "help") {
		fmt.Println("SOS")
		return
	}
	fmt.Println(sc.Text())
}
