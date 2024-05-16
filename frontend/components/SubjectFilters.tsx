import React from 'react'
import { semesters } from '../globals/globals'

export default function SubjectFilters(semester: number, setSemester: React.Dispatch<React.SetStateAction<number>>, subjectName: string, setSubjectName: React.Dispatch<React.SetStateAction<string>>) {
 return (
  <div className='subjects-filter'>
   <h1>Busca la materia</h1>
   <input type="" name="" className='subject-filter' value={subjectName} placeholder='Buscar materia' onChange={(e) => {
    setSubjectName(e.target.value)
   }} />
   <br />
   <div id='semester-filter'>
    <span>Semestre:</span>
    <select className='semester-filter' value={semester} onChange={(e) => setSemester(parseInt(e.target.value))}>
     {semesters.map((semesterText, index) => (<option value={index}>{semesterText}</option>))}
    </select>
   </div>
  </div>
 )
}

