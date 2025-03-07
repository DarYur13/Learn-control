import React, { useState } from "react";
import { Accordion, AccordionSummary, AccordionDetails, Typography, IconButton, Box, TextField } from "@mui/material";
import ExpandMoreIcon from "@mui/icons-material/ExpandMore";
import EditIcon from "@mui/icons-material/Edit";
import SaveIcon from "@mui/icons-material/Save";

export default function EmployeeCard({ employee }) {
  const [expanded, setExpanded] = useState(false);
  const [trainingDates, setTrainingDates] = useState(
    employee.trainings.reduce((acc, training) => {
      acc[training.name] = { 
        passDate: training.passDate || "", 
        rePassDate: training.rePassDate || "" 
      };
      return acc;
    }, {})
  );

  const [editing, setEditing] = useState(null);

  const handleExpand = () => {
    setExpanded(!expanded);
  };

  const handleEditClick = (trainingName) => {
    setEditing(trainingName);
  };

  const handleSaveClick = (trainingName) => {
    console.log(`Сохраняем новую дату для ${trainingName}:`, trainingDates[trainingName]);
    setEditing(null);
  };

  const handleDateChange = (trainingName, field, newDate) => {
    setTrainingDates((prev) => ({
      ...prev,
      [trainingName]: { ...prev[trainingName], [field]: newDate }
    }));
  };

  return (
    <Accordion expanded={expanded} onChange={handleExpand}>
      <AccordionSummary expandIcon={<ExpandMoreIcon />}>
        <Typography sx={{ flex: 1 }}>
          {employee.fullname} – {employee.department} – {employee.position}
        </Typography>
      </AccordionSummary>
      <AccordionDetails>
        <Box sx={{ display: "flex", flexDirection: "column", gap: 1 }}>
          {employee.trainings.length === 0 ? (
            <Typography color="textSecondary">Нет пройденных обучений</Typography>
          ) : (
            employee.trainings.map((training) => (
              <Box key={training.name} sx={{ display: "flex", flexDirection: "column", gap: 1, p: 1, border: "1px solid #ddd", borderRadius: 2 }}>
                <Typography sx={{ fontWeight: "bold" }}>{training.name}</Typography>

                {/* Блок с датой прохождения и кнопкой редактирования */}
                <Box sx={{ display: "flex", alignItems: "center", gap: 2 }}>
                  <Typography sx={{ flex: 1 }}>Прохождение:</Typography>
                  {editing === training.name ? (
                    <TextField
                      type="date"
                      size="small"
                      value={trainingDates[training.name]?.passDate || ""}
                      onChange={(e) => handleDateChange(training.name, "passDate", e.target.value)}
                    />
                  ) : (
                    <Typography>{trainingDates[training.name]?.passDate || "Не указано"}</Typography>
                  )}

                  <IconButton 
                    color={editing === training.name ? "success" : "primary"} 
                    onClick={() => (editing === training.name ? handleSaveClick(training.name) : handleEditClick(training.name))}
                  >
                    {editing === training.name ? <SaveIcon /> : <EditIcon />}
                  </IconButton>
                </Box>

                {/* Блок с датой перепрохождения */}
                <Box sx={{ display: "flex", alignItems: "center", gap: 2 }}>
                  <Typography sx={{ flex: 1 }}>Перепрохождение:</Typography>
                  <Typography>{trainingDates[training.name]?.rePassDate || "Не указано"}</Typography>
                </Box>

              </Box>
            ))
          )}
        </Box>
      </AccordionDetails>
    </Accordion>
  );
}
