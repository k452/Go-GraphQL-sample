// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type NewTask struct {
	Title string `json:"title"`
	Note  string `json:"note"`
}

type Task struct {
	ID        string `json:"id"`
	Title     string `json:"title"`
	Note      string `json:"note"`
	Completed int    `json:"completed"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
