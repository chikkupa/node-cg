package model

import (
	"io/ioutil"
	"os"
	"strings"
)

var schemaName string
var modelName string
var lowercaseName string

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

	content, err := generateHeaderContent(input)

	if err != nil {
		return err
	}

	if input.Settings.List {
		content += "\n\n" + getnerateGetFunction(input)
	}
	if input.Settings.Add {
		content += "\n\n" + generateCreateFunction(input)
	}
	if input.Settings.Details {
		content += "\n\n" + generateGetDetailsFunction(input)
	}
	if input.Settings.Update {
		content += "\n\n" + generateUpdateFunction(input)
	}
	if input.Settings.Delete {
		content += "\n\n" + generateDeleteFunction(input)
	}
	content += "\n\n" + generateExports(input)

	filename := strings.ToLower(input.Name) + ".js"

	err = ioutil.WriteFile(folderName+"/"+filename, []byte(content), 0644)

	if err != nil {
		return err
	}

	return nil
}

func generateHeaderContent(input Input) (string, error) {
	content := `const config = require("config");
const mongoose = require("mongoose");
	
mongoose.connect(config.get('mongo_connection'))
	.then(() => console.log("Connected to MongoDB"))
	.catch(err => console.log("Error connecting MongoDB"));
`
	content += "\n" + generateSchemaDefinition(input)
	return content, nil
}

func generateSchemaDefinition(input Input) string {
	lowercaseName = strings.ToLower(input.Name)
	schemaName = strings.ToLower(input.Name) + "Schema"
	schemaContent := "const " + schemaName + " = new mongoose.Schema({"

	for _, field := range input.Fields {
		schemaContent += "\n\t" + field.Name + " : " + field.Type + ","
	}
	schemaContent += "\n});"

	modelName = strings.Title(input.Name)

	schemaContent += "\n\nconst " + modelName + " = mongoose.model('" + modelName + "', " + schemaName + ");"

	return schemaContent
}

func getnerateGetFunction(input Input) string {
	content := "async function get" + modelName + "s(){"
	content += "\n\tconst result = await " + modelName + ".find()"
	content += "\n\t\t.select({"
	for index, field := range input.Fields {
		content += " " + field.Name + " : 1"
		if len(input.Fields) != index+1 {
			content += ","
		}
	}
	content += "});"
	content += "\n\treturn result;"
	content += "\n}"

	return content
}

func generateCreateFunction(input Input) string {
	content := "async function create" + modelName + "("
	for index, field := range input.Fields {
		content += field.Name
		if len(input.Fields) != index+1 {
			content += ", "
		}
	}
	content += "){"
	content += "\n\tconst " + lowercaseName + " = new " + modelName + "({"
	for index, field := range input.Fields {
		content += "\n\t\t" + field.Name + " : " + field.Name
		if len(input.Fields) != index+1 {
			content += ","
		}
	}
	content += "\n\t});"

	content += "\n\tconst result = await " + lowercaseName + ".save();"
	content += "\n\treturn result;"
	content += "\n}"
	return content
}

func generateGetDetailsFunction(input Input) string {
	content := "async function get" + modelName + "(id){"
	content += "\n\tif(!mongoose.Types.ObjectId.isValid(id))"
	content += "\n\t\treturn null;"
	content += "\n\n\tconst result = await " + modelName + ".find({_id : id})"
	content += "\n\t\t.select({"
	for index, field := range input.Fields {
		content += " " + field.Name + " : 1"
		if len(input.Fields) != index+1 {
			content += ","
		}
	}
	content += "});"
	content += "\n\tif(result.length == 0)"
	content += "\n\t\treturn null;"
	content += "\n\n\treturn result[0];"
	content += "\n}"

	return content
}

func generateUpdateFunction(input Input) string {
	content := "async function update" + modelName + "(id"
	for _, field := range input.Fields {
		if field.Name != "id" {
			content += ", " + field.Name
		}
	}
	content += "){"
	content += "\n\tif(!mongoose.Types.ObjectId.isValid(id))"
	content += "\n\t\treturn null;"
	content += "\n\tconst result = await " + modelName + ".findByIdAndUpdate(id, {"
	content += "\n\t\t$set : {"
	for index, field := range input.Fields {
		content += "\n\t\t" + field.Name + " : " + field.Name
		if len(input.Fields) != index+1 {
			content += ","
		}
	}
	content += "\n\t\t}"
	content += "\n\t});"
	content += "\n\n\tif(!result)"
	content += "\n\t\treturn null;"
	content += "\n\n\treturn result;"
	content += "\n}"

	return content
}

func generateDeleteFunction(input Input) string {
	content := "async function delete" + modelName + "(id){"
	content += "\n\tif(!mongoose.Types.ObjectId.isValid(id))"
	content += "\n\t\treturn null;"
	content += "\n\n\tconst " + lowercaseName + " = await " + modelName + ".findByIdAndDelete(id);"
	content += "\n\n\tif(!" + lowercaseName + ")"
	content += "\n\t\treturn null;"
	content += "\n\treturn " + lowercaseName + ";"
	content += "\n}"

	return content
}

func generateExports(input Input) string {
	content := ""

	if input.Settings.List {
		content += "module.exports.get" + modelName + "s = get" + modelName + "s;"
	}
	if input.Settings.Add {
		content += "\nmodule.exports.create" + modelName + " = create" + modelName + ";"
	}
	if input.Settings.Details {
		content += "\nmodule.exports.get" + modelName + " = get" + modelName + ";"
	}
	if input.Settings.Update {
		content += "\nmodule.exports.update" + modelName + " = update" + modelName + ";"
	}
	if input.Settings.Delete {
		content += "\nmodule.exports.delete" + modelName + " = delete" + modelName + ";"
	}

	return content
}
