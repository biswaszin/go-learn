// This is the main file.

package main

import (
	"fmt"
	"math"
)

const englishHelloPrefix = "Hello, "
const spanish = "Spanish"
const spanishHelloPrefix = "Hola, "

func Hello(name string, language string) string{
	if name == ""{
		name = "World"
	}
	if language == spanish{
		return spanishHelloPrefix + name
	}
	return englishHelloPrefix + name 
}

func justLoops() int {
	sum := 0
	for i := 0; i < 10 ; i++ {
		sum += 1
	}

	return sum
}

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}

	return fmt.Sprint(math.Sqrt(x))
}

func main(){
	fmt.Println(Hello("World", "English"))

	random := justLoops()
	fmt.Println(random)
}


