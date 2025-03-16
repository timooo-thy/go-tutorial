package main

import (
	"fmt"
	"sync"
	"time"
)

// Sets up a wait group for concurrency
var wg = sync.WaitGroup{}

// Sets up a mutex for thread safe processes, contains lock and unlock methods
var m = sync.RWMutex{}

var dbData = []string{"id1", "id2", "id3", "id4", "id5"}
var results = []string{}

func main() {
	t0 := time.Now()
	for i := range dbData {
		// Increment delta to the wait group
		wg.Add(1)
		go dbCall(i)
	}
	// Basically, awaiting for all tasks to be done before continuing
	wg.Wait()
	fmt.Printf("Total execution time: %v\n", time.Since(t0))
	fmt.Printf("The results are %v", results)
}

func dbCall(i int) {
	delay := 2000
	time.Sleep(time.Duration(delay) * time.Millisecond)
	fmt.Println("The result from the database is: ", dbData[i])
	// Writing to the same memory location at the same time could lead to corrupt memory; not thread-safe hence use a mutex
	// m.Lock()
	// results = append(results, dbData[i])
	// m.Unlock()
	save(dbData[i])
	log()
	// Decrement delta to the wait group
	wg.Done()
}

func save(result string) {
	// Full lock
	m.Lock()
	results = append(results, result)
	m.Unlock()
}

func log() {
	// Read lock
	m.RLock()
	fmt.Printf("The current results are: %v\n", results)
	m.RUnlock()
}
