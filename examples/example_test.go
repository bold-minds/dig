package examples_test

import (
	"encoding/json"
	"fmt"

	"github.com/bold-minds/dig"
)

func ExampleDig() {
	data := map[string]any{
		"user": map[string]any{
			"name":  "Alice",
			"email": "alice@example.com",
		},
	}

	name, ok := dig.Dig[string](data, "user", "name")
	fmt.Println(name, ok)
	// Output: Alice true
}

func ExampleDig_deepNested() {
	data := map[string]any{
		"org": map[string]any{
			"team": map[string]any{
				"lead": map[string]any{
					"email": "lead@example.com",
				},
			},
		},
	}

	email, _ := dig.Dig[string](data, "org", "team", "lead", "email")
	fmt.Println(email)
	// Output: lead@example.com
}

func ExampleDig_sliceIndex() {
	data := map[string]any{
		"users": []any{
			map[string]any{"name": "Alice"},
			map[string]any{"name": "Bob"},
		},
	}

	name, _ := dig.Dig[string](data, "users", 1, "name")
	fmt.Println(name)
	// Output: Bob
}

func ExampleDig_fromJSON() {
	raw := []byte(`{"count": 42, "items": ["a", "b", "c"]}`)

	var data any
	_ = json.Unmarshal(raw, &data)

	// JSON numbers unmarshal as float64
	count, _ := dig.Dig[float64](data, "count")
	fmt.Println(count)

	// JSON arrays unmarshal as []any
	item, _ := dig.Dig[string](data, "items", 1)
	fmt.Println(item)
	// Output:
	// 42
	// b
}

func ExampleDigOr() {
	cfg := map[string]any{
		"host": "example.com",
	}

	host := dig.DigOr(cfg, "localhost", "host")
	port := dig.DigOr(cfg, 8080, "port") // missing, fallback used

	fmt.Println(host, port)
	// Output: example.com 8080
}

func ExampleHas() {
	cfg := map[string]any{
		"features": map[string]any{
			"beta": true,
		},
	}

	fmt.Println(dig.Has(cfg, "features", "beta"))
	fmt.Println(dig.Has(cfg, "features", "alpha"))
	// Output:
	// true
	// false
}

func ExampleAt() {
	data := map[string]any{
		"value": "hello",
	}

	// Use At when you want to inspect a value's runtime type
	val, ok := dig.At(data, "value")
	if ok {
		switch v := val.(type) {
		case string:
			fmt.Println("string:", v)
		case []any:
			fmt.Println("slice:", v)
		}
	}
	// Output: string: hello
}
