package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func main(){

	fmt.Println("Welcome to basic counter in go ")

	var count int64
	var wg sync.WaitGroup

	//reader
	wg.Add(1)
	go func(){
		defer wg.Done()
		time.Sleep(10*time.Millisecond)
		fmt.Println("The value of count in reader is : ",atomic.LoadInt64(&count))
	}()

	//writer
	wg.Add(50)
	for i := 0 ; i < 50 ; i++{
		go func(){
			defer wg.Done()
			time.Sleep(10*time.Millisecond)
			atomic.AddInt64(&count,1)
		}()
	}

	wg.Wait()
	fmt.Println("The value of count is : ",count);
	

}