package model

import (
	"crudgengui/pkg"
	"fmt"
	"slices"
)

// Lookup is a string list
type Lookup struct {
	Name string
	List []string // list that contains the text values
}

// NewLookup creates a pointer to a new Lookup
func NewLookup(name string) *Lookup {
	return &Lookup{Name: name}
}

// CleanName removes all non-numeric and non-alphanumeric characters from the input string.
func (l Lookup) CleanName() string {
	return pkg.CleanString(l.Name)
}

// Add adds a text entry to the Lookup's list.
func (l *Lookup) Add(text string) {
	l.List = append(l.List, text)
}

// Delete removes a text entry at the specified index from the Lookup's list.
// It preserves the order of the remaining elements.
// Returns an error if the index is out of range.
func (l *Lookup) Delete(i int) error {
	if i < 0 || i >= len(l.List) {
		return fmt.Errorf("index %d out of range", i)
	}
	l.List = slices.Delete(l.List, i, i+1)
	return nil
}
