package main

import "fmt"


type MyInt struct {
    *int; // this is an anonymous field and will be named as int, the asterisk is ignored while naming
    val int
}

// *int implies that it is a pointer to int and hence while initialising its value we should be passing the pointer

// Functions with pointer receivers in Go (Similar to extension functions)
// https://go.dev/tour/methods/4
func (myInt *MyInt) multiply(multiPlier int) int {
    return (*myInt.int) * multiPlier
}


func main() {
    var myIntInt int = 5
    myInt := MyInt { int : &myIntInt } // over here we are passing a pointer to the int variable declared
    intPlusMyInt := 10 + *myInt.int
    fmt.Println("The res --> ", intPlusMyInt)

    fmt.Println("The mutliplication res --> ", myInt.multiply(5))
}
