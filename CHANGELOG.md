# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

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
