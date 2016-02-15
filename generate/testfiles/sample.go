package testfiles

type App struct {
	Id    string `json:"id"`
	Title string `json:"title"`
	Url   string `json:"url,omitempty"`
	Hours int
}

// Create a table for each database entity
type Table struct {
	Id      int      `json:"-" something:"else"`
	Name    string   `json:"name"`
	Columns []Column `random:"tag" json:"columns,omitempty"`
}

// Represents each column in the table
type Column struct {
	Name string `json:"name"`
	Type string `json:"type"`
}