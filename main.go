package main

import (
	"fmt"
	"github.com/Comcast/golang-money"
)

func main() {
	fmt.Println("this is a test")
	m := money.Money{}
	fmt.Println(m.ToString())
}
