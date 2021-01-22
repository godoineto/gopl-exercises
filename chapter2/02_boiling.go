//Show the boiling point of water
package main

import "fmt"

const boilingF = 212.0

func main() {
	c := (boilingF - 32) * 5 / 9
	fmt.Printf("Boiling point: %g°F or %g°C\n", boilingF, c)
}
