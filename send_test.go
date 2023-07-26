package main

import "testing"

func TestSendMail(t *testing.T) {
	err := Send()

	// Case: Should not returning error.
	if err != nil {
		t.Error("Failed to sending mail!")
	}
}
