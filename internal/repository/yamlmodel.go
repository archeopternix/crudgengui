package repository

import (
	"crudgengui/internal/model"
	"os"
)

// YAMLModel struct represents a model for handling YAML file operations.
type YAMLModel struct {
	yamlfile string
}

// NewYAMLModel creates a new instance of YAMLModel with the provided file name.
func NewYAMLModel(fname string) *YAMLModel {
	return &YAMLModel{yamlfile: fname}
}

// WriteModel writes the given model to the YAML file.
func (ym YAMLModel) WriteModel(m *model.Model) error {
	// Create a new file with the name stored in the YAMLModel struct.
	file, err := os.Create(ym.yamlfile)
	if err != nil {
		return err
	}
	defer file.Close() // Ensure file is closed after writing

	// Write the model data to the file in YAML format.
	return m.WriteYAML(file)
}

// ReadModel reads the model from the YAML file and populates the given model instance.
func (ym *YAMLModel) ReadModel(m *model.Model) error {
	// Open the file with the name stored in the YAMLModel struct.
	file, err := os.Open(ym.yamlfile)
	if err != nil {
		return err
	}
	defer file.Close() // Ensure file is closed after reading

	// Read the YAML data from the file and populate the model instance.
	return m.ReadYAML(file)
}
