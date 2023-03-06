package goautoneg

import (
	"math"
	"testing"
)

var chrome = "application/xml,application/xhtml+xml,text/html;q=0.9,text/plain;q=0.8,image/png,*/*;q=0.5"

func TestNegotiate(t *testing.T) {
	alternatives := []string{"text/html", "image/png"}
	contentType := Negotiate(chrome, alternatives)
	if contentType != "image/png" {
		t.Errorf("got %s expected image/png", contentType)
	}

	alternatives = []string{"text/html", "text/plain", "text/n3"}
	contentType = Negotiate(chrome, alternatives)
	if contentType != "text/html" {
		t.Errorf("got %s expected text/html", contentType)
	}

	alternatives = []string{"text/n3", "text/plain"}
	contentType = Negotiate(chrome, alternatives)
	if contentType != "text/plain" {
		t.Errorf("got %s expected text/plain", contentType)
	}

	alternatives = []string{"text/n3", "application/rdf+xml"}
	contentType = Negotiate(chrome, alternatives)
	if contentType != "text/n3" {
		t.Errorf("got %s expected text/n3", contentType)
	}
}

func TestParseAccept(t *testing.T) {
	actual := ParseAccept("application/xml,application/xhtml+xml,text/html;q=0.9,text/plain;q=0.8;otherParam=blah,image/png,*/*;q=0.5")
	expected := []Accept{
		{
			Type:    "application",
			SubType: "xml",
			Q:       1,
		},
		{
			Type:    "application",
			SubType: "xhtml+xml",
			Q:       1,
		},
		{
			Type:    "image",
			SubType: "png",
			Q:       1,
		},
		{
			Type:    "text",
			SubType: "html",
			Q:       0.9,
		},
		{
			Type:    "text",
			SubType: "plain",
			Q:       0.8,
		},
		{
			Type:    "*",
			SubType: "*",
			Q:       0.5,
		},
	}

	if len(actual) != len(expected) {
		t.Fatalf("expected %d entries, but got %d in %v", len(expected), len(actual), actual)
	}

	for i, expectedEntry := range expected {
		actualEntry := actual[i]

		qDiff := math.Abs(actualEntry.Q - expectedEntry.Q)

		if actualEntry.Type != expectedEntry.Type || actualEntry.SubType != expectedEntry.SubType || qDiff > 0.0001 {
			t.Fatalf("expected: %v\nactual: %v\nat position %d in %v", expectedEntry, actualEntry, i, actual)
		}
	}
}

func BenchmarkParseAccept(b *testing.B) {
	scenarios := []string{
		"",
		"application/json",
		"application/json,text/plain",
		"application/json;q=0.9,text/plain",
		chrome,
	}

	for _, scenario := range scenarios {
		b.Run(scenario, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = ParseAccept(scenario)
			}
		})
	}
}
