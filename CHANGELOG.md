# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Changed
- Package documentation now lists the full `map[any]any` key whitelist
  (`string`, `bool`, `int`/`int32`/`int64`, `uint`/`uint32`/`uint64`,
  `float32`/`float64`) and explains the rationale — the whitelist exists
  to preserve the never-panic guarantee against unhashable path values,
  not to minimize supported types.
- Package documentation now states that all functions are safe for
  concurrent use on immutable input, and that a nil `data` argument
  always returns an unsuccessful result (including for an empty path).
- `go.mod` now requires Go 1.26, matching the CHANGELOG "Requires" note.
- Benchmarks migrated to `b.Loop()` and now report allocations via
  `b.ReportAllocs()` so allocation regressions surface immediately.

### Added
- `FuzzWalk` exercises `Dig`, `DigOr`, `Has`, and `At` with fuzzer-driven
  paths — including unhashable and non-whitelisted key types — against a
  pool of nested structures. Protects the never-panic guarantee.
- Table-driven test covering every whitelisted `map[any]any` key type.
- Test for typed nil pointer leaves (`(*Foo)(nil)` in an `any` is not a
  literal nil and matches `T=*Foo` as `(nil, true)`).
- Test for `DigOr[any]` on a literal nil leaf, confirming it returns the
  real nil value rather than the fallback.

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
