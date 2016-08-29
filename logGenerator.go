// bulk_query_gen generates queries for various use cases. Its output will
// be consumed by query_benchmarker.
package main

import (
	"flag"
	"fmt"
	// "log"
	"math/rand"
	"os"
	// "sort"
	"time"
)

// Program option vars:
var (
	runtime   int
	rate int
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func RandStringBytes(n int) string {
    b := make([]byte, n)
    for i := range b {
        b[i] = letterBytes[rand.Intn(len(letterBytes))]
    }
    return string(b)
}

// Parse args:
func init() {
	// Change the Usage function to print the use case matrix of choices:
	oldUsage := flag.Usage
	flag.Usage = func() {
		oldUsage()

		fmt.Fprintf(os.Stderr, "\n")
		fmt.Fprintf(os.Stderr, "Test fluentd\n")

	}

	flag.IntVar(&runtime, "runtime", 300, "Run time fot testing")
	flag.IntVar(&rate, "rate", 500000, "Writing rate of log")

	flag.Parse()
	fmt.Println(runtime, rate)
}

func main() {
	var t  *time.Timer
	count := 0
	for i:=0; i<runtime; i++ {
		t = time.NewTimer(1 * time.Second)
		for j:=0; j < rate; j++ {
			fmt.Println(count, RandStringBytes(64))
			count++
		}
		<- t.C
	}
}
