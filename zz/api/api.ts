const API_BASE_URL = "localhost:8000";

export const getEmployees = async (name: string) => {
  // Строим тело запроса, как это сделано в curl
  const requestBody = JSON.stringify({ name });

  console.log('Request Body:', requestBody); // Логируем тело запроса

  // Отправляем POST-запрос
  const response = await fetch(`${API_BASE_URL}/employees/get`, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json', // Указываем, что передаем данные как JSON
    },
    body: requestBody, // Передаем тело запроса
  });

  // Проверяем, успешен ли запрос
  if (!response.ok) {
    const errorData = await response.json();
    throw new Error(`Failed to fetch employees: ${errorData.message || 'Unknown error'}`);
  }

  // Возвращаем данные из ответа сервера
  return response.json();
};

export const getEmployee = async (id: number) => {
  const response = await fetch(`${API_BASE_URL}/employee/${id}`);
  return response.json();
};
