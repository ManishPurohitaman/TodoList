package models

import "time"

type TaskStatus string

const (
    Pending    TaskStatus = "pending"
    Completed  TaskStatus = "completed"
    Cancelled  TaskStatus = "cancelled"
)

type Task struct {
    ID       int     `json:"id"`
    Title    string     `json:"title"`
    Status   TaskStatus `json:"status"`
    Deadline time.Time  `json:"deadline"` 
}

type TaskRequest struct {
	Title    string `json:"title"`
	Deadline string `json:"deadline"`
}

type TaskIDRequest struct {
	TaskID int `json:"taskId"`
}