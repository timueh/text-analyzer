package models

import (
	"encoding/json"
	"sort"
)

type RechartsPlotter struct {
	r Result
}

func NewRechartsPlotter(r Result) *RechartsPlotter {
	return &RechartsPlotter{r: r}
}

func (rp *RechartsPlotter) Convert() []Data {
	d := make([]Data, 0, len(rp.r))
	for key, val := range rp.r {
		d = append(d, Data{
			Key: key,
			Val: val,
		})
	}

	return rp.sort(d)
}

func (rp *RechartsPlotter) sort(d []Data) []Data {
	sort.Slice(d, func(i, j int) bool {
		return d[i].Val > d[j].Val
	})

	return d
}

func (rp *RechartsPlotter) MarshalJSON() ([]byte, error) {
	return json.Marshal(rp.Convert())
}
