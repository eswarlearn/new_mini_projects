import React from "react";
import { useStore } from "../store";

const TodoList = () => {
//   const todos = [];
const todos = useStore((state) => state.todos);
const toggleTodo = useStore((state) => state.toggleTodo);
const deleteTodo = useStore((state) => state.deleteTodo);
  return (
    <div>
      {todos.map((todos, index) => (
        <div className="todo-item" key={index}>
          <span className={todos.completed ? 'completed':''}>{todos.text}</span>
          <button className="btn btn-success" disabled={todos.completed} onClick={()=> toggleTodo(index)}> Complete </button>
          <button className="btn btn-danger"  onClick={()=> deleteTodo(index)}>Delete</button>
        </div>
      ))}
    </div>
  );
};

export default TodoList;
