package models

type Article struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Desc    string `json:"description"`
	Content string `json:"content"`
}	