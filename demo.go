package main

import (
	"fmt"
	"sync"
)

//var mu sync.Mutex

type Job struct {
	ID     int
	Worker int
}

func worker(id int, job int, results chan<- Job, wg *sync.WaitGroup) {
	defer wg.Done()
	//mu.Lock()
	//for jobID := range jobs {
	//time.Sleep(time.Duration(1000) * time.Millisecond) // Mô phỏng thời gian thực hiện công việc
	result := Job{ID: job, Worker: id}
	results <- result
	//}
	//mu.Unlock()

}
func workerPool() {
	numWorkers := 5
	numJobs := 21

	jobs := make(chan int, numJobs)
	results := make(chan Job, numJobs)
	var wg sync.WaitGroup
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	for i, j := 1, 1; i <= numWorkers && j <= numJobs; i, j = i+1, j+1 {
		//mu.Lock()
		//v, ok := <-jobs
		//if !ok {
		//	break
		//}
		wg.Add(1)
		go worker(i, <-jobs, results, &wg)
		if i == numWorkers {
			i = 0
		}
		//mu.Unlock()
	}

	close(jobs)

	wg.Wait()
	close(results)

	completedJobs := make(map[int][]int)
	for result := range results {
		completedJobs[result.Worker] = append(completedJobs[result.Worker], result.ID)
	}

	for workerID, jobIDs := range completedJobs {
		fmt.Printf("Worker %d completed jobs: %v\n", workerID, jobIDs)
	}
}

func main() {
	//WorkerPool
	workerPool()

	//WaitGroup
	//waitgroup.Demo()

	//Mutex
	//mutex.Demo()

	//Atomic
	//atomic.Demo()

	//Select
	//selectDemo.Demo()

	//Context
	//context.Demo()
}
