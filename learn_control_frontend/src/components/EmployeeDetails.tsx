import React, { useEffect, useState } from 'react';
import { getEmployeePersonalCard } from '../api/learnControlApi';

interface Training {
  name: string;
  passDate: string;
  rePassDate: string;
}

interface EmployeeDetailsProps {
  id: number;
}

const EmployeeDetails: React.FC<EmployeeDetailsProps> = ({ id }) => {
  const [employee, setEmployee] = useState<any>(null);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    const fetchEmployee = async () => {
      try {
        const data = await getEmployeePersonalCard(id);
        setEmployee(data || {});
        setError(null);
      } catch (err) {
        console.error('Ошибка загрузки данных сотрудника:', err);
        setError('Не удалось загрузить данные сотрудника.');
      }
    };
    fetchEmployee();
  }, [id]);

  if (error) return <p className="text-danger">{error}</p>;
  if (!employee) return <p>Загрузка...</p>;

  return (
    <div className="card mt-3">
      <div className="card-body">
        <h5 className="card-title">{employee.fullname || 'Неизвестно'}</h5>
        <p>Дата рождения: {employee.birthdate || 'Не указано'}</p>
        <p>Отдел: {employee.department || 'Не указано'}</p>
        <p>Должность: {employee.position || 'Не указано'}</p>
        <p>СНИЛС: {employee.snils || 'Не указано'}</p>

        <h6>Обучения:</h6>
        {employee.trainings && employee.trainings.length > 0 ? (
          <ul className="list-group">
            {employee.trainings.map((train: Training, index: number) => (
              <li key={index} className="list-group-item">
                <strong>{train.name}</strong><br />
                Пройдено: {train.passDate || '—'}<br />
                Перепрохождение: {train.rePassDate || '—'}
              </li>
            ))}
          </ul>
        ) : (
          <p>Нет данных по обучению</p>
        )}
      </div>
    </div>
  );
};

export default EmployeeDetails;
