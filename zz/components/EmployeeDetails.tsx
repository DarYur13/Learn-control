import { useEffect, useState } from "react";
import { getEmployee } from "../api/api";

// Определи типы для данных
interface Training {
  name: string;
  date: string;
  nextdate: string;
}

interface Employee {
  fullname: string;
  birthdate: string;
  department: string;
  position: string;
  snils: string;
  trainings: Training[];
}

const EmployeeDetails = ({ id }: { id: number }) => {
  const [employee, setEmployee] = useState<Employee | null>(null);

  useEffect(() => {
    if (id) {
      getEmployee(id).then(setEmployee);
    }
  }, [id]);

  if (!employee) return <div>Выберите сотрудника</div>;

  return (
    <div>
      <h2>{employee.fullname}</h2>
      <p>Дата рождения: {employee.birthdate}</p>
      <p>Отдел: {employee.department}</p>
      <p>Должность: {employee.position}</p>
      <p>СНИЛС: {employee.snils}</p>
      <h3>Обучения:</h3>
      <ul>
        {employee.trainings.map((t) => (
          <li key={t.name}>
            {t.name}: {t.date} (повтор: {t.nextdate})
          </li>
        ))}
      </ul>
    </div>
  );
};

export default EmployeeDetails;
