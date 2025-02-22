package services

import (
	"errors"
	"go-todo/models"
	"time"
)

var tasksMap = map[int]models.Task{}
var taskIDCounter int = 1 

func CreateTask(title string, deadline time.Time) (models.Task, error) {
	currentTime := time.Now()

	if deadline.Before(currentTime) {
		return models.Task{}, errors.New("invalid value of deadline: must be a future time")
	}

	if title == "" {
		return models.Task{}, errors.New("title cannot be empty")
	}

	newTask := models.Task{
		ID:       taskIDCounter, 
		Title:    title,
		Deadline: deadline,
		Status:   models.Pending,
	}
	taskIDCounter++;

	tasksMap[newTask.ID] = newTask;

	return newTask, nil
}

func CancelTask(taskID int) (models.Task, error) {
	task, exists := tasksMap[taskID]
	if !exists {
		return models.Task{}, errors.New("task does not exist")
	}

	task.Status = models.Cancelled
	tasksMap[taskID] = task
	return task, nil
}

func MarkTaskCompleted(taskID int) (models.Task, error) {
	task, exists := tasksMap[taskID]
	if !exists {
		return models.Task{}, errors.New("task does not exist")
	}

	if task.Status != models.Pending {
		return models.Task{}, errors.New("task must be in pending state")
	}
	task.Status = models.Completed
	tasksMap[taskID] = task;
	return task, nil
}

func GetPendingTasks() ([]models.Task, error) {
	var taskList []models.Task

	for _, task := range tasksMap {
		if task.Status == models.Pending {
			taskList = append(taskList, task)
		}
	}

	if len(taskList) == 0 {
		return nil, errors.New("no pending tasks found")
	}

	return taskList, nil
}

