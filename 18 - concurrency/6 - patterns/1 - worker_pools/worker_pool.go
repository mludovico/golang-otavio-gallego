package main

func main() {
	tasks := make(chan int, 45)
	results := make(chan int, 45)

	go worker(tasks, results)
	go worker(tasks, results)
	go worker(tasks, results)
	go worker(tasks, results)
	go worker(tasks, results)
	go worker(tasks, results)
	go worker(tasks, results)
	go worker(tasks, results)

	for i := 0; i < 45; i++ {
		tasks <- i
	}

	close(tasks)

	for r := range results {
		println(r)
	}
}

func worker(tasks <-chan int, results chan<- int) {
	for n := range tasks {
		results <- fibonacci(n)
	}
	close(results)
}

func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}
