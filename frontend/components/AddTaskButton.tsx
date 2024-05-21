import React from 'react'

export default function AssTaskButton(props) {
 const { taskCache, day, setTaskCache, addTask } = props.properties
 return (
  <button type="button" onClick={() => addTask(taskCache as Object, day, setTaskCache)}>+</button>
 )
}

