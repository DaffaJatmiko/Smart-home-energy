package model

type Inputs struct {
	Table map[string][]string `json:"table"`
	Query string              `json:"query"`
}

type InputQuery struct {
	Query string `json:"query"`
}
