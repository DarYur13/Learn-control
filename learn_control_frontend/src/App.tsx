import React from 'react';
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';
import Menu from './pages/Menu';
import Overview from './pages/Overview';
import PersonalCards from './pages/PersonalCards';

const App: React.FC = () => {
  return (
    <Router>
      <div className="container py-4">
        <Routes>
          <Route path="/" element={<Menu />} />
          <Route path="/overview" element={<Overview />} />
          <Route path="/personal-cards" element={<PersonalCards />} />
        </Routes>
      </div>
    </Router>
  );
};

export default App;
