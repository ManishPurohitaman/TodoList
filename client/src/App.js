import React from "react";
import "./components/styles/task.css";
import TaskBox from "./components/taskBox";
import CreateTask from "./components/createTask";

function App() {
  return (
    <div>
      <CreateTask/>
      <TaskBox/>
    </div>
  );
}

export default App;
