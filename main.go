package main

import (
	"bufio"
	"fmt"
	"os"
)

func run(source string) {
	scanner := CreateScanner(source)
	fmt.Println("scanner is", scanner)
	tokens := scanner.ScanTokens()

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
