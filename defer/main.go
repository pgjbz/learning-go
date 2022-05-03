package main

func main() {
	defer print(", world!")
	print("hello")

	for i := 0; i < 10; i++ {
		defer println(i)
	}
	/*
	   defer is a stack call
	*/

	println("done")
}
