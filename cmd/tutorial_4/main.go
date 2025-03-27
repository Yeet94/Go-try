package main
import "fmt"

func main(){
	intArr := [3]int{1, 2, 3}
	fmt.Println(intArr[0])
	fmt.Println(intArr[1:3])
	fmt.Println(&intArr[0])

	var myMap map[string]uint8 = make(map[string]uint8)
	myMap["key1"] = 1
	myMap["key2"] = 2
	fmt.Println(myMap)
}