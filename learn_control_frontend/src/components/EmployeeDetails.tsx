import React, { useEffect, useState } from 'react';
import { getEmployee } from '../api/learnControlApi';

interface Training {
  name: string;
  date: string;
  nextdate: string;
}

interface EmployeeDetailsProps {
  id: number;
}

const EmployeeDetails: React.FC<EmployeeDetailsProps> = ({ id }) => {
  const [employee, setEmployee] = useState<any>(null);

  useEffect(() => {
    const fetchEmployee = async () => {
      const data = await getEmployee(id);
      setEmployee(data);
    };
    fetchEmployee();
  }, [id]);

  if (!employee) return <p>Загрузка...</p>;

  return (
    <div className="card mt-3">
      <div className="card-body">
        <h5 className="card-title">{employee.fullname}</h5>
        <p>Дата рождения: {employee.birthdate}</p>
        <p>Отдел: {employee.department}</p>
        <p>Должность: {employee.position}</p>
        <p>СНИЛС: {employee.snils}</p>

        <h6>Обучения:</h6>
        <ul className="list-group">
          {employee.trainings.map((train: Training, index: number) => (
            <li key={index} className="list-group-item">
              <strong>{train.name}</strong><br />
              Пройдено: {train.date}<br />
              Перепрохождение: {train.nextdate}
            </li>
          ))}
        </ul>
      </div>
    </div>
  );
};

export default EmployeeDetails;