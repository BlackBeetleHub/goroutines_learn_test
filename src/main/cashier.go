package main

type Cashier struct {
	time int
}

func (casher *Cashier) GetTime() int {
	return casher.time
}
