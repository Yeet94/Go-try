package main

import "fmt"

func main(){
	var intNum int64 = 42
	fmt.Println(intNum)

	var floatNum float64 = 3.14159
	fmt.Println(floatNum)

	var floatNum2 float32 = 6.8888
	var intNum2 int32 = 89
	var result float32 = float32(intNum2) + floatNum2
	fmt.Println(result)

	var intNum3 int = 42
	var intNum4 int = 89
	fmt.Println(intNum3/intNum4)

	var myString string = "Hello World!"
	fmt.Println(myString)

	myVar := "Text"
	fmt.Println(myVar)
}