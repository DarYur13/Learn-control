import { useEffect, useState } from "react";
import { Task, TaskType } from "@/entities/task/types";
import { getTasks } from "@/shared/api/tasks";
import { TaskCard } from "@/widgets/Tasks/TaskCard";
import { Container, Typography, Tabs, Tab, Box } from "@mui/material";

export default function TasksPage() {
  const [tasks, setTasks] = useState<Task[]>([]);
  const [filter, setFilter] = useState<"all" | "done" | "not_done">("not_done");

  useEffect(() => {
    const fetchTasks = async () => {
      try {
        const doneParam = filter === "all" ? undefined : filter === "done";
        const data = await getTasks(doneParam);
        console.log("Ответ от getTasks:", data.tasks);
        setTasks(Array.isArray(data.tasks) ? data.tasks : []);
      } catch (err) {
        console.error("Ошибка загрузки задач:", err);
        setTasks([]);
      }
    };
    fetchTasks();
  }, [filter]);

  return (
    <Container>
      <Typography variant="h4" gutterBottom>
        Задачи
      </Typography>

      <Tabs value={filter} onChange={(e, val) => setFilter(val)}>
        <Tab value="not_done" label="Незавершенные" />
        <Tab value="done" label="Завершенные" />
        <Tab value="all" label="Все" />
      </Tabs>

      <Box mt={2} display="flex" flexDirection="column" gap={2}>
        {tasks.map((task) => (
          <TaskCard key={task.id} task={task} />
        ))}
      </Box>
    </Container>
  );
}