package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)


type calculator struct{
	res atomic.Value
}

func (c *calculator) add(n float64){
	fmt.Println("Inside add")
	c.res.Store(c.result() + n )
}

func (c *calculator) sub(n float64){
	c.res.Store(c.result() - n )
}

func (c *calculator) mul(n float64){
	c.res.Store(c.result() * n )
}

func (c *calculator) div(n float64){
	if n == 0 {
		panic("cannot divide by zero")
	}
	c.res.Store(c.result() / n )
}


func (c *calculator) result() float64{
	r , ok := c.res.Load().(float64)
	if !ok {
		panic("opearting with wrong type")
	}

	return r
}

func newCalculator() calculator{
	c := calculator{}
	c.res.Store(float64(0))//will remember it
	return c
}


func main(){

	fmt.Println("Welcome to basic calculator in go ")
	c := newCalculator()

	var wg sync.WaitGroup

	wg.Add(4)
	go func(){
		defer wg.Done()
		c.add(2)
	}()
	go func(){
		defer wg.Done()
		c.div(3)
	}()
	go func(){
		defer wg.Done()
		c.mul(5)
	}()
	go func(){
		defer wg.Done()
		c.div(7)
	}()

	wg.Wait()


	fmt.Println("res : ",c.res)

}