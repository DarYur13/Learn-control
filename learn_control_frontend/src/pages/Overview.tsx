import React, { useEffect, useState } from 'react';
import axios from 'axios';

interface FilterBaseInfo {
  id: number;
  name: string;
}

const Overview: React.FC = () => {
  const [filters, setFilters] = useState<{ departments: FilterBaseInfo[], positions: FilterBaseInfo[], trainings: FilterBaseInfo[] }>({
    departments: [],
    positions: [],
    trainings: [],
  });

  const [selectedDepartment, setSelectedDepartment] = useState<string | null>(null);
  const [selectedPosition, setSelectedPosition] = useState<string | null>(null);
  const [selectedTraining, setSelectedTraining] = useState<number | null>(null);
  const [onlyNotPassed, setOnlyNotPassed] = useState<boolean>(false);
  const [daysToRetraining, setDaysToRetraining] = useState<number>(30);

  useEffect(() => {
    axios.get('http://localhost:8080/employees/filters/get')
      .then(response => setFilters(response.data))
      .catch(error => console.error('Ошибка загрузки фильтров:', error));
  }, []);

  const handleSearch = () => {
    const queryParams = {
      department: selectedDepartment || undefined,
      position: selectedPosition || undefined,
      trainingID: selectedTraining || undefined,
      retrainingIn: daysToRetraining,
      trainigsPassed: onlyNotPassed,
    };

    axios.get('http://localhost:8080/employees/get/by_filters', { params: queryParams })
      .then(response => console.log('Результаты поиска:', response.data))
      .catch(error => console.error('Ошибка поиска:', error));
  };

  return (
    <div className="container py-4">
      <h2 className="mb-3">Обзор сотрудников</h2>

      <div className="row g-2">
        <div className="col-md">
          <select className="form-select" onChange={(e) => setSelectedDepartment(e.target.value || null)}>
            <option value="">Отдел</option>
            {filters.departments.map((dep) => (
              <option key={dep.id} value={dep.name}>{dep.name}</option>
            ))}
          </select>
        </div>

        <div className="col-md">
          <select className="form-select" onChange={(e) => setSelectedPosition(e.target.value || null)}>
            <option value="">Должность</option>
            {filters.positions.map((pos) => (
              <option key={pos.id} value={pos.name}>{pos.name}</option>
            ))}
          </select>
        </div>

        <div className="col-md">
          <select className="form-select" onChange={(e) => setSelectedTraining(Number(e.target.value) || null)}>
            <option value="">Обучение</option>
            {filters.trainings.map((train) => (
              <option key={train.id} value={train.id}>{train.name}</option>
            ))}
          </select>
        </div>

        <div className="col-md-auto d-flex align-items-center">
          <input type="checkbox" className="form-check-input me-2" id="notPassed" checked={onlyNotPassed} onChange={() => setOnlyNotPassed(!onlyNotPassed)} />
          <label className="form-check-label" htmlFor="notPassed">Без обучения</label>
        </div>
      </div>

      <div className="row mt-3 g-2">
        <div className="col-md d-flex align-items-center">
          <label className="me-2">Перепрохождение:</label>
          <input type="number" className="form-control w-auto" min="0" max="60" value={daysToRetraining} onChange={(e) => setDaysToRetraining(Number(e.target.value))} />
          <input type="range" className="form-range ms-2" min="0" max="60" value={daysToRetraining} onChange={(e) => setDaysToRetraining(Number(e.target.value))} />
        </div>

        <div className="col-md-auto">
          <button className="btn btn-primary" onClick={handleSearch}>🔍 Поиск</button>
        </div>
      </div>
    </div>
  );
};

export default Overview;
