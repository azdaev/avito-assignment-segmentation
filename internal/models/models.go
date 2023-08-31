package models

type User struct {
	ID int `json:"id"`
}

type Segment struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type SetUserSegmentsResponse struct {
	AddError    []string `json:"add_error,omitempty"`
	RemoveError []string `json:"remove_error,omitempty"`
}
