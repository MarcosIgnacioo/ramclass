import React from 'react'
import AddTaskButton from './AddTaskButton'

export default function TaskContainer(props) {
 const { day, tasks, taskCache, setTaskCache, addTask } = props
 return (
  <div className='task-container'>
   <span className='day' data-day={day}>
    {day}
    <AddTaskButton properties={{ taskCache, day, setTaskCache, addTask }} />
   </span>
   {tasks}
  </div>
 )
}

