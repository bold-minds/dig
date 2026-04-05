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
