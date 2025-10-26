package main //packages are ways to group functions

import "fmt" //text formatting and printing package
import "rsc.io/quote" //package for quoting strings

func main() {
    fmt.Println("Hello, World!")
	fmt.Println(quote.Go())
}