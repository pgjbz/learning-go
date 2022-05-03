package main

func split(num int) (x, y int) {
	x = num * 4 / 9
	y = num - x
	return
}

func main() {
	println(split(25))
}
