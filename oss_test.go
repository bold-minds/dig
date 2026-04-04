package oss

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestExampleFunction demonstrates basic unit testing with testify
func TestExampleFunction(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "normal input",
			input:    "hello",
			expected: "processed: HELLO",
		},
		{
			name:     "empty input",
			input:    "",
			expected: "empty input",
		},
		{
			name:     "mixed case",
			input:    "Hello World",
			expected: "processed: HELLO WORLD",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ExampleFunction(tt.input)
			assert.Equal(t, tt.expected, result, "ExampleFunction should return expected result")
		})
	}
}

// TestExampleStruct demonstrates struct testing with testify
func TestExampleStruct(t *testing.T) {
	t.Run("NewExampleStruct", func(t *testing.T) {
		example := NewExampleStruct("test", 42)
		assert.Equal(t, "test", example.Name, "Name should be set correctly")
		assert.Equal(t, 42, example.Value, "Value should be set correctly")
		assert.NotEmpty(t, example.ID, "ID should be generated")
		assert.Len(t, example.ID, 26, "ULID should be 26 characters long")
		require.NotNil(t, example, "NewExampleStruct should not return nil")
	})

	t.Run("Process", func(t *testing.T) {
		example := NewExampleStruct("test", 42)
		result := example.Process()
		// Should include ID in the output
		assert.Contains(t, result, "test (ID:", "Process should include name and ID")
		assert.Contains(t, result, "has value 42", "Process should include value")
		assert.Contains(t, result, "...", "Process should truncate ID")
	})

	t.Run("Validate - valid", func(t *testing.T) {
		example := NewExampleStruct("test", 42)
		err := example.Validate()
		assert.NoError(t, err, "Valid struct should pass validation")
	})

	t.Run("Validate - empty name", func(t *testing.T) {
		example := NewExampleStruct("", 42)
		err := example.Validate()
		assert.Error(t, err, "Empty name should cause validation error")
		assert.Contains(t, err.Error(), "name cannot be empty", "Error message should mention empty name")
	})

	t.Run("Validate - negative value", func(t *testing.T) {
		example := NewExampleStruct("test", -1)
		err := example.Validate()
		assert.Error(t, err, "Negative value should cause validation error")
		assert.Contains(t, err.Error(), "value cannot be negative", "Error message should mention negative value")
	})
}

// TestExampleFunction_EdgeCases demonstrates additional edge case testing with testify
func TestExampleFunction_EdgeCases(t *testing.T) {
	// Test with whitespace
	result := ExampleFunction("  hello  ")
	assert.Equal(t, "processed:   HELLO  ", result, "Should preserve whitespace")

	// Test with special characters
	result = ExampleFunction("hello@world!")
	assert.Equal(t, "processed: HELLO@WORLD!", result, "Should handle special characters")

	// Test with numbers
	result = ExampleFunction("test123")
	assert.Equal(t, "processed: TEST123", result, "Should handle alphanumeric input")
}

// TestExampleStruct_Constructor demonstrates constructor testing with testify
func TestExampleStruct_Constructor(t *testing.T) {
	t.Run("with positive value", func(t *testing.T) {
		example := NewExampleStruct("positive", 100)
		assert.NotNil(t, example)
		assert.Equal(t, "positive", example.Name)
		assert.Equal(t, 100, example.Value)
	})

	t.Run("with zero value", func(t *testing.T) {
		example := NewExampleStruct("zero", 0)
		assert.NotNil(t, example)
		assert.Equal(t, "zero", example.Name)
		assert.Equal(t, 0, example.Value)
	})

	t.Run("with negative value", func(t *testing.T) {
		example := NewExampleStruct("negative", -5)
		assert.NotNil(t, example)
		assert.Equal(t, "negative", example.Name)
		assert.Equal(t, -5, example.Value)
	})
}

// TestGenerateID tests the standalone ID generation function
func TestGenerateID(t *testing.T) {
	id := GenerateID()
	assert.NotEmpty(t, id, "GenerateID should return a non-empty ID")
	assert.Len(t, id, 26, "Generated ID should be 26 characters long")
}

// TestGetIDTimestamp tests timestamp extraction from IDs
func TestGetIDTimestamp(t *testing.T) {
	id := GenerateID()
	timestamp, err := GetIDTimestamp(id)
	assert.NoError(t, err, "GetIDTimestamp should not return an error for valid ID")
	assert.False(t, timestamp.IsZero(), "Timestamp should not be zero")
	
	// Test with invalid ID
	_, err = GetIDTimestamp("invalid-id")
	assert.Error(t, err, "GetIDTimestamp should return error for invalid ID")
}

// TestGetIDAge tests age calculation from IDs
func TestGetIDAge(t *testing.T) {
	id := GenerateID()
	age, err := GetIDAge(id)
	assert.NoError(t, err, "GetIDAge should not return an error for valid ID")
	assert.True(t, age >= 0, "Age should be non-negative")
	
	// Test with invalid ID
	_, err = GetIDAge("invalid-id")
	assert.Error(t, err, "GetIDAge should return error for invalid ID")
}

// TestCompareIDs tests ID comparison functionality
func TestCompareIDs(t *testing.T) {
	id1 := GenerateID()
	id2 := GenerateID()
	
	// Compare IDs
	cmp, err := CompareIDs(id1, id2)
	assert.NoError(t, err, "CompareIDs should not return an error for valid IDs")
	assert.True(t, cmp == -1 || cmp == 0 || cmp == 1, "Comparison result should be -1, 0, or 1")
	
	// Compare ID with itself
	cmp, err = CompareIDs(id1, id1)
	assert.NoError(t, err, "CompareIDs should not return an error for same ID")
	assert.Equal(t, 0, cmp, "Comparing ID with itself should return 0")
	
	// Test with invalid IDs
	_, err = CompareIDs("invalid-id1", "invalid-id2")
	assert.Error(t, err, "CompareIDs should return error for invalid IDs")
}
