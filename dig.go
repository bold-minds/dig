// Package dig provides safe nested data navigation for Go.
//
// It replaces chains of .(map[string]any) type assertions with a single
// function call. Every operation is nil-safe, never panics, and never
// modifies its input. All functions are safe for concurrent use, provided
// the caller does not mutate the input concurrently.
//
// A nil data argument always returns (zero, false) / false, regardless of
// path length, including an empty path.
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
// A literal nil leaf value is treated as a successful result when T is
// an interface type (such as any): Dig returns (nil, true). For concrete
// T, a nil leaf returns (zero, false). Note that a typed nil pointer
// (e.g. (*Foo)(nil) stored in an any) is not a literal nil and matches
// T=*Foo normally, returning (nil, true).
//
// Supported source types along the path:
//   - map[string]any (string keys)
//   - map[any]any (see key whitelist below)
//   - []any (non-negative int indices)
//
// For map[any]any, the path key must be one of the Go primitive types
// commonly produced by JSON/YAML unmarshalling, plus the signed/unsigned
// integer and floating-point widths: string, bool, int, int32, int64,
// uint, uint32, uint64, float32, float64. Other key types — including
// hashable ones like structs or int8 — are rejected with (zero, false)
// rather than used as a lookup key. This whitelist exists so that an
// unhashable path value (e.g. a slice) cannot trigger a runtime panic
// from Go's map lookup, preserving the never-panic guarantee without
// reflection.
//
// Typed containers like map[string]string or []string are not walked;
// use encoding/json or similar to unmarshal into any first.
func Dig[T any](data any, path ...any) (T, bool) {
	var zero T
	current, ok := walk(data, path)
	if !ok {
		return zero, false
	}
	// A nil leaf is a valid value when T is an interface type.
	// any(zero) == nil is true exactly when T is an interface type: the
	// zero value of a concrete type boxed in any carries a non-nil type
	// descriptor, while the zero value of an interface type boxed in any
	// is the nil interface itself. This lets us branch on "is T an
	// interface?" at runtime without reflection.
	if current == nil {
		if any(zero) == nil {
			return zero, true
		}
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
// Equivalent to Dig[any] but named for the outcome. A literal nil leaf
// returns (nil, true) — matching Dig[any] and Has.
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
			// Restrict keys to a whitelist of hashable primitive types.
			// Using an arbitrary value as a map key would panic at runtime
			// if the caller passed an unhashable value (e.g., a slice).
			// This preserves the "never panics" guarantee without reflection.
			switch key.(type) {
			case string, int, int64, int32, float64, float32, bool, uint, uint64, uint32:
			default:
				return nil, false
			}
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
