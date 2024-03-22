package main

import (
	"archive/zip"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
)

// TODO: I CAN ALSO ADD A TEST TO CHECK IF THE CONTENT OF THE FILE IS NOT CORRECTED

const OPEN_FILE = "Opening the file: "
const ZIP_FILE = "Zipping the file: "
const WRITE_FILE = "Writing file: "

const ERR_OPEN_FILE = "Couldn't open the file"
const ERR_CREATE_ZIP = "There was an error creating the output file: "
const ERR_CREATE_BUFFER = "There was an error creating the buffer"
const ERR_COPY_BUFFER = "There was an error copying the file into the buffer writer"
const ERR_READ_DIR = "There was an error reading the directory"

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
	err := filepath.WalkDir("test", func(path string, d fs.DirEntry, err error) error {
		fmt.Println(path, d.Name(), "directory?", d.IsDir())
		return nil
	})

	if err != nil {
		app.errlog.Fatalf("%s %v", ERR_READ_DIR, err)
	}
}
