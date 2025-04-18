import { Task, TaskType } from "@/entities/task/types";
import { BASE_URL } from "@/shared/api/baseURL";

type GetTasksResponse = { tasks: Task[] };

export const getTasks = async (done?: boolean): Promise<GetTasksResponse> => {
  const url = done === undefined ? "/tasks/get" : `/tasks/get?done=${done}`;
  const res = await fetch(url);
  return res.json();
};


export const closeAssignTask = async (
  taskID: number,
  taskType: TaskType
) => {
  return fetch(BASE_URL + "/tasks/close", {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({ taskID, taskType }),
  });
};

export const closeWithDate = async (
  taskID: number,
  date: string, // ISO string
  taskType: TaskType
) => {
  return fetch(BASE_URL + "/tasks/close_with_training_date_set", {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({ taskID, date, taskType }),
  });
};

export const closeWithProtocol = async (
  taskID: number
) => {
  return fetch(BASE_URL + "/tasks/close_with_training_protocol_confirm", {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({ taskID }),
  });
};

export const closeChooseTask = async (
  taskID: number,
  trainingsIDs: number[]
) => {
  return fetch(BASE_URL + "/tasks/close_with_position_trainings_set", {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({ taskID, trainingsIDs }),
  });
};
