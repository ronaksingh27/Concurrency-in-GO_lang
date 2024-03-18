package main

import "fmt"

func main(){
	var data int = 0
	go func(){
		data++;
	}()
	
	
	if( data == 0 ){
		fmt.Printf("the value of data is %v\n", data)
	}
}