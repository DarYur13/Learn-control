import {
    Box,
    Typography,
    IconButton,
    Stack,
  } from "@mui/material";
  import { Edit, Save, Close } from "@mui/icons-material";
  import { DatePicker } from "@mui/x-date-pickers";
  import dayjs, { Dayjs } from "dayjs";
  import { useState } from "react";
  import { Training } from "@/entities/employee/types";
  
  type Props = {
    training: Training;
    employeeID: number;
    onDateUpdate: (employeeID: number, trainingID: number, newDate: string) => Promise<{
      passDate: string;
      rePassDate: string;
    }>;
  };
  
  export default function TrainingCard({ training, employeeID, onDateUpdate }: Props) {
    const [edit, setEdit] = useState(false);
    const [newDate, setNewDate] = useState<Dayjs | null>(
      training.passDate ? dayjs(training.passDate) : null
    );
    const [loading, setLoading] = useState(false);
    const [localPass, setLocalPass] = useState(training.passDate);
    const [localRepass, setLocalRepass] = useState(training.rePassDate);
  
    const handleSave = async () => {
      if (!newDate) return;
      try {
        setLoading(true);
        const updated = await onDateUpdate(employeeID, training.id, newDate.toISOString());
  
        // Обновляем локальное отображение
        setLocalPass(updated.passDate);
        setLocalRepass(updated.rePassDate);
        setEdit(false);
      } catch (e) {
        console.error("Ошибка при обновлении даты:", e);
      } finally {
        setLoading(false);
      }
    };
  
    const format = (d: string) => {
      return !d || dayjs(d).isBefore("1900-01-01") ? "—" : dayjs(d).format("DD.MM.YYYY");
    };
  
    return (
      <Box
        sx={{
          border: "1px solid #ddd",
          borderRadius: 2,
          p: 2,
          boxShadow: 1,
          bgcolor: "#fafafa",
        }}
      >
        <Typography variant="subtitle1" gutterBottom>
          {training.name}
        </Typography>
  
        {/* Дата прохождения */}
        <Box display="flex" alignItems="center" justifyContent="space-between">
          <Typography><strong>Дата прохождения:</strong></Typography>
          {edit ? (
            <Stack direction="row" spacing={1} alignItems="center">
              <DatePicker
                value={newDate}
                onChange={(d) => setNewDate(d)}
                slotProps={{ textField: { size: "small", sx: { width: 130 } } }}
              />
              <IconButton onClick={handleSave} disabled={loading} size="small">
                <Save fontSize="small" />
              </IconButton>
              <IconButton onClick={() => setEdit(false)} size="small">
                <Close fontSize="small" />
              </IconButton>
            </Stack>
          ) : (
            <Box display="flex" alignItems="center" gap={1}>
              <Typography>{format(localPass)}</Typography>
              <IconButton onClick={() => setEdit(true)} size="small">
                <Edit fontSize="small" />
              </IconButton>
            </Box>
          )}
        </Box>
  
        {/* Повторная дата */}
        <Box display="flex" justifyContent="space-between" mt={1}>
          <Typography><strong>Повторная:</strong></Typography>
          <Typography>{format(localRepass)}</Typography>
        </Box>
  
        {/* Протокол */}
        <Box display="flex" justifyContent="space-between" mt={1}>
          <Typography><strong>Протокол:</strong></Typography>
          <Typography>{training.hasProtocol}</Typography>
        </Box>
      </Box>
    );
  }
  