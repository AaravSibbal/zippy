package main

import (
	"log"
	"os"
)


const HELP = `Usage of zippy:

    - How to use:

        zippy [option] input output

        Examples:

            **Files options**
            "zippy -f some.txt":
                option: file option
                - input: some.txt
                - output: output.zip

            "zippy -f":
                This does not work because the option is set for file.
                Which requires file input

            "zippy -f some.txt publish":
                option: file option
                - input: some.txt
                - output: publish.zip

            "zippy -f some.txt .":
                option: file option
                - input: some.txt
                - output: output.zip


            **Directory options**
            "zippy -d some":
                option: directory option
                - input: some
                    Here some is a directory inside the working directory
                - output: output.zip

            "zippy -d .":
                option: directory option
                - input: content inside the currrent working directory
                - output: output.zip

            "zippy -d":
                option: directory option
                - input: content inside the currrent working directory
                - output: output.zip
            "zippy -d some publish":
                option: directory option
                - input: some is a directory inside the working directory
                - output: publish.zip

    - help:
        Help

        expample: "zippy help", "zippy"

    - [Option]:
        There are 2 options to choose from:
            1. -f: to zip files
            2. -d: to zip directories.
                Note: all directories are recursively zipped

    - input
        This can be either a file or a directory based on the
        optiton that you hae set earlier

        takes in the path, ex:
            - "file.txt"
            - "../../file.txt"
            - "/dir"
            - "../../dir"
            - ".": uses the content of the current directory as input

    - output
        this is always a file, the extension of the file always is .zip
        Examples:
            - "file": file.zip
            - "file.zip": file.zip.zip
            - "": output.zip
`
type application struct {
	helpCache  string
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
		helpCache:  HELP,
		option:     "-d",
		input:      ".",
		outputFile: "output.zip",
		infoLog: infoLog,
		errlog: errLog,
	}
	
	app.run()
}
