import { BrowserRouter, Route, Routes } from "react-router-dom";
import TasksPage from "@/pages/TasksPage";

export default function App() {
  return (
    <BrowserRouter>
      <Routes>
        <Route path="/" element={<TasksPage />} />
      </Routes>
    </BrowserRouter>
  );
}
