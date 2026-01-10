Exercise:

1. Create a function worker(id int, jobs <-chan int, results chan<- int).
2. Start N worker goroutines.
3. Send a list of integers (e.g. 1–20) into a jobs channel.
4. Each worker:
    - Reads a number from jobs
    - Simulates work with time.Sleep
    - Sends back number * number on results

5. Collect and print all results in main.
6. Close channels correctly and avoid deadlocks.

