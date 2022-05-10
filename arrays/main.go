package main

import (
	"fmt"
	"strings"
)

func main() {
	var a [2]string
	a[0] = "Paulo"
	a[1] = "Gabriel"

	for idx, name := range a {
		println(idx, " = ", name)
	}

	for _, name := range a { //use _ to ignore value, like in Rust
		println(name)
	}

	numbers := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	var slice []int = numbers[2:5]
	fmt.Println(slice)
	/*
	   slices in go are same properties of slices in Rust
	   slices points to memory range of memory,
	   change element value in a slice, change value in array
	*/
	slice[0] = 20
	fmt.Println(numbers)

	/*
	   literal slices are arrays without inform the size,
	   to create a slice dont give a size of array,
	   Go make a matriz and automatic create a slice of
	*/
	fmt.Println([]bool{true, true, false, true, true})

	/*
	   on create a slice, you can inform the range of slice,
	   but, you can ommits the lowers and bigger positions, like in Rust
	   let array = [0; 10];
	   let slice_array = &array[..];
	   let slice_array = &array[0..];
	*/
	fmt.Println(numbers[:10])
	fmt.Println(numbers[:])

	/*
	   slices have capacity and size,
	   size are the total of elements in slice
	   capacity are capacity of a slice,
	   to get len of slice use built in method len(slice),
	   to get total cap of slice use built in method cap(slice);
	*/

	sliceWithMake := make([]int, 0, 5) //the 3th argument is the capacity of slice

	fmt.Println(len(sliceWithMake))
	fmt.Println(cap(sliceWithMake))

	sliceWithMakeWithoutCapacity := make([]int, 5)
	fmt.Println(cap(sliceWithMakeWithoutCapacity))
	fmt.Println(len(sliceWithMakeWithoutCapacity))

	board := [][]string{
		{"_", "_", "_"},
		{"_", "_", "_"},
		{"_", "_", "_"},
	}

	/*slices can have any type, including another slice*/
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}

	var s []uint8
	//append works with nil slices
	s = append(s, 10) //append method alloc another slice
	s = append(s, 20) //if slice cap cannot have capacity, another slice are alloc and
	s = append(s, 30) //point to the new slice
	fmt.Println(s)

}
