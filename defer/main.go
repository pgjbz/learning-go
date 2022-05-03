package main

func main() {
	defer print(", world!")
	print("hello")
}
