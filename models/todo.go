package models

type TodoInput struct {
    Todo string `json:"Todo"`
    Done bool   `json:"done"`
}
