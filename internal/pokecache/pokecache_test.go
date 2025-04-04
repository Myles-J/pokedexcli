package pokecache

import (
	"testing"
	"time"
)

func TestAdd(t *testing.T) {
	cache := NewCache(10 * time.Millisecond)

	cache.Add("key", []byte("value"))

	val, ok := cache.Get("key")
	if !ok {
		t.Error("expected key to be in cache")
	}

	if string(val) != "value" {
		t.Errorf("expected value to be 'value', got %s", string(val))
	}
}

func TestGet(t *testing.T) {
	cache := NewCache(10 * time.Millisecond)

	cache.Add("key", []byte("value"))

	val, ok := cache.Get("key")
	if !ok {
		t.Error("expected key to be in cache")
	}

	if string(val) != "value" {
		t.Errorf("expected value to be 'value', got %s", string(val))
	}
}

func TestReapLoop(t *testing.T) {
	cache := NewCache(5 * time.Millisecond)

	// Add multiple entries
	cache.Add("key1", []byte("value1"))
	cache.Add("key2", []byte("value2"))
	cache.Add("key3", []byte("value3"))

	// Verify entries exist initially
	if _, ok := cache.Get("key1"); !ok {
		t.Error("expected key1 to be in cache initially")
	}
	if _, ok := cache.Get("key2"); !ok {
		t.Error("expected key2 to be in cache initially")
	}
	if _, ok := cache.Get("key3"); !ok {
		t.Error("expected key3 to be in cache initially")
	}

	// Wait for entries to expire
	time.Sleep(10 * time.Millisecond)

	// Verify all entries are removed
	if _, ok := cache.Get("key1"); ok {
		t.Error("expected key1 to be expired")
	}
	if _, ok := cache.Get("key2"); ok {
		t.Error("expected key2 to be expired")
	}
	if _, ok := cache.Get("key3"); ok {
		t.Error("expected key3 to be expired")
	}
}
