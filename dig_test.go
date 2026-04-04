package dig_test

import (
	"testing"

	"github.com/bold-minds/dig"
)

// =============================================================================
// Dig
// =============================================================================

func TestDig_ShallowMap(t *testing.T) {
	data := map[string]any{"name": "alice"}
	got, ok := dig.Dig[string](data, "name")
	if !ok || got != "alice" {
		t.Errorf("got (%q, %v), want (alice, true)", got, ok)
	}
}

func TestDig_DeepMap(t *testing.T) {
	data := map[string]any{
		"user": map[string]any{
			"contact": map[string]any{
				"email": "alice@example.com",
			},
		},
	}
	got, ok := dig.Dig[string](data, "user", "contact", "email")
	if !ok || got != "alice@example.com" {
		t.Errorf("got (%q, %v), want (alice@example.com, true)", got, ok)
	}
}

func TestDig_SliceIndex(t *testing.T) {
	data := map[string]any{
		"items": []any{"first", "second", "third"},
	}
	got, ok := dig.Dig[string](data, "items", 1)
	if !ok || got != "second" {
		t.Errorf("got (%q, %v), want (second, true)", got, ok)
	}
}

func TestDig_MixedMapAndSlice(t *testing.T) {
	data := map[string]any{
		"users": []any{
			map[string]any{"name": "alice"},
			map[string]any{"name": "bob"},
		},
	}
	got, ok := dig.Dig[string](data, "users", 1, "name")
	if !ok || got != "bob" {
		t.Errorf("got (%q, %v), want (bob, true)", got, ok)
	}
}

func TestDig_MapAnyAny(t *testing.T) {
	data := map[any]any{42: "forty-two"}
	got, ok := dig.Dig[string](data, 42)
	if !ok || got != "forty-two" {
		t.Errorf("got (%q, %v), want (forty-two, true)", got, ok)
	}
}

func TestDig_MissingKey(t *testing.T) {
	data := map[string]any{"name": "alice"}
	if _, ok := dig.Dig[string](data, "missing"); ok {
		t.Error("expected ok=false for missing key")
	}
}

func TestDig_WrongIntermediateType(t *testing.T) {
	data := map[string]any{"name": "alice"} // string, not a map
	if _, ok := dig.Dig[string](data, "name", "first"); ok {
		t.Error("expected ok=false when digging into a string as a map")
	}
}

func TestDig_LeafTypeMismatch(t *testing.T) {
	data := map[string]any{"age": 30}
	if _, ok := dig.Dig[string](data, "age"); ok {
		t.Error("expected ok=false for int→string leaf mismatch")
	}
}

func TestDig_OutOfBoundsIndex(t *testing.T) {
	data := map[string]any{"items": []any{"a", "b"}}
	if _, ok := dig.Dig[string](data, "items", 5); ok {
		t.Error("expected ok=false for out-of-bounds index")
	}
}

func TestDig_NegativeIndexRejected(t *testing.T) {
	data := map[string]any{"items": []any{"a", "b"}}
	if _, ok := dig.Dig[string](data, "items", -1); ok {
		t.Error("expected ok=false for negative index")
	}
}

func TestDig_StringKeyOnSlice(t *testing.T) {
	data := map[string]any{"items": []any{"a", "b"}}
	if _, ok := dig.Dig[string](data, "items", "first"); ok {
		t.Error("expected ok=false for string key on slice")
	}
}

func TestDig_IntKeyOnStringMap(t *testing.T) {
	data := map[string]any{"0": "first"}
	if _, ok := dig.Dig[string](data, 0); ok {
		t.Error("expected ok=false for int key on map[string]any")
	}
}

func TestDig_NilData(t *testing.T) {
	if _, ok := dig.Dig[string](nil, "any", "path"); ok {
		t.Error("expected ok=false for nil data")
	}
}

func TestDig_EmptyPathMatchingType(t *testing.T) {
	got, ok := dig.Dig[string]("hello")
	if !ok || got != "hello" {
		t.Errorf("got (%q, %v), want (hello, true)", got, ok)
	}
}

func TestDig_EmptyPathTypeMismatch(t *testing.T) {
	if _, ok := dig.Dig[int]("hello"); ok {
		t.Error("expected ok=false for empty path with wrong type")
	}
}

func TestDig_AnyType(t *testing.T) {
	data := map[string]any{"mixed": 42}
	got, ok := dig.Dig[any](data, "mixed")
	if !ok || got != 42 {
		t.Errorf("got (%v, %v), want (42, true)", got, ok)
	}
}

// =============================================================================
// DigOr
// =============================================================================

func TestDigOr_Hit(t *testing.T) {
	data := map[string]any{"name": "alice"}
	if got := dig.DigOr(data, "unknown", "name"); got != "alice" {
		t.Errorf("got %q, want alice", got)
	}
}

func TestDigOr_MissingKey(t *testing.T) {
	data := map[string]any{"name": "alice"}
	if got := dig.DigOr(data, "default", "missing"); got != "default" {
		t.Errorf("got %q, want default", got)
	}
}

func TestDigOr_TypeMismatch(t *testing.T) {
	data := map[string]any{"age": 30}
	if got := dig.DigOr(data, "unknown", "age"); got != "unknown" {
		t.Errorf("got %q, want unknown", got)
	}
}

func TestDigOr_NilData(t *testing.T) {
	if got := dig.DigOr[int](nil, 42, "any"); got != 42 {
		t.Errorf("got %d, want 42", got)
	}
}

func TestDigOr_DeepPath(t *testing.T) {
	data := map[string]any{
		"config": map[string]any{
			"timeout": 30,
		},
	}
	if got := dig.DigOr(data, 60, "config", "timeout"); got != 30 {
		t.Errorf("got %d, want 30", got)
	}
	if got := dig.DigOr(data, 60, "config", "missing"); got != 60 {
		t.Errorf("got %d, want 60", got)
	}
}

// =============================================================================
// Has
// =============================================================================

func TestHas_ShallowHit(t *testing.T) {
	data := map[string]any{"name": "alice"}
	if !dig.Has(data, "name") {
		t.Error("expected Has to return true")
	}
}

func TestHas_DeepHit(t *testing.T) {
	data := map[string]any{
		"user": map[string]any{"email": "alice@example.com"},
	}
	if !dig.Has(data, "user", "email") {
		t.Error("expected Has to return true for deep path")
	}
}

func TestHas_Miss(t *testing.T) {
	data := map[string]any{"name": "alice"}
	if dig.Has(data, "missing") {
		t.Error("expected Has to return false for missing key")
	}
}

func TestHas_NilData(t *testing.T) {
	if dig.Has(nil, "any") {
		t.Error("expected Has to return false for nil data")
	}
}

func TestHas_LeafTypeDoesNotMatter(t *testing.T) {
	// Has returns true regardless of leaf type — it only checks existence.
	data := map[string]any{
		"intVal":  42,
		"strVal":  "text",
		"boolVal": true,
		"nilVal":  nil,
	}
	for _, key := range []string{"intVal", "strVal", "boolVal", "nilVal"} {
		if !dig.Has(data, key) {
			t.Errorf("expected Has(%q) to return true", key)
		}
	}
}

// =============================================================================
// At
// =============================================================================

func TestAt_ReturnsRawValue(t *testing.T) {
	data := map[string]any{"count": 42}
	val, ok := dig.At(data, "count")
	if !ok || val != 42 {
		t.Errorf("got (%v, %v), want (42, true)", val, ok)
	}
}

func TestAt_Miss(t *testing.T) {
	data := map[string]any{"name": "alice"}
	if _, ok := dig.At(data, "missing"); ok {
		t.Error("expected ok=false for missing key")
	}
}

func TestAt_NilData(t *testing.T) {
	if _, ok := dig.At(nil, "any"); ok {
		t.Error("expected ok=false for nil data")
	}
}

func TestAt_InspectRuntimeType(t *testing.T) {
	// At lets callers inspect a value's type at runtime
	data := map[string]any{
		"value": map[string]any{"nested": true},
	}
	val, ok := dig.At(data, "value")
	if !ok {
		t.Fatal("expected ok=true")
	}
	if _, isMap := val.(map[string]any); !isMap {
		t.Error("expected At to return a map[string]any")
	}
}

func TestAt_SliceValue(t *testing.T) {
	data := map[string]any{
		"tags": []any{"go", "web"},
	}
	val, ok := dig.At(data, "tags")
	if !ok {
		t.Fatal("expected ok=true")
	}
	tags, isSlice := val.([]any)
	if !isSlice {
		t.Fatal("expected At to return a []any")
	}
	if len(tags) != 2 {
		t.Errorf("got %d tags, want 2", len(tags))
	}
}

// =============================================================================
// Integration test (mirrors the real-world pain pattern dig replaces)
// =============================================================================

func TestIntegration_ThreeLevelNestedPattern(t *testing.T) {
	// Mirrors the three-level nested type assertion pattern common in
	// codebases dealing with JSON responses:
	//
	//     saveRecs["customer"].(map[string]any)["address"].(map[string]any)["tvzr_key"].(string)
	saveRecs := map[string]any{
		"customer": map[string]any{
			"name": "Acme Corp",
			"address": map[string]any{
				"street":   "123 Main St",
				"city":     "Springfield",
				"tvzr_key": "01HQZX3T7K9W2B4N5F8G6P1M0S",
			},
		},
	}
	addressKey, ok := dig.Dig[string](saveRecs, "customer", "address", "tvzr_key")
	if !ok {
		t.Fatal("expected ok=true")
	}
	if addressKey != "01HQZX3T7K9W2B4N5F8G6P1M0S" {
		t.Errorf("got %q, want 01HQZX3T7K9W2B4N5F8G6P1M0S", addressKey)
	}
}
