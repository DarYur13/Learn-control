import { useState } from "react";
import EmployeeSearch from "./components/EmployeeSearch";
import EmployeeDetails from "./components/EmployeeDetails";

export default function App() {
  const [selectedEmployeeId, setSelectedEmployeeId] = useState<number | null>(null);

  return (
    <div className="p-4">
      <h1 className="text-2xl">Поиск сотрудников</h1>
      <EmployeeSearch onSelect={setSelectedEmployeeId} />
      {selectedEmployeeId && <EmployeeDetails id={selectedEmployeeId} />}
    </div>
  );
}
