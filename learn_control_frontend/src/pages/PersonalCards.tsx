import React, { useState } from 'react';
import { useNavigate } from 'react-router-dom';
import EmployeeSearch from '../components/EmployeeSearch';
import EmployeeDetails from '../components/EmployeeDetails';

const PersonalCards: React.FC = () => {
  const [selectedId, setSelectedId] = useState<number | null>(null);
  const navigate = useNavigate();

  return (
    <div className="container py-4">
      <button className="btn btn-secondary mb-3" onClick={() => navigate('/')}>⬅ Назад</button>
      <h2 className="mb-4">Личные карты сотрудников</h2>
      <EmployeeSearch onSelect={setSelectedId} />
      {selectedId && <EmployeeDetails id={selectedId} />}
    </div>
  );
};

export default PersonalCards;
