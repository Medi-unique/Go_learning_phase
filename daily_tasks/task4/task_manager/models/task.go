package models

type Task struct {
	ID       string `json:"id" validate:"required"`
	Name     string `json:"name" validate:"required min=3 max=20" `
	Detail   string `json:"detail"`
}
