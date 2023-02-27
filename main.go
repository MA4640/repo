package main

import "fmt"

func main() {
	my_function(1, 2)
}
func my_function(first, second int) {
	fmt.Print(first*second + 10)

}