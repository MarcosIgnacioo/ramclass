import React from 'react'

export default function Subject({ grade, group, period, semester, state, subject_name, teacher, turn, type }) {
 return (
  <div className='subject'>
   <div className='subject-data'>
    <span className='subject-name'>
     {semester}
    </span>
    <span className='subject-name'>
     {subject_name}
    </span>
    <span className='grade'>
     {grade}
    </span>
   </div>
   <div className='dropdown'>
    <span className='teacher'>{teacher}</span>
    <span className='group'>{group}</span>
    <span className='period'>{period}</span>
    <span className='state'>{state}</span>
    <span className='turn'>{turn}</span>
    <span className='type'>{type}</span>
   </div>
  </div>
 )
}

