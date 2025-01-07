package model

import (
    "testing"
    "time"

    "github.com/stretchr/testify/assert"
)

// Mock Field struct for testing
type Field struct {
    Name string
}

// Test for NewEntity
func TestNewEntity(t *testing.T) {
    e := NewEntity()
    assert.NotNil(t, e)
    assert.Empty(t, e.Name)
    assert.Empty(t, e.Type)
    assert.Empty(t, e.Fields)
}

// Test for CleanName
func TestCleanName(t *testing.T) {
    e := Entity{Name: "Hello@World!"}
    cleanedName := e.CleanName()
    assert.Equal(t, "HelloWorld", cleanedName)
}

// Test for TimeStamp
func TestTimeStamp(t *testing.T) {
    e := Entity{}
    timestamp := e.TimeStamp()
    _, err := time.Parse("01.02.2006 15:04:05", timestamp)
    assert.NoError(t, err)
}

// Test for Add
func TestAdd(t *testing.T) {
    e := NewEntity()
    field := Field{Name: "Field1"}
    e.Add(field)
    assert.Equal(t, 1, len(e.Fields))
    assert.Equal(t, "Field1", e.Fields[0].Name)
}

// Test for GetFieldIndexByName
func TestGetFieldIndexByName(t *testing.T) {
    e := Entity{Fields: []Field{{Name: "Field1"}, {Name: "Field2"}}}
    index := e.GetFieldIndexByName("Field2")
    assert.Equal(t, 1, index)
    index = e.GetFieldIndexByName("Field3")
    assert.Equal(t, -1, index)
}

// Test for GetFieldByName
func TestGetFieldByName(t *testing.T) {
    e := Entity{Fields: []Field{{Name: "Field1"}, {Name: "Field2"}}}
    field := e.GetFieldByName("Field2")
    assert.NotNil(t, field)
    assert.Equal(t, "Field2", field.Name)
    field = e.GetFieldByName("Field3")
    assert.Nil(t, field)
}

// Test for DeleteFieldByName
func TestDeleteFieldByName(t *testing.T) {
    e := Entity{Fields: []Field{{Name: "Field1"}, {Name: "Field2"}}}
    e.DeleteFieldByName("Field1")
    assert.Equal(t, 1, len(e.Fields))
    assert.Equal(t, "Field2", e.Fields[0].Name)
}

// Test for IsValid
func TestIsValid(t *testing.T) {
    e := Entity{Name: "Entity1", Type: "Type1"}
    valid, errormap := e.IsValid()
    assert.True(t, valid)
    assert.Empty(t, errormap)

    e = Entity{Name: "Entity1", Type: ""}
    valid, errormap = e.IsValid()
    assert.False(t, valid)
    assert.NotEmpty(t, errormap)
    assert.Contains(t, errormap, "Type")

    e = Entity{Name: "Entity1!", Type: "Type1"}
    valid, errormap = e.IsValid()
    assert.False(t, valid)
    assert.NotEmpty(t, errormap)
    assert.Contains(t, errormap, "Name")
}
