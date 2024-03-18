package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"sync"
)

func main(){
	total,max := 10,3
	var wg sync.WaitGroup

	for i := 0 ; i < total ; i+= max{

		limit := max;
		if( i + max > total ){
			limit = total - i
		}

		wg.Add(limit)
		for j := 0 ; j < limit ; j++{
			go func(j int){
				defer wg.Done()
				conn , err := net.Dial("tcp",":2708")
				if err != nil{
					panic(err)
				}

				bs , err := ioutil.ReadAll(conn)
				if err != nil{
					panic(err)
				}

				if string(bs) != "success"{
					panic(err)
				}

				fmt.Printf("Request %d success\n", i + 1 + j)
			}(j)
		}
		wg.Wait()
	}

}