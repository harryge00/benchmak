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

	flag.IntVar(&runtime, "runtime", 10, "Run time fot testing")
	flag.IntVar(&rate, "rate", 5000, "Writing rate of log")

	flag.Parse()
}

func printLogs(count, num int) {
	for j:=0; j < num; j++ {
		fmt.Printf("{\"log\":\"%d_%s\", \"stream\":\"stdout\",\"time\":\"%s\"}\n", count, RandStringBytes(64), time.Now())
		count++
	}
}
func main() {
	count := 0
	ticker := time.NewTicker(time.Duration(1000000000/rate) * time.Nanosecond)
	for j := 0; j < runtime; j++ {
		i := 0
		for range ticker.C {
			fmt.Printf("{\"log\":\"%d_%s\", \"stream\":\"stdout\",\"time\":\"%s\"}\n", count, RandStringBytes(64), time.Now().Format("2006-01-02T15:04:05.999999999Z"))
			count++
			i++
			if i >= rate {
				break
			}
		}
	}
}
