package main

import (
	"fmt"
	"sync"
)


type request func()

func main(){
	requests := map[int]request{}

	for i := 0 ; i < 100 ; i++{
		f := func(n int ) request{
			return func(){
				fmt.Println("request : ",n)
			}
		}
		requests[i]  = f(i)
	}

	var wg sync.WaitGroup
	max := 7
	for i := 0 ; i < len(requests) ; i += max{
		for j := i ; j < max + i ; j++{
			wg.Add(1)
			go func(r request){
				defer wg.Done()
				r()
			}(requests[j])
		}
		wg.Wait()
		fmt.Println("Request processed for ",max," items")
	}
}