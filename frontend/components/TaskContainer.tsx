import React from 'react'
import AddTaskButton from './AddTaskButton'

// day, setTaskCache, addTask, tasks
export default function TaskContainer(props) {
 const { day, setTaskCache, addTask, tasks } = props
 return (
  <div className='task-container'>
   <span className='day' data-day={day}>
    {day}
    <AddTaskButton properties={{ day, setTaskCache, addTask }} />
   </span>
   {tasks}
  </div>
 )
}

