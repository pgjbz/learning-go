package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

func main() {
	m := make(map[string]Vertex) //map[keyType]valueType
	/*
	   map has to be initialized with make
	   the default value in map is nil,
	   and nil map cannot hold values
	*/
	m["banana"] = Vertex{Lat: 41.2, Long: 44.50}
	/*
	   to set a value in a map, simple use mapName["name"]=value
	   to access use same method bug not sign a value
	*/
	fmt.Println(m["banana"])
	/*
	   maps can be literal
	*/

	anotherMap := map[string]Vertex{
		"banana": {Lat: 41.2, Long: 44.50},
		"eoq":    {44.5, 22.9}, //declare type in this are redudant
	}
	fmt.Println(anotherMap)

	delete(anotherMap, "eoq")        //delete key
	elem, ok := anotherMap["banana"] //get key if present ok is true, use := if variables is not declared yet
	fmt.Println(elem)
	fmt.Println(ok)
	fmt.Println(anotherMap)
}
