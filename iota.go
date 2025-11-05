package main

import "fmt"

//iota resets after the const block finishes executing
const (
	_ = iota //can pass blank values to start iota count (0)
	monday // iota = 1
	tuesday
	wednesday
	thursday
	friday
	saturday
	sunday
)

func main(){
	fmt.Println("Hello world")
	fmt.Println("Today is", monday)
}