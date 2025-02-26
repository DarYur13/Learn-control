import { useState, useEffect } from "react";
import { getEmployees } from "../api/api";

// Интерфейс для сотрудников
interface EmployeeBaseInfo {
  id: number;
  fullname: string;
  birthdate: string;
}

const EmployeeSearch = ({ onSelect }: { onSelect: (id: number) => void }) => {
  const [query, setQuery] = useState("");
  const [employees, setEmployees] = useState<EmployeeBaseInfo[]>([]);

  useEffect(() => {
    if (query.length > 1) {
      getEmployees(query).then(setEmployees);
    } else {
      setEmployees([]);
    }
  }, [query]);

  return (
    <div>
      <input
        type="text"
        placeholder="Введите имя"
        value={query}
        onChange={(e) => setQuery(e.target.value)}
        className="w-full p-2 border rounded"
      />
      {employees.length > 0 && (
        <ul className="mt-2 border rounded">
          {employees.map((e) => (
            <li
              key={e.id}
              className="p-2 hover:bg-gray-200 cursor-pointer"
              onClick={() => onSelect(e.id)}
            >
              {e.fullname} ({e.birthdate})
            </li>
          ))}
        </ul>
      )}
    </div>
  );
};

export default EmployeeSearch;
