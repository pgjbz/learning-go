package main

type Person struct {
	name string
	age  int8
}

func (p *Person) printName() {
	println(p.name)
}

func (p *Person) printAge() {
	println(p.age)
}

func (p *Person) changeName(name string) {
	p.name = name
}

func (p Person) maybeChangeName(name string) { //receive a copy of value
	p.name = name
}

/*
  methos are functions, and can be writer without the "receptor argument"
  (receptor arugment is argument between func and name of function)

  this function:

  func printAge(person *Person) {
    println(person.name)
  }

  has the same effect as printName with receptor argument

*/

func main() {
	person := Person{"Paulo", 22}
	person.printAge()
	person.printName()
	person.changeName("My name")
	person.printName()
	person.maybeChangeName("Paulo")
	person.printName()
}
