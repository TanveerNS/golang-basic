package main

import (
	"fmt"
)

type Person struct {
	Name  string
	Age   int
	Alive bool
}

func main() {
	var a int = 10
	var b int32 = 20
	var c int64 = 30
	fmt.Println("Integer types:")
	fmt.Println("a (int):", a)
	fmt.Println("b (int32):", b)
	fmt.Println("c (int64):", c)

	var f1 float32 = 3.14
	var f2 float64 = 2.71828
	fmt.Println("\nFloating point types:")
	fmt.Println("f1 (float32):", f1)
	fmt.Println("f2 (float64):", f2)

	var str string = "Hello, Go!"
	fmt.Println("\nString:")
	fmt.Println("str:", str)

	var isTrue bool = true
	var isFalse bool = false
	fmt.Println("\nBoolean types:")
	fmt.Println("isTrue:", isTrue)
	fmt.Println("isFalse:", isFalse)

	var arr [3]int = [3]int{1, 2, 3}
	fmt.Println("\nArray:")
	fmt.Println("arr:", arr)

	var slice []int = []int{10, 20, 30, 40}
	fmt.Println("\nSlice:")
	fmt.Println("slice:", slice)

	var m map[string]int = map[string]int{
		"apple":  5,
		"banana": 3,
	}
	fmt.Println("\nMap:")
	fmt.Println("m:", m)

	p := Person{Name: "John", Age: 30, Alive: true}
	fmt.Println("\nStruct:")
	fmt.Println("p:", p)

	var ptr *int = &a
	fmt.Println("\nPointer:")
	fmt.Println("Pointer to a:", ptr)
	fmt.Println("Value pointed to by ptr:", *ptr)
}
