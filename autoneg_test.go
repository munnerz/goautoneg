package goautoneg

import (
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
