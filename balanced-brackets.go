package main

import "fmt"
import "strconv"
import "log"

func main() {
	//Enter your code here. Read input from STDIN. Print output to STDOUT
	lines, err := getLines()
	if err != nil || lines < 1 || lines >= 1000 {
		log.Println("Number of strings must be >= 1 and <= 10^3")
	}

	for line := 0; line < lines; line++ {
		var lineString string
		fmt.Scanln(&lineString)
		log.Println(len(lineString))

		if balancedBrackets(lineString) {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}

func getLines() (lines int, err error) {
	var countString string

	fmt.Scanln(&countString)
	lines, err = strconv.Atoi(countString)

	return
}

func balancedBrackets(lineString string) bool {
	r := []rune(lineString)
	if len(r)%2 != 0 {
		return false
	}

	brackets := make(map[rune]rune)

	brackets[40] = 41   // ()
	brackets[123] = 125 // {}
	brackets[91] = 93   // []

	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		if brackets[r[i]] != r[j] {
			return false
		}
	}
	return true
}
