package dynamicconnectivity

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// FileReader is a client that reads in the first line
// the number of elements that will create and after
// pairs of numbers
func FileReader(file string, dc UnionFind) {
	f, e := os.Open(file)
	errorChecker(e)

	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		points := strings.Split(scanner.Text(), " ")
		p, err := strconv.ParseUint(points[0], 10, 64)
		errorChecker(err)
		q, err := strconv.ParseUint(points[1], 10, 64)
		errorChecker(err)

		if !dc.Connected(p, q) {
			dc.Union(p, q)
			fmt.Printf("%v ---> %v \n", p, q)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}

func errorChecker(e error) {
	if e != nil {
		panic(e)
	}
}
