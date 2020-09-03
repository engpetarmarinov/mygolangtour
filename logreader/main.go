package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	level := flag.String("level", "ERROR", "This will filter out ERROR loglevel messages from your log.")

	flag.Parse()

	f, err := os.Open("test.log")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	r := bufio.NewReader(f)
	for {
		line, err := r.ReadString('\n')
		if err != nil {
			break
		}
		if strings.Contains(line, *level) {
			fmt.Print(line)
		}
	}
}
