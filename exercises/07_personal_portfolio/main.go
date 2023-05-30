package main

import (
	"encoding/json"
	"html/template"
	"io/ioutil"
	"log"
	"os"
)

// TODO: Fill out the personal_data.json file in the data directory

// TODO: Create a Project struct to hold the data for each project
type Project struct {
	Title       string `json:"Title"`
	Description string `json:"Description"`
	Link        string `json:"Link"`
}

// TODO: Create a Network struct to hold the data for each network
type Network struct {
	Name string `json:"Name"`
	Link string `json:"Link"`
}

// TODO: Create a Portfolio struct to hold the data for the entire portfolio
type Portfolio struct {
	Name     string    `json:"Name"`
	Title    string    `json:"Title"`
	About    string    `json:"About"`
	Projects []Project `json:"Projects"`
	Networks []Network `json:"Networks"`
}

func main() {
	portfolio := handleData()
	createHTML(portfolio)
}

func handleData() Portfolio {
	// TODO: Read the data from the JSON file
	file, err := ioutil.ReadFile("data/personal_data.json")
	if err != nil {
		log.Fatal("Failed to read personal_data.json:", err)
	}

	// TODO: Unmarshal the JSON data into a Portfolio struct
	var portfolio Portfolio
	err = json.Unmarshal(file, &portfolio)
	if err != nil {
		log.Fatal("Failed to unmarshal JSON data:", err)
	}

	// TODO: Go into templates/portfolio.gohtml and call the data from the Portfolio struct everywhere it says "TODO"

	return portfolio
}

func createHTML(portfolio Portfolio) {
	// TODO: Create the index.html file
	file, err := os.Create("index.html")
	if err != nil {
		log.Fatal("Failed to create index.html file:", err)
	}
	defer file.Close()

	// TODO: Parse the template file
	tmpl, err := template.ParseFiles("templates/portfolio.gohtml")
	if err != nil {
		log.Fatal("Failed to parse template file:", err)
	}

	// TODO: Execute the template and write the output to the index.html
	err = tmpl.Execute(file, portfolio)
	if err != nil {
		log.Fatal("Failed to execute template:", err)
	}
}
