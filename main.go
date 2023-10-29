package main

import (
	"fmt"
	"strings"
)

func main() {

	var myVariable, my1Variable, my3Variable, my4Variable = "welome", "newJourney", 1, -5.5
	fmt.Printf("%s\n", myVariable)
	fmt.Printf("The type of myVariable is : %T\n", myVariable)

	fmt.Printf("%s\n", my1Variable)
	fmt.Printf("The type of my1Variable is : %T\n", my1Variable)

	fmt.Printf("%d\n", my3Variable)
	fmt.Printf("The type of my3Variable is : %T\n", my3Variable)
	fmt.Printf("%f\n", my4Variable)

	fmt.Printf("The type of my4Variable is : %T\n", my4Variable)

	//take input

	var a int = 50
	if a%2 == 0 {

		fmt.Printf("a is even number \n")
	} else {

		fmt.Printf("\n ,a  is odd number")
	}

	var b string = "welcome"
	var c string = "Welcome"
	var d string = "Welcome"

	if b == c {
		fmt.Printf("%s\n", "b and c are equal")
	} else if b == d {
		fmt.Printf("%s\n", "b and d are equal")
	} else if c == d {
		fmt.Printf("%s\n", "c and d are equal")
	} else {
		fmt.Printf("all are not equal")
	}

	//switch fallthrough

	dayOfWeek := 3

	switch dayOfWeek {
	case 1:
		fmt.Println("Sunday")

	case 2:
		fmt.Println("Monday")

	case 3:
		fmt.Println("Tuesday")
		fallthrough

	case 4:
		fmt.Println("Wednesday")

	case 5:
		fmt.Println("Thursday")

	case 6:
		fmt.Println("Friday")

	case 7:
		fmt.Println("Saturday")

	default:
		fmt.Println("Invalid day")
	}

	//strings

	// var message = "vandhana"

	// var stringLength = len(message)

	// fmt.Println("Length of a string is:", len(stringLength))
	string1 := "vandhana"
	string2 := "Vandhana"
	string3 := "vandhana"

	//compare strings
	compareName := string1 == string3
	fmt.Println(compareName) // returns true or false

	fmt.Println(strings.Compare(string1, string2)) //-1
	fmt.Println(strings.Compare(string2, string3)) //1
	fmt.Println(strings.Compare(string1, string3)) //0
	text := "Vandhana go to new journey"
	fmt.Println(strings.Contains(text, string3))       //returns true or false  cheak whether the item is inside a string or not
	replacedText := strings.Replace(text, "a", "n", 1) //replace a value

	fmt.Println("New String:", replacedText)

	fmt.Println(strings.ToUpper(string1))
	fmt.Println(strings.ToLower(text))

	fmt.Println(strings.Split(text, " "))

	words := []string{"I", "love", "Golang"}

	//join each element of the slice
	add := strings.Join(words, " ")
	fmt.Println(add)

	//arrays

	var fruits = [...]string{"apple", "mango", "banana", "orange"} //inferredlength
	fmt.Println(fruits)
	// length of an array
	fmt.Println(len(fruits))
	//slicing an array
	mySlice := fruits[0:4] //index,length

	fmt.Printf("mySlice = %v\n", mySlice)

	var personAge = [5]int{12, 13, 14, 13} //specified length
	fmt.Println(personAge)

	// slice provides various builtin functions

	// adding elements to an array
	var addNum = []int{2, 3, 6}

	res := append(addNum, 5, 7)

	fmt.Println("Add:", res)

	//accessing an array
	fmt.Println(personAge[0])
	fmt.Println(fruits[1])

	//change the elements of an array

	fruits[1] = "grapes"
	fmt.Println(fruits)

	//intialize specific elements in an array

	arr1 := [5]int{1: 10, 2: 40}

	fmt.Println(arr1)

	var arrayOfIntegers = [...]int{1, 5, 8, 0, 3, 10}

	// find the length of array using len()
	length := len(arrayOfIntegers)

	fmt.Println("The length of array is", length)

	//throws error
	// numbers := [5]int{11, 22, 33, 44, 55}

	// // index variable is declared but not used
	// for index, item := range numbers {
	// 	// throws an error
	// 	fmt.Println(item)
	// }

	// avoid this error we can use blank identifier in this case second par of the array item indicates

	numbersBlank := [5]int{11, 22, 33, 44, 55}
	for _, item := range numbersBlank {
		fmt.Println(item)
	}

	//multidimenshional array
	array2D := [2][2]int{{1, 2}, {3, 4}}

	// access the values of 2d array
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			fmt.Println(array2D[i][j])
		}
	}

	numbers := [8]int{10, 20, 30, 40, 50, 60, 70, 80}

	// create slice from an array
	sliceNumbers := numbers[4:7]

	fmt.Println(sliceNumbers)

	// go language contains only forloop

	for i := 0; i < 20; i++ {

		fmt.Printf("\n helloworld")
	}

	//forloop acts as infinite loop

	// for {
	// 	fmt.Printf("\n hi")
	// }
	//infinite loop no condition

	//for loop acts as while loop

	// i := 0
	// for i < 3 {
	// 	i += 6

	// }

	arr := []string{"apple", "mango", "banana"}

	for i, j := range arr {
		fmt.Println(i, j)
	}

	// using numbers
	numbers_range := [5]int{21, 24, 27, 30, 33}

	for index, item := range numbers_range {
		fmt.Printf("numbers[%d] = %d \n", index, item)
	}

	//using strings

	string := "Balu"
	fmt.Println("Index: Character")

	for i, item := range string {
		fmt.Printf("%d= %c \n", i, item)
	}

	// var name string

	// fmt.Print("Enter your Name:")
	// fmt.Scan("&name")

	// fmt.Printf(name)

	// var Name string
	// var age int

	// // take name and age input
	// fmt.Println("Enter your name and age:")
	// fmt.Scan(Name, &age)

	// // print input values
	// fmt.Printf("Name: %s\nAge: %d", Name, age)

}
