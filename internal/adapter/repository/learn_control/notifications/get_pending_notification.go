package learncontrol

import (
	"context"

	"github.com/DarYur13/learn-control/internal/domain"
)

const queryGetPendingNotifications = `
	SELECT 
		nq.id,
		nq.notification_type,
		nq.employee_id,
		nq.training_id,
		ntt.subject_template,
		ntt.body_template,
		i.full_name AS instructor_name,
		i.email AS instructor_email,
		i.position AS instructor_position,
		e.full_name AS employee_name,
		e.birth_date,
		e.position AS employee_position,
		e.department,
		t.training_type,
		string_agg(la.act_name, E'\n') AS acts,
		et.retraining_date
	FROM notifications_queue nq
	JOIN notification_types_templates ntt ON ntt.notification_type = nq.notification_type
	JOIN employees e ON e.id = nq.employee_id
	JOIN employees i ON i.department = e.department AND i.is_leader = TRUE
	JOIN trainings t ON t.id = nq.training_id
	LEFT JOIN acts_trainings at ON at.training_id = nq.training_id
	LEFT JOIN local_acts la ON la.id = at.local_act_id
	LEFT JOIN employee_trainings et ON et.employee_id = nq.employee_id AND et.training_id = nq.training_id
	WHERE nq.is_sent = FALSE
	AND nq.created_at <= NOW()
	GROUP BY 
		nq.id,
		nq.notification_type,
		nq.employee_id,
		nq.training_id,
		ntt.subject_template,
		ntt.body_template,
		i.full_name,
		i.email,
		i.position,
		e.full_name,
		e.birth_date,
		e.position,
		e.department,
		t.training_type,
		et.retraining_date
	ORDER BY nq.created_at ASC
	LIMIT 5;
	`

func (ns *NotificationsStorage) GetPendingNotifications(ctx context.Context) ([]domain.PendingNotification, error) {
	rows, err := ns.db.QueryContext(ctx, queryGetPendingNotifications)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []domain.PendingNotification
	for rows.Next() {
		var n domain.PendingNotification
		if err := rows.Scan(
			&n.ID,
			&n.Type,
			&n.EmployeeID,
			&n.TrainingID,
			&n.Subject,
			&n.Body,
			&n.InstructorName,
			&n.InstructorEmail,
			&n.InstructorPosition,
			&n.EmployeeName,
			&n.EmployeeBirthDate,
			&n.EmployeePosition,
			&n.EmployeeDepartment,
			&n.TrainingType,
			&n.Acts,
			&n.ReTrainingDate,
		); err != nil {
			return nil, err
		}

		result = append(result, n)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return result, nil
}
