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

type Coordinates struct {
	Lat, Long float64
}

type Vertex struct{
	Lat, Long float64
}

var m map[string]Vertex

func justMaps(){
	m = make(map[string]Vertex)
	m["Adamas"] = Vertex{
		40.11290387, -85.91023809,
	}
	fmt.Println(m["Adamas"])
}

func justMakeMaps(){
	m := make(map[string]int)

	m["Answer"] = 42
	fmt.Println("The value is: ", m["Answer"])

	m["Answer"] = 48
	fmt.Println("The value is: ", m["Answer"])

	delete(m, "Answer")
	fmt.Println("The value is: ", m["Answer"])

	v, ok := m["Answer"]
	fmt.Println("The value is: ", v, "Present?", ok)

	// let's add another key value pair 
	
	m["AU/2025/0003743"] = 333
	// To see the key and the value using range, we do
	for key, value := range m {
		fmt.Println(key,value)
	}
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

func justMake(){
	a := make([]int, 4, 5) //make is another way to initialize slices, like
	printSlices("a", a) // var := make(data type, length, capacity)
	b := make([]int, 5) // length only 5, and so is capacity 5 by default
	c := b[:2]
	printSlices("c", c)
}

func justAppend(){
	var slice_eg []int
	slice_eg = append(slice_eg, 5)
	printSlices("slice_eg", slice_eg)
}

func justAnonymousFunc() {
	add := func (x, y int) int{
		return x+y
	}

	fmt.Printf("Addition: %v\n", add(5,6))
}


func printSlices(s string, x []int) {
	fmt.Printf("Var: %s, Len: %v, Cap: %v, Var: %v\n", s, len(x), cap(x), x)
}

// Understanding the Methods in Go
// So far, it's kinda like member functions from classes in C++
// You can essentially do the same here in Go. For example

type Position struct{
	X int
	Y int
}

func (point Position) ChangeInPosition() Position{
	return Position{
		point.X + 10,
		point.Y + 8,
	}
}

func main(){
	fmt.Println(Hello("World", "English"))
	random := justLoops()
	fmt.Println(random)
	fmt.Println(justDefer())
	fmt.Println(justPointers())
	
	point := Coordinates{}
	point.Lat = 2.1231
	point.Long = 4.4123
	
	fmt.Println(point.Lat, point.Long)
	another_point := &point
	another_point.Lat, another_point.Long = 4,8
	fmt.Println(point.Lat, point.Long)

	fmt.Println(justArrays())
	fmt.Println(justSlices())

	// setting up the git so making a random commit
	
	justMake()
	justAppend()
	justMaps()
	justMakeMaps()
	justAnonymousFunc()

	// For the basic ass physics engine (it's far from engine. i jus wanna see how methods work)
	pointA := Position{5,8}
	fmt.Println(pointA.ChangeInPosition())
}


