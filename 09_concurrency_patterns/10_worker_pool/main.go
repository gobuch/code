package main

import (
	"fmt"
	"runtime"
	"time"
)

type job func(chan result)

type result struct {
	name string
	sum  int
	err  error
}

func startWorker(jobs chan job, results chan result) {
	fmt.Println("starte Worker")
	go func() {
		for j := range jobs {
			j(results)
		}
		fmt.Println("Worker beendet")
	}()
}

func main() {
	jobs := make(chan job)
	results := make(chan result)
	// Ausgabe der Ergebnisse in einer eigenen
	// Goroutine
	go func() {
		for r := range results {
			fmt.Println(r)
		}
	}()
	numWorkers := runtime.NumCPU()
	for i := 0; i < numWorkers; i++ {
		startWorker(jobs, results)
	}
	job1 := func(resCh chan result) {
		r := result{
			"10 + 5", 10 + 5, nil,
		}
		resCh <- r
	}
	counter := 0
	job2 := func(resCh chan result) {
		counter++
		r := result{
			"Counter: ", counter, nil,
		}
		resCh <- r
	}
	fmt.Println("Mit Ctrl+C abbrechen")
	for {
		jobs <- job1
		jobs <- job2
		time.Sleep(time.Second)
	}
}
