# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Changed
- **Behavior:** `Dig[any](nil)`, `At(nil)`, and `Has(nil)` with an empty
  path now return `(nil, true)` / `true` — a nil data argument with an
  empty path is treated as a nil leaf, consistent with the existing
  "nil leaf is a real value for interface T" rule applied to nested
  nils. `Dig[string](nil)` (concrete `T`) still returns `(zero, false)`.
  Nil data with a **non-empty** path still fails, as before.
- Package documentation now lists the full `map[any]any` key whitelist
  (`string`, `bool`, `int`/`int32`/`int64`, `uint`/`uint32`/`uint64`,
  `float32`/`float64`) and explicitly calls out non-whitelisted hashable
  types (`int8`, `int16`, `uint8`, `uint16`, `uintptr`, `complex64`,
  `complex128`, structs, arrays, pointers) as silent-miss hazards, with
  guidance to convert narrow integer keys to `int`/`int64` first.
- Package documentation now states that all functions are safe for
  concurrent use on immutable input.
- README softens the "zero allocations on every path" claim: zero
  allocations apply to the walk itself; variadic `path ...any` literals
  typically stay stack-allocated thanks to escape analysis, but a
  caller-constructed `[]any` slice is allocated at the construction
  site, not inside `dig`.
- README version floor, Go version badge link, and `scripts/validate.sh`
  updated to Go 1.26 to match `go.mod`. The CI workflow Go version is
  also bumped to 1.26.
- `go.mod` now requires Go 1.26, matching the CHANGELOG "Requires" note.
- Benchmarks migrated to `b.Loop()` and now report allocations via
  `b.ReportAllocs()` so allocation regressions surface immediately.

### Added
- `FuzzWalk` exercises `Dig`, `DigOr`, `Has`, and `At` with fuzzer-driven
  paths — including unhashable, non-whitelisted, and nil path elements —
  against a pool of nested structures. Protects the never-panic guarantee.
- Table-driven test covering every whitelisted `map[any]any` key type,
  plus a negative-enumeration test that locks the whitelist boundary in
  place for `int8`, `int16`, `uint8`, `uint16`, `uintptr`, `complex64`,
  `complex128`, structs, arrays, and pointers.
- Test for typed nil pointer leaves (`(*Foo)(nil)` in an `any` is not a
  literal nil and matches `T=*Foo` as `(nil, true)`).
- Test for `DigOr[any]` on a literal nil leaf, confirming it returns the
  real nil value rather than the fallback.
- Test for the nil-data + empty-path edge case across `Dig[any]`, `At`,
  `Has`, and `Dig[string]`.
- `ExampleDig_mapAnyAny` runnable example demonstrating navigation
  through the `map[any]any` containers that YAML decoders produce.
- Benchmarks for the `map[any]any` path (string key, int key, miss) so
  regressions in the key-whitelist switch are caught alongside the
  hot `map[string]any` path.
- `govulncheck` now runs in CI on Linux as an additional security
  posture check.

## [0.1.0] — Initial release

### Added
- `Dig[T any](data any, path ...any) (T, bool)` — extract a typed value from nested data
- `DigOr[T any](data any, fallback T, path ...any) T` — extract with fallback on failure
- `Has(data any, path ...any) bool` — existence check regardless of leaf type
- `At(data any, path ...any) (any, bool)` — extract raw value without type matching
- Support for navigation through `map[string]any`, `map[any]any`, and `[]any`
- Strict leaf type matching (no automatic conversion; chain with `bold-minds/to` for coercion)
- Nil-safe, zero-panic guarantees across all functions
- 97%+ test coverage including a real-world integration test
- Zero-allocation benchmarks for all operations (5–22 ns/op)
- Zero external dependencies — pure stdlib

### Requires
- Go 1.26 or later
