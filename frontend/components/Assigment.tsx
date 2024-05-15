import React from 'react'

export default function Assigment({ class_subject, date, link, title }) {
 if (!title) {
  title = "No hay nada que hacer"
 }
 return (
  <div className='assigment'>
   <h1 className='class_subject'>{class_subject}</h1>
   <div className='date'>
    <div className='date-format'>
     <span>Fecha de entrega: </span>
     <span className='day'>{date.day}/</span>
     <span className='month'>{date.month}/</span>
     <span className='year'>{date.year}</span>
    </div>
    <div className='date-hour'>
     <span>A las: </span>
     <span className='hour'>{date.hour}</span>
    </div>
   </div>
   <div>
    <br />
    <div className='assigment-info'>
     <span className='title'>{title}</span>
     <br />
     <a className='link' href={link}>Ver tarea</a>
    </div>
   </div>
  </div>
 )
}

