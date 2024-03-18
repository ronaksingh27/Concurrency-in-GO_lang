package main

import (
	"fmt"
	"sync"
)

func main(){
	fmt.Println("Welcome to WaitGroup")

	// var wg sync.WaitGroup

	// wg.Add(1)
	// go func(){
	// 	defer wg.Done()
	// 	fmt.Println("Calling from func1")
	// 	time.Sleep(2)
	// }()

	// wg.Add(1)
	// go func(){
	// 	defer wg.Done()
	// 	fmt.Println("Calling from func2")
	// 	time.Sleep(1)
	// }()

	// wg.Wait()


	hello := func(wg *sync.WaitGroup , id int){
		defer wg.Done()
		fmt.Printf("Hello from %v\n", id)
	}


	const numGreeters = 5
	var wg sync.WaitGroup

	wg.Add(numGreeters)
	for i := 1 ; i <= numGreeters ; i++ {
		go hello(&wg,i)
	}

	wg.Wait()

}