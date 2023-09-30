package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sync"
)

const THREAD_COUNT = 5

func main() {

	input := readFile()
	var inputMonitor *Monitor = CreateMonitor(10)
	// var outputMonitor *Monitor = CreateMonitor(len(input))
	MainMethod(inputMonitor, input[0:10])

	// var wg sync.WaitGroup
	// wg.Add(THREAD_COUNT+1)
	// wg.Add(1)

	// for i:=0; i < THREAD_COUNT; i++ {
	// 	go WorkerMethod(inputMonitor, outputMonitor, input, &wg)
	// }

	// wg.Wait()


}

func readFile() []Payment {
	file, err := os.Open("input.json")

	if err != nil {
		fmt.Println("error opening file: ", err)
		return nil
	}

	defer file.Close()
	var payment []Payment

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&payment)

	if err != nil {
		fmt.Println("error while decoding: ", err)
	}

	return payment

}

func MainMethod(inputMonitor *Monitor, input []Payment) {

	for len(input) > 0 {
		inputMonitor.Add(input[0])
		input = input[1:]
	}
}

func WorkerMethod(inputMonitor *Monitor, outputMonitor *Monitor, input []Payment, wg *sync.WaitGroup) {

	for len(input) > 0 || inputMonitor.GetCurrentLength() > 0 {
		item := inputMonitor.Remove()
		outputMonitor.Add(item)
	}
	wg.Done()
}
