package storage

import (
	"encoding/csv"
	"os"
	"strconv"
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
	records, _ := reader.ReadAll()

	var tasks []model.Task

	for _, r := range records {
		id, _ := strconv.Atoi(r[0])
		completed, _ := strconv.ParseBool(r[2])

		tasks = append(tasks, model.Task{
			ID:          id,
			Title:       r[1],
			Completed:   completed,
			CreatedAt:   r[3],
			CompletedAt: r[4],
		})
	}
	return tasks, nil
}

func (s *CSVStore) WriteTasks(tasks []model.Task) error {
	file, err := os.Create((s.FilePath))

	if err != nil {
		return err
	}

	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, t := range tasks {
		record := []string{
			strconv.Itoa(t.ID),
			t.Title,
			strconv.FormatBool(t.Completed),
			t.CreatedAt,
			t.CompletedAt,
		}
		writer.Write(record)
	}
	return nil
}

 