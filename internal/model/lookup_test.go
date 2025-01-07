package model

import (
    "testing"

    "github.com/stretchr/testify/assert"
)

// Test for NewLookup
func TestNewLookup(t *testing.T) {
    l := NewLookup("TestLookup")
    assert.NotNil(t, l)
    assert.Equal(t, "TestLookup", l.Name)
    assert.Empty(t, l.List)
}

// Test for CleanName
func TestCleanName(t *testing.T) {
    l := Lookup{Name: "Hello@World!"}
    cleanedName := l.CleanName()
    assert.Equal(t, "HelloWorld", cleanedName)
}

// Test for Add
func TestAdd(t *testing.T) {
    l := NewLookup("TestLookup")
    l.Add("Entry1")
    assert.Equal(t, 1, len(l.List))
    assert.Equal(t, "Entry1", l.List[0])
    l.Add("Entry2")
    assert.Equal(t, 2, len(l.List))
    assert.Equal(t, "Entry2", l.List[1])
}

// Test for Delete
func TestDelete(t *testing.T) {
    l := NewLookup("TestLookup")
    l.Add("Entry1")
    l.Add("Entry2")

    err := l.Delete(1)
    assert.NoError(t, err)
    assert.Equal(t, 1, len(l.List))
    assert.Equal(t, "Entry1", l.List[0])

    err = l.Delete(0)
    assert.NoError(t, err)
    assert.Empty(t, l.List)

    // Test out of range
    err = l.Delete(0)
    assert.Error(t, err)
    assert.Equal(t, "index 0 out of range", err.Error())
}
