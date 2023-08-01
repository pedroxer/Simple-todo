// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0
// source: tasksFromList.sql

package sqlc

import (
	"context"
)

const addTaskToList = `-- name: AddTaskToList :one
INSERT INTO "tasks_to_list" (task_id, list_id)
VALUES ($1, $2) RETURNING task_id, list_id
`

type AddTaskToListParams struct {
	TaskID int32 `json:"task_id"`
	ListID int32 `json:"list_id"`
}

func (q *Queries) AddTaskToList(ctx context.Context, arg AddTaskToListParams) (TasksToList, error) {
	row := q.db.QueryRowContext(ctx, addTaskToList, arg.TaskID, arg.ListID)
	var i TasksToList
	err := row.Scan(&i.TaskID, &i.ListID)
	return i, err
}

const changeListForTask = `-- name: ChangeListForTask :exec
UPDATE "tasks_to_list" SET list_id = $1 
where task_id = $2
`

type ChangeListForTaskParams struct {
	ListID int32 `json:"list_id"`
	TaskID int32 `json:"task_id"`
}

func (q *Queries) ChangeListForTask(ctx context.Context, arg ChangeListForTaskParams) error {
	_, err := q.db.ExecContext(ctx, changeListForTask, arg.ListID, arg.TaskID)
	return err
}

const deleteTaskFromList = `-- name: DeleteTaskFromList :exec
DELETE FROM "tasks_to_list" where task_id = $1
`

func (q *Queries) DeleteTaskFromList(ctx context.Context, taskID int32) error {
	_, err := q.db.ExecContext(ctx, deleteTaskFromList, taskID)
	return err
}