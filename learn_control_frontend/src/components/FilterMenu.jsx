import React, { useState, useEffect } from "react";
import { Box, FormControl, InputLabel, Select, MenuItem, Button, Typography, Checkbox, Slider, TextField, Popover, FormControlLabel} from "@mui/material";
import { Close } from "@mui/icons-material";

const convertToTimestamp = (dateString) => {
  const date = new Date(dateString);
  return {
    seconds: Math.floor(date.getTime() / 1000),
    nanos: 0,
  };
};

export default function FilterMenu({ onApplyFilters }) {
  const availableFiltersMap = {
    "department": "Отдел",
    "position": "Должность",
    "training": "Обучение",
    "notPassed": "Обучение не пройдено",
    "datePeriod": "Период прохождения",
    "daysUntilRetraining": "Осталось дней до перепрохождения"
  };
  
  const [availableFilters, setAvailableFilters] = useState(Object.keys(availableFiltersMap));
  const [selectedFilters, setSelectedFilters] = useState([]);
  const [filters, setFilters] = useState({ daysUntilRetraining: 30 });
  const [departments, setDepartments] = useState([]);
  const [positions, setPositions] = useState([]);
  const [trainings, setTrainings] = useState([]);
  const [anchorEl, setAnchorEl] = useState(null);

  useEffect(() => {
    fetch("http://localhost:8000/departments/get")
      .then((res) => res.json())
      .then((data) => setDepartments(data.departments || []));

    fetch("http://localhost:8000/positions/get")
      .then((res) => res.json())
      .then((data) => setPositions(data.positions || []));

    fetch("http://localhost:8000/trainings/get")
      .then((res) => res.json())
      .then((data) => setTrainings(data.trainings || []));
  }, []);

  const openFilterMenu = (event) => {
    setAnchorEl(event.currentTarget);
  };

  const closeFilterMenu = () => {
    setAnchorEl(null);
  };

  const addFilter = (filter) => {
    if (!selectedFilters.includes(filter)) {
      setSelectedFilters([...selectedFilters, filter]);
      setAvailableFilters(availableFilters.filter(f => f !== filter));
      setFilters((prevFilters) => ({ ...prevFilters, [filter]: filter === "notPassed" ? false : "" }));
    }
    closeFilterMenu();
  };

  const removeFilter = (filter) => {
    setSelectedFilters(selectedFilters.filter(f => f !== filter));
    setAvailableFilters([...availableFilters, filter]);
    const updatedFilters = { ...filters };
    delete updatedFilters[filter];
    setFilters(updatedFilters);
  };

  const handleApplyFilters = () => {
    const requestBody = {
      department: selectedFilters.includes("department") ? filters.department : undefined,
      position: selectedFilters.includes("position") ? filters.position : undefined,
      trainingID: selectedFilters.includes("training") ? filters.training : undefined,
      retrainingIn: selectedFilters.includes("daysUntilRetraining") ? parseInt(filters.daysUntilRetraining, 10) || 30 : undefined,
      trainigsNotPassed: selectedFilters.includes("notPassed") ? filters.notPassed : undefined,
      dateFrom: selectedFilters.includes("datePeriod") && filters.dateFrom ? convertToTimestamp(filters.dateFrom) : undefined,
      dateTo: selectedFilters.includes("datePeriod") && filters.dateTo ? convertToTimestamp(filters.dateTo) : undefined,
    };
  
    console.log("Отправляемый запрос:", requestBody);
  
    onApplyFilters(requestBody);
  };
  

  return (
    <Box sx={{ display: "flex", flexDirection: "column", gap: 2, mt: 2, mb: 2 }}>
      <Button variant="contained" onClick={openFilterMenu} disabled={availableFilters.length === 0}>+</Button>
      <Popover open={Boolean(anchorEl)} anchorEl={anchorEl} onClose={closeFilterMenu} anchorOrigin={{ vertical: "bottom", horizontal: "left" }}>
        <Box sx={{ p: 2 }}>
          {availableFilters.map((filter, index) => (
            <Button key={index} onClick={() => addFilter(filter)} sx={{ display: "block", textAlign: "left" }}>{availableFiltersMap[filter]}</Button>
          ))}
        </Box>
      </Popover>
      {selectedFilters.map((filter, index) => (
      <Box key={index} sx={{ display: "flex", alignItems: "center", gap: 1, width: "100%" }}>
        <FormControl fullWidth variant="outlined">
          {["department", "position", "training"].includes(filter) && (
            <InputLabel id={`filter-label-${filter}`}>{availableFiltersMap[filter]}</InputLabel>
          )}

          {filter === "notPassed" ? (
            <FormControlLabel
              control={
                <Checkbox
                  checked={filters.notPassed || false}
                  onChange={(e) => setFilters({ ...filters, notPassed: e.target.checked })}
                />
              }
              label={
                <Typography
                  sx={{
                    color: filters.notPassed ? "#1976d2" : "#757575", // Синий при активации, серый в обычном состоянии
                    fontSize: "1rem",
                    fontWeight: 400,
                  }}
                >
                  {availableFiltersMap[filter]}
                </Typography>
              }
            />
          ) : filter === "datePeriod" ? (
            <Box sx={{ display: "flex", flexDirection: "column", gap: 1 }}>
              <Typography sx={{ fontSize: "1rem", fontWeight: 400, color: "#757575" }}>{availableFiltersMap[filter]}</Typography>
              <Box sx={{ display: "flex", gap: 2 }}>
                <TextField 
                  type="date" 
                  fullWidth 
                  variant="outlined" 
                  value={filters.dateFrom || ""}
                  onChange={(e) => {
                    const newDateFrom = e.target.value;
                    setFilters({
                      ...filters,
                      dateFrom: newDateFrom,
                      dateTo: filters.dateTo && filters.dateTo < newDateFrom ? newDateFrom : filters.dateTo, // Если dateTo раньше dateFrom, делаем их одинаковыми
                    });
                  }}
                  InputLabelProps={{ shrink: true }}
                />
                <TextField 
                  type="date" 
                  fullWidth 
                  variant="outlined" 
                  value={filters.dateTo || ""}
                  onChange={(e) => {
                    const newDateTo = e.target.value;
                    setFilters({
                      ...filters,
                      dateTo: newDateTo,
                      dateFrom: filters.dateFrom && filters.dateFrom > newDateTo ? newDateTo : filters.dateFrom, // Если dateFrom позже dateTo, делаем их одинаковыми
                    });
                  }}
                  InputLabelProps={{ shrink: true }}
                />
              </Box>
            </Box>
          ) : filter === "daysUntilRetraining" ? (
            <Box sx={{ display: "flex", flexDirection: "column", gap: 1, width: "100%" }}>
              <Box sx={{ display: "flex", alignItems: "center", gap: 2, width: "100%" }}>
                <Typography sx={{ fontSize: "1rem", fontWeight: 400, color: "#757575", whiteSpace: "nowrap" }}>
                  {availableFiltersMap[filter]}
                </Typography>
                <TextField
                  type="number"
                  variant="outlined"
                  value={filters.daysUntilRetraining || 30}
                  onChange={(e) => {
                    const value = Math.min(60, Math.max(1, Number(e.target.value)));
                    setFilters({ ...filters, daysUntilRetraining: value });
                  }}
                  sx={{ flex: 1 }}
                />
              </Box>
              <Slider
                min={1}
                max={60}
                value={Number(filters.daysUntilRetraining) || 30}
                onChange={(e, value) => setFilters({ ...filters, daysUntilRetraining: value })}
                sx={{ flex: 1, width: "100%" }}
              />
            </Box>
          ) : (
            <Select
              labelId={`filter-label-${filter}`}
              value={filters[filter] || ""}
              onChange={(e) => setFilters({ ...filters, [filter]: e.target.value })}
              label={availableFiltersMap[filter]}
            >
              {(filter === "training"
                ? trainings.map((train) => (
                    <MenuItem key={train.id} value={train.id}>{train.name}</MenuItem>
                  ))
                : (filter === "department" ? departments : positions).map((item, i) => (
                    <MenuItem key={i} value={item}>{item}</MenuItem>
                  ))
              )}
            </Select>
          )}
        </FormControl>
        <Button
          onClick={() => removeFilter(filter)}
          sx={{
            minWidth: "30px",
            height: "30px",
            p: 0,
            color: "#9e9e9e", // Бледно-серый цвет крестика
            background: "transparent",
            marginTop: filter === "daysUntilRetraining" ? "-35px" : filter === "datePeriod" ? "35px" : "0px",
            "&:hover": { color: "#d32f2f", background: "rgba(211, 47, 47, 0.1)" },
          }}
        >
          <Close fontSize="small" />
        </Button>
      </Box>
    ))}
      <Button variant="contained" color="primary" onClick={handleApplyFilters}>
        Поиск
      </Button>
    </Box>
  );
}

