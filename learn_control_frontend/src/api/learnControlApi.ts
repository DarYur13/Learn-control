import axios from 'axios';

// Базовый экземпляр Axios
const api = axios.create({ 
  baseURL: 'https://6dnqnvhj-8000.uks1.devtunnels.ms', // Единая точка входа
});

export const getEmployeesByName = async (name: string) => {
  const { data } = await api.post('/employees/get/by_name', { name });
  return data.employees;
};

export const getEmployeePersonalCard = async (id: number) => {
  const { data } = await api.post(`/employee/personal_card/get`, { id });
  return data;
};

export const getFilters = async () => {
  const { data } = await api.get('/employees/filters/get');
  return data;
};
