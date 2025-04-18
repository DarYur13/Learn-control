import { BASE_URL } from "@/shared/api/baseURL";

export type Training = {
  id: number;
  name: string;
};
  
export const getTrainings = async (): Promise<Training[]> => {
  const res = await fetch(BASE_URL + "/trainings/get");
  const data = await res.json();
  return data.trainings || [];
};

export const updateTrainingDate = async (
    employeeID: number,
    trainingID: number,
    date: string
  ): Promise<{ passDate: string; rePassDate: string }> => {
    const res = await fetch(BASE_URL + "/employees/update_training_date", {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({ employeeID, trainingID, date }),
    });
  
    if (!res.ok) {
      throw new Error("Ошибка при обновлении даты");
    }
  
    const data = await res.json();
    return {
      passDate: data.passDate,
      rePassDate: data.rePassDate,
    };
  };
  
