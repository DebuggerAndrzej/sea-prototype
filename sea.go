package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	switch numberOfArgsPassed := len(os.Args); numberOfArgsPassed {
	case 2:
		filepath := os.Args[1]
		executeFile(filepath)
	case 1:
		enterPrompt()
	default:
		fmt.Println("Usage: sea or sea <filename>")
		os.Exit(64)
	}
}

func executeFile(filepath string) {
	contents, err := os.ReadFile(filepath)
	if err != nil {
		fmt.Print(err)
	}
	run(string(contents))
}

func enterPrompt() {
	fmt.Print("> ")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		fmt.Print("> ")
		run(scanner.Text())
	}

	if scanner.Err() != nil {
		fmt.Println("Encountered error during reading line")
	}
}

func run(codeToExecute string) {
	tokens := getTokensFromText(codeToExecute)
	fmt.Println(tokens)
}

func reportError(lineNumber int, errorMessage string) {
	fmt.Printf("[%d] Error at %s -> %s", lineNumber, "", errorMessage)
	os.Exit(65)
}
