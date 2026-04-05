package dig_test

import (
	"testing"

	"github.com/bold-minds/dig"
)

var (
	shallowData = map[string]any{
		"name": "alice",
	}

	deepData = map[string]any{
		"level1": map[string]any{
			"level2": map[string]any{
				"level3": map[string]any{
					"level4": map[string]any{
						"level5": "deep-value",
					},
				},
			},
		},
	}

	mixedData = map[string]any{
		"users": []any{
			map[string]any{"name": "alice", "role": "admin"},
			map[string]any{"name": "bob", "role": "editor"},
			map[string]any{"name": "carol", "role": "admin"},
		},
	}

	// mapAnyAnyData exercises the map[any]any path, which pays the cost
	// of the extra key-type whitelist switch. Benchmarked so regressions
	// in that guard are caught alongside the hot string-map path.
	mapAnyAnyData = map[any]any{
		"name":       "alice",
		int(1):       "one",
		int64(2):     "two",
		float64(3.5): "three-point-five",
		true:         "yes",
	}
)

func BenchmarkDig_Shallow(b *testing.B) {
	b.ReportAllocs()
	for b.Loop() {
		_, _ = dig.Dig[string](shallowData, "name")
	}
}

func BenchmarkDig_Deep(b *testing.B) {
	b.ReportAllocs()
	for b.Loop() {
		_, _ = dig.Dig[string](deepData, "level1", "level2", "level3", "level4", "level5")
	}
}

func BenchmarkDig_Mixed(b *testing.B) {
	b.ReportAllocs()
	for b.Loop() {
		_, _ = dig.Dig[string](mixedData, "users", 1, "name")
	}
}

func BenchmarkDig_Miss(b *testing.B) {
	b.ReportAllocs()
	for b.Loop() {
		_, _ = dig.Dig[string](shallowData, "missing")
	}
}

func BenchmarkDigOr_Hit(b *testing.B) {
	b.ReportAllocs()
	for b.Loop() {
		_ = dig.DigOr(shallowData, "default", "name")
	}
}

func BenchmarkDigOr_Miss(b *testing.B) {
	b.ReportAllocs()
	for b.Loop() {
		_ = dig.DigOr(shallowData, "default", "missing")
	}
}

func BenchmarkHas_Hit(b *testing.B) {
	b.ReportAllocs()
	for b.Loop() {
		_ = dig.Has(shallowData, "name")
	}
}

func BenchmarkHas_Miss(b *testing.B) {
	b.ReportAllocs()
	for b.Loop() {
		_ = dig.Has(shallowData, "missing")
	}
}

func BenchmarkAt_Hit(b *testing.B) {
	b.ReportAllocs()
	for b.Loop() {
		_, _ = dig.At(shallowData, "name")
	}
}

func BenchmarkAt_Miss(b *testing.B) {
	b.ReportAllocs()
	for b.Loop() {
		_, _ = dig.At(shallowData, "missing")
	}
}

func BenchmarkDig_MapAnyAny_StringKey(b *testing.B) {
	b.ReportAllocs()
	for b.Loop() {
		_, _ = dig.Dig[string](mapAnyAnyData, "name")
	}
}

func BenchmarkDig_MapAnyAny_IntKey(b *testing.B) {
	b.ReportAllocs()
	for b.Loop() {
		_, _ = dig.Dig[string](mapAnyAnyData, int(1))
	}
}

func BenchmarkDig_MapAnyAny_Miss(b *testing.B) {
	b.ReportAllocs()
	for b.Loop() {
		_, _ = dig.Dig[string](mapAnyAnyData, "missing")
	}
}

// BenchmarkDig_MapAnyAny_NonWhitelistedKey measures the failure path
// through the whitelist switch. If a regression accidentally reflects
// on the key type (instead of using the compile-time type switch), the
// alloc count here will jump from 0 and catch it.
func BenchmarkDig_MapAnyAny_NonWhitelistedKey(b *testing.B) {
	b.ReportAllocs()
	type k struct{ ID int }
	key := k{ID: 1}
	for b.Loop() {
		_, _ = dig.Dig[string](mapAnyAnyData, key)
	}
}
