package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("---> START <---")
	// go startOne()
	// time.Sleep(1 * time.Second)
	// go startTwo()
	// time.Sleep(1 * time.Second)
	// fmt.Println("Application started...")

	// sync.WaitGroup
	var wg sync.WaitGroup
	wg.Add(2)
	go waitGroup1(&wg)
	go waitGroup2(&wg)
	wg.Wait()

}

func waitGroup2(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Wait Groupd 1")
}

func waitGroup1(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Wait Groupd 2")
}
