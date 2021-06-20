package services

import (
	"database/sql"

	_ "github.com/lib/pq"
	"github.com/picolloo/todo-cli/models"
)

type TaskService struct {
	db *sql.DB
}

func NewTaskService(db *sql.DB) *TaskService {
	return &TaskService{
		db: db,
	}
}

func (s *TaskService) SaveTask(task *models.Task) error {
	tx, err := s.db.Begin()
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare("insert into tasks (description, status) values ($1, $2)")
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(&task.Description, &task.Status)
	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}
