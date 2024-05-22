import React from 'react'

export default function AssTaskButton(props) {
 const { day, setTaskCache, addTask } = props.properties
 return (
  <button type="button" onClick={() => addTask(day, setTaskCache)}>+</button>
 )
}

