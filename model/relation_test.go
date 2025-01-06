package model

import (
    "testing"

    "github.com/stretchr/testify/assert"
)

func TestContainsEntity(t *testing.T) {
    relation := Relation{
        Name:        "relation1",
        Type:        "One-to-Many",
        Source:      "Entity1",
        Destination: "Entity2",
    }

    tests := []struct {
        entityName string
        expected   bool
    }{
        {"Entity1", true},
        {"Entity2", true},
        {"Entity3", false},
    }

    for _, test := range tests {
        result := relation.ContainsEntity(test.entityName)
        assert.Equal(t, test.expected, result, "ContainsEntity(%s) = %v; want %v", test.entityName, result, test.expected)
    }
}
