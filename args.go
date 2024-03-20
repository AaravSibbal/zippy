package main

import "os"

func (app *application) getArgs(){
	len := len(os.Args)
	
	switch len{
		case 0:
			app.errlog.Fatalln("No agrument given, \nrun 'zippy help' for not having skill issues")
		case 1:
			app.zipFile(app.input, app.outputFile)
		case 2:
			option := os.Args[1]
			switch option {
			case "help":	
				
			}
	}
}