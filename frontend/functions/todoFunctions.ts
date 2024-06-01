import { TaskClass } from "../classes/Tasks"
import { getCacheOf, storeInLocal } from "./store"

export function addTask(day: string, setTaskCache) {
 const newTaskCache = { ...getCacheOf("tasks") as Object }
 if (newTaskCache[day] === undefined || newTaskCache[day] === null) newTaskCache[day] = []
 newTaskCache[day].push({ is_done: false, task_description: "" })
 storeInLocal(newTaskCache, "tasks")
 setTaskCache(newTaskCache)
}


export function createTasksCollection() {
 const tasksContainers = document.querySelectorAll('.task-container')
 const tasks: Object = {}
 tasksContainers.forEach(taskContainer => {
  const tasksInsideContainer = taskContainer.querySelectorAll(".task")
  const day = (taskContainer.querySelector(".day") as HTMLSpanElement).dataset.day || ""
  tasksInsideContainer.forEach(task => {
   const taskDescription = ((task.querySelector(".task-text") as HTMLInputElement) || { value: "" }).value
   const isDone = ((task.querySelector(".check-button") as HTMLInputElement) || { checked: false }).checked
   if (tasks[day] === undefined) tasks[day] = []
   tasks[day].push(new TaskClass(taskDescription, isDone))
  })
 })
 return tasks
}

export function startDraggable() {
 const tasksDraggable = document.querySelectorAll(".task")
 const containers = document.querySelectorAll(".task-container")
 tasksDraggable.forEach(task => {
  task.classList.add("test")
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
    console.log(draggingTask)
    container.insertBefore(draggingTask!, closest)
   }
   const tasksCollection = createTasksCollection()
   storeInLocal(tasksCollection, "tasks")
  })
 })
}

export function getDragAfterElement(container, y) {
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

export function deleteTask({ taskCache, day, index, setTaskCache }) {
 taskCache[day][index].is_deleted = true
 const newCache = { ...taskCache }
 storeInLocal(newCache, "tasks")
 setTaskCache(newCache)
}
