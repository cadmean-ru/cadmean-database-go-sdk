package database

type Query struct {
	Path      string      `json:"path"`
	Filters   []Filter    `json:"filters"`
	QueryType string      `json:"query_type"`
	Data 	  interface{} `json:"data"`
}
