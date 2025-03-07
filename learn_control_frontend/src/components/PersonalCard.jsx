import React, { useState } from "react";
import { Paper, Typography, TextField, Box, Divider, IconButton, Container } from "@mui/material";
import EditIcon from "@mui/icons-material/Edit";
import SaveIcon from "@mui/icons-material/Save";

const formatDateForInput = (dateString) => {
  if (!dateString) return "";
  const parts = dateString.split(".");
  return parts.length === 3 ? `${parts[2]}-${parts[1]}-${parts[0]}` : dateString;
};

const formatDateForDisplay = (dateString) => {
  if (!dateString) return "Не пройдено";
  const parts = dateString.split("-");
  return parts.length === 3 ? `${parts[2]}.${parts[1]}.${parts[0]}` : dateString;
};

export default function PersonalCard({ employee, onSaveDate }) {
  const [editingTrainingId, setEditingTrainingId] = useState(null);
  const [editedDates, setEditedDates] = useState({});

  const handleEditClick = (trainingId) => {
    setEditingTrainingId((prevId) => (prevId === trainingId ? null : trainingId));
    setEditedDates((prevDates) => ({
      ...prevDates,
      [trainingId]: formatDateForInput(employee.trainings.find((t) => t.id === trainingId)?.passDate) || "",
    }));
  };

  const handleSaveDate = (trainingId) => {
    if (editedDates[trainingId]) {
      const formattedDate = formatDateForDisplay(editedDates[trainingId]);
      onSaveDate(employee.id, trainingId, formattedDate);
      setEditingTrainingId(null);
    }
  };

  return (
    <Container maxWidth="md" sx={{ mt: 4 }}>
      <Paper elevation={3} sx={{ p: 4, borderRadius: 3 }}>
        <Typography variant="h5" fontWeight="bold" align="center" gutterBottom>
          Личная карта сотрудника
        </Typography>
        <Divider sx={{ mb: 3 }} />

        <Box sx={{ display: "flex", flexDirection: "column", gap: 1 }}>
          <Typography variant="h6">{employee.fullname}</Typography>
          <Typography variant="body1">📅 Дата рождения: {employee.birthdate}</Typography>
          <Typography variant="body1">📄 СНИЛС: {employee.snils}</Typography>
          <Typography variant="body1">🏢 Отдел: {employee.department}</Typography>
          <Typography variant="body1">💼 Должность: {employee.position}</Typography>
          <Typography variant="body1">📆 Дата приёма: {employee.employmentDate}</Typography>
        </Box>

        <Divider sx={{ my: 3 }} />

        <Typography variant="h6" fontWeight="bold" gutterBottom>
          Обучения:
        </Typography>

        {employee.trainings.map((training) => (
          <Box
            key={training.id}
            sx={{
              display: "flex",
              alignItems: "center",
              justifyContent: "space-between",
              p: 2,
              borderRadius: 2,
              backgroundColor: "#f9f9f9",
              mb: 1,
            }}
          >
            <Box>
              <Typography variant="body1" fontWeight="bold">
                {training.name}
              </Typography>
              <Typography variant="body2" color="text.secondary">
                Перепрохождение: {training.rePassDate || "-"}
              </Typography>
            </Box>

            <Box sx={{ display: "flex", alignItems: "center", gap: 1 }}>
              {editingTrainingId === training.id ? (
                <TextField
                  type="date"
                  size="small"
                  value={editedDates[training.id] || ""}
                  onChange={(e) =>
                    setEditedDates((prevDates) => ({
                      ...prevDates,
                      [training.id]: e.target.value,
                    }))
                  }
                />
              ) : (
                <Typography sx={{ minWidth: 100, textAlign: "right" }}>
                  {formatDateForDisplay(training.passDate)}
                </Typography>
              )}

              <IconButton
                color={editingTrainingId === training.id ? "success" : "primary"}
                onClick={(e) => {
                  e.stopPropagation(); // Останавливаем всплытие события
                  editingTrainingId === training.id ? handleSaveDate(training.id) : handleEditClick(training.id);
                }}
              >
                {editingTrainingId === training.id ? <SaveIcon /> : <EditIcon />}
              </IconButton>
            </Box>
          </Box>
        ))}
      </Paper>
    </Container>
  );
}
