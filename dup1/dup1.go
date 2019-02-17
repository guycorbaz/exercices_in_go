// Dup1 prints the text of each line that appears more than
// once in the standard input, preceded by its count.
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)  // Create an empty map: set of key/value
	input := bufio.NewScanner(os.Stdin) // Read the lines and remove newline character from the end of the line
	fmt.Println("--> Begin")
	for input.Scan() {
		counts[input.Text()]++ // Increment value correspondif to key
	}

	// NOTE: ignoring potential errors from input.Err()
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

