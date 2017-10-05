package main

import (
	"math/rand"

	"time"
)

type Bank struct {
	clients  []int
	cashiers []int
}

func initBankPart(count, maxTimeWait int) ([]int) {
	res := make([]int, count)
	for i:=0; i< count; i++ {
		res[i] = rand.Int() % maxTimeWait
	}
	return res
}

func (bank *Bank) InitClient(count, maxTimeWait int) {
	bank.clients = initBankPart(count, maxTimeWait)
}

func (bank *Bank) InitCashiers(count, maxTimeWait int) {
	bank.cashiers = initBankPart(count, maxTimeWait)
}

func (bank *Bank) processClient(index chan int, indexCashier int){


	for {
		res := <-index
		index <- res + 1
		if(res > len(bank.clients)){
			break
		}
		time.Sleep(bank.cashiers[indexCashier] * time.Millisecond)
	}



}

func (bank *Bank) Process(){
	index := make(chan int)
	for i:=0; i< len(bank.cashiers); i++ {
		go bank.processClient(index, i)
	}
}