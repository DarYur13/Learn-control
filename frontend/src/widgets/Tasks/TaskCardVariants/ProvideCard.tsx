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
import Grid from "@mui/material/Grid";
import { useState } from "react";

export default function ProvideCard({ task }: { task: Task }) {
  const [loading, setLoading] = useState(false);
  const [done, setDone] = useState(task.done);

  const handleComplete = async () => {
    try {
      setLoading(true);
      await closeAssignTask(task.id, TaskType.PROVIDE);
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
          {/* Верх: заголовок и статус */}
          <Stack direction="row" justifyContent="space-between" alignItems="center">
            <Typography variant="h6">{task.description}</Typography>
            <Chip
              label={done ? "Завершено" : "Не завершено"}
              color={done ? "success" : "warning"}
              variant="outlined"
            />
          </Stack>

          <Divider />

          {/* Информация по строкам */}
          <Stack spacing={1}>
            {[
              { label: "Сотрудник", value: task.employee },
              { label: "Обучение", value: task.training },
              { label: "Отдел", value: task.department },
              { label: "Должность", value: task.position },
            ].map(({ label, value }) => (
              <Box key={label} display="flex" justifyContent="space-between">
                <Typography><strong>{label}:</strong></Typography>
                <Typography>{value}</Typography>
              </Box>
            ))}
          </Stack>

          {/* Кнопки */}
          {!done && (
            <Stack direction="row" spacing={2} justifyContent="flex-end">
              {task.downloadFileLink && (
                <Button
                  variant="outlined"
                  color="primary"
                  href={task.downloadFileLink}
                  target="_blank"
                >
                  Скачать лист регистрации
                </Button>
              )}
              <Button
                variant="contained"
                disabled={loading}
                onClick={handleComplete}
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
