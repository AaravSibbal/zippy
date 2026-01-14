package main

const UNEXPECTED_BEHAVIOUR_ERR = "Unexpected behaviour, typical skill issues\nrun 'zippy help' to do better"
const INPUT_FILE_ERR = "Input file name not given"

func (app *application) run() {
	if len(app.inputFiles) == 0 {
		app.help();
		return;
	}

	app.outputFile = createFileName(app.outputFile);
	if(app.isDir){
		app.zipDirectory(app.inputFiles[0], app.outputFile)
	} else {
		app.zipMultipleFiles(app.inputFiles, app.outputFile)
	}

	app.verifyArchive(app.outputFile);
}


func createFileName(outputFile string) string{
	outputFile += ".zip"
	return outputFile
}