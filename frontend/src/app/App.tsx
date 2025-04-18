import { BrowserRouter, Route, Routes } from "react-router-dom";
import TasksPage from "@/pages/TasksPage";
import { LocalizationProvider } from "@mui/x-date-pickers";
import { AdapterDayjs } from "@mui/x-date-pickers/AdapterDayjs";
import 'dayjs/locale/ru';
import PersonalCardPage from "@/pages/PersonalCardPage";

export default function App() {
  return (
    <LocalizationProvider dateAdapter={AdapterDayjs} adapterLocale="ru">
      <BrowserRouter>
        <Routes>
          <Route path="/" element={<TasksPage />} />
          <Route path="/cards" element={<PersonalCardPage />} />
        </Routes>
      </BrowserRouter>
    </LocalizationProvider>
  );
}

