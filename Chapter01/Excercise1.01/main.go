package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

/* Prints random number of stars (*) */
func main()  {	
	rand.Seed(time.Now().UnixNano())
	r := rand.Intn(5) + 1
	stars := strings.Repeat("*",r)
	fmt.Println(stars)
	fmt.Println( "test ")
}