import React, { useState } from "react";
import "./styles/createTask.css";

const CreateTask = () => {
  const [task, setTask] = useState({
    title: "",
    description: "",
    deadline: "",
  });

  const handleChange = (e) => {
    setTask({ ...task, [e.target.name]: e.target.value });
  };

  const handleSubmit = (e) => {
    e.preventDefault();
    onSubmit(task);
    console.log("Task Created:", task);
    setTask({ title: "", description: "", deadline: "" });
  };

  const onSubmit = async (task) => {
    try {
      const { title, deadline } = task;
      const formattedDeadline = new Date(deadline).toISOString();
      await fetch("http://localhost:8080/todo-list-service/task/create", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({
          title,
          deadline: formattedDeadline,
        }),
      });
    } catch (error) {
      console.log("Error :: ", error);
    }
  };

  return (
    <div className="create-task-container">
      <h2>Create Task</h2>
      <form onSubmit={handleSubmit}>
        <div className="form-group">
          <label>Title</label>
          <input
            type="text"
            name="title"
            value={task.title}
            onChange={handleChange}
            required
          />
        </div>
        <div className="form-group">
          <label>Description</label>
          <textarea
            name="description"
            value={task.description}
            onChange={handleChange}
            required
          />
        </div>
        <div className="form-group">
        <label>Deadline</label>
        <input
            type="datetime-local"
            name="deadline"
            value={task.deadline}
            onChange={handleChange}
            required
        />
        </div>
        <button type="submit">Create Task</button>
      </form>
    </div>
  );
};

export default CreateTask;
