package models

type User struct {
	ID int `json:"id"`
}

type Segment struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
