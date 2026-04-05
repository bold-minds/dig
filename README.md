# dig

[![Go Reference](https://pkg.go.dev/badge/github.com/bold-minds/dig.svg)](https://pkg.go.dev/github.com/bold-minds/dig)
[![Build](https://img.shields.io/github/actions/workflow/status/bold-minds/dig/test.yaml?branch=main&label=tests)](https://github.com/bold-minds/dig/actions/workflows/test.yaml)
[![Go Version](https://img.shields.io/github/go-mod/go-version/bold-minds/dig)](go.mod)

**Nested data navigation for Go, one line.**

Every Go codebase that touches JSON, YAML, config files, or API responses ends up writing chains of `.(map[string]any)` type assertions. `dig` replaces the chains with one call.

```go
// Before — unsafe: panics on any mismatch
addr := data.(map[string]any)["user"].(map[string]any)["contact"].(map[string]any)["address"].(string)

// Before — safe but verbose: nested , ok := ... .(map[string]any) chains
user, ok := data.(map[string]any)
if !ok { return }
contact, ok := user["user"].(map[string]any)
if !ok { return }
address, ok := contact["contact"].(map[string]any)
if !ok { return }
addr, ok := address["address"].(string)
if !ok { return }

// After
addr, _ := dig.Dig[string](data, "user", "contact", "address")
```

## ✨ Why dig?

- 🎯 **One call, not four assertions** — walk any nested structure in a single line
- 🛡️ **Nil-safe and zero-panic** — missing paths and type mismatches return `(zero, false)`, never panic
- 🧭 **Mixed map and slice paths** — string keys navigate maps, `int` keys navigate slices, in the same path
- ⚡ **No reflection** — concrete type switches, measured in nanoseconds; the walk itself never allocates
- 🪶 **Tiny** — four functions, one file, zero dependencies
- 🔗 **Pairs with [`bold-minds/to`](https://github.com/bold-minds/to)** — if you need type coercion at the leaf, chain them

## 📦 Installation

```bash
go get github.com/bold-minds/dig
```

Requires Go 1.26 or later (see `go.mod`).

## 🚀 Quick Start

```go
package main

import (
    "encoding/json"
    "fmt"

    "github.com/bold-minds/dig"
)

func main() {
    raw := []byte(`{
        "user": {
            "name": "Alice",
            "age": 30,
            "posts": [
                {"title": "First post"},
                {"title": "Second post"}
            ]
        }
    }`)

    var data any
    _ = json.Unmarshal(raw, &data)

    // Typed extraction — use the Go type json.Unmarshal produces.
    // JSON numbers → float64, JSON strings → string, JSON arrays → []any.
    name, _ := dig.Dig[string](data, "user", "name")
    fmt.Println(name) // "Alice"

    age, _ := dig.Dig[float64](data, "user", "age")
    fmt.Println(age) // 30

    // Mix map keys and slice indices in the same path
    title, _ := dig.Dig[string](data, "user", "posts", 1, "title")
    fmt.Println(title) // "Second post"

    // With a fallback
    role := dig.DigOr(data, "member", "user", "role")
    fmt.Println(role) // "member"

    // Existence check
    if dig.Has(data, "user", "posts") {
        fmt.Println("user has posts")
    }
}
```

## 🔧 Core Features

### Typed extraction with `Dig`

Specify the destination type at the call site. `Dig` walks the path and returns the leaf value if it matches type `T` exactly.

```go
name, ok   := dig.Dig[string](data, "user", "name")
age, ok    := dig.Dig[float64](data, "user", "age")
active, ok := dig.Dig[bool](data, "user", "active")
```

Returns `(zero, false)` if the path is missing, the traversal hits the wrong type, or the leaf value cannot be converted to `T`.

**`dig` uses strict type matching at the leaf** — no automatic conversion. If the stored value is `float64` and you ask for `int`, you get `(0, false)`. If you need type coercion (e.g., JSON numbers as `int`), chain with [`bold-minds/to`](https://github.com/bold-minds/to):

```go
raw, ok := dig.At(data, "user", "age")
if !ok { return }
age := to.Int(raw) // to handles the numeric conversion
```

### Fallback with `DigOr`

When you don't want to check `ok` — just give me a value or a sensible default:

```go
timeout := dig.DigOr(cfg, 30, "server", "timeout_seconds")
host    := dig.DigOr(cfg, "localhost", "server", "host")
debug   := dig.DigOr(cfg, false, "logging", "debug")
```

Returns the fallback on any failure: missing path, wrong intermediate type, or leaf type mismatch.

### Existence check with `Has`

When you only want to know whether the path exists:

```go
if dig.Has(cfg, "features", "beta_enabled") {
    enableBeta()
}

if !dig.Has(user, "email") {
    return errors.New("user must have an email")
}
```

Returns `true` if the path resolves to any value, `false` on any navigation failure. Leaf type is not checked.

### Raw value access with `At`

When you need to inspect a value's actual type before deciding how to handle it:

```go
val, ok := dig.At(data, "user", "preferences")
if !ok { return }

switch v := val.(type) {
case string:
    handleSinglePref(v)
case []any:
    handlePrefList(v)
case map[string]any:
    handlePrefMap(v)
}
```

`At` returns the raw `any` value at the path without attempting type matching. Equivalent to `Dig[any](data, path...)` but named for the outcome.

### Mixed map and slice navigation

Path elements are navigated based on the current node's type. String keys navigate maps (`map[string]any`, `map[any]any`). Integer keys navigate slices (`[]any`). Both work in the same path:

```go
// Third user's second post's tags
tags, _ := dig.Dig[[]any](data, "users", 2, "posts", 1, "tags")
```

## 🛡️ Safety guarantees

- **Never panics.** Nil inputs, wrong intermediate types, out-of-bounds indices, missing keys, leaf type mismatches, and unhashable `map[any]any` keys all return `(zero, false)` or fall through to the fallback.
- **Immutable.** `dig` never modifies input data.
- **No reflection.** Uses concrete type switches for `map[string]any`, `map[any]any`, and `[]any`.
- **Zero dependencies.** Pure stdlib.

## 🏎️ Performance

Measured on Go 1.26 (Intel Ultra 9 275HX). **Zero allocations inside the walk itself**, including deep nested navigation. The variadic `path ...any` is a slice header passed by the caller; in the benchmarks below (and in typical code that passes path literals), the compiler's escape analysis keeps that slice on the caller's stack, so the whole call costs zero heap allocations. If you construct a `[]any` from dynamic data and pass it with `path...`, the allocation happens at the construction site, not inside `dig`.

```
BenchmarkDig_Shallow-24    201598330    5.77 ns/op    0 B/op    0 allocs/op
BenchmarkDig_Deep-24        54915938   22.47 ns/op    0 B/op    0 allocs/op
BenchmarkDig_Mixed-24      100000000   10.96 ns/op    0 B/op    0 allocs/op
BenchmarkDig_Miss-24       184699155    6.58 ns/op    0 B/op    0 allocs/op
BenchmarkDigOr_Hit-24      204561072    5.81 ns/op    0 B/op    0 allocs/op
BenchmarkDigOr_Miss-24     184626600    6.39 ns/op    0 B/op    0 allocs/op
BenchmarkHas_Hit-24        229360736    5.23 ns/op    0 B/op    0 allocs/op
BenchmarkHas_Miss-24       200012004    6.01 ns/op    0 B/op    0 allocs/op
BenchmarkAt_Hit-24         236363647    5.12 ns/op    0 B/op    0 allocs/op
BenchmarkAt_Miss-24        196614397    5.96 ns/op    0 B/op    0 allocs/op
```

A 5-level deep dig runs in ~22 nanoseconds. A shallow dig runs in ~5 nanoseconds. Both with zero heap allocations.

## 🧪 Testing

```bash
# Run all tests
go test ./...

# Run with race detection
go test -race ./...

# Run benchmarks
go test -bench=. -benchmem ./...
```

Current coverage: 97.2%.

## 📚 API Reference

```go
// Dig extracts a value of type T at the given path from nested data.
// Returns (zero, false) if any path element is missing, the traversal
// hits the wrong type, or the leaf value's type is not exactly T.
//
// Dig uses strict type matching at the leaf — no automatic conversion.
// For type coercion, chain with bold-minds/to.
func Dig[T any](data any, path ...any) (T, bool)

// DigOr extracts a value of type T at the given path, returning fallback
// on any failure (missing path, wrong intermediate type, or leaf mismatch).
func DigOr[T any](data any, fallback T, path ...any) T

// Has reports whether the given path resolves to a value in data,
// regardless of that value's type.
func Has(data any, path ...any) bool

// At returns the raw value at the given path without type matching.
// Use when you need to inspect a value's actual type before handling it.
func At(data any, path ...any) (any, bool)
```

### Path element rules

Each path element is interpreted based on the current node's Go type:

| Current node type  | Path element type                       | Action                        |
|--------------------|-----------------------------------------|-------------------------------|
| `map[string]any`   | `string`                                | Look up the key               |
| `map[any]any`      | `string`, `int`/`int64`/`int32`, `uint`/`uint64`/`uint32`, `float64`/`float32`, `bool` | Look up the key               |
| `[]any`            | `int` (non-negative, exact type)        | Index into the slice          |
| anything else      | anything                                | Navigation fails, return zero |

If a path element's type doesn't match the current node's expected key or index type, navigation fails cleanly without panic.

**Strict type matching for path elements:**

- Slice indices must be exactly `int` — `int64`, `uint`, `float64`, and friends are rejected. If your index came from a JSON number (which `encoding/json` decodes as `float64`), a `len()`-derived `uint`, or any other typed integer, convert to `int` first:

  ```go
  // JSON-decoded index: float64 → int
  idxF, _ := dig.Dig[float64](data, "selected")
  item, _ := dig.Dig[string](data, "items", int(idxF))

  // Typed loop variable: int64 → int
  var i int64 = 2
  tag, _ := dig.Dig[string](data, "tags", int(i))
  ```
- `map[any]any` keys are restricted to a whitelist of hashable primitive types (see table above). This guarantees `dig` never panics even if a caller passes an unhashable value like a slice. Non-whitelisted key types — including hashable ones like `int8`, `int16`, `uint8`, `uint16`, `uintptr`, `complex64`, `complex128`, structs, arrays, and pointers — fail cleanly with `(zero, false)` **even when the map literally contains such a key**. If your decoder produces narrow integers (some YAML libraries decode small ints as `uint8`), convert the key to `int` or `int64` before calling `Dig`.

**Supported container types:**

`dig` walks exactly three container types: `map[string]any`, `map[any]any`, and `[]any`. These are the types produced by `encoding/json` (and most YAML decoders) when unmarshaling into `any`. Typed containers such as `map[string]string`, `map[string]int`, or `[]string` are **not** walked — navigation fails cleanly at the typed container. If your data uses typed containers, convert to `any`-based containers first, or use reflection-based alternatives.

## 🤝 Contributing

See [CONTRIBUTING.md](CONTRIBUTING.md). Bold Minds Go libraries follow a shared set of design principles; read [PRINCIPLES.md](https://github.com/bold-minds/oss/blob/main/PRINCIPLES.md) before opening a PR.

## 📄 License

MIT. See [LICENSE](LICENSE).

## 🔗 Related Projects

- [`bold-minds/to`](https://github.com/bold-minds/to) — Go value conversion. Chain with `dig` when you need type coercion at the leaf (e.g., `to.Int(dig.At(data, "count"))`).
- [`tidwall/gjson`](https://github.com/tidwall/gjson) — JSON path queries directly from `[]byte`. Different problem space: `gjson` parses JSON strings; `dig` walks already-unmarshaled Go values.
- Go standard library type assertions — `.(map[string]any)` chains. The pattern `dig` replaces.
- [`samber/lo`](https://github.com/samber/lo) — General-purpose Go utility library with ~200 helpers. `dig` focuses narrowly on nested path navigation.
