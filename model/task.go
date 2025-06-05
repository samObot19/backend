package Models

import (
	"time"
)

type Task struct{
	ID			string		`json: id`
	Title			string 		`json: title`
	Description		string		`json: description`
	DueDate			time.Time	`json: duedate`
	Completed			bool		`json: completed`
	Status          string     `json: status`

}
