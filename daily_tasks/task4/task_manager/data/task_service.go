package data

import (
	"example.com/task_manager/models"
	"errors"
	"fmt"

	"github.com/go-playground/validator/v10"
)

var Data = map[string]models.Task{
	"1": {
		ID:       "1",
		Name:     "Task 1",
		Detail:   "Detail for Task 1",
	},
	"2": {
		ID:       "2",
		Name:     "Task 2",
		Detail:   "Detail for Task 2",
	},
	"3": {
		ID:       "3",
		Name:     "Task 3",
		Detail:   "Detail for Task 3",
	},
	"4": {
		ID:       "4",
		Name:     "Task 4",
		Detail:   "Detail for Task 4",
	},
}

func AddTask(task models.Task) (models.Task, error) {

	validator := validator.New()
	err := validator.Struct(task)
	if err != nil {
		return models.Task{}, errors.New("a task must have an id")
	}
	_, exists := Data[task.ID]
	if !exists {
		Data[task.ID] = task
		return task, nil
	}

	fmt.Println("no Data")
	return models.Task{}, errors.New("a task with this id exists")
}
func GetAllTask() map[string]models.Task {
	return Data
}
func GetTask(id string) (models.Task, error) {
	val, exists := Data[id]
	if !exists {
		fmt.Println("no Data")
		return models.Task{}, errors.New("no task with this ID")
	}
	return val, nil
}
func DeleteTask(id string) (models.Task, error) {
	if val, ok := Data[id]; ok {
		delete(Data, id)
		return val, nil
	}
	return models.Task{}, errors.New("NO such Task ID")
}
func UpdateTask(task models.Task, id string) (models.Task, error) {

	if val, ok := Data[id]; ok {
		if task.Detail != "" {
			val.Detail = task.Detail
		}
	
		if task.ID != "" {
			val.ID = task.ID
		}
		if task.Name != "" {
			val.Name = task.Name
		}
		
		Data[id] = val
		return Data[id], nil
	}
	return models.Task{}, errors.New("NO Data Such ID")
}
