package models

type Users struct {
	ID   int     `json:"id"`
	SID  string  `json:"sid"`
	Name string  `json:"name"`
	CGPA float32 `json:"cgpa"`
}
