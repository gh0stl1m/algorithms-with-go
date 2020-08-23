package main

import (
	"fmt"

	dc "github.com/gh0stl1m/algorithms-with-go/dynamicconnectivity"
)

func main() {
	quickFind := new(dc.QuickFind)
	quickFind.Init(10)
	quickUnion := new(dc.QuickUnion)
	quickUnion.Init(10)
	dc.FileReader("union-find-test.txt", quickFind)
	dc.FileReader("union-find-test.txt", quickUnion)

	fmt.Println("Are points linked using qf? ", quickFind.Connected(9, 3))
	fmt.Println("Are points linked unding qf?", quickFind.Connected(0, 4))
	fmt.Println("Are points linked using qu? ", quickUnion.Connected(9, 3))
	fmt.Println("Are points linked unding qu?", quickUnion.Connected(0, 4))

}
