package main

import (
	"text/template"
	"os"
)

type DataAction struct {
	Classname string
}

func main() {
	fileName := "ActionTemplate.php"

	file, _ := os.Create(fileName)

	fileTemplate := "./Action.php"
	tmpl, err := template.ParseFiles(fileTemplate)

	if err != nil {
		panic(err.Error())
	}

	data := DataAction{
		Classname: "Action",
	}

	err = tmpl.Execute(file, data)

	if err != nil {
		panic(err.Error())
	}
}
