package main

import "fmt"

func (app *application) help() {
	fmt.Println(app.helpCache)
}