import React from 'react'
import { SemestersMap } from '../globals/globals'

export default function Subject({ grade, group, period, semester, state, subject_name, teacher, turn, type }) {
 const semesterTag = (semester === 0) ? "Optativa" : "Semestre:"
 const semesterNo = (semester === 0) ? "" : semester
 return (
  <div className={`subject ${SemestersMap[semester]}`}>
   <div className='subject-data'>
    <div className='semester-container'>
     <div>
      <span>{semesterTag}</span>
      <span className='semester'>
       {semesterNo}
      </span>
     </div>
     <div>
      <span>Grupo:</span>
      <span className='group'>
       {group}
      </span>
     </div>
     <div>
      <span>Turno:</span>
      <span className='turn'>{turn} </span>
     </div>
    </div>
    <br />
    <div className='subject-result'>
     <div>
      <span>Materia:</span>
      <br />
      <span className='subject-name'>
       {subject_name}
      </span>
     </div>
     <div>
      <span>Calificación:</span>
      <br />
      <span className='grade'>
       {grade}
      </span>
     </div>

    </div>
    <br />
   </div>
   <div className='dropdown'>
    <span className='teacher'>{teacher} </span>
    <span className='period'>{period} </span>
    <span className='state'>{state} </span>
    <span className='type'>{type}</span>
   </div>
  </div>
 )
}

