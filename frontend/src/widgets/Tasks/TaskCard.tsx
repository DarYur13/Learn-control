import { Task, TaskType } from "@/entities/task/types";
import ProvideCard from "./TaskCardVariants/ProvideCard";
// подключай остальные карточки по мере реализации

export const TaskCard = ({ task }: { task: Task }) => {
  switch (task.type) {
    case TaskType.PROVIDE:
      return <ProvideCard task={task} />;
    // case TaskType.ASSIGN:
    //   return <AssignCard task={task} />;
    default:
      return <div>Неизвестная задача типа: {task.type}</div>;
  }
};
