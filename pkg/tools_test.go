package pkg

import (
	"os"
	"testing"
)

// Test for IsFirstLetterUppercase
func TestIsFirstLetterUppercase(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"Hello", true},
		{"hello", false},
		{"", false},
		{"123", false},
	}

	for _, test := range tests {
		result := IsFirstLetterUppercase(test.input)
		if result != test.expected {
			t.Errorf("IsFirstLetterUppercase(%s) = %v; want %v", test.input, result, test.expected)
		}
	}
}

// Test for CheckMkdir
func TestCheckMkdir(t *testing.T) {
	dir := "test_dir"
	defer os.RemoveAll(dir) // Clean up

	err := CheckMkdir(dir)
	if err != nil {
		t.Errorf("CheckMkdir(%s) failed: %v", dir, err)
	}

	// Try creating the same directory again to check for DirectoryExistError
	err = CheckMkdir(dir)
	if _, ok := err.(*DirectoryExistError); !ok {
		t.Errorf("CheckMkdir(%s) did not return DirectoryExistError; got %v", dir, err)
	}
}

// Test for fileExist
func TestFileExist(t *testing.T) {
	file := "test_file.txt"
	defer os.Remove(file) // Clean up

	// File should not exist initially
	err := fileExist(file)
	if err != nil {
		t.Errorf("fileExist(%s) failed: %v", file, err)
	}

	// Create the file
	_, err = os.Create(file)
	if err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}

	// File should exist now
	err = fileExist(file)
	if _, ok := err.(*FileExistError); !ok {
		t.Errorf("fileExist(%s) did not return FileExistError; got %v", file, err)
	}
}

// Test for CopyFile
func TestCopyFile(t *testing.T) {
	source := "source_file.txt"
	dest := "dest_file.txt"
	defer os.Remove(source) // Clean up
	defer os.Remove(dest)   // Clean up

	// Create source file
	content := []byte("Hello, World!")
	err := os.WriteFile(source, content, 0644)
	if err != nil {
		t.Fatalf("Failed to create source file: %v", err)
	}

	// Copy to destination
	err = CopyFile(source, dest)
	if err != nil {
		t.Errorf("CopyFile(%s, %s) failed: %v", source, dest, err)
	}

	// Check if content matches
	destContent, err := os.ReadFile(dest)
	if err != nil {
		t.Fatalf("Failed to read destination file: %v", err)
	}
	if string(destContent) != string(content) {
		t.Errorf("Content mismatch: got %s; want %s", destContent, content)
	}
}

// Test for StringYAML
func TestStringYAML(t *testing.T) {
	obj := map[string]interface{}{
		"name": "test",
		"age":  25,
	}

	yamlStr, err := StringYAML(obj)
	if err != nil {
		t.Errorf("StringYAML(%v) failed: %v", obj, err)
	}

	expected := "age: 25\nname: test\n"
	if yamlStr != expected {
		t.Errorf("StringYAML(%v) = %s; want %s", obj, yamlStr, expected)
	}
}

// Test for CleanString
func TestCleanString(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"hello", "hello"},
		{"héllo", "hello"},
		{"äöü", "aeoeue"},
		{"123", "123"},
		{"!@#", ""},
	}

	for _, test := range tests {
		result := CleanString(test.input)
		if result != test.expected {
			t.Errorf("CleanString(%s) = %s; want %s", test.input, result, test.expected)
		}
	}
}

// Test for CleanID
func TestCleanID(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"hello", "Hello"},
		{"Hello", "Hello"},
		{"héllo", "Hello"},
		{"äöü", "Aeoeue"},
		{"123", "123"},
	}

	for _, test := range tests {
		result := CleanID(test.input)
		if result != test.expected {
			t.Errorf("CleanID(%s) = %s; want %s", test.input, result, test.expected)
		}
	}
}
