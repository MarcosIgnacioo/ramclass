import React from 'react'

export default function Empty() {
 return (
  <div className='no assigment'>
   <h1 className='class_subject'>No hay nada aquí</h1>
   <div className='date'>
    <div className='date-format'>
     <span></span>
     <span className='day'></span>
     <span className='month'></span>
     <span className='year'></span>
    </div>
    <div className='date-hour'>
     <span></span>
     <span className='hour'></span>
    </div>
   </div>
   <div className='assigment-info-container'>
    <div className='assigment-info'>
     <span className='title'></span>
     <a className='link' href=''>Quizás mas tarde...</a>
    </div>
   </div>
  </div>
 )
}

