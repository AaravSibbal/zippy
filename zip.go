package main

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
)

// TODO: I CAN ALSO ADD A TEST TO CHECK IF THE CONTENT OF THE FILE IS NOT CORRECTED

const OPEN_FILE = "Opening the file: "
const ZIP_FILE = "Zipping the file: "
const WRITE_FILE = "Writing file: "

const ERR_OPEN_FILE = "Couldn't open the file"
const ERR_CREATE_ZIP = "There was an error creating the output file: "
const ERR_CREATE_BUFFER = "There was an error creating the buffer"
const ERR_COPY_BUFFER = "There was an error copying the file into the buffer writer"

func (app *application) zipFile(input, outputFile string) {
	// TODO: GET ZIP FILE WORKING
	archive, err := os.Create(outputFile)
	if err != nil {
		app.errlog.Fatalf("%s %s", ERR_CREATE_ZIP, outputFile)
	}
	defer archive.Close()
	zipWrirter := zip.NewWriter(archive)

	app.infoLog.Printf("%s %s", OPEN_FILE, input)
	file, err := os.Open(input)
	if err != nil {
		app.errlog.Fatalf("%s %v", ERR_OPEN_FILE, err)
	}
	defer file.Close()

	app.infoLog.Printf("%s %s", WRITE_FILE, input)
	writer, err := zipWrirter.Create(input)
	if err != nil {
		app.errlog.Fatalf("%s %v", ERR_CREATE_BUFFER, err)
	}

	if _, err = io.Copy(writer, file); err != nil {
		app.errlog.Fatalf("%s %v", ERR_COPY_BUFFER, err)
	}
	zipWrirter.Close()
}

func (app *application) zipDirectory(input, outputFile string) {
	fmt.Println("add the zip dir code here")
	// TODO: GET THE ZIP DIRECTORY FUNCTION WORKING
	fmt.Println(input, outputFile)

}

func (app *application) getDir(input string) {
	// TODO: GET THE GET DIR FUNCTION WORKING

}
