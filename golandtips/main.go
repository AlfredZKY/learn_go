package main

import (

	"html/template"
	"io/ioutil"
	"os"
)

type Page struct {
	Title string
	Users []user
}

type user struct {
	Username string
}

func main() {
	tpl := funcName()

	var data = Page{
		Title: "Demo file",
		Users: []user{
			{Username: "Florin"},
		},
	}
	err := tpl.Execute(os.Stdout, data)
	if err != nil {
		panic(err)
	}

}

func funcName() *template.Template {
	file, err := ioutil.ReadFile("tpl.gohtml")
	if err != nil {
		panic(err)
	}

	tpl, err := template.New("mytemplete").Parse(string(file))
	if err != nil {
		panic(err)
	}
	return tpl
}
