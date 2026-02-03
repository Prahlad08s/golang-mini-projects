package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

/*
Note: Multiple fmt.Println calls from different goroutines are serialized via an internal mutex,
     so they effectively queue.
*/

func main() {

	fmt.Printf("=== Worker Pool Demo ===\n")
	fmt.Println("--------------------------------")

	scanner := bufio.NewScanner(os.Stdin)

	// Get number of workers from user
	fmt.Print("Enter the number of workers: ")
	scanner.Scan()
	workersStr := strings.TrimSpace(scanner.Text())
	numberOfWorkers, err := strconv.Atoi(workersStr)
	if err != nil || numberOfWorkers <= 0 {
		fmt.Println("Invalid input. Please enter a positive number.")
		return
	}

	// Get number of jobs from user
	fmt.Print("Enter the number of jobs: ")
	scanner.Scan()
	jobsStr := strings.TrimSpace(scanner.Text())
	numberOfJobs, err := strconv.Atoi(jobsStr)
	if err != nil || numberOfJobs <= 0 {
		fmt.Println("Invalid input. Please enter a positive number.")
		return
	}

	// Channel size is fixed at 5 (hardcoded)
	channelSize := 5

	fmt.Println("--------------------------------")
	fmt.Printf("Starting worker pool with %d workers and %d jobs\n", numberOfWorkers, numberOfJobs)
	fmt.Printf("Channel size: %d (fixed)\n", channelSize)
	fmt.Println("--------------------------------")

	// step 1:  creating the channel
	// for what thing do i want to create the channel for? -> to send job ids to the workers
	// what data will be sent through the channel? ->  Job Ids
	jobs := make(chan int, channelSize) // buffered channel

	// step 3.1: wait for all the workers to complete their jobs
	var wg sync.WaitGroup
	wg.Add(numberOfWorkers)

	//step 4: start the workers
	for workerId := 0; workerId < numberOfWorkers; workerId++ {
		go worker(workerId, jobs, &wg) //Pass Waitgroup by pointer
	}

	// step 2: send jobs to the channel
	for jobId := 1; jobId <= numberOfJobs; jobId++ {
		jobs <- jobId
		fmt.Printf("Job %d sent to channel\n", jobId)
	}

	// Close channel after all jobs are sent
	close(jobs)

	//step 3.2 : wait for all the workers to complete their jobs
	fmt.Println("Waiting for all workers to complete their jobs...")
	wg.Wait() // Block here until all workers call Done()

	fmt.Println("All jobs completed!")
	fmt.Println("--------------------------------")

	// completion message
	fmt.Println("Worker pool system completed successfully")
	fmt.Println("--------------------------------")
}

func worker(workerId int, jobs chan int, wg *sync.WaitGroup) {

	//Critical section: Call Done() when the worker is done
	defer wg.Done()

	fmt.Printf("Worker %d started\n", workerId)

	for jobId := range jobs {
		fmt.Println("--------------------------------")
		fmt.Printf("Worker %d started job %d\n", workerId, jobId)
		time.Sleep(time.Second * 2)
		fmt.Printf("Worker %d completed job %d\n", workerId, jobId)
		fmt.Println("--------------------------------")
	}

	fmt.Printf("Worker %d completed all jobs\n", workerId)
	fmt.Println("---------------END-----------------")
}
