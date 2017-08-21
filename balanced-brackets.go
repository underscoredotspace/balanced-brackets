package main

import "fmt"
import "strconv"
import "log"

type stack []rune
type bracketTypeList map[rune]rune

func (s *stack) push(newTop rune) {
	*s = append(*s, newTop)
}

func (s *stack) pop() rune {
	stackLen := len(*s) - 1
	topElement := (*s)[stackLen]
	*s = (*s)[:stackLen]
	return topElement
}

func main() {
	lines, err := getLines()
	if err != nil || lines < 1 || lines > 1000 {
		log.Println("Number of strings must be >= 1 and <= 10^3")
		return
	}

	brackets := make(bracketTypeList)
	brackets[0x0028] = 0x0029
	brackets[0x007b] = 0x007d
	brackets[0x005b] = 0x005d

	for line := 0; line < lines; line++ {
		var lineString string
		fmt.Scanln(&lineString)

		if brackets.balanced(lineString) {
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

func (bracketTypes bracketTypeList) balanced(lineString string) bool {
	lineRunes := []rune(lineString)
	if len(lineRunes)%2 != 0 {
		return false
	}

	var bracketStack stack

	for _, r := range lineRunes {
		if match, ok := bracketTypes[r]; ok {
			bracketStack.push(match)
			continue
		}

		if bracketStack.pop() != r {
			return false
		}
	}

	if len(bracketStack) != 0 {
		return false
	}
	return true
}
