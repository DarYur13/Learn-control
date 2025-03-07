import React, { useState } from "react";
import { Routes, Route } from "react-router-dom";
import { Container, Tabs, Tab } from "@mui/material";
import { useNavigate, useLocation } from "react-router-dom";
import EmployeeList from "./pages/EmployeesList";
import PersonalCards from "./pages/PersonalCards";

export default function App() {
  const navigate = useNavigate();
  const location = useLocation();
  const [tab, setTab] = useState(location.pathname);

  const handleChange = (event, newValue) => {
    setTab(newValue);
    navigate(newValue);
  };

  return (
    <Container maxWidth="md" sx={{ mt: 4 }}>
      <Tabs value={tab} onChange={handleChange} centered>
        <Tab label="Обзор сотрудников" value="/employees" />
        <Tab label="Личные карты" value="/personal_cards" />
      </Tabs>
      <Routes>
        <Route path="/employees" element={<EmployeeList />} />
        <Route path="/personal_cards" element={<PersonalCards />} />
      </Routes>
    </Container>
  );
}
