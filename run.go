package main

import "os"

const UNEXPECTED_BEHAVIOUR_ERR = "Unexpected behaviour, typical skill issues\nrun 'zippy help' to do better" 
const INPUT_FILE_ERR = "Input file name not given"
func (app *application) run(){
	len := len(os.Args)
	
	switch len{
		case 0:
			app.errlog.Fatalln(UNEXPECTED_BEHAVIOUR_ERR)
		case 1:
			app.zipDirectory(app.input, app.outputFile)
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
					app.errlog.Fatalln(INPUT_FILE_ERR)
				default:
					app.errlog.Fatalln(UNEXPECTED_BEHAVIOUR_ERR)
			}
		case 3:
			// That is zippy [option] input
			option := os.Args[1]
			input := os.Args[2]

			switch option {
				case "-d":
					switch input{
						case "":
							app.zipDirectory(app.input, app.outputFile)
						case ".":
							app.zipDirectory(app.input, app.outputFile)
						
						default:
							app.zipDirectory(input, app.outputFile)
					}
				case "-f":
					switch input{

						case "":
							app.errlog.Fatalln(INPUT_FILE_ERR)
						default:
							app.zipFile(input, app.outputFile)
					}
				default:
					app.errlog.Fatalln(UNEXPECTED_BEHAVIOUR_ERR)
			}
		case 4:

			option := os.Args[1]
			input := os.Args[2]
			output := os.Args[3]
			switch option{
				case "-d":
					if(input == "." && output == "."){
						app.zipDirectory(input, app.outputFile)
					} else if(input == "." && output != "."){
						app.zipDirectory(input, output)
					} else if (input != "." && output != "."){
						app.zipDirectory(input, output)
					} else {
						app.errlog.Fatalln(UNEXPECTED_BEHAVIOUR_ERR)
					}

					
				case "-f":
					if(input == "." && output == "."){
						app.errlog.Fatalln(INPUT_FILE_ERR)
					} else if (input != "." && output != "."){
						app.zipFile(input, app.outputFile)
					}else {
						app.errlog.Fatalln(UNEXPECTED_BEHAVIOUR_ERR)
					}
				default:
					app.errlog.Fatalln(UNEXPECTED_BEHAVIOUR_ERR)
			}
		default:
			app.errlog.Fatalln(UNEXPECTED_BEHAVIOUR_ERR)			
	}
}

// func twoArgs(input, output string){

// }