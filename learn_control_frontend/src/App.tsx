import React, { useState } from 'react';
import EmployeeSearch from './components/EmployeeSearch';
import EmployeeDetails from './components/EmployeeDetails';

const App: React.FC = () => {
  const [selectedId, setSelectedId] = useState<number | null>(null);

  return (
    <div className="container py-4">
      <h2 className="mb-4">Поиск сотрудников</h2>
      <EmployeeSearch onSelect={setSelectedId} />
      {selectedId && <EmployeeDetails id={selectedId} />}
    </div>
  );
};

export default App;