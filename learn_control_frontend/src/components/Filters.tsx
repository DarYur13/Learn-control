import React, { useEffect, useState } from 'react';
import { getFilters } from '../api/learnControlApi';

interface FiltersProps {
  onSearch: (filters: any) => void;
}

const Filters: React.FC<FiltersProps> = ({ onSearch }) => {
  const [filters, setFilters] = useState<{ departments: string[], positions: string[], trainings: { id: number, name: string }[] }>({
    departments: [],
    positions: [],
    trainings: [],
  });

  const [activeFilters, setActiveFilters] = useState<string[]>([]);
  const [selectedDepartment, setSelectedDepartment] = useState<string | null>(null);
  const [selectedPosition, setSelectedPosition] = useState<string | null>(null);
  const [selectedTraining, setSelectedTraining] = useState<number | null>(null);
  const [onlyNotPassed, setOnlyNotPassed] = useState<boolean>(false);
  const [trainingStartDate, setTrainingStartDate] = useState<string | null>(null);
  const [trainingEndDate, setTrainingEndDate] = useState<string | null>(null);
  const [daysToRetraining, setDaysToRetraining] = useState<number>(30);

  useEffect(() => {
    getFilters()
      .then(data => setFilters(data))
      .catch(err => console.error('Ошибка загрузки фильтров:', err));
  }, []);

  const handleSearch = () => {
    const queryParams = {
      department: activeFilters.includes('department') ? selectedDepartment || undefined : undefined,
      position: activeFilters.includes('position') ? selectedPosition || undefined : undefined,
      trainingID: activeFilters.includes('training') ? selectedTraining || undefined : undefined,
      onlyNotPassed: activeFilters.includes('notPassed') ? onlyNotPassed || undefined : undefined,
      trainingStartDate: activeFilters.includes('trainingPeriod') ? trainingStartDate || undefined : undefined,
      trainingEndDate: activeFilters.includes('trainingPeriod') ? trainingEndDate || undefined : undefined,
      retrainingIn: activeFilters.includes('retraining') ? daysToRetraining || undefined : undefined,
    };

    onSearch(queryParams);
  };

  // Добавление фильтра
  const addFilter = (filter: string) => {
    if (!activeFilters.includes(filter)) {
      setActiveFilters([...activeFilters, filter]);
    }
  };

  // Удаление фильтра
  const removeFilter = (filter: string) => {
    setActiveFilters(activeFilters.filter(f => f !== filter));
  };

  // Доступные фильтры (те, которые еще не выбраны)
  const availableFilters = [
    { key: 'department', label: 'Отдел' },
    { key: 'position', label: 'Должность' },
    { key: 'training', label: 'Обучение' },
    { key: 'notPassed', label: 'Не прошел обучение' },
    { key: 'trainingPeriod', label: 'Прошел обучение в период' },
    { key: 'retraining', label: 'Дней до переобучения' }
  ].filter(f => !activeFilters.includes(f.key));

  return (
    <div className="container">
      {/* 🔹 Кнопка добавления фильтров */}
      <div className="text-center mb-3">
        <div className="dropdown">
          <button
            className="btn btn-outline-primary dropdown-toggle"
            type="button"
            data-bs-toggle="dropdown"
          >
            ➕ Добавить фильтр
          </button>
          <ul className="dropdown-menu">
            {availableFilters.map(f => (
              <li key={f.key}>
                <button className="dropdown-item" onClick={() => addFilter(f.key)}>
                  {f.label}
                </button>
              </li>
            ))}
          </ul>
        </div>
      </div>

      {/* 🔹 Фильтры */}
      <div className="row g-4 align-items-end">
        <div className="col-md-6">
          {activeFilters.includes('department') && (
            <div className="mb-3">
              <label className="form-label">Отдел</label>
              <select className="form-select" onChange={(e) => setSelectedDepartment(e.target.value || null)}>
                <option value="">Выберите отдел</option>
                {filters.departments.map((dep) => (
                  <option key={dep} value={dep}>{dep}</option>
                ))}
              </select>
              <button className="btn btn-sm btn-outline-danger mt-1" onClick={() => removeFilter('department')}>❌ Удалить</button>
            </div>
          )}

          {activeFilters.includes('position') && (
            <div className="mb-3">
              <label className="form-label">Должность</label>
              <select className="form-select" onChange={(e) => setSelectedPosition(e.target.value || null)}>
                <option value="">Выберите должность</option>
                {filters.positions.map((pos) => (
                  <option key={pos} value={pos}>{pos}</option>
                ))}
              </select>
              <button className="btn btn-sm btn-outline-danger mt-1" onClick={() => removeFilter('position')}>❌ Удалить</button>
            </div>
          )}
        </div>

        <div className="col-md-6">
          {activeFilters.includes('notPassed') && (
            <div className="mb-3 d-flex align-items-center">
              <input type="checkbox" className="form-check-input me-2" checked={onlyNotPassed} onChange={() => setOnlyNotPassed(!onlyNotPassed)} />
              <label className="form-check-label">Не прошел обучение</label>
              <button className="btn btn-sm btn-outline-danger ms-2" onClick={() => removeFilter('notPassed')}>❌</button>
            </div>
          )}
        </div>
      </div>

      {/* 🔹 Кнопка поиска */}
      <div className="col-12 text-center mt-3">
        <button className="btn btn-primary" onClick={handleSearch}>🔍 Поиск</button>
      </div>
    </div>
  );
};

export default Filters;
