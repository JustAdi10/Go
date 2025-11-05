package main

import "fmt"

//happens when variables inside different scopes have the same names
func main(){
	x := 10 //outer scope
	fmt.Println(x)

	//entering inner scope
	if true {
		x := 5 //takes precedence over the main fxn declaration of the variable, the main fxn variable is temporarily disabled till the scope is changed (in this case the "if" is exited)
		fmt.Println(x)
	}

	fmt.Println(x)
}