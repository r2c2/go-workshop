package main

import (
	"fmt"
	"time"
)

func main() {
	//Declare pointer
	var count1 *int

	count2 := new(int)

	/*Create pointer from variable
	You can't take the address of a literal number.
	Create a temporary variable to hold a number: */
	countTemp := 5
	count3 := &countTemp

	//It's possible to create a pointer from some types without a temporary variable
	t := &time.Time{}

	fmt.Printf("count1: %#v\n", count1)
	fmt.Printf("count2: %#v\n", count2)
	fmt.Printf("count3: %#v\n", count3)
	fmt.Printf("time : %#v\n", t)

}
