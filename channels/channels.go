package channels

import (
	"fmt"
	"math/rand"
	"time"
)

const workersNumber int = 2

var processedJobs [workersNumber]int = [workersNumber]int{}

func worker(id int, jobs <-chan int, results chan<- int) {
	//     - Reads a number from jobs
	for jobNumber := range jobs {
		fmt.Printf("> Worker %d is processing job %d. Channel still contains %d items\n", id, jobNumber, len(jobs))

		// - Simulates work with time.Sleep
		delay := rand.Intn(100) + 1
		fmt.Printf("> Worker %d processing job %d during %d milliseconds\n", id, jobNumber, delay)
		time.Sleep(time.Duration(delay) * time.Millisecond)

		// - Sends back number * number on results
		results <- jobNumber * jobNumber
		fmt.Printf("> Worker %d completed job %d\n", id, jobNumber)

		// Increase this worker's number of processed jobs
		processedJobs[id] = processedJobs[id] + 1
	}
	fmt.Printf("> Worker %d, no more jobs!\n", id)

}

func launch() {

	fmt.Println("Launch!")

	jobs := createJobs(16)
	jobsChannel := make(chan int, len(jobs))
	resultsChannel := make(chan int, len(jobs))

	// start workers
	for i := range workersNumber {
		go worker(i, jobsChannel, resultsChannel)
	}

	// Send job to the input channel
	for _, job := range jobs {
		fmt.Println("Adding job ", job)
		jobsChannel <- job
	}

	// no more jobs
	close(jobsChannel)

	for range jobs {
		result := <-resultsChannel
		fmt.Println("Result:", result)
	}

	close(resultsChannel)

	fmt.Printf("Processed Jobs %v\n", processedJobs)
	fmt.Println("Launch complete. Bye bye")
}

func createJobs(quantity int) []int {
	jobs := make([]int, quantity)
	for i := range quantity {
		jobs[i] = i
	}
	return jobs
}
