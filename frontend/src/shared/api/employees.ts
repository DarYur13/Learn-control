import { EmployeeBaseInfo, PersonalCard } from "@/entities/employee/types";
import { BASE_URL } from "./baseURL";

export const getEmployeesByName = async (name: string): Promise<EmployeeBaseInfo[]> => {
  const res = await fetch(BASE_URL + "/employees/get_list_by_name", {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({ name }),
  });

  const data = await res.json();
  return data.employees;
};

export const getPersonalCard = async (id: number): Promise<PersonalCard> => {
  const res = await fetch(BASE_URL + "/employees/get_personal_card", {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({ id }),
  });

  const data = await res.json();
  return data;
};
