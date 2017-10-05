package main

import (
	"fmt"
	"time"
	"math/rand"
)

func generate(res chan int){
	for  {
		time.Sleep(2000 * time.Millisecond)
		res <- rand.Int()
	}
}

func main() {
	ch := make(chan int)

	go generate(ch)
	go generate(ch)
	go generate(ch)
	go generate(ch)
	go generate(ch)
	go generate(ch)

	for  {
		new_res := <-ch
		fmt.Println("NewGen!")
		fmt.Println(new_res)
	}
	return
}
