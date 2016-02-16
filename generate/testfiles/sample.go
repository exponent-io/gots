package testfiles

import (
	"encoding/json"
	"time"
)

type App struct {
	Id     string `json:"id"`
	Title  string `json:"title"`
	Url    string `json:"url,omitempty"`
	Hours  int
	Config json.RawMessage `json:"config"`
	Extend interface{}     `json:"exten"`
	Blob   []byte          `json:"blob"`
}

// Create a table for each database entity
type Table struct {
	Id      int       `json:"-" something:"else"`
	Name    string    `json:"name"`
	Columns []Column  `random:"tag" json:"columns,omitempty"`
	Created time.Time `json:"created"`
}

// Represents each column in the table
type Column struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

type ItemValue string

const (
	ItemA ItemValue = "A"
	ItemB           = "B"
	ItemC           = "C"
)

type OtherValue string

const OtherItem OtherValue = "Other"

const StringVal = "string"
