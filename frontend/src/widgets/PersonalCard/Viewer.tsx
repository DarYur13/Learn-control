import {
    Box,
    Typography,
    Stack,
    Paper,
    Divider,
  } from "@mui/material";
  import { PersonalCard, Training } from "@/entities/employee/types";
  import TrainingCard from "./TrainingCard";
  import { updateTrainingDate } from "@/shared/api/trainings";
  import dayjs from "dayjs";
  
  export default function PersonalCardViewer({
    card,
  }: {
    card: PersonalCard;
  }) {
    return (
      <Paper sx={{ p: 3, borderRadius: 3 }}>
        <Typography variant="h5" gutterBottom>
          Личная карта сотрудника
        </Typography>
  
        <Divider sx={{ mb: 2 }} />
  
        {/* Основная информация */}
        <Stack spacing={1} mb={3}>
          {[
            { label: "ФИО", value: card.fullname },
            { label: "Дата рождения", value: dayjs(card.birthdate).format("DD.MM.YYYY") },
            { label: "Отдел", value: card.department },
            { label: "Должность", value: card.position },
            { label: "Дата трудоустройства", value: dayjs(card.employmentDate).format("DD.MM.YYYY") },
            { label: "СНИЛС", value: card.snils },
          ].map(({ label, value }) => (
            <Box key={label} display="flex" justifyContent="space-between">
              <Typography><strong>{label}:</strong></Typography>
              <Typography>{value}</Typography>
            </Box>
          ))}
        </Stack>
  
        {/* Обучения */}
        <Typography variant="h6" gutterBottom>
          Обучения
        </Typography>
  
        <Stack spacing={2}>
          {card.trainings.map((t: Training) => (
            <TrainingCard
              key={t.id}
              training={t}
              employeeID={card.employeeID}
              onDateUpdate={async (empId, trainingId, newDate) => {
                return await updateTrainingDate(empId, trainingId, newDate);
              }}
            />
          ))}
        </Stack>
      </Paper>
    );
  }
  