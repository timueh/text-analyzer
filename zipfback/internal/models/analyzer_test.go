package models

import (
	"reflect"
	"testing"
)

func TestNewLetters(t *testing.T) {
	cases := []struct {
		desc string
		val  string
	}{
		{"no spaces, all lowercase", "helloworld"},
		{"no spaces, all uppercase", "HELLOWORLD"},
		{"no spaces, camel case", "HelloWorld"},
		{"spaces, camel case", "Hello World"},
		{"weird spaces, camel case", " He llo W orl    d"},
		{"punctuation, camel case", " He llo--'! W.orl    d!"},
		{"numbers, camel case", " He123ll93o--'! W.orl    d!"},
	}

	want := Result{
		"h": 1,
		"e": 1,
		"l": 3,
		"o": 2,
		"w": 1,
		"r": 1,
		"d": 1,
	}

	aa := []Analyzer{NewLetterFrequency()}

	for _, a := range aa {
		for _, c := range cases {
			t.Run(c.desc, func(t *testing.T) {
				got := a.Run(c.val)
				if !reflect.DeepEqual(got, want) {
					t.Errorf("got %v, want %v", got, want)
				}
			})
		}
	}
}

func TestWordFrequency(t *testing.T) {
	cases := []struct {
		desc string
		val  string
	}{
		{"happy path", "hello world"},
		{"with punctuation", "!hello, world?"},
	}

	want := Result{"hello": 1, "world": 1}
	for _, a := range []Analyzer{NewWordFrequency()} {
		for _, c := range cases {
			t.Run(c.desc, func(t *testing.T) {
				got := a.Run(c.val)
				if !reflect.DeepEqual(got, want) {
					t.Errorf("got %v, want %v", got, want)
				}
			})
		}
	}
}

func assertValue[K comparable](got, want K, t *testing.T) {
	t.Helper()
	if got != want {
		t.Fatalf("got %v, want %v", got, want)
	}
}
