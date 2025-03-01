import { create } from "zustand";  
// import { devtools } from "zustand/middleware";


export const useStore = create((set) => ({
    todos: [],
    addTodo: (todo) => set((state) => ({ todos: [...state.todos, todo] })),
    deleteTodo: (index) =>
      set((state) => ({
        todos: state.todos.filter((_, i) => i !== index),
      })),
    toggleTodo: (index) =>set((state) => ({todos: state.todos.map((todo, i) => i === index ? {...todo, completed: !todo.completed} : todo) })),
}));