import React, { useRef, useState } from 'react'
import updateCurrentLocation from '../functions/location'
import "../css/todo.css"
import { TaskClass } from '../classes/Tasks'
import Task from '../components/Task'
import getUser, { getCacheOf, storeInLocal } from '../functions/store'
import { useUser } from '../components/UserContext'
import UserData from '../classes/UserData'
import useTasks from '../functions/useTasks'
import useGetTasks from '../functions/useGetTasks'
import Error from '../components/Error'
import Loading from '../components/Loading'
import Message from '../components/Message'
import Success from '../components/Success'

const days = ["Lunes", "Martes", "Miércoles", "Jueves", "Viernes", "Sábado", "Domingo"]

export default function Todo() {
 updateCurrentLocation()
 const userLocal = ((useUser().username == "") ? getUser() : useUser() as UserData)
 const identifier = userLocal?.username as string
 let [taskCache, setTaskCache] = useState(getCacheOf("tasks"))
 const [taskCacheUpdate, setTaskCacheUpdate]: any = useState()
 let successMessage = true
 let ErrorMessage = true
 const successMessageRef = useRef()
 const errorMessageRef = useRef()
 const tasksResponse = useGetTasks(identifier, taskCache)
 const savingTask = useTasks(taskCacheUpdate, identifier)


 startDraggable()

 if (taskCache === null && tasksResponse.isError) {
  return (<main>
   <Error />
  </main>)
 }
 //
 if (taskCache === null && tasksResponse.isLoading) {
  return (<main>
   <Loading />
  </main>)
 }
 //
 if (taskCache === null && tasksResponse.isSuccess) {
  taskCache = tasksResponse.data
 }
 //
 if (taskCache === null) {
  return (<main>
   <Error />
  </main>)
 }
 //
 if (savingTask.isSuccess) {
  console.log(successMessageRef)
  console.log(successMessageRef.current)
  console.log(errorMessageRef.current)
 }
 //
 if (savingTask.isError) {
  console.log(errorMessageRef.current)
 }
 //
 return (
  <div onInputCapture={() => {
   const localChanges = createTasksCollection()
   storeInLocal(localChanges, "tasks")
  }}>
   <div className='todo-header' >
    {(days).map(day => {
     if (taskCache !== undefined && taskCache[day] !== undefined) {
      //
      const tasksInsideContainer = (taskCache[day] as Array<any>).map((task, index) => {
       if (!task.is_deleted) {
        return (
         <Task day={day} setCache={setTaskCache} cache={taskCache} index={index} is_done={task.is_done} task_description={task.task_description} is_deleted={task.is_deleted} />
        )
       }
      })
      //
      return (
       <div className='task-container'>
        <span className='day'>
         {day}
        </span>
        {tasksInsideContainer}
       </div>)
     }
     //
     else {
      return (
       <div className='task-container'>
        <span className='day'>
         {day}
        </span>
        <Task cache={taskCache} is_done={false} task_description={""} />
       </div>)
     }
    }

    )}
   </div>
   <div>
    <button type="button" className='cloud-button' onClick={() => { setTaskCacheUpdate(createTasksCollection()) }}>Guardar en la nube</button>
    <Message ref={successMessageRef} message="Se guardó con éxito" isHidden={false} class="message success" />
    <Message ref={errorMessageRef} message="Ocurrió un error innesperado" isHidden={false} class="message error" />
   </div>
  </div>
 )
}



function addTask(taskCache: object, day: string, setTaskCache) {
 const newTaskCache = { ...taskCache }
 newTaskCache[day].push({ is_done: false, task_description: "" })
 storeInLocal(newTaskCache, "tasks")
 setTaskCache(newTaskCache)
}


function createTasksCollection() {
 const tasksContainers = document.querySelectorAll('.task-container')
 const tasks: Object = {}
 tasksContainers.forEach(taskContainer => {
  const tasksInsideContainer = taskContainer.querySelectorAll(".task")
  const day = (taskContainer.querySelector(".day") as HTMLSpanElement).innerText.replace("+", "")
  tasksInsideContainer.forEach(task => {
   const taskDescription = ((task.querySelector(".task-text") as HTMLInputElement) || { value: "" }).value
   const isDone = ((task.querySelector(".check-button") as HTMLInputElement) || { checked: false }).checked
   if (tasks[day] === undefined) tasks[day] = []
   tasks[day].push(new TaskClass(taskDescription, isDone))
  })
 })
 return tasks
}

function startDraggable() {
 const tasksDraggable = document.querySelectorAll(".task")
 const containers = document.querySelectorAll(".task-container")
 tasksDraggable.forEach(task => {
  task.addEventListener('dragstart', () => {
   task.classList.add("dragging")
  })
  task.addEventListener('dragend', () => {
   task.classList.remove("dragging")
  })
 })
 containers.forEach(container => {
  container.addEventListener('dragover', (e) => {
   e.preventDefault()
   const draggingTask = document.querySelector(".dragging")
   const closest = getDragAfterElement(container, e.clientY)

   if (closest === null) {
    container.appendChild(draggingTask!)
   } else {
    container.insertBefore(draggingTask!, closest)
   }
   const tasksCollection = createTasksCollection()
   storeInLocal(tasksCollection, "tasks")
  })
 })
}

function getDragAfterElement(container, y) {
 const tasksInContainer = [...container.querySelectorAll(".task:not(.dragging)")]
 return tasksInContainer.reduce((closest, task) => {
  const box = task.getBoundingClientRect()
  const offset = y - box.top - (box.height / 2)
  if (offset < 0 && offset > closest.offset) {
   return { offset: offset, element: task }
  }
  else { return closest }
 }, { offset: Number.NEGATIVE_INFINITY }).element
}
