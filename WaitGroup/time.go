package main

import (
	"fmt"
	"sync"
	"time"
)

func main(){

	now := time.Now()
	var wg sync.WaitGroup;

	n := 1000
	wg.Add(n)
	for i := 0 ; i < n ; i++{
		go work(&wg,i+1)
	}

	wg.Wait()
	fmt.Println(time.Since(now))

}

func work(wg *sync.WaitGroup , id int){
	defer wg.Done()

	fmt.Println("id : ",id,"is working")
}

/* 
	Wait groups executes go routines in random order

*/