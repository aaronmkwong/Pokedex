package pokecache

import (
	"bytes"
	"testing"
	"time"
)

// Test that a cached value can be added and retrieved
func TestAddGet(t *testing.T) {
	cache := NewCache(5 * time.Second)

	key := "https://pokeapi.co/api/v2/location-area"
	val := []byte("hello")

	cache.Add(key, val)

	got, ok := cache.Get(key)
	if !ok {
		t.Fatal("expected key to exist in cache")
	}

	if !bytes.Equal(got, val) {
		t.Fatalf("expected %q, got %q", val, got)
	}
}

// Test that entries expire and are removed by the reap loop
func TestReapLoopRemovesExpiredEntries(t *testing.T) {
	interval := 50 * time.Millisecond
	cache := NewCache(interval)

	key := "test-key"
	val := []byte("test-value")

	cache.Add(key, val)

	// Confirm entry exists immediately
	got, ok := cache.Get(key)
	if !ok {
		t.Fatal("expected key to exist immediately after Add")
	}

	if !bytes.Equal(got, val) {
		t.Fatalf("expected %q, got %q", val, got)
	}

	// Wait long enough for:
	// 1. the entry to become expired
	// 2. the ticker to fire and reap it
	time.Sleep(150 * time.Millisecond)

	_, ok = cache.Get(key)
	if ok {
		t.Fatal("expected key to be removed after expiration")
	}
}