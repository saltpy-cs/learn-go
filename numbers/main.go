package main

import "fmt"

func main() {
	ns := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for _, n := range ns {
		if n%2 == 0 {
			fmt.Println(n, "is even")
		} else {
			fmt.Println(n, "is odd")
		}
	}
}
