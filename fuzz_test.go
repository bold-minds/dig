package dig_test

import (
	"testing"

	"github.com/bold-minds/dig"
)

// FuzzWalk is the load-bearing safety test for dig's "never panics"
// guarantee. It walks a small pool of nested structures — including
// map[any]any with exotic key types and deeply nested mixed containers —
// using fuzzer-provided path elements of varying types. Any panic from
// Dig, Has, At, or DigOr is a regression of the core contract.
//
// The fuzzer cannot produce arbitrary Go values directly, so we derive
// path elements and a structure selector from the raw byte/string inputs
// it does provide. Each byte in the path blob picks one element from a
// fixed menu of Go types, including unhashable ones (slices) that would
// panic a naive map[any]any lookup.
func FuzzWalk(f *testing.F) {
	// Seed corpus: a handful of hand-picked inputs so the fuzzer starts
	// from realistic shapes before mutating.
	f.Add(uint8(0), "name", []byte{0})
	f.Add(uint8(1), "user.contact.email", []byte{0, 0, 0})
	f.Add(uint8(2), "0", []byte{1})
	f.Add(uint8(3), "k", []byte{0, 7, 8}) // includes unhashable-key index
	f.Add(uint8(4), "", []byte{})
	f.Add(uint8(0), "", []byte{9, 9, 9})

	// The structure pool. Each fuzzer input selects one of these via
	// (structSel % len(pool)).
	pool := []any{
		// 0: simple flat map[string]any
		map[string]any{
			"name": "alice",
			"age":  30,
			"nil":  nil,
		},
		// 1: deeply nested map[string]any
		map[string]any{
			"user": map[string]any{
				"contact": map[string]any{
					"email": "alice@example.com",
				},
			},
		},
		// 2: slice of maps
		map[string]any{
			"items": []any{
				map[string]any{"k": "v0"},
				map[string]any{"k": "v1"},
			},
		},
		// 3: map[any]any with whitelisted key types
		map[any]any{
			"k":           "string-key",
			int(1):        "int-key",
			int64(2):      "int64-key",
			float64(1.5):  "float-key",
			true:          "bool-key",
			uint32(7):     "uint32-key",
		},
		// 4: nil — should never panic regardless of path
		nil,
	}

	// Path-element menu. Indices 7 and 8 are deliberately unhashable or
	// non-whitelisted values that MUST be rejected without panicking.
	pathMenu := func(b byte, strHint string) any {
		switch b % 10 {
		case 0:
			return strHint
		case 1:
			return int(int(b))
		case 2:
			return int32(b)
		case 3:
			return int64(b)
		case 4:
			return uint(b)
		case 5:
			return float64(b)
		case 6:
			return true
		case 7:
			return []int{int(b)} // unhashable — must not panic the map lookup
		case 8:
			return struct{ N int }{N: int(b)} // hashable but not whitelisted
		default:
			return nil
		}
	}

	f.Fuzz(func(t *testing.T, structSel uint8, strHint string, pathBlob []byte) {
		// Cap path length so the fuzzer stays fast and stack-bounded.
		if len(pathBlob) > 16 {
			pathBlob = pathBlob[:16]
		}
		data := pool[int(structSel)%len(pool)]
		path := make([]any, len(pathBlob))
		for i, b := range pathBlob {
			path[i] = pathMenu(b, strHint)
		}

		// None of these may panic. We don't care about the return
		// values — only that every entry point stays panic-free.
		_, _ = dig.Dig[any](data, path...)
		_, _ = dig.Dig[string](data, path...)
		_, _ = dig.Dig[int](data, path...)
		_ = dig.DigOr[any](data, "fallback", path...)
		_ = dig.Has(data, path...)
		_, _ = dig.At(data, path...)
	})
}
