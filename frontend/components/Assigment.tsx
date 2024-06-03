import React from 'react'

export default function Assigment({ class_subject, date, link, title }) {
 const showDate = ((date.day === ""))
 const showMonth = ((date.month === ""))
 if (!title) {
  title = "No hay nada que hacer"
 }
 return (
  <div className='assigment'>
   <h1 className='class_subject'>{class_subject}</h1>
   <div className='date'>
    <div className='date-format' hidden={showDate}>
     <span>Fecha de entrega: </span>
     <span className='day-assigment'>{date.day} </span>
     <span className='month' hidden={showMonth}>/ {date.month} /</span>
     <span className='year'> {date.year}</span>
    </div>
    <div className='date-hour' hidden={showDate}>
     <span>A las: </span>
     <span className='hour'>{date.hour}</span>
    </div>
   </div>
   <div className='assigment-info-container'>
    <div className='assigment-info'>
     <p className='title'>{title}</p>
     <a className='link' href={link}>Ver tarea</a>
    </div>
   </div>
  </div>
 )
}

