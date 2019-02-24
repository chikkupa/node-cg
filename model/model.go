package model

import (
	"io/ioutil"
	"os"
	"strings"
)

// GenerateModel to generate model file
func GenerateModel(input Input) error {
	folderName := "model"
	// Create model folder if not exists
	if _, err := os.Stat(folderName); os.IsNotExist(err) {
		err = os.MkdirAll(folderName, 0755)
		if err != nil {
			return err
		}
	}

	headerContent, err := getHeaderContent(input)

	if err != nil {
		return err
	}

	filename := strings.ToLower(input.Name) + ".js"

	err = ioutil.WriteFile(folderName+"/"+filename, []byte(headerContent), 0644)

	if err != nil {
		return err
	}

	return nil
}

func getHeaderContent(input Input) (string, error) {
	content := `const config = require("config");
const mongoose = require("mongoose");
	
mongoose.connect(config.get('mongo_connection'))
	.then(() => console.log("Connected to MongoDB"))
	.catch(err => console.log("Error connecting MongoDB"));
`
	return content, nil
}
