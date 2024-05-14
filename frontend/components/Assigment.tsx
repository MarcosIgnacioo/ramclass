import React from 'react'

export default function Assigment({ class_subject, date, link, title }) {
 return (
  <div className='assigment'>
   <h1 className='class_subject'>{class_subject}</h1>
   <div className='date'>
    <span className='hour'>{date.hour}</span>
    <span className='day'>{date.day}</span>
    <span className='month'>{date.month}</span>
    <span className='year'>{date.year}</span>
   </div>
   <div>
    <span className='title'>{title}</span>
    <a className='link' href={link}>{link}</a>
   </div>
  </div>
 )
}

