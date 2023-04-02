package main

import (
	"ch16/ex16.2/publicpkg"
	"fmt"
)

func main() {
	fmt.Println("PI:", publicpkg.PI)

	publicpkg.PublicFunc()

	var myint publicpkg.MyInt = 10
	fmt.Println("myint:", myint)

	//	var mystruct = publicpkg.MyStruct{Age: 18}
	var mystruct = publicpkg.MyStruct{name: "highschool"}

	fmt.Println("mystruct:", mystruct)
}
