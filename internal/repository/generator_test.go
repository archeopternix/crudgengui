package repository

import (
    "io/ioutil"
    "os"
    "path/filepath"
    "testing"

    "crudgengui/model"
    "github.com/stretchr/testify/assert"
    "gopkg.in/yaml.v3"
)

// Test for NewGenerator
func TestNewGenerator(t *testing.T) {
    gen1 := NewGenerator()
    gen2 := NewGenerator()
    assert.Equal(t, gen1, gen2)
    assert.NotNil(t, gen1.Modules)
}

// Test for AddModules
func TestAddModules(t *testing.T) {
    gen := NewGenerator()
    err := gen.AddModules("testdata/module1.yaml")
    assert.NoError(t, err)
    assert.Contains(t, gen.Modules, "Module1")
}

// Test for LoadFromFile
func TestLoadFromFile(t *testing.T) {
    gen := NewGenerator()
    err := gen.LoadFromFile("testdata/generator.yaml")
    assert.NoError(t, err)
    assert.Contains(t, gen.Modules, "Module1")
}

// Test for SaveToFile
func TestSaveToFile(t *testing.T) {
    gen := NewGenerator()
    gen.Modules["Module1"] = Module{Name: "Module1"}
    err := gen.SaveToFile("testdata/save_test.yaml")
    assert.NoError(t, err)
    defer os.Remove("testdata/save_test.yaml")

    data, err := ioutil.ReadFile("testdata/save_test.yaml")
    assert.NoError(t, err)
    var loadedGen Generator
    err = yaml.Unmarshal(data, &loadedGen)
    assert.NoError(t, err)
    assert.Contains(t, loadedGen.Modules, "Module1")
}

// Test for GenerateAll
func TestGenerateAll(t *testing.T) {
    gen := NewGenerator()
    gen.LoadFromFile("testdata/generator.yaml")
    app := model.NewModel()
    err := gen.GenerateAll(app, "testdata/gen")
    assert.NoError(t, err)
}

// Test for GenerateModule
func TestGenerateModule(t *testing.T) {
    gen := NewGenerator()
    gen.LoadFromFile("testdata/generator.yaml")
    module := gen.Modules["Module1"]
    app := model.NewModel()
    err := module.GenerateModule(app, "testdata/gen")
    assert.NoError(t, err)
}

// Test for AddTask
func TestAddTask(t *testing.T) {
    module := Module{path: "testdata"}
    task := Task{Kind: "copy", Source: []string{"testdata/source.txt"}, Target: "testdata/target"}
    module.AddTask(&task)
    assert.Equal(t, 1, len(module.Tasks))
    assert.Equal(t, filepath.Join("testdata", "testdata/source.txt"), module.Tasks[0].Source[0])
}

// Test for CleanPaths
func TestCleanPaths(t *testing.T) {
    task := Task{Source: []string{"source.txt"}, Target: "target"}
    cleanedTask := task.CleanPaths("testdata")
    assert.Equal(t, filepath.Join("testdata", "source.txt"), cleanedTask.Source[0])
    assert.Equal(t, "target", cleanedTask.Target)
}
