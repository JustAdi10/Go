package main

import "fmt"

var variableOutsideFunction = 56

func main(){
	var x = 32
	const y = 23
	z := 4545 // shorter way to decalare variables, this only works inside a function
	z = 1
	variableOutsideFunction = 100 // once updated the variable changes everywhere even if the updation was done inside a function
	fmt.Println(x,y,z,variableOutsideFunction)
	printSomething()
	zeroValueDeclarations()
}

func printSomething(){
	fmt.Println(variableOutsideFunction)
}

func zeroValueDeclarations(){
	var i int //defaults to 0
	var j string // defaults to ""
	var k bool // defaults to false
	var l float64 // defaluts to 0
	fmt.Println(i,j,k,l)
	fmt.Printf("%v %v %v %q /n",i,j,k,l) // this lets us control what type of variable output we want (ex : we want the o/p to be a float) 
	// and also lets us format the output (ex : strings are printed with quotes)
}