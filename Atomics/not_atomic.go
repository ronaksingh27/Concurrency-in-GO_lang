package main

import (
	"fmt"
	"sync/atomic"
)


type student struct{
	grades map[string]int
}


func main(){
	fmt.Println("Welcome to student grades ")

	var wg sync.WaitGroup
	var val atomic.Value
	val.Store(student{grades: map[string]int{} })

	wg.Add(3)
	go func(){
		defer wg.Done()
		s := val.Load().(student)//getting student object
		m := s.grades
		m["physics"] = 10
		val.Store(student{grades: m})
	}()
	go func(){
		defer wg.Done()
		s := val.Load().(student)
		m := s.grades
		m["maths"] = 10
		val.Store(student{grades: m})
	}()
	go func(){
		defer wg.Done()
		s := val.Load().(student)
		m := s.grades
		m["chemistry"] = 10
		val.Store(student{grades: m})
	}()

	
}