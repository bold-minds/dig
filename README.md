# dig

[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![Go Reference](https://pkg.go.dev/badge/github.com/bold-minds/dig.svg)](https://pkg.go.dev/github.com/bold-minds/dig)
[![Go Version](https://img.shields.io/endpoint?url=https://raw.githubusercontent.com/bold-minds/dig/main/.github/badges/go-version.json)](https://golang.org/doc/go1.26)
[![Latest Release](https://img.shields.io/github/v/release/bold-minds/dig?logo=github&color=blueviolet)](https://github.com/bold-minds/dig/releases)
[![Last Updated](https://img.shields.io/endpoint?url=https://raw.githubusercontent.com/bold-minds/dig/main/.github/badges/last-updated.json)](https://github.com/bold-minds/dig/commits)
[![golangci-lint](https://img.shields.io/endpoint?url=https://raw.githubusercontent.com/bold-minds/dig/main/.github/badges/golangci-lint.json)](https://github.com/bold-minds/dig/actions/workflows/test.yaml)
[![Coverage](https://img.shields.io/endpoint?url=https://raw.githubusercontent.com/bold-minds/dig/main/.github/badges/coverage.json)](https://github.com/bold-minds/dig/actions/workflows/test.yaml)
[![Dependabot](https://img.shields.io/endpoint?url=https://raw.githubusercontent.com/bold-minds/dig/main/.github/badges/dependabot.json)](https://github.com/bold-minds/dig/security/dependabot)

**Nested data navigation for Go, one line.**

Every Go codebase that touches JSON, YAML, config files, or API responses ends up writing chains of `.(map[string]any)` type assertions. `dig` replaces the chains with one call.

```go
// Before
addr := data.(map[string]any)["user"].(map[string]any)["contact"].(map[string]any)["address"].(string)

// After
addr, _ := dig.Dig[string](data, "user", "contact", "address")
```

## ✨ Why dig?

- 🎯 **One call, not four assertions** — walk any nested structure in a single line
- 🛡️ **Nil-safe and zero-panic** — missing paths and type mismatches return `(zero, false)`, never panic
- 🧭 **Mixed map and slice paths** — string keys navigate maps, `int` keys navigate slices, in the same path
- ⚡ **No reflection** — concrete type switches, measured in nanoseconds, zero allocations
- 🪶 **Tiny** — four functions, one file, zero dependencies
- 🔗 **Pairs with [`bold-minds/to`](https://github.com/bold-minds/to)** — if you need type coercion at the leaf, chain them

## 📦 Installation

```bash
go get github.com/bold-minds/dig
```

Requires Go 1.26 or later.

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

- **Never panics.** Nil inputs, wrong intermediate types, out-of-bounds indices, missing keys, and leaf type mismatches all return `(zero, false)` or fall through to the fallback.
- **Immutable.** `dig` never modifies input data.
- **No reflection.** Uses concrete type switches for `map[string]any`, `map[any]any`, and `[]any`.
- **Zero dependencies.** Pure stdlib.

## 🏎️ Performance

Measured on Go 1.26 (Intel Ultra 9 275HX). **Zero allocations on every path**, including deep nested navigation.

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

| Current node type  | Path element type | Action                        |
|--------------------|-------------------|-------------------------------|
| `map[string]any`   | `string`          | Look up the key               |
| `map[any]any`      | any               | Look up the key               |
| `[]any`            | `int` (non-neg)   | Index into the slice          |
| anything else      | anything          | Navigation fails, return zero |

If a path element's type doesn't match the current node's expected key or index type, navigation fails cleanly without panic.

## 🤝 Contributing

See [CONTRIBUTING.md](CONTRIBUTING.md). Bold Minds Go libraries follow a shared set of design principles; read [PRINCIPLES.md](https://github.com/bold-minds/oss/blob/main/PRINCIPLES.md) before opening a PR.

## 📄 License

MIT. See [LICENSE](LICENSE).

## 🔗 Related Projects

- [`bold-minds/to`](https://github.com/bold-minds/to) — Go value conversion. Chain with `dig` when you need type coercion at the leaf (e.g., `to.Int(dig.At(data, "count"))`).
- [`tidwall/gjson`](https://github.com/tidwall/gjson) — JSON path queries directly from `[]byte`. Different problem space: `gjson` parses JSON strings; `dig` walks already-unmarshaled Go values.
- Go standard library type assertions — `.(map[string]any)` chains. The pattern `dig` replaces.
- [`samber/lo`](https://github.com/samber/lo) — General-purpose Go utility library with ~200 helpers. `dig` focuses narrowly on nested path navigation.
