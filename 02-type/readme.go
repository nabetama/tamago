package main

import (
	"log"
)

type Money struct {
	amount   uint
	currency string
}

func (this *Money) ToEmpty() {
	this.amount = 0
}

func main() {
	money := &Money{240, "$"}
	money.ToEmpty()
	log.Printf("%+v", money)
}
