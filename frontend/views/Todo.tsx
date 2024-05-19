import React, { DragEvent, DragEventHandler, useEffect, useState } from 'react'
import updateCurrentLocation from '../functions/location'
import "../css/todo.css"
import { Navigate, useLocation, useNavigation } from 'react-router-dom'
import Tasks, { TaskClass } from '../classes/Tasks'
import Task from '../components/Task'

const days = ["Lunes", "Martes", "Miércoles", "Jueves", "Viernes", "Sábado", "Domingo"]
export default function Todo() {
 const [text, setText] = React.useState('');

 updateCurrentLocation()

 startDraggable()
 return (
  <div>
   <div className='todo-header' >
    {days.map(day => (
     <div className='task-container'>
      <span className='day'>{day}</span>
      <Task />
     </div>))}
   </div>
   <button type="button" onClick={() => { createTasksCollection("marcos") }}>AAAA</button>
  </div>
 )
}

function createTasksCollection(identifier: string) {
 const tasksContainers = document.querySelectorAll('.task-container')
 const caca: TaskClass[] = []
 tasksContainers.forEach(taskContainer => {
  const taskDescription = (taskContainer.querySelector(".task-text") as HTMLInputElement).value
  const day = (taskContainer.querySelector(".day") as HTMLSpanElement).innerText
  const isDone = (taskContainer.querySelector(".check-button") as HTMLInputElement).checked
  caca.push(new TaskClass(day, taskDescription, isDone))
 })
 console.log(caca)
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
