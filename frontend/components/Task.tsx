import React, { useState } from 'react'
import DeleteButton from './DeleteButton'

export default function Task(props) {
 const { dragTask, draggedOverTask, handleSort, task_description, is_done, index, day } = props
 const [task, setTask] = useState(task_description)
 const [isDone, setIsDone] = useState(is_done)
 const deleteButton = (task !== "") ? <DeleteButton props={props} /> : ""

 return (
  <div className='task'>
   <h5>{day}</h5>
   <div className='task-header'

    onDragStart={() => {
     dragTask.current = { day: day, index: index }
     console.log("dragging", dragTask.current)
    }}

    onDragEnter={() => {
     draggedOverTask.current = { day: day, index: index }
     console.log("over:", draggedOverTask.current)
    }}
    onDragEnd={handleSort}

    draggable={true}>
    :::
   </div>
   <div className='task-content'>
    <input className='check-button' checked={isDone} onChange={(e) => setIsDone(e.target.checked)} type="checkbox" name="" value="" />
    <input className='task-text' type="text" name="" value={task} onChange={(e) => {
     setTask(e.target.value)
    }} />
    {deleteButton}
   </div>
  </div>
 )
}

