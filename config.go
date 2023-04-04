package main

import (
  "fmt"
  "io"
  "time"
	"gopkg.in/yaml.v3"
)

type Version struct {
  Major int
  Minor int
  Date time.Time
}

func (v Version) VersionString() string {
  return fmt.Sprintf("%2d.%2d",v.Major,v.Minor)
}

func (v *Version)IncrementMinor() {
  v.Minor+=1
  v.Date = time.Now()
}

type Config struct {
  Version
  FieldDefs FieldTypeDefinitions
}

func NewConfig() (*Config) {
  c:= new(Config)
  c.Version.Major = 0
  c.Version.Minor = 1
  c.Date = time.Now()
  c.FieldDefs=make(FieldTypeDefinitions)
  return c
}

func (c *Config) ReadYAML(reader io.Reader) error {
	data, err := io.ReadAll(reader)
	if err != nil {
		return fmt.Errorf("error in config read, %v", err)
	}
	err = yaml.Unmarshal([]byte(data), c)
	if err != nil {
		return fmt.Errorf("error in config read, %v", err)
	}
	return nil
}

func (c *Config) WriteYAML(writer io.Writer) error {
	data, err := yaml.Marshal(c)
	if err != nil {
		return fmt.Errorf("error in config write, %v", err)
	}
	_, err = writer.Write(data)
	if err != nil {
		return fmt.Errorf("error in config write, %v", err)
	}
	return nil
}


type FieldTypeDefinition struct {
  Name string
  Optional bool
  HasLength bool
  DefaultLength int
  HasSize bool
  DefaultSize int
  // numerical values
  HasMin bool
  DefaultMin int
  HasMax bool
  DefaultMax int
  HasStep bool
  DefaultStep float32
  HasDigits bool
  DefaultDigits int
  Format string // uses fmt.Printf format and special format for dates e.g. "2006-01-02 15:04:05"
  BaseType string // int, bool, string, time.date, float
  Icon string // e.g. fa-euro-sign shown in entry fields

  Version int
  ChangeDate time.Time
}

type FieldTypeDefinitions map[string]FieldTypeDefinition