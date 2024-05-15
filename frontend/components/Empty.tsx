import React from 'react'

export default function Empty() {
 return (
  <div className='no assigment'>
   <h1 className='class_subject'>No hay nada que hacer</h1>
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
   <div>
    <br />
    <div className='assigment-info'>
     <span className='title'></span>
     <a className='link' href=''>No hay tarea!</a>
    </div>
   </div>
  </div>
 )
}

