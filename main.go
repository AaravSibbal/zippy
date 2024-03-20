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
const EXPECTED_ARGS = 4

type application struct {
	input string
	outputFile string
	option string
	errlog  *log.Logger
	infoLog *log.Logger
}

func main() {

	option := "-d"
	input := "."
	outputFile := "."

	argsLen := len(os.Args)
	WRONG_ARGS_LEN := fmt.Sprintf("Expected %d got %d", EXPECTED_ARGS, argsLen)

	errLog := log.New(os.Stdout, "Error: \t", log.Ldate|log.Ltime)
	infoLog := log.New(os.Stdout, "Info: \t", log.Ldate|log.Ltime)

	app := &application{
		option: "-d",
		input: ".",
		outputFile: "output.zip",
		infoLog: infoLog,
		errlog:  errLog,
	}

	if argsLen < EXPECTED_ARGS {
		errLog.Fatal(WRONG_ARGS_LEN)
		return
	}
	if os.Args[1] == "" {

	}

	option = os.Args[1]
	input = os.Args[2]
	outputFile = os.Args[3]

	if option == "-f" {
		infoLog.Println("this option is for a file")
		app.zipFile(input, outputFile)
	}
	// os.Exit(1)
}
