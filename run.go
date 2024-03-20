package main

import "os"

const UNEXPECTED_BEHAVIOUR = "Unexpected behaviour, typical skill issues\nrun 'zippy help' to do better" 

func (app *application) run(){
	len := len(os.Args)
	
	switch len{
		case 0:
			app.errlog.Fatalln(UNEXPECTED_BEHAVIOUR)
		case 1:
			app.zipFile(app.input, app.outputFile)
			return
		case 2:
			option := os.Args[1]
			switch option {
				case "help":	
					app.help()
					return		
				case "-d":
					app.zipDirectory(app.input, app.outputFile)
				case "-f":
					app.errlog.Fatalln("Input file name not given")
				default:
					app.errlog.Fatalln(UNEXPECTED_BEHAVIOUR)
			}
		case 3:
				
	}
}