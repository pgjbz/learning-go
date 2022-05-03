package main

var java, rust, golang bool          //default values is false false false
var node, python, deno int = 1, 2, 3 //initial value for each variable

func main() {
	var i int                                          //defaul value is 0
	var banana, anyNumber, quiabo = "fruit", 10, false //type inference for each variable
	println(java, rust, golang, node, python, deno, i)
	println(banana, anyNumber, quiabo)
	j := 10 //sintaxe sugar for variable declaration
	//:= operator cannot use outside functions
	println(j)
}
