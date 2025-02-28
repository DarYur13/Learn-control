import React from 'react';
import { Link } from 'react-router-dom';

const Menu: React.FC = () => {
  return (
    <div className="container py-4">
      <h2 className="mb-4">Меню</h2>
      <div className="d-grid gap-3">
        <Link to="/overview" className="btn btn-primary">Обзор</Link>
        <Link to="/personal-card" className="btn btn-secondary">Личные карты</Link>
      </div>
    </div>
  );
};

export default Menu;
