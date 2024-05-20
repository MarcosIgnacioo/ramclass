import React, { ButtonHTMLAttributes, DragEvent, DragEventHandler, useEffect, useState } from 'react'
import updateCurrentLocation from '../functions/location'
import "../css/todo.css"
import { Navigate, useLocation, useNavigation } from 'react-router-dom'
import Tasks, { TaskClass } from '../classes/Tasks'
import Task from '../components/Task'
import getUser, { checkBothCache, getCacheOf, storeInLocal } from '../functions/store'
import { useUser } from '../components/UserContext'
import UserData from '../classes/UserData'
import useTasks from '../functions/useTasks'
import useGetTasks from '../functions/useGetTasks'

const days = ["Lunes", "Martes", "Miércoles", "Jueves", "Viernes", "Sábado", "Domingo"]

export default function Todo() {
 checkBothCache
 const [taskCache, setTaskCache] = useState(getCacheOf("tasks"))
 console.log(taskCache)
 updateCurrentLocation()
 const userLocal = ((useUser().username == "") ? getUser() : useUser() as UserData)?.username as string
 useTasks(taskCache as Object, userLocal)
 useGetTasks(userLocal)
 startDraggable(userLocal)

 return (
  <div onInputCapture={() => storeInLocal(createTasksCollection(), "tasks")}>
   <div className='todo-header' >
    {(days).map(day => {
     if (taskCache !== null && taskCache[day] !== undefined) {
      //
      const tasksInsideContainer = taskCache[day].map((task, index) => {
       if (!task.is_deleted) {
        // This makes completly sense i swear !!!
        return (<div>
         <Task day={day} setCache={setTaskCache} cache={taskCache} index={index} is_done={task.is_done} task_description={task.task_description} is_deleted={task.is_deleted} />
        </div>)
       }
      })
      //
      return (
       <div className='task-container'>
        <span className='day'>
         {day}
        </span>
        {tasksInsideContainer}
        <button type="button" onClick={() => addTask(taskCache as Object, day, setTaskCache)}>+</button>
       </div>)
     }
     //
     else {
      return (<div className='task-container'>
       <span className='day'>{day}</span>
       <Task cache={taskCache} is_done={false} task_description={""} />
       <button type="button" onClick={() => addTask(taskCache as Object, day, setTaskCache)}>+</button>
      </div>)
     }
    }
    )}
   </div>
   <button type="button" className='cloud-button' onClick={() => { setTaskCache(createTasksCollection()) }}>Guardar en la nube</button>
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
  const day = (taskContainer.querySelector(".day") as HTMLSpanElement).innerText
  tasksInsideContainer.forEach(task => {
   const taskDescription = ((task.querySelector(".task-text") as HTMLInputElement) || { value: "" }).value
   const isDone = ((task.querySelector(".check-button") as HTMLInputElement) || { checked: false }).checked
   if (tasks[day] === undefined) tasks[day] = []
   tasks[day].push(new TaskClass(taskDescription, isDone))
  })
 })
 return tasks
}

function startDraggable(userLocal) {
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
   const tasksCollection = createTasksCollection(userLocal)
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
