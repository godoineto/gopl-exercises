package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	fmt.Println(os.Args[0])
	fmt.Println(strings.Join(os.Args[1:], "-"))
	secs := time.Since(start).Seconds()
	fmt.Println(secs)
}
