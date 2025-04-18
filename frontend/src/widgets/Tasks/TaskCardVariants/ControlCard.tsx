import { Task, TaskType } from "@/entities/task/types";
import { closeAssignTask } from "@/shared/api/tasks";
import {
  Card,
  CardContent,
  Typography,
  Button,
  CircularProgress,
  Stack,
  Box,
  Divider,
  Chip,
} from "@mui/material";
import { useState } from "react";

export default function ControlCard({ task }: { task: Task }) {
  const [loading, setLoading] = useState(false);
  const [done, setDone] = useState(task.done);

  const handleComplete = async () => {
    try {
      setLoading(true);
      await closeAssignTask(task.id, TaskType.CONTROL);
      setDone(true);
    } catch (err) {
      console.error("Ошибка при завершении задачи", err);
    } finally {
      setLoading(false);
    }
  };

  return (
    <Card variant="outlined" sx={{ borderRadius: 3, boxShadow: 2, p: 2 }}>
      <CardContent>
        <Stack spacing={2}>
          {/* Заголовок и статус */}
          <Stack direction="row" justifyContent="space-between" alignItems="center">
            <Typography variant="h6">{task.description}</Typography>
            <Chip
              label={done ? "Завершено" : "Не завершено"}
              color={done ? "success" : "warning"}
              variant="outlined"
            />
          </Stack>

          <Divider />

          {/* Инфо о задаче */}
          <Stack spacing={1}>
            {[
              { label: "Сотрудник", value: task.employee },
              { label: "Обучение", value: task.training },
              { label: "Отдел", value: task.department },
              { label: "Должность", value: task.position },
              { label: "Исполнитель", value: task.executor },
            ].map(({ label, value }) => (
              <Box key={label} display="flex" justifyContent="space-between">
                <Typography><strong>{label}:</strong></Typography>
                <Typography>{value}</Typography>
              </Box>
            ))}
          </Stack>

          {/* Кнопка завершения */}
          {!done && (
            <Stack direction="row" justifyContent="flex-end">
              <Button
                variant="contained"
                onClick={handleComplete}
                disabled={loading}
              >
                {loading ? <CircularProgress size={20} /> : "Завершить"}
              </Button>
            </Stack>
          )}
        </Stack>
      </CardContent>
    </Card>
  );
}
