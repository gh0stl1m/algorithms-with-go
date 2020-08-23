package main

import (
	"fmt"

	dc "github.com/gh0stl1m/algorithms-with-go/dynamicconnectivity"
)

func main() {
	quickFind := new(dc.QuickFind)
	quickFind.Init(10)
	dc.FileReader("union-find-test.txt", quickFind)

	fmt.Println("Are points linked? ", quickFind.Connected(9, 3))
	fmt.Println("Are points linked? ", quickFind.Connected(0, 4))

}
