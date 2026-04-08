package service

import (
	"time"
	"todo-cli/internal/model"
	"todo-cli/internal/storage"
)

type TaskService struct {
	Store *storage.CSVStore
}

var defaultService = &TaskService{
	Store: &storage.CSVStore{
		FilePath: "tasks.csv",
	},
}

func AddTask(title string) error {
	return defaultService.AddTask(title)
}

func ListTasks() ([]model.Task, error) {
	return defaultService.ListTasks()
}

func (s *TaskService) AddTask(title string) error {
	tasks, _:= s.Store.ReadTasks()

	id := len(tasks)+1

	task:= model.Task{
		ID:id,
		Title:title,
		CreatedAt:time.Now().Format(time.RFC3339),
	}

	tasks=append(tasks, task)
	return s.Store.WriteTasks(tasks)

}


func (s *TaskService) ListTasks() ([]model.Task,error) {
	return  s.Store.ReadTasks()
}

func(s *TaskService) CompleteTask(id int) error {
	tasks,_ := s.Store.ReadTasks()

	for i := range tasks {
		if tasks[i].ID==id{
			tasks[i].Completed=true
			tasks[i].CompletedAt=time.Now().Format(time.RFC3339)
		}
	}
	return s.Store.WriteTasks(tasks)
} 

func (s *TaskService) DeleteTask(id int) error {
	tasks,_:=s.Store.ReadTasks()

	var updatedTasks []model.Task
	for _,t :=range tasks{
		if t.ID != id {
			updatedTasks=append(updatedTasks,t)
		}
	}
	return s.Store.WriteTasks(updatedTasks)
}