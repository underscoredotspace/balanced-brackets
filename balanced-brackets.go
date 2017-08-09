package main

import "fmt"
import "strconv"
import "log"

type stack []string

func (s stack) Push(newString string) (stack stack) {
	stack = append(s, newString)
	return
}

func (s stack) Pop() (stackTop string, stack stack) {
	l := len(s)
	if l == 0 {
		return
	}
	stackTop = s[l-1]
	stack = s[:l-1]
	return
}

func main() {
	lines, err := getLines()
	if err != nil || lines < 1 || lines > 1000 {
		log.Println("Number of strings must be >= 1 and <= 10^3")
		return
	}

	for line := 0; line < lines; line++ {
		var lineString string
		fmt.Scanln(&lineString)

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
	lineRunes := []rune(lineString)
	if l := len(lineRunes); l%2 != 0 || l > 1000 || l < 1 {
		return false
	}

	rightLeft := make(map[string]string)

	rightLeft["("] = ")"
	rightLeft["["] = "]"
	rightLeft["{"] = "}"

	var bracketStack stack

	for _, r := range lineRunes {
		if match, ok := rightLeft[string(r)]; ok {
			bracketStack = bracketStack.Push(match)
			continue
		}

		if l := len(bracketStack); l > 0 && string(r) == bracketStack[l-1] {
			_, bracketStack = bracketStack.Pop()
		} else {
			return false
		}

	}

	if len(bracketStack) != 0 {
		return false
	}
	return true
}
