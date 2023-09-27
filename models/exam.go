package models

type Exam struct {
	Patient Patient `json:"patient"`
	Doctor  Doctor  `json:"doctor"`
	Token   string  `json:"token"`
	Date    string  `json:"date"`
	Type    string  `json:"type"`
	Limits  string  `json:"limits"`
	Results string  `json:"results"`
}
