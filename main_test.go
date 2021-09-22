package main

import "testing"

//TestConnection tests for database connection made in the function
func TestConnection(t *testing.T) {

	if _, got := connection(); got != nil {
		t.Error("Expected  nil ", "Got", got)
	}
}
