package types

import (
	"github.com/renacargnelutti/golang-mongodb-api/models"
)

type TasksData struct {
	Success bool              `json:"success"`
	Tasks   []models.ToDoList `json:"tasks"`
}

type TaskData struct {
	Success bool            `json:"success"`
	Task    models.ToDoList `json:"task"`
}
