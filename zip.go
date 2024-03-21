package main

import (
	// "archive/zip"
	"fmt"
	"os"
)
// TODO: I CAN ALSO ADD A TEST TO CHECK IF THE CONTENT OF THE FILE IS NOT CORRECTED

func (app *application) zipFile(input, outputFile string) {
	// TODO: GET ZIP FILE WORKING 
	archive, err := os.Create("output.zip")
	if err != nil {
		app.errlog.Fatalln(err)
	}


	// zip.Writer()
	fmt.Println("add the zip file code here")

	fmt.Println(input, outputFile)
}

func (app *application) zipDirectory(input, outputFile string) {
	fmt.Println("add the zip dir code here")
	// TODO: GET THE ZIP DIRECTORY FUNCTION WORKING
	fmt.Println(input, outputFile)

}

func (app *application) getDir(input string){
	// TODO: GET THE GET DIR FUNCTION WORKING

}