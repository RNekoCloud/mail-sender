package main

import "testing"

func TestSendMail(t *testing.T) {
	err := Send()

	if err != nil {
		t.Error("Failed to sending mail!")
	}
}
