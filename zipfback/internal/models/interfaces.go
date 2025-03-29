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
	Convert() []Data
	sort(d []Data) []Data
	json.Marshaler
}

type AnalyzerPlotter interface {
	Analyzer
	Plotter
}
