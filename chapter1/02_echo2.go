package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	start := time.Now()
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		fmt.Print(index)
		fmt.Println(": " + arg)
		s += sep + arg
		sep = "-"
	}

	fmt.Println(os.Args[0])
	fmt.Println(s)
	secs := time.Since(start).Seconds()
	fmt.Println(secs)
}
