package main

import (
	"fmt"
	"golang.org/x/example/stringutil"
)

func main() {
  fmt.Println(stringutil.Reverse("hello"))
  fmt.Println(stringutil.ToUpper("hello"))
  learnVariables()
  learnOutput()
  learnArray()
  learnSlice()
  learnIfElse()
  learnSwitchCase()
  learnLoops()
  fmt.Println(learnUnnamedReturnTypes())
  fmt.Println(learnNamedReturnTypes())
  learnStruct()
  learnMap()
}

func learnMap() {
  // We all know what a map is. Here is how you can declare it in Go
  // Syntax map_var = map[keyType][valueType] { key:value, key:value ...}
  firstMap := map[string] string { "name": "Yogesh", "surname": "Nikam" }
  secondMap := map[int] bool { 0: false, 1: true }
  fmt.Println("\nThe firstMap --> ", firstMap)
  fmt.Println("The secondMap --> ", secondMap)

  fmt.Println("The firstMap name key value --> ", firstMap["name"])
  fmt.Println("The secondMap 1 key value --> ", secondMap[1])
  
  fmt.Println("The length of firstMap --> ", len(firstMap))
  fmt.Println("The keys in secondMap --> ", secondMap)

  thirdMap := make(map[int]int)
  thirdMap[1] = 10
  thirdMap[2] = 20
  thirdMap[3] = 30
  fmt.Println("The thirdMap using make --> ", thirdMap)
  
  // in Go we have `nil` which is equivalent to `null` in other 
}

func learnStruct() {
  // Struct is basically Class in Go
  // Structs can be passed around in functions as arguments
  // Struct can be initialized in a class like format as well
  // Struct fields are mutable by default
  type Student struct {
    name string;
    age int;
    isSchoolStudent bool;
  }

  var student Student
  student.name = "Yogesh"
  student.age = 25
  student.isSchoolStudent = false

  fmt.Println("\nThe student is --> ", student)

  var anotherStudent Student = Student{ "Aditya", 10, true }
  fmt.Println("Another student with class like initialization syntax --> ", anotherStudent)
}

func learnLoops() {
  fmt.Println("\nLearning for loop")
  for i := 0; i < 5; i++ {
    fmt.Println("The num := ", i)
  }

  fmt.Println("\nIterating over an array/slice")
  arr := [3]string { "apple", "banana", "orange" }

  for idx, val := range(arr) {
    fmt.Println("The index - ", idx, "\tThe value - ", val)
  }
}

func learnIfElse() {
  // If statement
  if (10 > 5) {
    fmt.Println("\nIf statement result --> 10 is greater than 5")
  }
  
  if(5 > 10) {

  } else {
    fmt.Println("If else statement result --> 5 is not greater than 10")
  }
  
  name := "Yogesh"
  if(name == "Aditya") {
    fmt.Println("My name is Aditya")
  } else if(name == "Reva") {
    fmt.Println("My name is Reva")
  } else {
    fmt.Println("If else if else statement result --> My name is Yogesh")
  }
}

func learnSwitchCase() {
  val := 10
  result := -1
  switch(val) {
    case 20 : result = 200
    case 10 : result = 100
    default : result = 0
  } 
  fmt.Println("\nThe switch case result --> ", result)
}

func learnSlice() {
  // Slice is similar to ArrayList/Vector in Go
  // The length of a slice is not defined and can grow or shrink based on the needs
  first_slice := []int {100, 200}
  fmt.Println("\nThe first slice --> ", first_slice)
  fmt.Println("The length of first slice --> ", len(first_slice))
  fmt.Println("The capacity of first slice --> ", cap(first_slice))

  // creating a slice from an array
  arr := []int {400, 500, 600, 700, 800}
  second_slice := arr[1:4]
  fmt.Println("The second slice created using an array --> ", second_slice)

  // creating a slice using make method => make(type, len, capacity)
  third_slice := make([]string, 2, 5)
  fmt.Println("The third slice --> ", third_slice)
  fmt.Println("The third slice len --> ", len(third_slice))
  fmt.Println("The third slice cap --> ", cap(third_slice))

  // adding an element to the slice
  second_slice = append(second_slice, 800, 900, 1000)
  fmt.Println("The second slice post append --> ", second_slice)

  // append one slice to another
  fourth_slice := append(first_slice, second_slice...)
  fmt.Println("The fourth slice post merging first and second slice --> ", fourth_slice)

  // copy a slice into another slice
  copy_this_slice := []int {10,20,30,40}
  paste_to_this_slice := make([]int, 4)
  fmt.Println("Paste to this slice before --> ", paste_to_this_slice)
  copy(paste_to_this_slice, copy_this_slice)
  fmt.Println("Paste to this slice after --> ", paste_to_this_slice)
}

func learnArray() {
  // array is declared in go using the syntax [length] type {...values}
  // If we don't define the length, it will be inferred based on how many values we provide during declaration
  var first_array = [2]int {1,2}
  fmt.Println("\nThe first array --> ", first_array)

  second_array := []string {"one", "two"}
  fmt.Println("The second array --> ", second_array)
  
  // accessing an element from array
  fmt.Println("The second element from above array --> ", second_array[1])

  // changing element value from array
  second_array[0] = "three"
  fmt.Println("The second array post mutation --> ", second_array)

  // initialise only specific values during declaration
  third_array := []int { 2: 30, 5: 60 }
  fmt.Println("The third array with specific elements initialised --> ", third_array)
  fmt.Println("The length of the third array --> ", len(third_array))
}

func learnVariables() {
  // Variables in Go 
  // by default the variables defined below are mutable
  // once the go lang has inferred the type of the variable, it cannot be assigned a value of some other type
  var first_var int = 10 // Explicitly define the type of variable like we have done for int 
  second_var := 10       // The type is inferred by the language automatically based on the value provided.
  
  // The := operator can only be used within a function.
  // It cannot be used outside functions. We have rely on the var and const keywords
  // if we want to declare a variable outside the functions.
   
  // we can define a variable as a const making it immutable.
  // using const is what is preferred in general programming.
  const first_const string = "first_const"

  // multiple variable declaration of same type
  const firstMulti, secondMulti int = 100,200
  // multiple variable declartion of different types
  const thirdMulti, fourthMulti = 300, "fourth_multi"

  // variables declaration within a block
  const (
    blockVarOne = 400
    blockVarTwo = "block_var_two"
  ) 

  fmt.Println(first_var, second_var, first_const, firstMulti, secondMulti, thirdMulti, fourthMulti)
  fmt.Println(blockVarOne, blockVarTwo)
}

func learnUnnamedReturnTypes() int {
 return 3 
}

func learnNamedReturnTypes() (returnVal int) {
  returnVal = 5
  returnVal += 5
  return
}

func learnOutput() {
  // The fmt library is used for output operations in Go.
  // Print, Println and Printf are the mostly used methods.
  // The Print and Println functions work the same way as they do in kotlin.
  // Printf can be used to print interpolated output.

  // %v to print any kind of value
  fmt.Printf("\n%v, %v, %v\n", 10, "ten", true)

  // Below are some formatters for specific types.
  // These are most commonly used ones. There are several others present out there which can be explored.
  fmt.Printf("%d\n", 100)
  fmt.Printf("%f\n", 2.34)
  fmt.Printf("%s\n", "Hundred")
  fmt.Printf("%t\n", false)
  
  // %T can be used to print the type of a value
  fmt.Printf("The type of the given value is : %T\n", 100)
}

