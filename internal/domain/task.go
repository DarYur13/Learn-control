package domain

import "database/sql"

type TaskType string

const (
	TaskTypeAssign  TaskType = "ASSIGN"
	TaskTypeSet     TaskType = "SET"
	TaskTypeProvide TaskType = "PROVIDE"
	TaskTypeConfirm TaskType = "CONFIRM"
	TaskTypeControl TaskType = "CONTROL"
	TaskTypeChoose  TaskType = "CHOOSE"
)

type TaskBaseInfo struct {
	Type       TaskType
	TrainingID sql.NullInt64
	EmployeeID sql.NullInt64
	ExecutorID sql.NullInt64
	PositionID sql.NullInt64
}

type Task struct {
	ID          int
	Type        TaskType
	Description string
	Employee    sql.NullString
	Training    sql.NullString
	Position    sql.NullString
	Department  sql.NullString
	Executor    sql.NullString
	Done        bool
}
