package gpa_test

import (
	"testing"

	"example.com/Learning-Go/gpa"
)

// my basic test function whether a student is yodyci or not
func TestMyGPA(t *testing.T) {
	if gpa.TellMeMyGPA() < 3.5 {
		t.Fatal("Yodyci abim merhaba")
	}
}
