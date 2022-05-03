package main

func swap[T interface{}](x, y T) (T, T) {
	return y, x
}

func fakeFibo() (int8, int8, int8) {
	return 0, 1, 1
}

func main() {
	hello, world := swap("world", "hello")
	println(hello, world)
	zero, one, zeroPlusOne := fakeFibo()
	println(zero, one, zeroPlusOne)
}
