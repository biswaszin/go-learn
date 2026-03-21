// This is the main file.

package main

import "fmt"

func Hello(word string) string{
	return "Hello, " + word
}

func main(){
	fmt.Println(Hello("Chris"))
}
