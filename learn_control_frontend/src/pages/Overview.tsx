import React from 'react';
import { useNavigate } from 'react-router-dom';
import Filters from '../components/Filters';

const Overview: React.FC = () => {
  const navigate = useNavigate();

  const handleSearch = (filters: any) => {
    console.log('Фильтры:', filters);
    // Здесь можно сделать запрос на бэкенд
  };

  return (
    <div className="container py-4">
      <button className="btn btn-secondary mb-3" onClick={() => navigate('/')}>⬅ Назад</button>
      <h2 className="mb-3">Обзор сотрудников</h2>
      <Filters onSearch={handleSearch} />
    </div>
  );
};

export default Overview;
