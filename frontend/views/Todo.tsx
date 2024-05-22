import React, { useRef, useState } from 'react'
import updateCurrentLocation from '../functions/location'
import "../css/todo.css"
import { TaskClass, TaskDrag, TaskList } from '../classes/Tasks'
import Task from '../components/Task'
import getUser, { getCacheOf, storeInLocal } from '../functions/store'
import { useUser } from '../components/UserContext'
import UserData from '../classes/UserData'
import useTasks from '../functions/useTasks'
import useGetTasks from '../functions/useGetTasks'
import Error from '../components/Error'
import Loading from '../components/Loading'
import Message from '../components/Message'
import TaskContainer from '../components/TaskContainer'
import { createTasksCollection, addTask } from '../functions/todoFunctions'
import LoadingTodo from '../components/LoadingTodo'
import MessagePopUp from '../components/MessagePopUp'
import tasks from '../functions/fetchs/tasks'

const days = ["Lunes", "Martes", "Miércoles", "Jueves", "Viernes", "Sábado", "Domingo"]

export default function Todo() {

 const dragTask = useRef<TaskDrag>({ day: "", index: 0 })
 const draggedOverTask = useRef<TaskDrag>({ day: "", index: 0 })
 updateCurrentLocation()
 const userLocal = ((useUser().username == "") ? getUser() : useUser() as UserData)
 const identifier = userLocal?.username as string
 let [taskCache, setTaskCache] = useState(getCacheOf("tasks"))
 console.log(taskCache)
 const [taskCacheUpdate, setTaskCacheUpdate]: any = useState()
 const successMessageRef = useRef()
 const errorMessageRef = useRef()
 const tasksResponse = useGetTasks(identifier, taskCache)
 const savingTask = useTasks(taskCacheUpdate, identifier)

 let message = <MessagePopUp savingTask={savingTask} successMessageRef={successMessageRef} errorMessageRef={errorMessageRef} />

 function handleSort() {
  const tasksClone: TaskList = { ...taskCache as object }
  console.log("SORTING")
  console.log("dragging", dragTask.current)
  console.log("over", draggedOverTask.current)
  console.log("taskCache", tasksClone)
  const temp = tasksClone[dragTask.current.day][dragTask.current.index]
  tasksClone[dragTask.current.day][dragTask.current.index] = tasksClone[draggedOverTask.current.day][draggedOverTask.current.index]
  tasksClone[draggedOverTask.current.day][draggedOverTask.current.index]
   = temp
  setTaskCache(tasksClone)
  storeInLocal(tasksClone, "tasks")
 }


 LoadingTodo(taskCache, tasksResponse)
 //
 //

 if (taskCache === null) {
  return (<main>
   <Error />
  </main>)
 }

 return (
  <main className='todo-main' onInputCapture={() => {
   const localChanges = createTasksCollection()
   storeInLocal(localChanges, "tasks")
  }}>
   {days.map(day => (
    // TOFIX en el backedn en vez de madnar un objeto vacio mandar uno ya con cosas
    // day, setTaskCache, addTask, tasks
    <TaskContainer day={day} setTaskCache={setTaskCache} addTask={addTask} tasks={taskCache}>
     {
      taskCache[day].map((task, index) => (
       <Task day={day} dragTask={dragTask} draggedOverTask={draggedOverTask} taskCache={taskCache} setTaskCache={setTaskCache} index={index} is_done={task.is_done} task_description={task.task_description} is_deleted={task.is_deleted} handleSort={handleSort} />
      ))
     }
    </TaskContainer>
   ))}
   <div className='save-container'>
    <button type="button" className='cloud-button' onClick={() => { setTaskCacheUpdate(createTasksCollection()) }}>Guardar en la nube</button>
    {message}
   </div>
  </main>
 )
}

// <div className='todo-container' >
//     {(days).map(day => {
//      if (taskCache !== undefined && taskCache[day] !== undefined) {
//       //
//       const tasksInsideContainer = (taskCache[day] as Array<any>).map((task, index) => {
//        if (!task.is_deleted) {
//         return (
//          <Task day={day} taskCache={taskCache} setTaskCache={setTaskCache} index={index} is_done={task.is_done} task_description={task.task_description} is_deleted={task.is_deleted} />
//         )
//        }
//       })
//       //
//       return (
//        <TaskContainer day={day} setTaskCache={setTaskCache} addTask={addTask} tasks={tasksInsideContainer} />
//       )
//      }
//      //
//      else {
//       const tasks = <Task day={day} taskCache={taskCache} setCache={setTaskCache} is_done={false} task_description={""} />
//       return (
//        <TaskContainer day={day} taskCache={taskCache} setTaskCache={setTaskCache} addTask={addTask} tasks={tasks} />
//       )
//      }
//     }
//     )}
//    </div>
