package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
)

// Computing square value and passing the values through the channel.

func computeSquareValue(number int, wg *sync.WaitGroup, ch chan int) {
	defer wg.Done()
	squareValue := number * number
	ch <- squareValue
}

func aggregateSquareValues(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done() // Signal that aggregator is done

	aggregatedSquareValues := 0
	for squareValue := range ch {
		aggregatedSquareValues += squareValue
	}

	fmt.Println("Aggregated square values:", aggregatedSquareValues)
}

func main() {

	wg := sync.WaitGroup{}    // WaitGroup for producers
	aggWg := sync.WaitGroup{} // WaitGroup for aggregator

	fmt.Println("Starting the square values generator...")
	fmt.Println("--------------------------------")

	fmt.Print("Enter numbers separated by spaces (e.g., 2 3 4 5): ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := strings.TrimSpace(scanner.Text())

	if input == "" {
		fmt.Println("No numbers provided. Please enter at least one number.")
		return
	}

	// Split input by spaces and convert to integers
	numberStrings := strings.Fields(input)
	numbers := make([]int, 0, len(numberStrings))

	for _, numStr := range numberStrings {
		number, err := strconv.Atoi(numStr)
		if err != nil {
			fmt.Printf("Invalid number '%s'. Please enter only valid integers.\n", numStr)
			return
		}
		numbers = append(numbers, number)
	}

	if len(numbers) == 0 {
		fmt.Println("No valid numbers found. Please enter at least one number.")
		return
	}

	fmt.Printf("Processing %d numbers: %v\n", len(numbers), numbers)
	fmt.Println("--------------------------------")

	ch := make(chan int, len(numbers))

	// Start ONLY ONE aggregator goroutine (outside the loop!)
	aggWg.Add(1) // Add aggregator to its WaitGroup
	go aggregateSquareValues(ch, &aggWg)

	// Start all worker goroutines
	for _, number := range numbers {
		wg.Add(1)
		go computeSquareValue(number, &wg, ch)
	}

	// Wait for all workers to finish sending values
	wg.Wait()

	// Close channel after all values are sent (this signals aggregator to stop)
	close(ch)

	// Wait for aggregator to finish processing and printing
	aggWg.Wait()
}
