import React, { useState, useEffect } from "react";
import { Container, Typography, Paper, CircularProgress } from "@mui/material";
import FilterMenu from "../components/FilterMenu";
import EmployeeCard from "../components/CompactCard";

export default function EmployeeList() {
  const [employees, setEmployees] = useState([]);
  const [loading, setLoading] = useState(false);

  const fetchEmployees = (filters) => {
    setLoading(true);
    fetch("http://localhost:8000/employees/get_list_by_filters", {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify(filters),
    })
      .then((res) => res.json())
      .then((data) => {
        setEmployees(data.employees || []);
        setLoading(false);
      })
      .catch(() => setLoading(false));
  };

  return (
    <Container maxWidth="md" sx={{ mt: 4 }}>
      <Typography variant="h5" gutterBottom align="center">Обзор сотрудников</Typography>
      <FilterMenu onApplyFilters={fetchEmployees} />
      {loading ? (
        <CircularProgress sx={{ display: "block", mx: "auto", mt: 2 }} />
      ) : (
        <Paper elevation={3} sx={{ p: 2, mt: 2 }}>
          {employees.length === 0 ? (
            <Typography color="textSecondary" align="center">Нет сотрудников</Typography>
          ) : (
            employees.map((employee, index) => <EmployeeCard key={index} employee={employee} />)
          )}
        </Paper>
      )}
    </Container>
  );
}
