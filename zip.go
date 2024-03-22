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
	archive, err := createArchive(outputFile)
	if err != nil {
		app.errlog.Fatalf("%s %s", ERR_CREATE_ZIP, outputFile)
	}
	defer archive.Close()
	zipWriter := zip.NewWriter(archive)
	defer zipWriter.Close()

	app.writeFileToArchive(zipWriter, input)

}

func createArchive(name string) (*os.File, error) {
	archive, err := os.Create(name)
	if err != nil {
		return nil, err
	}

	return archive, nil
}

func (app *application) writeFileToArchive(zipWriter *zip.Writer, input string) {

	app.infoLog.Printf("%s %s", OPEN_FILE, input)
	file, err := os.Open(input)
	if err != nil {
		app.errlog.Fatalf("%s %v", ERR_OPEN_FILE, err)
	}

	app.infoLog.Printf("%s %s", WRITE_FILE, input)
	writer, err := zipWriter.Create(input)
	if err != nil {
		app.errlog.Fatalf("%s %v", ERR_CREATE_BUFFER, err)
	}

	if _, err = io.Copy(writer, file); err != nil {
		app.errlog.Fatalf("%s %v", ERR_COPY_BUFFER, err)
	}
	file.Close()
}

func changePath(path string) *string {
	newPath := ""

	for i := 0; i < len(path); i++ {
		if string(path[i]) == `\` {
			newPath += "/"
		} else {
			newPath += string(path[i])
		}
	}

	return &newPath
}

func (app *application) zipDirectory(input, outputFile string) {
	archive, err := os.Create(outputFile)
	if err != nil {
		app.errlog.Fatalf("%s %v", ERR_CREATE_ZIP, err)
	}
	zipWriter := zip.NewWriter(archive)
	defer zipWriter.Close()

	err = filepath.WalkDir(input, func(path string, d fs.DirEntry, err error) error {
		fmt.Println(path, d.Name(), "directory?", d.IsDir())
		if d.IsDir() {
			// if d.Name()
			// do nothing
		} else {
			// app.writeFileToArchive(zipWriter, path)
			file, err := os.Open(path)
			pathForWriter := changePath(path)
			if err != nil {
				app.errlog.Fatalf("%s %v", ERR_OPEN_FILE, err)
			}
			defer file.Close()

			writer, err := zipWriter.Create(*pathForWriter)
			if err != nil {
				app.errlog.Fatalf("%s %v", ERR_CREATE_BUFFER, err)
			}
			_, err = io.Copy(writer, file)
			if err != nil {
				app.errlog.Fatalf("%s %v", ERR_COPY_BUFFER, err)
			}
			file.Close()

		}
		return nil
	})

	if err != nil {
		app.errlog.Fatalf("%s %v", ERR_READ_DIR, err)
	}
}
