package internal

import (
    "bytes"
    "html/template"
    "testing"

    "github.com/labstack/echo/v4"
    "github.com/stretchr/testify/assert"
)

// Mock echo.Context for testing purposes
type MockContext struct{}

func (mc *MockContext) Deadline() (deadline time.Time, ok bool) {
    return time.Time{}, false
}
func (mc *MockContext) Done() <-chan struct{} {
    return nil
}
func (mc *MockContext) Err() error {
    return nil
}
func (mc *MockContext) Value(key interface{}) interface{} {
    return nil
}

func TestNewTemplateRegistry(t *testing.T) {
    tr := NewTemplateRegistry()
    assert.NotNil(t, tr)
    assert.NotNil(t, tr.templates)
    assert.NotNil(t, tr.funcMap)
}

func TestGetFuncMap(t *testing.T) {
    tr := NewTemplateRegistry()
    funcMap := tr.getFuncMap()
    assert.NotNil(t, funcMap)
    assert.Equal(t, strings.Title("hello"), funcMap["title"].(func(string) string)("hello"))
    assert.Equal(t, strings.ToLower("HELLO"), funcMap["lowercase"].(func(string) string)("HELLO"))
    assert.Equal(t, strings.ToUpper("hello"), funcMap["uppercase"].(func(string) string)("hello"))
}

func TestAddTemplate(t *testing.T) {
    tr := NewTemplateRegistry()
    err := tr.addTemplate("base", "", "testdata/base.html")
    assert.NoError(t, err)
    assert.Contains(t, tr.templates, "base")

    err = tr.addTemplate("child", "base", "testdata/child.html")
    assert.NoError(t, err)
    assert.Contains(t, tr.templates, "child")
}

func TestAddTemplateOrPanic(t *testing.T) {
    tr := NewTemplateRegistry()
    assert.NotPanics(t, func() {
        tr.AddTemplateOrPanic("base", "", "testdata/base.html")
    })
    assert.Panics(t, func() {
        tr.AddTemplateOrPanic("", "", "testdata/base.html")
    })
}

func TestRender(t *testing.T) {
    tr := NewTemplateRegistry()
    tr.AddTemplateOrPanic("base", "", "testdata/base.html")
    tr.AddTemplateOrPanic("child", "base", "testdata/child.html")

    var buf bytes.Buffer
    err := tr.Render(&buf, "child", map[string]interface{}{"Name": "World"}, &MockContext{})
    assert.NoError(t, err)
    assert.Contains(t, buf.String(), "Hello, World!")
}
