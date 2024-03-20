package main

import (
	"log"
	"os"
)

/*
TODO: make the general file zip
TODO: make dir zip
TODO: add this to path
*/
const EXPECTED_ARGS = 4

type application struct {
	input      string
	outputFile string
	option     string
	errlog     *log.Logger
	infoLog    *log.Logger
}

/*
	sets the defaults and runs the application
*/

func main() {

	errLog := log.New(os.Stdout, "Error: \t", log.Ldate|log.Ltime)
	infoLog := log.New(os.Stdout, "Info: \t", log.Ldate|log.Ltime)

	app := &application{
		option:     "-d",
		input:      ".",
		outputFile: "output.zip",
		infoLog:    infoLog,
		errlog:     errLog,
	}

	app.run()
}
