package model

// Input The structure to parse input JSON file
type Input struct {
	Name     string   `json:"name"`
	Fields   []Field  `json:"fields"`
	Settings Settings `json:"settings"`
}

// Field Representation of required fields
type Field struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

// Settings List of required operations
type Settings struct {
	List    bool `json:"list"`
	Add     bool `json:"add"`
	Details bool `json:"details"`
	Update  bool `json:"update"`
	Delete  bool `json:"delete"`
}
