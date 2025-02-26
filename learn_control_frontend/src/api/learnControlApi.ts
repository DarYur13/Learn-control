import axios from 'axios';

const api = axios.create({ baseURL: 'http://localhost:8000' });

export const getEmployees = async (name: string) => {
  const { data } = await api.post('/employees/get', { name });
  return data.employees;
};

export const getEmployee = async (id: number) => {
  const { data } = await api.get(`/employee/${id}`);
  return data;
};