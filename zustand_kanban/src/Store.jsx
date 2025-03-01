import { create } from 'zustand';
import { devtools,persist } from 'zustand/middleware';
/*store is a arrow function
*set is a parameter basically a state setter for the whole store
* it returns a object returns everything in a store*/ 
const store = (set)=> ({
    tasks:[],
    draggedTask:null,
    /*
    *addTask is a funtion with args it alter our set function tp manipulate our store
    *in set we will have a object to manuplate*/
    addTask: (title, state) =>set((store)=>({tasks: [...store.tasks,{title,state}]}),false,"addTask"),

    deleteTask: (title) =>set((store)=>({tasks: store.tasks.filter((task)=>task.title !== title)})),
    setDraggedTask: (title) => set({draggedTask:title}),
    moveTask: (title,state) => set(store => ({tasks: store.tasks.map(task => task.title === title ? {title,state}: task)})),
});

const log =(config) =>(set,get,api)=>config(
    (...args)=>{
        // const current= get()
        // if(!current){

        // }
        console.log(args);
        set(...args)        
    },
    get,
    api
)

export const useStore = create(log(persist(devtools(store),{name:"store"})));

// export const useStore = create(log(devtools(store),{name:"store"}));