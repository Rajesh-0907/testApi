package models

type Question struct {
	ID      int      `json:"id"`
	Question string   `json:"question"`
	Options  []string `json:"options"`
	InputID []string `json:"inputId"`
}