package oss

import (
	"testing"
)

// BenchmarkExampleFunction demonstrates benchmark testing for functions
func BenchmarkExampleFunction(b *testing.B) {
	input := "benchmark test input"
	b.ResetTimer()
	
	for i := 0; i < b.N; i++ {
		_ = ExampleFunction(input)
	}
}

// BenchmarkExampleStruct_Process demonstrates benchmark testing for methods
func BenchmarkExampleStruct_Process(b *testing.B) {
	example := NewExampleStruct("benchmark", 100)
	b.ResetTimer()
	
	for i := 0; i < b.N; i++ {
		_ = example.Process()
	}
}

// BenchmarkExampleStruct_Validate demonstrates benchmark testing for validation
func BenchmarkExampleStruct_Validate(b *testing.B) {
	example := NewExampleStruct("benchmark", 100)
	b.ResetTimer()
	
	for i := 0; i < b.N; i++ {
		_ = example.Validate()
	}
}

// BenchmarkNewExampleStruct demonstrates benchmark testing for constructors
func BenchmarkNewExampleStruct(b *testing.B) {
	b.ResetTimer()
	
	for i := 0; i < b.N; i++ {
		_ = NewExampleStruct("benchmark", i)
	}
}