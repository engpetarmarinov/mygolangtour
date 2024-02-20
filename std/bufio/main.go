package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	const input = "Sjhklahdlkjh 2178936287163"

	r := bufio.NewReader(strings.NewReader(input))
	io.Copy(os.Stdout, r)
	fmt.Println()
	s := bufio.NewScanner(strings.NewReader(input))
	s.Split(bufio.ScanWords)
	for s.Scan() {
		fmt.Printf("%s-\n", s.Bytes())
	}
}
