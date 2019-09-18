package main

import "fmt"

func main() {

	sum := Calculatesum(5, 10)
	fmt.Println(sum)

}

func Calculatesum(a int, b int) int {
	return (a + b)
}
