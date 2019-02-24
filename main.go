package main

import (
	"fmt"
	"sync"
	"time"
)

func printChar(c chan int) {
	for i := 0; i < 4; i++ {
		fmt.Println(<-c)
	}
	//close(c)
}

func printDigit(num []int) {
	fmt.Print("start printDigit")
	for {
	}
}

func printNum(wg *sync.WaitGroup, i int) {
	time.Sleep(time.Millisecond)
	fmt.Println(i)
	wg.Done()
}

func numWorker(wg *sync.WaitGroup, tasks <-chan int, results chan<- int, instance int) {
	for task := range tasks {
		time.Sleep(time.Millisecond)
		fmt.Printf("Worked %v picked up the task %v\n", instance, task)
		results <- task
	}
	wg.Done()
}

var inc int

func raceCondition(wg *sync.WaitGroup, m *sync.Mutex, instance int) {
	m.Lock()
	inc++
	//fmt.Printf("Worker %v incremented inc to %v\n", instance, inc)
	m.Unlock()
	wg.Done()
}

func main() {
	start := time.Now()
	fmt.Println("start")

	var wg sync.WaitGroup
	var m sync.Mutex

	//tasks := make(chan int, 10)
	//results := make(chan int, 10)

	for i := 1; i < 10000; i++ {
		wg.Add(1)
		//go numWorker(&wg, tasks, results, i)
		go raceCondition(&wg, &m, i)
	}

	wg.Wait()

	fmt.Println("Value of inc", ":", inc)
	/*for j := 1; j < 7; j++ {
		tasks <- j
	}

	fmt.Println("Finished pushing tasks")

	close(tasks)

	wg.Wait()

	for k := 1; k < 7; k++ {
		fmt.Printf("Main received back %v\n", <-results)
	}*/

	//go printChar(c)
	//wg.Wait()

	fmt.Println("finished", time.Since(start))

}
