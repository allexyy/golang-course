package main

import (
	"fmt"
	"html/template"
	"log"
	"os"
)

type Todo struct {
	Title string
	Done  bool
}

type TodoPageData struct {
	PageTitle string
	Todos     []Todo
}

func main() {
	tmpl, err := template.ParseFiles("template.html")
	if err != nil {
		log.Fatal(err)
	}

	file, err := os.Create("index.html")
	if err != nil {
		fmt.Println("Unable to create file:", err)
		os.Exit(1)
	}
	defer file.Close()

	data := TodoPageData{
		PageTitle: "TODO list",
		Todos: []Todo{
			{Title: "Task 1", Done: false},
			{Title: "Task 2", Done: true},
			{Title: "Task 3", Done: true},
		},
	}

	tmpl.Execute(file, data)
}
