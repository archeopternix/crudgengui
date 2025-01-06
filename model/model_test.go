package model

import (
    "bytes"
    "testing"
    "time"

    "github.com/stretchr/testify/assert"
)

// Test for NewEUROSettings
func TestNewEUROSettings(t *testing.T) {
    settings := NewEUROSettings()
    assert.Equal(t, "€", settings.CurrencySymbol)
    assert.Equal(t, ",", settings.DecimalSeparator)
    assert.Equal(t, ".", settings.ThousendSeparator)
    assert.Equal(t, "15:04:05", settings.TimeFormat)
    assert.Equal(t, "02.01.2006", settings.DateFormat)
}

// Test for NewUSSettings
func TestNewUSSettings(t *testing.T) {
    settings := NewUSSettings()
    assert.Equal(t, "$", settings.CurrencySymbol)
    assert.Equal(t, ".", settings.DecimalSeparator)
    assert.Equal(t, ",", settings.ThousendSeparator)
    assert.Equal(t, "15:04:05", settings.TimeFormat)
    assert.Equal(t, "01/02/2006", settings.DateFormat)
}

// Test for NewModel
func TestNewModel(t *testing.T) {
    model := NewModel()
    assert.NotNil(t, model)
    assert.Equal(t, NewEUROSettings(), model.Settings)
    assert.NotNil(t, model.Entities)
    assert.NotNil(t, model.Relations)
    assert.NotNil(t, model.Lookups)
}

// Test for CleanName
func TestCleanName(t *testing.T) {
    model := Model{Name: "Hello@World!"}
    cleanedName := model.CleanName()
    assert.Equal(t, "HelloWorld", cleanedName)
}

// Test for TimeStamp
func TestTimeStamp(t *testing.T) {
    model := Model{Settings: NewEUROSettings()}
    timestamp := model.TimeStamp()
    _, err := time.Parse("02.01.2006 15:04:05", timestamp)
    assert.NoError(t, err)
}

// Test for ReadYAML
func TestReadYAML(t *testing.T) {
    yamlData := `
name: TestModel
entities:
  Entity1:
    name: Entity1
relations:
  Relation1:
    name: Relation1
lookups:
  Lookup1:
    name: Lookup1
`
    model := NewModel()
    err := model.ReadYAML(bytes.NewReader([]byte(yamlData)))
    assert.NoError(t, err)
    assert.Equal(t, "TestModel", model.Name)
    assert.Contains(t, model.Entities, "Entity1")
    assert.Contains(t, model.Relations, "Relation1")
    assert.Contains(t, model.Lookups, "Lookup1")
}

// Test for WriteYAML
func TestWriteYAML(t *testing.T) {
    model := NewModel()
    model.Name = "TestModel"
    var buffer bytes.Buffer
    err := model.WriteYAML(&buffer)
    assert.NoError(t, err)

    expectedYAML := `name: TestModel
settings:
  currency_symbol: "€"
  decimal_separator: ","
  thousend_separator: "."
  time_format: "15:04:05"
  date_format: "02.01.2006"
entities: {}
relations: {}
lookups: {}
`
    assert.Equal(t, expectedYAML, buffer.String())
}

// Test for EntityInRealtions
func TestEntityInRealtions(t *testing.T) {
    model := NewModel()
    model.Relations["Relation1"] = Relation{Source: "Entity1", Destination: "Entity2"}
    assert.True(t, model.EntityInRealtions("Entity1"))
    assert.True(t, model.EntityInRealtions("Entity2"))
    assert.False(t, model.EntityInRealtions("Entity3"))
}

// Test for deepCopy
func TestDeepCopy(t *testing.T) {
    model := NewModel()
    model.Name = "TestModel"
    model.Entities["Entity1"] = Entity{Name: "Entity1"}
    model.Relations["Relation1"] = Relation{Name: "Relation1"}
    model.Lookups["Lookup1"] = Lookup{Name: "Lookup1"}

    copy := model.deepCopy()
    assert.Equal(t, model, copy)
}

// Test for ParseDependencies
func TestParseDependencies(t *testing.T) {
    model := NewModel()
    model.Entities["Entity1"] = Entity{Name: "Entity1", Fields: []Field{{Name: "Field1", Type: "Lookup", Lookup: "lookup1"}}}
    model.Lookups["lookup1"] = Lookup{Name: "Lookup1"}

    copy, err := model.ParseDependencies()
    assert.NoError(t, err)
    assert.Contains(t, copy.Entities["Entity1"].Fields, Field{Name: "ID", Type: "Integer", Required: true, Auto: true})
    assert.Equal(t, "lookup1", copy.Entities["Entity1"].Fields[0].Object)
}

// Test for WriteToFile
func TestWriteToFile(t *testing.T) {
    model := NewModel()
    model.Name = "TestModel"
    filename := "test_model.yaml"
    err := WriteToFile(model, filename)
    assert.NoError(t, err)
    defer os.Remove(filename) // Clean up

    fileContent, err := os.ReadFile(filename)
    assert.NoError(t, err)

    expectedYAML := `name: TestModel
settings:
  currency_symbol: "€"
  decimal_separator: ","
  thousend_separator: "."
  time_format: "15:04:05"
  date_format: "02.01.2006"
entities: {}
relations: {}
lookups: {}
`
    assert.Equal(t, expectedYAML, string(fileContent))
}
