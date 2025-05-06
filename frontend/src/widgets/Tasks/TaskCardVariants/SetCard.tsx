import { Task, TaskType } from "@/entities/task/types";
import { closeWithDate } from "@/shared/api/tasks";
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
import { DatePicker } from "@mui/x-date-pickers";
import dayjs, { Dayjs } from "dayjs";
import "dayjs/locale/ru"; // Локализация
import utc from "dayjs/plugin/utc";
import { useState } from "react";

dayjs.extend(utc);

export default function SetCard({ task }: { task: Task }) {
  const [date, setDate] = useState<Dayjs | null>(null);
  const [done, setDone] = useState(task.done);
  const [loading, setLoading] = useState(false);

  const handleComplete = async () => {
    if (!date) return;

    try {
      setLoading(true);
      await closeWithDate(
        task.id,
        dayjs.utc(date.format("YYYY-MM-DD")).toISOString(),
        TaskType.SET
      );
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
            ].map(({ label, value }) => (
              <Box key={label} display="flex" justifyContent="space-between">
                <Typography><strong>{label}:</strong></Typography>
                <Typography>{value}</Typography>
              </Box>
            ))}
          </Stack>

          {/* Кнопки + календарь */}
          {!done && (
            <Stack direction="row" spacing={2} justifyContent="flex-end" alignItems="center">
              <DatePicker
                label="Дата обучения"
                value={date}
                onChange={(newDate) => setDate(newDate)}
                disableFuture={false}
                slotProps={{ textField: { size: "small", sx: { width: 220 } } }}
              />
              <Button
                variant="contained"
                disabled={!date || loading}
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
