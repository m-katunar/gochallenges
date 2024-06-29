package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Golang sorting tutorial")

	myInt := []int{1, 4, 2, 3, 5, 8, 7, 6, 9}

	fmt.Println("This stuff is unsorted.")

	fmt.Println(myInt)

	fmt.Println("All sorted out now.")
	sort.Ints(myInt)
	fmt.Println(myInt)

}
