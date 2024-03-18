package main

import (
	"fmt"
	"net"
	"sync/atomic"
	"time"
)

func main(){
	fmt.Println("Welcome to TCP Server")
	
	listener ,err := net.Listen("tcp",":2708")
	if err != nil{
		panic(err)
	}
	defer listener.Close()//close connection at the end of this function

	var connections int32
	for{
		conn , err := listener.Accept()
		if err != nil{
			panic(err)
		}
		connections++;

		go func(){

			//run the function to close connection at the end of above function
			defer func(){
				_ = conn.Close()
				atomic.AddInt32(&connections,-1)
			}()

			if( atomic.LoadInt32(&connections) > 3 ){
				return
			} 

			time.Sleep(time.Second)
			_ , err := conn.Write([]byte("success"))
			if err != nil{
				panic(err)
			}


		}()
	}
	

}