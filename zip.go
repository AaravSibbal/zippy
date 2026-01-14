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


func (app *application) writeFileToArchive(zipWriter *zip.Writer, input string) {

	app.infoLog.Printf("%s %s", OPEN_FILE, input)
	file, err := os.Open(input)
	if err != nil {
		app.errlog.Fatalf("%s %v", ERR_OPEN_FILE, err)
	}

	defer file.Close();

	app.infoLog.Printf("%s %s", WRITE_FILE, input)
	header := &zip.FileHeader{
		Name: input,
		Method: zip.Deflate,
	}
	writer, err := zipWriter.CreateHeader(header)
	if err != nil {
		app.errlog.Fatalf("%s %v", ERR_CREATE_BUFFER, err)
	}

	if _, err = io.Copy(writer, file); err != nil {
		app.errlog.Fatalf("%s %v", ERR_COPY_BUFFER, err)
	}
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
	absOutput, err := filepath.Abs(outputFile)
	if err != nil {
		app.errlog.Fatalf("Error calculating the path: %s,\n%s", outputFile, err.Error())
	}
	archive, err := os.Create(outputFile)
	if err != nil {
		app.errlog.Fatalf("%s %v", ERR_CREATE_ZIP, err)
	}
	defer archive.Close();

	zipWriter := zip.NewWriter(archive)
	defer zipWriter.Close()

	err = filepath.WalkDir(input, func(path string, d fs.DirEntry, err error) error {
		if d.IsDir(){
			if d.Name() == ".git" || d.Name() == "node_modules"{
				return filepath.SkipDir;
			}
			return nil;
		}
		if err != nil {
			return err;
		}
		fmt.Println(path, d.Name(), "directory?", d.IsDir());
		absPath, _ := filepath.Abs(path);
		if absPath == absOutput {
			return nil;
		}

		if !d.IsDir() {
			app.writeFileToArchive(zipWriter, path)
		}
		return nil
	})

	if err != nil {
		app.errlog.Fatalf("%s %v", ERR_READ_DIR, err)
	}
}

func (app *application) zipMultipleFiles(files []string, outputFile string){
	archive, err := os.Create(outputFile);
	if err != nil {
		app.errlog.Fatalf("Something went wrong with creating the zip file\n%s", err.Error())
		return
	}
	defer archive.Close();

	zipWriter := zip.NewWriter(archive);
	
	defer zipWriter.Close();

	for _, file := range files {
		app.writeFileToArchive(zipWriter, file);
	}

}

func (app *application) verifyArchive(path string){
	app.infoLog.Printf("Verifying the archive integrity: %s\n", path);

	reader, err := zip.OpenReader(path);
	if err != nil {
		app.errlog.Fatalf("CORRUPTION DETECTED: the zip file is invalid\n%s", err.Error())
	}
	reader.Close();
	app.infoLog.Println("Archive verified successfully!");
}