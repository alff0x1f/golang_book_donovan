// Exercise 1.4: Modify dup2 to print the names of all files in which each
// duplicated line occurs

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	countsPerFile := make(map[string]map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countsPerFile[arg] = make(map[string]int)
			countLines(f, countsPerFile[arg])
			for line, n := range countsPerFile[arg] {
				counts[line] += n
			}
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			var filenames, sep string
			for name, n2 := range countsPerFile {
				if n2[line] > 0 {
					filenames += sep + name
					sep = " "
				}
			}
			fmt.Printf("%d\t%s\t%s\n", n, line, filenames)
		}
	}
}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
		// countsPerFile[filename][input.Text()]++
	}
	// ignoring potential errors from input.Err()
}
