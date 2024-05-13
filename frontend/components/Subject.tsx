import React from 'react'

export default function Subject({ grade, group, period, semester, state, subject_name, teacher, turn, type }) {
 return (
  <div className='subject'>
   <h1>
    {subject_name}
   </h1>
  </div>
 )
}

