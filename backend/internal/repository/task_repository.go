package repository

import (
	"database/sql"

	"granet-task-manager/backend/internal/model"
)

type TaskRepository struct {
	db *sql.DB
}

func NewTaskRepository(db *sql.DB) *TaskRepository {
	return &TaskRepository{db: db}
}

func (r *TaskRepository) Create(task *model.Task) error {
	query := `
		INSERT INTO tasks (title, description, due_date, status)
		VALUES (?, ?, ?, ?)
	`

	result, err := r.db.Exec(
		query,
		task.Title,
		task.Description,
		task.DueDate,
		task.Status,
	)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	task.ID = id
	return nil
}

func (r *TaskRepository) GetAll() ([]model.Task, error) {
	query := `
		SELECT id, title, description, due_date, status
		FROM tasks
		ORDER BY id DESC
	`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	tasks := make([]model.Task, 0)

	for rows.Next() {
		var task model.Task

		err := rows.Scan(
			&task.ID,
			&task.Title,
			&task.Description,
			&task.DueDate,
			&task.Status,
		)
		if err != nil {
			return nil, err
		}

		tasks = append(tasks, task)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return tasks, nil
}

func (r *TaskRepository) Update(task *model.Task) error {
	query := `
		UPDATE tasks
		SET title = ?, description = ?, due_date = ?, status = ?
		WHERE id = ?
	`

	_, err := r.db.Exec(
		query,
		task.Title,
		task.Description,
		task.DueDate,
		task.Status,
		task.ID,
	)

	return err
}

func (r *TaskRepository) Delete(id int64) error {
	query := `DELETE FROM tasks WHERE id = ?`

	_, err := r.db.Exec(query, id)
	return err
}
