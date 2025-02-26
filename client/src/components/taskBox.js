import React, { useEffect, useState } from "react";
import Task from "./task"; 
import "./styles/taskbox.css"; 

const TaskBox = () => {
  const [tasks, setTasks] = useState([]); 

  useEffect(() => {
    const fetchTasks = async () => {
      try {
        const response = await fetch("http://localhost:8080/todo-list-service/task/getPending", {
          method: "POST",
          headers: { "Content-Type": "application/json" },
          body: JSON.stringify({}),
        });

        console.log(response);
        if (!response.ok) {
          throw new Error(`HTTP error! Status: ${response.status}`);
        }

        const data = await response.json();
        setTasks(data);
      } catch (error) {
        console.error("Error fetching tasks:", error);
      }
    };

    fetchTasks();
  }, []); 

  return (
    <div className="task-box">
      <h2>Pending Tasks</h2>
      {tasks.length === 0 ? (
        <p>No pending tasks</p>
      ) : (
        tasks.map((task, index) => (
          <Task
            key={index}
            title={task.title}
            description={task.description}
            deadline={task.deadline}
            status={task.status || "Pending"} 
          />
        ))
      )}
    </div>
  );
};

export default TaskBox;
