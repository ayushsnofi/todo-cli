package storage

import (
	"encoding/csv"
	"os"
	"todo-cli/internal/model"
)

type CSVStore struct {
	FilePath string
}

func (s *CSVStore) ReadTasks() ([]model.Task, error) {
	file, err := os.OpenFile(s.FilePath, os.O_CREATE|os.O_RDWR, 0644)

	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)

}
