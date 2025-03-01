import { useState } from 'react'
import reactLogo from './assets/react.svg'
import viteLogo from '/vite.svg'
import './App.css'
import 'bootstrap/dist/css/bootstrap.min.css'
import Todoform from './components/Todoform'  
import TodoList from './components/TodoList'


function App() {
  const [count, setCount] = useState(0)

  return (
   <div className='cotainer'>
    <h1 className='my-4'> Zustand Todo List</h1>
    <Todoform/>
    <TodoList/>
   </div>
   
  )
}

export default App
