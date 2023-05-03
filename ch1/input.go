package main

import "fmt"

func main() {
	var str1, str2 string
	n, _ := fmt.Scan(&str1, &str2)
	fmt.Println("input cnt:", n)
	fmt.Println(str1, str2)
}
