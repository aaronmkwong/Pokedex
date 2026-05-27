package pokeapi

import (
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/aaronMkwong/Pokedex/internal/pokecache"
)

const baseURL = "https://pokeapi.co/api/v2"

// Client handles requests to the PokeAPI and caches responses
type Client struct {
	httpClient http.Client
	cache      *pokecache.Cache
}

// NewClient creates a new PokeAPI client
func NewClient(cache *pokecache.Cache) *Client {
	return &Client{
		httpClient: http.Client{
			Timeout: 10 * time.Second,
		},
		cache: cache,
	}
}

// Get fetches data from the PokeAPI.
// It checks the cache first before making an HTTP request.
func (c *Client) Get(endpoint string) ([]byte, error) {
	fullURL := baseURL + endpoint

	// Check cache first
	if cachedData, found := c.cache.Get(fullURL); found {
		return cachedData, nil
	}

	// Make HTTP request
	resp, err := c.httpClient.Get(fullURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Validate status code
	if resp.StatusCode > 299 {
		return nil, fmt.Errorf("request failed with status code: %d", resp.StatusCode)
	}

	// Read response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// Save to cache
	c.cache.Add(fullURL, body)

	return body, nil
}