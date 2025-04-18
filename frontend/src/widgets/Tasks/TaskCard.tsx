import { Task, TaskType } from "@/entities/task/types";
import ProvideCard from "./TaskCardVariants/ProvideCard";
import SetCard from "./TaskCardVariants/SetCard";
import ConfirmCard from "./TaskCardVariants/ConfirmCard";
import ChooseCard from "./TaskCardVariants/ChooseCard";
import AssignCard from "./TaskCardVariants/AssignCard";
import ControlCard from "./TaskCardVariants/ControlCard";

export const TaskCard = ({ task }: { task: Task }) => {
  switch (task.type) {
    case TaskType.PROVIDE:
      return <ProvideCard task={task} />;
    case TaskType.SET:
      return <SetCard task={task} />;
    case TaskType.CONFIRM:
      return <ConfirmCard task={task} />;
    case TaskType.CHOOSE:
      return <ChooseCard task={task} />;
    case TaskType.ASSIGN:
      return <AssignCard task={task} />;
    case TaskType.CONTROL:
      return <ControlCard task={task} />;

    default:
      return <div>Неизвестная задача типа: {task.type}</div>;
  }
};
