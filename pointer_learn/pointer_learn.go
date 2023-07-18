package main

import "fmt"

// Although named as pointer learn, we will cover interfaces as part of this module

// This is an interface. Now whoever implements the method signature written
// in this interface will be considered of type Animal
type Animal interface {
    Run() string
    Sleep() string
}

type Dog struct {}
func (dog *Dog) Run() string {
    return "Runs fast!"
}
func (dog *Dog) Sleep() string {
    return "Sleeps sound!"
}

type Cat struct {}
func (cat *Cat) Run() string {
    return "Doesn't run hehe"
}
func (cat *Cat) Sleep() string {
    return "Always eepy..."
}

type NonAnimal struct{}
func (nonAnimal *NonAnimal) Run() string {
    return "run"
}

func main() {
    var dog Animal
    
    fmt.Println("Assigning variable of Dog type to a variable of Animal type...")
    fmt.Println("It doesn't complain since Dog has implemented all the methods defined as part of Animal interface!")
    // We have to provide the pointer reference to dog during assignment
    // because we have implemented the methods of Animal interface using the pointer reference to Dog
    dog = &Dog {}
    fmt.Println("\nHere's what dog does when told to do the respective..")
    fmt.Println(dog.Run())
    fmt.Println(dog.Sleep())

    var cat Animal = &Cat{}
    fmt.Println("\nHere's what cat does when told to do the respective..")
    fmt.Println(cat.Run())
    fmt.Println(cat.Sleep())
    
    // The below does not work, since the NonAnimal struct does not implement ALL the methods of Animal interface
    // var nonAnimal Animal = &NonAnimal{}
}

/* 
    Note that, we did not have to explicitly write any keyword like `implements` to implement the interface defined.
    Rather, the interface got implicitly implemented just by implementing its methods.
    Similar to the interfaces in other languages, the interface itself can never be instantiated and its methods
    can never be called unless they are implemented by someone
*/












