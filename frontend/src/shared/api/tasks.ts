import { Task, TaskType } from "@/entities/task/types";

const BASE_URL = "http://localhost:8000";

type GetTasksResponse = { tasks: Task[] };

export const getTasks = async (done?: boolean): Promise<GetTasksResponse> => {
  const url = done === undefined ? "/tasks/get" : `/tasks/get?done=${done}`;
  const res = await fetch(url);
  return res.json();
};


export const closeAssignTask = async (
  taskID: number,
  employeeID: number,
  trainingID: number,
  taskType: TaskType
) => {
  return fetch(BASE_URL + "/tasks/close", {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({ taskID, employeeID, trainingID, taskType }),
  });
};

export const closeWithDate = async (
  taskID: number,
  employeeID: number,
  trainingID: number,
  date: string, // ISO string
  taskType: TaskType
) => {
  return fetch(BASE_URL + "/tasks/close_with_training_date_set", {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({ taskID, employeeID, trainingID, date, taskType }),
  });
};

export const closeWithProtocol = async (
  taskID: number,
  employeeID: number,
  trainingID: number
) => {
  return fetch(BASE_URL + "/tasks/close_with_training_protocol_confirm", {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({ taskID, employeeID, trainingID }),
  });
};

export const closeChooseTask = async (
  taskID: number,
  positionID: number,
  trainingsIDs: number[]
) => {
  return fetch(BASE_URL + "/tasks/close_with_position_trainings_set", {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({ taskID, positionID, trainingsIDs }),
  });
};
