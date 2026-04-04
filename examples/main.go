// Package main demonstrates usage of the ossgo package
// This serves as an example for users of your library
package main

import (
	"fmt"
	"log"
	"time"

	"github.com/bold-minds/oss"
)

func main() {
	fmt.Println("=== ossgo Package Example ===")
	fmt.Printf("Package Version: %s\n\n", oss.Version)

	// Example 1: Basic function usage
	fmt.Println("1. Basic Function Usage:")
	result1 := oss.ExampleFunction("hello world")
	fmt.Printf("   Input: %q\n", "hello world")
	fmt.Printf("   Output: %q\n\n", result1)

	result2 := oss.ExampleFunction("")
	fmt.Printf("   Input: %q (empty)\n", "")
	fmt.Printf("   Output: %q\n\n", result2)

	// Example 2: Struct creation and usage (with bold-minds/id integration)
	fmt.Println("2. Struct Usage with ID Generation:")
	example := oss.NewExampleStruct("demo", 42)
	fmt.Printf("   Created: Name=%s, Value=%d, ID=%s\n", example.Name, example.Value, example.ID)
	fmt.Printf("   Process(): %q\n\n", example.Process())

	// Example 2b: Demonstrate ID operations using bold-minds/id
	fmt.Println("2b. ID Operations (bold-minds/id package):")
	generatedID := oss.GenerateID()
	fmt.Printf("   Generated ID: %s\n", generatedID)
	
	if timestamp, err := oss.GetIDTimestamp(generatedID); err == nil {
		fmt.Printf("   ID Timestamp: %s\n", timestamp.Format("2006-01-02 15:04:05"))
	}
	
	if age, err := oss.GetIDAge(generatedID); err == nil {
		fmt.Printf("   ID Age: %v\n\n", age)
	}

	// Example 3: Validation
	fmt.Println("3. Validation Examples:")
	
	// Valid example
	validExample := oss.NewExampleStruct("valid", 100)
	if err := validExample.Validate(); err != nil {
		log.Printf("   Validation error: %v\n", err)
	} else {
		fmt.Println("   ✓ Valid example passed validation")
	}

	// Invalid example - empty name
	invalidExample1 := oss.NewExampleStruct("", 50)
	if err := invalidExample1.Validate(); err != nil {
		fmt.Printf("   ✗ Invalid example (empty name): %v\n", err)
	}

	// Invalid example - negative value
	invalidExample2 := oss.NewExampleStruct("test", -10)
	if err := invalidExample2.Validate(); err != nil {
		fmt.Printf("   ✗ Invalid example (negative value): %v\n", err)
	}

	// Example 4: ID Comparison (demonstrates bold-minds/id capabilities)
	fmt.Println("4. ID Comparison:")
	id1 := oss.GenerateID()
	time.Sleep(1 * time.Millisecond) // Small delay to ensure different timestamps
	id2 := oss.GenerateID()
	
	if cmp, err := oss.CompareIDs(id1, id2); err == nil {
		switch {
		case cmp < 0:
			fmt.Printf("   %s was created before %s\n", id1[:8]+"...", id2[:8]+"...")
		case cmp > 0:
			fmt.Printf("   %s was created after %s\n", id1[:8]+"...", id2[:8]+"...")
		default:
			fmt.Printf("   %s and %s were created at the same time\n", id1[:8]+"...", id2[:8]+"...")
		}
	}

	fmt.Println("\n=== Example Complete ===")
	fmt.Println("This template demonstrates integration with bold-minds/id package!")
	fmt.Println("Replace this example with your actual package usage!")
}
