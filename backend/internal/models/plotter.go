package models

import (
	"encoding/json"
	"sort"
)

type RechartsPlotter struct {
	r Result
}

func NewRechartsPlotter() *RechartsPlotter {
	return &RechartsPlotter{}
}

func (rp *RechartsPlotter) Populate(r Result) {
	rp.r = r
}

func (rp *RechartsPlotter) Get() Result {
	return rp.r
}

func (rp *RechartsPlotter) convert() []Data {
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
		if d[i].Val != d[j].Val {
			return d[i].Val > d[j].Val
		}

		return d[i].Key < d[j].Key
	})

	return d
}

func (rp *RechartsPlotter) MarshalJSON() ([]byte, error) {
	return json.Marshal(rp.convert())
}
