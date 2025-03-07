import React, { useState, useEffect } from "react";
import { Container, TextField, Typography, Autocomplete, CircularProgress } from "@mui/material";
import PersonalCard from "../components/PersonalCard";

export default function PersonalCards() {
  const [name, setName] = useState("");
  const [employee, setEmployee] = useState(null);
  const [suggestions, setSuggestions] = useState([]);
  const [loading, setLoading] = useState(false);

  useEffect(() => {
    if (name.length > 1) {
      setLoading(true);
      fetch("http://localhost:8000/employees/get_list_by_name", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ name }),
      })
        .then((res) => res.json())
        .then((data) => setSuggestions(data.employees || []))
        .catch(() => setSuggestions([]))
        .finally(() => setLoading(false));
    } else {
      setSuggestions([]);
    }
  }, [name]);

  const handleSelect = (event, newValue) => {
    if (newValue) {
      setLoading(true);
      fetch("http://localhost:8000/employees/get_personal_card", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ id: newValue.id }),
      })
        .then((res) => res.json())
        .then((data) => setEmployee({ ...data, trainings: [...data.trainings] })) // Глубокая копия
        .catch(() => setEmployee(null))
        .finally(() => setLoading(false));
    }
  };

  const handleSaveDate = (employeeID, trainingID, newDate) => {
    fetch("http://localhost:8000/employees/update_training_date", {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({ employeeID, trainingID, date: newDate }),
    })
      .then(() => {
        setEmployee((prev) => {
          if (!prev) return prev;

          // Создаем новый массив, чтобы React понял, что состояние изменилось
          const updatedTrainings = prev.trainings.map((t) =>
            t.id === trainingID ? { ...t, passDate: newDate } : { ...t }
          );

          return { ...prev, trainings: updatedTrainings };
        });
      })
      .catch(() => alert("Ошибка обновления даты"));
  };

  return (
    <Container maxWidth="md" sx={{ mt: 4 }}>
      <Typography variant="h5" gutterBottom>
        Личные карты сотрудников
      </Typography>

      <Autocomplete
        options={suggestions}
        getOptionLabel={(option) => `${option.fullname} (${option.birthdate})`}
        onInputChange={(event, newInputValue) => setName(newInputValue)}
        onChange={handleSelect}
        loading={loading}
        renderInput={(params) => (
          <TextField {...params} label="Введите ФИО сотрудника" variant="outlined" />
        )}
      />

      {loading && <CircularProgress sx={{ display: "block", mx: "auto", mt: 2 }} />}
      {employee && <PersonalCard employee={employee} onSaveDate={handleSaveDate} />}
    </Container>
  );
}
