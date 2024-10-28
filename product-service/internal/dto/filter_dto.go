package dto

type Filter struct {
	Field string      `json:"field"`
	Value interface{} `json:"value"`
}

type FilterRequest struct {
	Filters []Filter `json:"filters"`
}
