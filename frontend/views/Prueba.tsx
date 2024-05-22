import React, { useState, useRef } from "react"
import Person from "../components/Person"
import Task from "../components/Task"
import { TaskDrag, TaskList } from "../classes/Tasks"
import tasks from "../functions/fetchs/tasks";
import "../css/todo.css"

export default function Home() {

 const initialTasks: TaskList = {
  "Lunes": [{ "task_description": "esta es la task 1 del lunes", "is_done": false, "is_deleted": false }],
  "Martes": [{ "task_description": "esta es la task 1 del martes", "is_done": false, "is_deleted": false }],
  "Miércoles": [{ "task_description": "culooo", "is_done": false, "is_deleted": false }],
  "Jueves": [{ "task_description": "hola", "is_done": false, "is_deleted": false }, { "task_description": "b", "is_done": false, "is_deleted": false }],
  "Viernes": [{ "task_description": "viernes gamer", "is_done": false, "is_deleted": false }],
  "Sábado": [{ "task_description": "sabadazo", "is_done": false, "is_deleted": false }, { "task_description": "", "is_done": false, "is_deleted": false }],
  "Domingo": [{ "task_description": "domingowo", "is_done": false, "is_deleted": false }]
 };

 const [taskCache, setTaskCache] = useState<TaskList>(initialTasks)
 console.log(taskCache)

 const days = ["Lunes", "Martes", "Miércoles", "Jueves", "Viernes", "Sábado", "Domingo"]
 const dragTask = useRef<TaskDrag>({ day: "", index: 0 })
 const draggedOverTask = useRef<TaskDrag>({ day: "", index: 0 })

 function handleSort() {
  console.log("SORTING")
  console.log("dragging", dragTask.current)
  console.log("over", draggedOverTask.current)
  const tasksClone: TaskList = { ...taskCache }
  const temp = tasksClone[dragTask.current.day][dragTask.current.index]
  tasksClone[dragTask.current.day][dragTask.current.index] = tasksClone[draggedOverTask.current.day][draggedOverTask.current.index]
  tasksClone[draggedOverTask.current.day][draggedOverTask.current.index]
   = temp
  setTaskCache(tasksClone)
 }

 return (
  <main className="flex min-h-screen flex-col items-center space-y-4">
   {days.map(day => (
    taskCache[day].map((task, index) => (
     <Task day={day} dragTask={dragTask} draggedOverTask={draggedOverTask} taskCache={taskCache} setTaskCache={setTaskCache} index={index} is_done={task.is_done} task_description={task.task_description} is_deleted={task.is_deleted} handleSort={handleSort} />
    ))
   ))}
  </main>
 )
}
