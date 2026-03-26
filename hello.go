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

func justDefer() string {
	fmt.Println("Counting")
	
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
	return "Done"
}

func justPointers() string {
	fmt.Println("Learning Pointers")

	i, j := 31, 68

	p := &i //pointing towards i
	// now to print the data of the i itself, we do
	fmt.Println(*p)

	// to change the value of i from p
	*p = 23
	fmt.Println(i, j) 

	return "Done"
}

// to define a struct 

type coordinates struct {
	x int
	y int
}

func justArrays() string{
	arr := [2]string{}
	arr[0] = "Hello"
	arr[1] = "World"

	fmt.Println(arr[0], arr[1])

	return "Done"
}

func justSlices() string{
	namesArray := [4]string {
		"Emma Stone",
		"Monu",
		"Biswaszin",
		"Dig",
	}
	setOfNamesA := namesArray[1:3]
	setOfNamesB := namesArray[0:2]
	fmt.Println(setOfNamesA, setOfNamesB)
	setOfNamesB[1] = "Fornite"
	fmt.Println(setOfNamesB)
	return "Done"
}

func main(){
	fmt.Println(Hello("World", "English"))
	random := justLoops()
	fmt.Println(random)
	fmt.Println(justDefer())
	fmt.Println(justPointers())
	
	point := coordinates{}
	point.x = 2
	point.y = 4
	
	fmt.Println(point.x, point.y)
	another_point := &point
	another_point.x, another_point.y = 4,8
	fmt.Println(point.x, point.y)

	fmt.Println(justArrays())
	fmt.Println(justSlices())

	// setting up the git so making a random commit
}


