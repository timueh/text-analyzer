package models

import "encoding/json"

type Result map[string]float64

type Data struct {
	Key string  `json:"name"`
	Val float64 `json:"value"`
}

type Analyzer interface {
	Run(s string) Result
}

type Plotter interface {
	Populate(r Result)
	json.Marshaler
}
