package main

import "fmt"

func main(){
 fmt.Println("Addition", add(3,4))
 fmt.Println("Subtraction", subtract(3,4))
}

func add(a int, b int) int {
    return a + b;
}

func subtract(a int, b int) int {
    return a - b;
}