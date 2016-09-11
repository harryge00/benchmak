// bulk_query_gen generates queries for various use cases. Its output will
// be consumed by query_benchmarker.
package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
	"strconv"
)

// Program option vars:
var (
	runtime int
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

func main() {
	rand.Seed(time.Now().UTC().UnixNano());
	runtime, _ = strconv.Atoi(os.Getenv("runtime"))
	rate, _ = strconv.Atoi(os.Getenv("rate"))
	var count uint64
	count = 0
	ticker := time.NewTicker(time.Duration(1000000000/rate) * time.Nanosecond)
	for j := 0; j < runtime; j++ {
		i := 0
		for range ticker.C {
			fmt.Printf("{\"log\":\"%d_%s\", \"stream\":\"stdout\",\"time\":\"%s\"}\n", count, RandStringBytes(64), time.Now().Format("2006-01-02T15:04:05.999999999Z"))
			// fmt.Printf("%d_%s\n", count, RandStringBytes(64))
			count++
			i++
			if i >= rate {
				break
			}
		}
	}
	select{}
}
