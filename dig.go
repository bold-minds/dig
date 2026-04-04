// Package dig provides safe nested data navigation for Go.
//
// It replaces chains of .(map[string]any) type assertions with a single
// function call. Every operation is nil-safe, never panics, and never
// modifies its input.
//
// For documentation and examples, see https://github.com/bold-minds/dig.
package dig

// Dig extracts a value of type T at the given path from nested data.
// It returns (zero, false) on any failure: a missing path element, a
// wrong intermediate type, an out-of-bounds index, or a leaf value whose
// type is not exactly T.
//
// Dig uses strict type matching at the leaf — no automatic conversion.
// For type coercion, chain with bold-minds/to:
//
//	raw, _ := dig.At(data, "user", "age")
//	age := to.Int(raw)
//
// Supported source types along the path:
//   - map[string]any (string keys)
//   - map[any]any (any comparable keys)
//   - []any (non-negative int indices)
func Dig[T any](data any, path ...any) (T, bool) {
	var zero T
	current, ok := walk(data, path)
	if !ok {
		return zero, false
	}
	result, ok := current.(T)
	if !ok {
		return zero, false
	}
	return result, true
}

// DigOr extracts a value of type T at the given path, returning fallback
// on any failure. Equivalent to Dig[T] with the (zero, false) case
// replaced by the fallback value.
func DigOr[T any](data any, fallback T, path ...any) T {
	if v, ok := Dig[T](data, path...); ok {
		return v
	}
	return fallback
}

// Has reports whether the given path resolves to a value in data,
// regardless of that value's type. Returns false for nil data, missing
// keys, wrong intermediate types, and out-of-bounds indices.
func Has(data any, path ...any) bool {
	_, ok := walk(data, path)
	return ok
}

// At returns the raw value at the given path without type matching.
// Use when you need to inspect a value's actual type before handling it.
// Equivalent to Dig[any] but named for the outcome.
func At(data any, path ...any) (any, bool) {
	return walk(data, path)
}

// walk navigates a path through nested data structures. It returns the
// final value and true on success, or (nil, false) on any navigation
// failure: nil input, missing key, wrong intermediate type, or
// out-of-bounds index.
//
// Path element types are matched against the current node type:
// string keys navigate maps, non-negative int indices navigate slices.
func walk(data any, path []any) (any, bool) {
	if data == nil {
		return nil, false
	}
	current := data
	for _, key := range path {
		switch v := current.(type) {
		case map[string]any:
			strKey, ok := key.(string)
			if !ok {
				return nil, false
			}
			val, exists := v[strKey]
			if !exists {
				return nil, false
			}
			current = val
		case map[any]any:
			val, exists := v[key]
			if !exists {
				return nil, false
			}
			current = val
		case []any:
			idx, ok := key.(int)
			if !ok || idx < 0 || idx >= len(v) {
				return nil, false
			}
			current = v[idx]
		default:
			return nil, false
		}
	}
	return current, true
}
