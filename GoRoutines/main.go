package main

import (
	"fmt"
	"sync"
	"time"
)

func AsyncPrint() {
	intsToPrint := []int{1, 2, 3, 4, 5}

	for _, intToPrint := range intsToPrint {
		fmt.Println(intToPrint)
	}
}

func AsyncPrintWithWaitGroup(wg *sync.WaitGroup) {
	intsToPrint := []int{1, 2, 3, 4, 5}

	for _, intToPrint := range intsToPrint {
		fmt.Println("Running in waitgroup")
		fmt.Println(intToPrint)
	}

	defer wg.Done()
}

func main() {

	AsyncPrint()

	// Invoking the function with go runs it asynchronously
	fmt.Println("Running Async")
	go AsyncPrint()

	time.Sleep(time.Second)

	// Running routines with waitgroup

	var waitGroup sync.WaitGroup

	waitGroup.Add(1)

	go AsyncPrintWithWaitGroup(&waitGroup)

	waitGroup.Wait()

	fmt.Println("Wait Group done")
}
