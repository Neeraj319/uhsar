package main

import (
	"bufio"
	"fmt"
	"os"
)

func run(source string) {
	scanner := CreateScanner(source)
	tokens := scanner.ScanTokens()
	parser := CreateParser(tokens)
	expression := parser.Parse()
	fmt.Println(Interpret(expression))
}

func runThroughFile(fileName string) {
	data, error := ReadFile(fileName)
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
