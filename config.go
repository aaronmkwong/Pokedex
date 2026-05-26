package main

// config stores application state shared between commands
type config struct {
	Next     *string
	Previous *string
}