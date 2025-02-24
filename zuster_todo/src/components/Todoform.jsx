import { useState } from "react";
import React from "react";
import "bootstrap/dist/css/bootstrap.min.css";
import { useStore } from "../store";

export default function Todoform() {
  const [text, setText] = useState("");

  const addTodo = useStore((todo) => todo.addTodo);

  const handleSubmit = (e) => {
    e.preventDefault();
    if (text.trim() !== "") {
        addTodo({   text: text, completed: false });
      setText('');
    }
  };

  return (
    <>
      <form onSubmit={handleSubmit}>
        <div className="container d-flex">
          <input
            className="form-control"
            type="text"
            value={text}
            onChange={(e) => {
              setText(e.target.value);
            }}
            placeholder="Enter new Tool"
          />

          <button className="btn btn-primary mt-2" type="submit">
            Add Todo
          </button>
        </div>
      </form>
    </>
  );
}
