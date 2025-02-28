import React, { useState, useEffect } from 'react';
import { getEmployeesByName } from '../api/learnControlApi';

interface Employee {
  id: number;
  fullname: string;
  birthdate: string;
}

interface Props {
  onSelect: (id: number) => void;
}

const EmployeeSearch: React.FC<Props> = ({ onSelect }) => {
  const [query, setQuery] = useState('');
  const [results, setResults] = useState<Employee[]>([]);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    if (query.length < 2) {
      setResults([]);
      return;
    }

    const fetchEmployees = async () => {
      try {
        const employees = await getEmployeesByName(query);
        setResults(employees || []);
        setError(null);
      } catch (err) {
        console.error('Ошибка при загрузке сотрудников:', err);
        setError('Не удалось загрузить сотрудников');
        setResults([]);
      }
    };

    fetchEmployees();
  }, [query]);

  return (
    <div className="mb-4">
      <input
        className="form-control"
        placeholder="Введите ФИО сотрудника..."
        value={query}
        onChange={(e) => setQuery(e.target.value)}
      />
      {error && <p className="text-danger mt-2">{error}</p>}
      {results.length > 0 && (
        <ul className="list-group mt-2">
          {results.map((emp) => (
            <li
              key={emp.id}
              className="list-group-item list-group-item-action"
              onClick={() => onSelect(emp.id)}
              style={{ cursor: 'pointer' }}
            >
              {emp.fullname} ({emp.birthdate})
            </li>
          ))}
        </ul>
      )}
    </div>
  );
};

export default EmployeeSearch;
