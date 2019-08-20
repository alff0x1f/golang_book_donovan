/*Exercese 1.3: Experiment to measure the difference in running time between
our potentially inefficient versions and the one that uses strings.Join
(Section 1.5 illustrates part of the time package, and Section 11.4 show how to
write benchmark tests for systematic perfomance evaluation)
*/
package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
	secs1 := time.Since(start).Seconds()

	start = time.Now()
	fmt.Println(strings.Join(os.Args[1:], " "))
	secs2 := time.Since(start).Seconds()
	fmt.Printf("Time uneficient: %f, time by join func: %f, diff %.1f times\n", secs1, secs2, secs1/secs2)
}
