import { useStore } from "../Store";
import "./Task.css";
import classNames from "classnames";
import trash from "../assets/trash.svg";

const STATUS = "PLANNED";

export default function Task({ title }) {
  const task = useStore((store) =>
    store.tasks.find((task) => task.title === title)
  );

  const setDraggedTask=useStore((store)=>store.setDraggedTask);
  const deleteTask =useStore((store)=> store.deleteTask)

  return (
    <div className="task" draggable 
    onDragStart={()=>{setDraggedTask(task.title)}}>
      <div>{title}</div>
      <div className="bottomWrapper">
        <div>
          <img src={trash} onClick={()=>deleteTask(task.title)}/>
        </div>
        <div className={classNames("status", task.state)}>{task.state}</div>
      </div>
    </div>
  );
}
