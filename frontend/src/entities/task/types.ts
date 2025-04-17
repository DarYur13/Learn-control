export enum TaskType {
  PROVIDE = "PROVIDE",
  ASSIGN = "ASSIGN",
  CHOOSE = "CHOOSE",
  SET = "SET",
  CONFIRM = "CONFIRM",
  CONTROL = "CONTROL",
  UNKNOWN = "UNKNOWN_TASK",
}

export interface Task {
  id: number;
  type: TaskType;
  description: string;
  employee: string;
  training: string;
  position: string;
  department: string;
  executor: string;
  downloadFileLink: string;
  done: boolean;
}
