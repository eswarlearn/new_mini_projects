// import { useMemo } from 'react';
// import { useStore } from '../Store'
// import './Column.css'
// import Task from './Task'
// import { shallow } from 'zustand/shallow';

// export default function Column({state}) {
//   const tasks =useStore(store=>
//     store.tasks.filter((task) => task.state === state), //.filter will cause more re-rendering
//     shallow
//   );
// // const filtered = useMemo[()=> tasks.filter((task) => task.state === state),[tasks,state]]

//     return (
//       <div className="column">
//         <p>{state}</p>
//         {tasks.map((task)=><Task
//         title={task.title} key={task.title}/>)}
//         {/* <Task title='Todo'/> */}
//       </div>
//     )
// }

import { useMemo, useState } from "react";
import { useStore } from "../Store";
import "./Column.css";
import Task from "./Task";
import { shallow } from "zustand/shallow";
import classNames from "classnames/bind";

export default function Column({ state }) {

  const [text,setText] = useState('');
  const [open,setOpen] = useState(false);  
  const [drop,setDrop] = useState(false);

  // Get all tasks from Zustand (without filtering)
  const tasks = useStore((store) => store.tasks, shallow); // ✅ Zustand only updates if task list changes

  // Filter tasks efficiently using useMemo
  const filteredTasks = useMemo(
    () => tasks.filter((task) => task.state === state),
    [tasks, state]
  );

  const addTask =useStore((store) => store.addTask)
  const setDraggedTask=useStore((store)=>store.setDraggedTask);
  const draggedTask=useStore((store)=>store.draggedTask);
  const moveTask=useStore((store)=>store.moveTask);


  return (
    <div className={classNames("column",{drop: drop})} 
    onDragOver={e=>{
      setDrop(true);
      e.preventDefault();
    }}

    onDragLeave={e=>{
      setDrop(false);
      e.preventDefault();
    }}

    onDrop={e=>{
      console.log(draggedTask);
      setDrop(false);
      moveTask(draggedTask, state)
      setDraggedTask(null);
    }}>
      <div className="titleWrapper">
        <p>{state}</p>
        <button onClick={()=>setOpen(true)}>Add</button>
      </div>

      {filteredTasks.map((task) => (
        <Task title={task.title} key={task.title} />
      ))}
      {open && (<div className="Modal">
        <div className="modalContent">
          <input onChange={e=>setText(e.target.value)} value={text}/>
          <button onClick={()=>{
            addTask(text,state);
            setText('');
            setOpen(false);
          }}> Submit </button>
        </div>
      </div>)}
    </div>
  );
}
