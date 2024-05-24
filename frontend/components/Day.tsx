import React from 'react'

export default function Day(props) {
 const { day, events, type } = props
 return (
  <div className={`day-calendar ${type}`}>
   {day}
   {events.map(event => (<h1 className={`event-text ${type}`}>{event}</h1>))}
  </div>
 )
}

