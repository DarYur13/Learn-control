package domain

import "database/sql"

const (
	TaskTypeProvide = "PROVIDE"
	TaskTypeAssign  = "ASSIGN"
	TaskTypeChoose  = "CHOOSE"
	TaskTypeSet     = "SET"
	TaskTypeConfirm = "CONFIRM"
)

type TaskBaseInfo struct {
	Type       string
	TrainingID sql.NullInt64
	EmployeeID sql.NullInt64
	ExecutorID sql.NullInt64
	PositionID sql.NullInt64
}

type Task struct {
	ID          int
	Type        string
	Description string
	Employee    sql.NullString
	Training    sql.NullString
	Position    sql.NullString
	Department  sql.NullString
	Executor    sql.NullString
	Done        bool
}
