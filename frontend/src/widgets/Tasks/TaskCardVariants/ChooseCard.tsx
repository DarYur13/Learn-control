import { Task, Training } from "@/entities/task/types";
import { closeChooseTask } from "@/shared/api/tasks";
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
  Dialog,
  DialogTitle,
  DialogContent,
  DialogActions,
  Checkbox,
  FormControlLabel,
} from "@mui/material";
import { useEffect, useState } from "react";
import { getTrainings } from "@/shared/api/trainings";

export default function ChooseCard({ task }: { task: Task }) {
  const [done, setDone] = useState(task.done);
  const [loading, setLoading] = useState(false);
  const [open, setOpen] = useState(false);
  const [trainings, setTrainings] = useState<Training[]>([]);
  const [selectedTrainings, setSelectedTrainings] = useState<number[]>([]);

  // Подгружаем обучения один раз
  useEffect(() => {
    if (!done) {
      getTrainings()
        .then(setTrainings)
        .catch((err) => console.error("Ошибка при загрузке обучений", err));
    }
  }, [done]);

  const handleToggle = (id: number) => {
    setSelectedTrainings((prev) =>
      prev.includes(id)
        ? prev.filter((t) => t !== id)
        : [...prev, id]
    );
  };

  const handleSave = async () => {
    try {
      setLoading(true);
      await closeChooseTask(task.id, selectedTrainings);
      setDone(true);
      setOpen(false);
    } catch (err) {
      console.error("Ошибка при сохранении обучений", err);
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
              { label: "Отдел", value: task.department },
              { label: "Должность", value: task.position },
            ].map(({ label, value }) => (
              <Box key={label} display="flex" justifyContent="space-between">
                <Typography><strong>{label}:</strong></Typography>
                <Typography>{value}</Typography>
              </Box>
            ))}
          </Stack>

          {/* Кнопка */}
          {!done && (
            <Stack direction="row" justifyContent="flex-end">
              <Button variant="contained" onClick={() => setOpen(true)}>
                Выбрать обучения
              </Button>
            </Stack>
          )}
        </Stack>
      </CardContent>

      {/* Модалка */}
      <Dialog open={open} onClose={() => setOpen(false)} maxWidth="sm" fullWidth>
        <DialogTitle>Выберите подходящие обучения</DialogTitle>
        <DialogContent>
          <Stack>
            {trainings.map((training) => (
              <FormControlLabel
                key={training.id}
                control={
                  <Checkbox
                    checked={selectedTrainings.includes(training.id)}
                    onChange={() => handleToggle(training.id)}
                  />
                }
                label={training.name}
              />
            ))}
          </Stack>
        </DialogContent>
        <DialogActions>
          <Button onClick={() => setOpen(false)}>Отмена</Button>
          <Button
            variant="contained"
            onClick={handleSave}
            disabled={selectedTrainings.length === 0 || loading}
          >
            {loading ? <CircularProgress size={20} /> : "Сохранить"}
          </Button>
        </DialogActions>
      </Dialog>
    </Card>
  );
}
