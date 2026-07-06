package main

import (
	"fmt"
	"math"
)

// func worker(id int, messages chan string) {
// 	fmt.Printf("Worker %d starting\n", id)
// 	time.Sleep(time.Second) // Simulate work
// 	messages <- fmt.Sprintf("Worker %d finished", id)
// }

func worker(id int, messages chan string) {
	fmt.Printf("Worker %d starting\n", id)

	// Simulate CPU bound task: calc the sum of sq roots
	sum := 0.0
	for i := 0; i < 100000000; i++ {
		sum += math.Sqrt(float64(i))
	}

	messages <- fmt.Sprintf("Worker %d finished (computed: %.2f)", id, sum)
}

func main() {
	messages := make(chan string)

	// Launch 3 goroutines
	for i := 1; i <= 10000; i++ {
		go worker(i, messages)
	}

	// Receive from channel 3 times
	for i := 1; i <= 10000; i++ {
		fmt.Println(<-messages)
	}
	fmt.Println("All workers done!")
}
