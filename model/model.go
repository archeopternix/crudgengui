package model

import (
	"fmt"
	"io"

	"gopkg.in/yaml.v3"
)

type Entity struct {
	Name string `json:"name" form:"entity-name"`
	Type string `json:"type" form:"entity-type"` // 'Entity' || 'Key-Values'
}

type Relation struct {
	Name        string `json:"name" form:"relation-name"`
	Type        string `json:"type" form:"relation-type"` // 'One-to-Many' | 'Many-to-Many'
	Source      string `json:"source" form:"relation-source"`
	Destination string `json:"destination" form:"relation-destination"`
}

type Model struct {
	Entities  map[string]Entity
	Relations map[string]Relation
}

func NewModel() *Model {
	m := new(Model)
	m.Entities = make(map[string]Entity)
	m.Relations = make(map[string]Relation)
	return m
}

func (m *Model) ReadYAML(reader io.Reader) error {
	data, err := io.ReadAll(reader)
	if err != nil {
		return fmt.Errorf("error in model read, %v", err)
	}
	err = yaml.Unmarshal([]byte(data), m)
	if err != nil {
		return fmt.Errorf("error in model read, %v", err)
	}
	return nil
}

func (m *Model) WriteYAML(writer io.Writer) error {
	data, err := yaml.Marshal(m)
	if err != nil {
		return fmt.Errorf("error in model write, %v", err)
	}
	_, err = writer.Write(data)
	if err != nil {
		return fmt.Errorf("error in model write, %v", err)
	}
	return nil
}
