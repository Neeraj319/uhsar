package main

import (
	"bufio"
	"fmt"
	"os"
)

type Scanner struct {
	Source string
}

func (s *Scanner) scanTokens() {

}

func error(line int, where string, message string) {
	report(line, where, message)

}
func report(line int, where string, message string) {
	error_string := fmt.Errorf("line %d Error %s : %s ", line, where, message)
	fmt.Println(error_string)

}
func run(source string) {
	scanner := Scanner{source}
	tokens := scanner.scanTokens()

	for _, token := range tokens {
		fmt.Println("token", token)
	}

}

func readFile(fileName string) (string, error) {
	data, err := os.ReadFile(fileName)
	if err != nil {
		error_string := fmt.Sprintf("Error reading file %s", fileName)
		return "", fmt.Errorf(error_string)
	}
	return string(data), nil
}

func runThroughFile(fileName string) {
	data, error := readFile(fileName)
	if error != nil {
		fmt.Println(error)
		return
	}
	run(data)

}

func runThroughInput() {
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print(">>> ")
		text, _ := reader.ReadString('\n')
		run(text)
	}

}

func main() {
	args := os.Args
	if len(args) == 2 {
		runThroughFile(args[1])
	} else if len(args) == 1 {
		runThroughInput()

	}

}
