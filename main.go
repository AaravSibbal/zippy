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
	errlog *log.Logger
	infoLog *log.Logger
}

func main() {

	argsLen := len(os.Args)
	WRONG_ARGS_LEN := fmt.Sprintf("Expected %d or got %d", EXPECTED_ARGS, argsLen)


	errLog := log.New(os.Stdout, "Error: \t", log.Ldate|log.Ltime)
	infoLog := log.New(os.Stdout, "Info: \t", log.Ldate|log.Ltime)

	app := &application{
		infoLog: infoLog,
		errlog: errLog,

	}
	if argsLen == 3 || argsLen > 3{
		errLog.Fatal(WRONG_ARGS_LEN)
		return
	}
	
	option := os.Args[1]
	inputFile := os.Args[2]
	outputFile := os.Args[3]

	if option == "-f"	 {
		infoLog.Println("this option is for a file")
		app.zipFile(inputFile, outputFile)
	}
	os.Exit(1)
}