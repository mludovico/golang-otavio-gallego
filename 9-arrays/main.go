package main

import (
	"fmt"
	"reflect"
)

func main() {
	var array1 [5]int
	array1[0] = 1
	fmt.Println(array1)

	array2 := [5]int{1, 2, 3, 4, 5}
	fmt.Println(array2)

	array3 := [...]int{1, 2, 3, 4, 5}
	fmt.Printf("array3: %v\n", array3)

	slice := []int{1, 2, 3, 4, 5, 6}
	slice = append(slice, 7)
	fmt.Println(slice)

	fmt.Println(reflect.TypeOf(array1), reflect.TypeOf(slice))

	slice = append(slice, 18)

	fmt.Println(slice)

	slice2 := array2[1:3]
	fmt.Println(slice2)

	array2[1] = 20
	fmt.Println(array2)

	slice3 := make([]int, 10, 12)

	fmt.Println(slice3, len(slice3), cap(slice3))

	slice3 = append(slice3, 1)
	fmt.Println(slice3, len(slice3), cap(slice3))

	slice3 = append(slice3, 2)
	fmt.Println(slice3, len(slice3), cap(slice3))

	slice3 = append(slice3, 3)
	fmt.Println(slice3, len(slice3), cap(slice3))
}
