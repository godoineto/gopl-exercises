package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	start := time.Now()
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		fmt.Print(i)
		fmt.Println(": " + os.Args[i])
		s += sep + os.Args[i]
		sep = "-"
	}
	fmt.Println(os.Args[0])
	fmt.Println(s)
	secs := time.Since(start).Seconds()
	fmt.Println(secs)
}
