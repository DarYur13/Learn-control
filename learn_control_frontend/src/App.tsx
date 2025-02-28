import React from 'react';
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';
import Menu from './pages/Menu';
import Overview from './pages/Overview';
import EmployeeDetails from './components/EmployeeDetails';

const App: React.FC = () => {
  return (
    <Router>
      <div className="container py-4">
        <Routes>
          <Route path="/" element={<Menu />} />
          <Route path="/overview" element={<Overview />} />
          <Route path="/personal-card" element={<EmployeeDetails id={1} />} />
        </Routes>
      </div>
    </Router>
  );
};

export default App;
