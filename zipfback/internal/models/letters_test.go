package models

import (
	"encoding/json"
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

	want := Letters{
		'h': 1,
		'e': 1,
		'l': 3,
		'o': 2,
		'w': 1,
		'r': 1,
		'd': 1,
	}

	for _, c := range cases {
		t.Run(c.desc, func(t *testing.T) {
			got := NewLetters(c.val)
			if !reflect.DeepEqual(got, want) {
				t.Errorf("got %v, want %v", got, want)
			}
		})
	}

}

func TestMarshalling(t *testing.T) {
	l := Letters{
		'h': 1,
		'e': 3,
	}
	data, err := json.Marshal(l)
	if err != nil {
		t.Fatal(err.Error())
	}
	var got []LetterValue
	if err := json.Unmarshal(data, &got); err != nil {
		t.Fatal(err)
	}

	wantJson := `[{"name":"e","value":3},{"name": "h","value":1}]`
	var want []LetterValue
	if err := json.Unmarshal([]byte(wantJson), &want); err != nil {
		t.Fatal(err)
	}

	assertValue(len(want), len(l), t)
	assertValue(len(got), len(l), t)

	for i := 0; i < len(l); i++ {
		assertValue(want[i].Letter, got[i].Letter, t)
		assertValue(want[i].Value, got[i].Value, t)
	}

}

func assertValue[K comparable](got, want K, t *testing.T) {
	t.Helper()
	if got != want {
		t.Fatalf("got %v, want %v", got, want)
	}
}
