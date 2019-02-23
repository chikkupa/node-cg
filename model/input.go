package model

import (
	"encoding/json"
)

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

// BuildInput Building json file input
func BuildInput(data string) (Input, error) {
	var inputObject Input

	var raw map[string]interface{}
	err := json.Unmarshal([]byte(data), &raw)

	if err != nil {
		return inputObject, err
	}
	inputObject.Name, _ = (raw["name"].(string))

	fields, _ := raw["fields"].([]interface{})

	inputObject.Fields = make([]Field, len(fields))

	for i := 0; i < len(fields); i++ {
		field, _ := fields[i].(map[string]interface{})
		fieldName, _ := field["name"].(string)
		fieldType, _ := field["type"].(string)
		inputField := Field{fieldName, fieldType}
		inputObject.Fields[i] = inputField
	}

	settings, _ := raw["settings"].(map[string]interface{})
	settingsList, _ := settings["list"].(bool)
	settingsAdd, _ := settings["add"].(bool)
	settingsDetails, _ := settings["details"].(bool)
	settingsUpdate, _ := settings["update"].(bool)
	settingsDelete, _ := settings["delete"].(bool)

	inputSettings := Settings{settingsList, settingsAdd, settingsDetails, settingsUpdate, settingsDelete}

	inputObject.Settings = inputSettings

	return inputObject, nil
}
