/* Hello World with GoLang */

package main

import (
	"fmt"
	"reflect"
	)

func main() {
	/* Hello World */
	fmt.Println("Hello World")

	/* Variables */
	var Num_A int = 10 // variable declaration
	var Byte_B byte = 'b'
	Str_C := "I am C"
	fmt.Println(Num_A)					  // print variable Num_A: 10
	fmt.Println(string(Byte_B))			  // print variable Str_B: b
	fmt.Println(reflect.TypeOf(Byte_B))	  // print variable Str_B type: uint8我又想起当年错过某提交时间 通过修改本机时间提交的骚操作
	fmt.Println("%d, %c\n", Byte_B, Byte_B) // print variable Str_B: byte
	fmt.Println(Str_C)					  // print variable Str_C: I am C

	/* Array & Slice */
	var Arr_A [5]int = [5]int{1, 2, 3, 4, 5}
	//or Arr_A := [5]int{1, 2, 3, 4, 5}
	for i:=0; i < len(Arr_A); i++{
		fmt.Println(Arr_A[i])
		Arr_A[i] += 100
	}
	// Slice is a reference to an array
	for i:=0; i < len(Arr_A); i++{
		fmt.Println(Arr_A[:i])
		fmt.Println(Arr_A[i:])
		fmt.Println(Arr_A[i:i+1])
	}

	/* Dictionary/Map */
	// Map_A := make(map[string]int)
	Map_B := map[string]int{"A": 1, "B": 2, "C": 2}
	Map_B["D"] = 4 // add new key-value pair
	Map_B["B"] = 0 // modify key-value pair
	delete(Map_B, "C") // delete key-value pair
	fmt.Println(Map_B)

	/* Pointer */ 
	Str_D := "Golang"
	var Pointer_Str_D *string = &Str_D
	fmt.Println(Str_D)
	*Pointer_Str_D = "Hello World" // change value of Str_D
	fmt.Println(Str_D)

	/* flow control */
	// if else
	var Age int
	fmt.Println("Please input the age:")
	fmt.Scanf("%d", &Age) // scanning and reading the input
	if Age <= 12 {
		fmt.Println("Primary School")
	} else if Age <= 15 {
		fmt.Println("Middle School")
	} else { 
		fmt.Println("High School")
	}

	// switch
	type Animal int8 // define a type called Gender with two constants - MALE/FEMALE
	const (
		DOG Animal = 1
		CAT Animal = 0
	)
	animal := DOG
	switch animal {
		case DOG:
			fmt.Println("dog")
		case CAT:
			fmt.Println("cat")
		default:
			fmt.Println("unknown")
	}

	// for loop
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
	// for loop with array
	for i:=0; i < len(Arr_A); i++{
		fmt.Println(Arr_A[i])
		Arr_A[i] += 100
	}
	// Slice is a reference to an array
	for i:=0; i < len(Arr_A); i++{
		fmt.Println(Arr_A[:i])
		fmt.Println(Arr_A[i:])
		fmt.Println(Arr_A[i:i+1])
	}

	/* function: addition */
	fmt.Println(addition_form_one(1, 2))
	fmt.Println(addition_form_two(2, 3))
}
