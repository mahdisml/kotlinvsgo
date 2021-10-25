package main

import "fmt"

func main() {
	var x []string = []string{"salam", "alo", "alo"}
	x = append(x, "3")
	for i := range x {
		if i == 1 {
			x = append(x[:1], x[2:]...)
		}
	}
	for i, s := range x {
		fmt.Println(i, ":", s)
	}
}
