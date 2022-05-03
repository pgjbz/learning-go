package main

const PI = 3.14

const (
	Big   = 1 << 100 //declare constants
	Small = Big >> 99
)

func main() {

	println(PI, (Big >> 80), Small)

}
