// Package pkg provides various utility functions and custom error types
// for handling file operations, string manipulations, and YAML conversions.
// 
// The utilities include functions for checking if the first letter of a string
// is uppercase, creating directories, copying files, and cleaning strings by
// removing or converting non-alphanumeric characters. It also includes custom
// error types for file and directory existence errors, and functions for
// converting data structures to YAML strings.
package pkg

import (
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
	"unicode"

	"gopkg.in/yaml.v3"
)

// IsFirstLetterUppercase checks if the first letter of the input string is uppercase.
func IsFirstLetterUppercase(str string) bool {
	if len(str) == 0 {
		return false
	}
	return unicode.IsUpper(rune(str[0]))
}

// DirectoryExistError will bne thrown by CheckMkdir when a directory already exists
type DirectoryExistError struct {
	Dir string
	Err error
}

// Error implements the error interface for DirectoryExistError
func (r *DirectoryExistError) Error() string {
	return fmt.Sprintf("directory '%s' %v", r.Dir, r.Err)
}

// FileExistError will be thrown when a file already exists and it is overwritten
// by copyFile
type FileExistError struct {
	File string
	Err  error
}

// Error implements the error interface for FileExistError
func (r *FileExistError) Error() string {
	return fmt.Sprintf("file '%s' %v", r.File, r.Err)
}

// CheckMkdir checks and creates a directory with given path when not yet exists
// when directory exists a DirectoryExistError will be thrown, in cas ea new
// directory will be created it returns nil
func CheckMkdir(path string) error {
	// throw error when directory already exists
	if _, err := os.Lstat(path); err == nil {
		return &DirectoryExistError{
			Dir: path,
			Err: errors.New("already exists"),
		}
	}

	if err := os.MkdirAll(path, os.ModePerm); err != nil {
		return fmt.Errorf("directory %s: err %v", path, err)
	}
	return nil
}

// fileExist returns whether the given file exists. Returns nil when file does
// not exist, FileExistError when files exist or the error when something went wrong
func fileExist(fname string) error {
	_, err := os.Stat(fname)
	if err == nil {
		return &FileExistError{
			File: fname,
			Err:  errors.New("already exists"),
		}
	}
	if os.IsNotExist(err) {
		return nil
	}
	return err
}

// CopyFile copies the content from sourcefile to destfile. If the file already
// exists, the file will be overwritten and an FileExistError error will be thrown
func CopyFile(sourcefile, destfile string) error {
	var source, dest *os.File
	var err error

	// open source file
	source, err = os.Open(sourcefile)
	if err != nil {
		return err
	}
	defer source.Close()

	// overwrite or new file
	exist := fileExist(destfile)

	// create target file
	dest, err = os.Create(destfile)
	if err != nil {
		return err
	}
	defer dest.Close()
	_, err = io.Copy(dest, source)
	if err != nil {
		return err
	}

	// when file exists an FileExistError will be thrown
	if _, ok := exist.(*FileExistError); ok {
		return exist
	}

	return nil
}

// StringYAML returns a YAML string of the data structure 'obj' or an error when
// something went wrong
func StringYAML(obj interface{}) (string, error) {
	data, err := yaml.Marshal(&obj)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

// CleanString converts or removes all non-numeric and non-alphanumeric characters from the input string.
func CleanString(input string) string {
	// Map for converting non-ASCII characters to ASCII equivalents
	conversions := map[rune]string{
		'ä': "ae",
		'ö': "oe",
		'ü': "ue",
		'Ä': "Ae",
		'Ö': "Oe",
		'Ü': "Ue",
		'ß': "ss",
		'é': "e",
		'è': "e",
		'ê': "e",
		'ç': "c",
		'ñ': "n",
		// Add more conversions as needed
	}

	var result strings.Builder
	for _, r := range input {
		if val, ok := conversions[r]; ok {
			result.WriteString(val)
		} else {
			if unicode.IsLetter(r) || unicode.IsDigit(r) {
				result.WriteRune(r)
			}
		}
	}
	return result.String()
}

// CleanID converts any name into a clean ID with just unicode letters and the
// first letter is uppercase
func CleanID(name string) string {
	if len(name) == 0 {
		return ""
	}
	txt := CleanString(name)
	if !IsFirstLetterUppercase(txt) {
		return strings.Title(txt)
	}
	return txt
}
