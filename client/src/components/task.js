import React from "react";

const Task = ({ title, description, deadline, status = "Pending" }) => {
  return (
    <div className="task-card">
      <h2>{title}</h2>
      <p>{description}</p>
      <p><strong>Deadline:</strong> {deadline}</p>
      <p><strong>Status:</strong> {status}</p>
    </div>
  );
};

export default Task;
