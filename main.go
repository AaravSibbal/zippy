package main

import (
	"fmt"
	"log"
	"os"
)

/*
	TODO: make the general file zip
	TODO: make dir zip
	TODO: make a nice man
	TODO: add this to path
*/
const EXPECTED_ARGS = 3

func main() {

	argsLen := len(os.Args)
	WRONG_ARGS_LEN := fmt.Sprintf("Expected %d or got %d", EXPECTED_ARGS, argsLen)
	errLog := log.New(os.Stdout, "Error: ", log.Ldate|log.Ltime)

	if argsLen == 3 || argsLen > 3{
		errLog.Fatal(WRONG_ARGS_LEN)
		return
	}
	inputFile := os.Args[1]
	outputFile := os.Args[2]

	fmt.Println(inputFile, outputFile)
	os.Exit(1)
}