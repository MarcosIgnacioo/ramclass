import React, { useRef, useState } from 'react'
import updateCurrentLocation from '../functions/location'
import "../css/todo.css"
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
import SignIn from './SignIn'
import Title from '../components/Title'

const days = ["Lunes", "Martes", "Miércoles", "Jueves", "Viernes", "Sábado", "Domingo"]

export default function Todo() {
 updateCurrentLocation()
 const userLocal = ((useUser().username == "") ? getUser() : useUser() as UserData)
 if (!userLocal) return (<SignIn />)

 const identifier = userLocal?.username as string
 let [taskCache, setTaskCache] = useState(getCacheOf("tasks"))
 const [taskCacheUpdate, setTaskCacheUpdate]: any = useState()
 const successMessageRef = useRef()
 const errorMessageRef = useRef()
 const tasksResponse = useGetTasks(identifier, taskCache)
 const savingTask = useTasks(taskCacheUpdate, identifier)
 let message: React.JSX.Element = <div></div>

 if (taskCache === null && tasksResponse.isError) {
  return (<main>
   <Error error={"Ha ocurrido un error inesperado, puedes volver a intentar lo que querías hacer o actualizar la página."} refreshButton={true} />
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
   <Error error={"Ha ocurrido un error inesperado, puedes volver a intentar lo que querías hacer o actualizar la página."} refreshButton={true} />
  </main>)
 }
 //

 if (savingTask.isSuccess) {
  message = <Message ref={successMessageRef} message="Se guardó con éxito" class="message success" />
 }

 //
 if (savingTask.isError) {
  message = <Message ref={errorMessageRef} message="Ocurrió un error innesperado" class="message error" />
 }
 //

 return (
  <main className='todo-main' onInputCapture={() => {
   const localChanges = createTasksCollection()
   storeInLocal(localChanges, "tasks")
  }}>
   <Title title='To Do' to='#' />
   <div className='todo-container' >
    {(days).map(day => {
     if ((taskCache !== undefined && taskCache !== null) && taskCache[day] !== undefined) {
      const tasksInsideContainer = (taskCache[day] as Array<any>).map((task, index) => {
       if (!task.is_deleted) {
        return (
         <Task
          day={day}
          taskCache={taskCache}
          setTaskCache={setTaskCache}
          index={index}
          is_done={task.is_done}
          task_description={task.task_description}
          is_deleted={task.is_deleted}
         />
        )
       }
      })
      return (
       <TaskContainer
        day={day}
        setTaskCache={setTaskCache}
        addTask={addTask}
        tasks={tasksInsideContainer}
       />
      )
     }
     else {
      const tasks =
       <Task
        day={day}
        taskCache={taskCache}
        setCache={setTaskCache}
        is_done={false}
        task_description={""}
       />
      return (
       <TaskContainer
        day={day}
        taskCache={taskCache}
        setTaskCache={setTaskCache}
        addTask={addTask}
        tasks={tasks} />
      )
     }
    }
    )}
   </div>
   <div className='save-container'>
    <button type="button" className='cloud-button' onClick={() => { setTaskCacheUpdate(createTasksCollection()) }}>Guardar en la nube</button>
    {message}
   </div>
  </main>
 )
}

