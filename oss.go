// Package ossgo provides [BRIEF_DESCRIPTION_OF_YOUR_PACKAGE]
//
// Example usage:
//
//	package main
//
//	import (
//		"fmt"
//		"github.com/bold-minds/oss"
//	)
//
//	func main() {
//		// Your example usage here
//		result := oss.ExampleFunction("hello")
//		fmt.Println(result)
//	}
package oss

import (
	"fmt"
	"strings"
	"time"

	"github.com/bold-minds/id"
)

// Version represents the current version of the package
const Version = "v0.1.0"

// ExampleStruct demonstrates a typical struct in your package
// It includes an ID field using the bold-minds/id package for unique identification
type ExampleStruct struct {
	ID    string // ULID from bold-minds/id package
	Name  string
	Value int
}

// NewExampleStruct creates a new instance of ExampleStruct with a unique ID
// Demonstrates integration with bold-minds/id package
func NewExampleStruct(name string, value int) *ExampleStruct {
	gen := id.NewGenerator()
	return &ExampleStruct{
		ID:    gen.Generate(),
		Name:  name,
		Value: value,
	}
}

// ExampleFunction demonstrates a typical function in your package
// Replace this with your actual functionality
func ExampleFunction(input string) string {
	if input == "" {
		return "empty input"
	}
	return fmt.Sprintf("processed: %s", strings.ToUpper(input))
}

// Process demonstrates a method on ExampleStruct
// Shows how to work with IDs from bold-minds/id package
func (e *ExampleStruct) Process() string {
	return fmt.Sprintf("%s (ID: %s) has value %d", e.Name, e.ID[:8]+"...", e.Value)
}

// Validate demonstrates input validation including ID validation
// Shows integration with bold-minds/id validation features
func (e *ExampleStruct) Validate() error {
	if e.ID == "" {
		return fmt.Errorf("ID cannot be empty")
	}
	// Validate ULID format using bold-minds/id package
	gen := id.NewGenerator()
	if !gen.IsIdValid(e.ID) {
		return fmt.Errorf("ID is not a valid ULID: %s", e.ID)
	}
	if e.Name == "" {
		return fmt.Errorf("name cannot be empty")
	}
	if e.Value < 0 {
		return fmt.Errorf("value cannot be negative")
	}
	return nil
}

// GenerateID demonstrates standalone ID generation using bold-minds/id
// This shows how other bold-minds packages can be used together
func GenerateID() string {
	gen := id.NewGenerator()
	return gen.Generate()
}

// GetIDTimestamp demonstrates extracting timestamp from an ID
// Shows time-based operations available in bold-minds/id
func GetIDTimestamp(idStr string) (time.Time, error) {
	gen := id.NewGenerator()
	return gen.ExtractTimestamp(idStr)
}

// GetIDAge demonstrates extracting age from an ID
// Shows time-based operations available in bold-minds/id
func GetIDAge(idStr string) (time.Duration, error) {
	gen := id.NewGenerator()
	return gen.Age(idStr)
}

// CompareIDs demonstrates ID comparison functionality
// Shows chronological ordering capabilities of bold-minds/id
func CompareIDs(id1, id2 string) (int, error) {
	gen := id.NewGenerator()
	return gen.Compare(id1, id2)
}
