package main

import (
	"fmt"
	"os"
)

func (app *application) help() {
	fmt.Println(app.helpCache)
}

func getHelpString() (string, error) {
	file, err := os.ReadFile("./out/help.txt")
	if err != nil {
		return "", err
	}
	data := string(file)
	return data, nil
}
