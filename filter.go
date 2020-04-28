package database

type Filter struct {
	Path     string      `json:"path"`
	Operator string      `json:"operator"`
	Value    interface{} `json:"value"`
}

func NewFilter(path string, operator string, value interface{}) Filter {
	return Filter{
		Path:     path,
		Operator: operator,
		Value:    value,
	}
}