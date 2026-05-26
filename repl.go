package main

import (
	"strings"
)

func CleanInput(text string) []string {
    // Trim leading/trailing whitespace and convert to lowercase
	lowered := strings.ToLower(strings.TrimSpace(text))
	
	// Handle cases where the input was empty (avoid returning [""])
	if lowered == "" {
		return []string{}
	}
	
	// Split by whitespace
	return strings.Fields(lowered)
}

 