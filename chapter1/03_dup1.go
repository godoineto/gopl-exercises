// Dup1 show text from duplicated lines
// that appers in standard input, with its count
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)

	fmt.Println("Write 'exit' to stop the scanner")
	for input.Scan() {
		text := input.Text()
		if text == "exit" {
			break
		}
		counts[text]++
	}
	// Error is ignored at this point

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
