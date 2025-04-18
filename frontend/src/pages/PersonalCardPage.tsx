import { useEffect, useState } from "react";
import {
  Box,
  Typography,
  TextField,
  Autocomplete,
  CircularProgress,
  Paper,
} from "@mui/material";
import { getEmployeesByName, getPersonalCard } from "@/shared/api/employees";
import { EmployeeBaseInfo, PersonalCard } from "@/entities/employee/types";
import PersonalCardViewer from "@/widgets/PersonalCard/Viewer";

export default function PersonalCardPage() {
  const [query, setQuery] = useState("");
  const [options, setOptions] = useState<EmployeeBaseInfo[]>([]);
  const [selected, setSelected] = useState<EmployeeBaseInfo | null>(null);
  const [card, setCard] = useState<PersonalCard | null>(null);
  const [loading, setLoading] = useState(false);

  // Автопоиск по имени
  useEffect(() => {
    if (query.length < 2) return;
    getEmployeesByName(query).then(setOptions);
  }, [query]);

  // Получение личной карты
  useEffect(() => {
    if (!selected) return;
    setLoading(true);
    getPersonalCard(selected.id)
      .then(setCard)
      .catch(console.error)
      .finally(() => setLoading(false));
  }, [selected]);

  return (
    <Box p={3}>
      <Typography variant="h5" mb={2}>
        Личные карты сотрудников
      </Typography>

      <Autocomplete
        fullWidth
        options={options}
        getOptionLabel={(option) => option.fullname}
        onInputChange={(_, value) => setQuery(value)}
        onChange={(_, value) => setSelected(value)}
        renderInput={(params) => (
          <TextField {...params} label="Поиск по имени" variant="outlined" />
        )}
      />

      <Box mt={4}>
        {loading && <CircularProgress />}
        {card && <PersonalCardViewer card={card} />}
      </Box>
    </Box>
  );
}
