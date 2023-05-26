package main

// TODO: Fill out the personal_data.json file in the data directory

// TODO: Create a Project struct with json tags to hold the data for each project

// TODO: Create a Network struct with json tags to hold the data for each network

// TODO: Create a Portfolio struct with json tags to hold the data for the entire portfolio

func main() {
	portfolio := handleData()
	createHTML(portfolio)
}

func handleData() Portfolio {
	// TODO: Read the data from the JSON file

	// TODO: Unmarshal the JSON data into a Portfolio struct

	// TODO: Go into templates/portfolio_tmpl.gohtml and call the data from the Portfolio struct everywhere it says "TODO"

	return portfolio
}

func createHTML(portfolio Portfolio) {
	// TODO: Create the index.html file

	// TODO: Parse the template file

	// TODO: Execute the template and write the output to the index.html

}
