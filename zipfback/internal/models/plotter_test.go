package models

import (
	"encoding/json"
	"testing"
)

func TestPlotter(t *testing.T) {
	cases := []struct {
		desc     string
		r        Result
		wantJson string
	}{
		{
			"letter frequency",
			Result{
				"h": 1,
				"e": 3,
			},
			`[{"name":"e","value":3},{"name": "h","value":1}]`,
		},
		{
			"word frequency",
			Result{
				"hello": 1,
				"world": 300,
			},
			`[{"name":"world","value":300},{"name": "hello","value":1}]`,
		},
	}

	for _, c := range cases {
		t.Run(c.desc, func(t *testing.T) {
			p := NewRechartsPlotter(c.r)

			data, err := json.Marshal(p)
			if err != nil {
				t.Fatal(err.Error())
			}
			var got []Data
			if err := json.Unmarshal(data, &got); err != nil {
				t.Fatal(err)
			}

			var want []Data
			if err := json.Unmarshal([]byte(c.wantJson), &want); err != nil {
				t.Fatal(err)
			}

			assertValue(len(want), len(c.r), t)
			assertValue(len(got), len(c.r), t)

			for i := 0; i < len(c.r); i++ {
				assertData(got[i], want[i], t)
			}
		})
	}

}

func assertData(got, want Data, t *testing.T) {
	t.Helper()
	assertValue(got.Key, want.Key, t)
	assertValue(got.Val, want.Val, t)
}
