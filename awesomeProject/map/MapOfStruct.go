package main

import (
	"fmt"
)

type Vertex struct {
	Latitude, Longitude float64
}

var structMap = map[string]Vertex{
	"rokomari": Vertex{
		Latitude:  12.0120120,
		Longitude: 12.212121,
	},
	"aust": Vertex{
		Latitude:  10.10101,
		Longitude: 11.11111,
	},
}

func main() {
	fmt.Println(structMap)
	fmt.Println(structMap["aust"])

	for key, value := range structMap {
		fmt.Println("Key : ", key, " Value : ", value)
	}

}
