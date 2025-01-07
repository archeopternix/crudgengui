package repository

import (
	model "crudgengui/internal/model"
	"fmt"
	"io"
	"os"

	"gopkg.in/yaml.v3"
)

// ModelToFile struct represents a model for handling file operations (as YAML).
type ModelToFile struct {
	yamlfile string
}

// NewModelToFile creates a new instance of ModelToFile with the provided file name.
func NewModelToFile(fname string) *ModelToFile {
	return &ModelToFile{yamlfile: fname}
}

func (ym ModelToFile) WriteModel(m *model.Model) error {
	return WriteModelToFile(m, ym.yamlfile)
}

// WriteModel writes the given model to the YAML file.
func WriteModelToFile(m *model.Model, filename string) error {
	// Create a new file with the name stored in the ModelToFile struct.
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close() // Ensure file is closed after writing

	// Write the model data to the file in YAML format.
	data, err := yaml.Marshal(m)
	if err != nil {
		return fmt.Errorf("error in model write, %v", err)
	}
	if _, err = file.Write(data); err != nil {
		return fmt.Errorf("error in model write, %v", err)
	}
	return nil
}

// ReadModel reads the model from the YAML file and populates the given model instance.
func (ym *ModelToFile) ReadModel(m *model.Model) error {
	return ReadModelFromFile(m, ym.yamlfile)
}

// ReadModel reads the model from the YAML file and populates the given model instance.
func ReadModelFromFile(m *model.Model, filename string) error {
	// Open the file with the name stored in the ModelToFile struct.
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close() // Ensure file is closed after reading

	// Read the YAML data from the file and populate the model instance.
	data, err := io.ReadAll(file)
	if err != nil {
		return fmt.Errorf("error in model read, %v", err)
	}
	if err = yaml.Unmarshal([]byte(data), m); err != nil {
		return fmt.Errorf("error in model read, %v", err)
	}

	if m.Relations == nil {
		m.Relations = make(map[string]model.Relation)
	}
	if m.Entities == nil {
		m.Entities = make(map[string]model.Entity)
	}

	return nil
}
