package main

import (
	"testing"

	"github.com/bold-minds/oss"
)

// BenchmarkExampleUsage demonstrates benchmarking the example usage
func BenchmarkExampleUsage(b *testing.B) {
	b.ResetTimer()
	
	for i := 0; i < b.N; i++ {
		// Benchmark the main example operations
		_ = oss.ExampleFunction("benchmark input")
		example := oss.NewExampleStruct("benchmark", i)
		_ = example.Process()
		_ = example.Validate()
	}
}

// BenchmarkIDGeneration demonstrates benchmarking ID operations
func BenchmarkIDGeneration(b *testing.B) {
	b.ResetTimer()
	
	for i := 0; i < b.N; i++ {
		id := oss.GenerateID()
		_, _ = oss.GetIDTimestamp(id)
		_, _ = oss.GetIDAge(id)
	}
}

// BenchmarkIDComparison demonstrates benchmarking ID comparison
func BenchmarkIDComparison(b *testing.B) {
	id1 := oss.GenerateID()
	id2 := oss.GenerateID()
	b.ResetTimer()
	
	for i := 0; i < b.N; i++ {
		_, _ = oss.CompareIDs(id1, id2)
	}
}
