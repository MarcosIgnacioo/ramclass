import React, { useState } from 'react'
import { useLocationContext, useLocationUpdateContext, useUser } from '../components/UserContext';
import useLogin from '../functions/useLogin';
import getUser, { LSK, checkBothCache } from '../functions/store';
import SubjectClass from '../classes/Subject.ts';
import Subject from '../components/Subject.tsx';
import SubjectFilters from '../components/SubjectFilters.tsx';
import UserData from '../classes/UserData.ts';
import useKardex from '../functions/useKardex.ts';
import State from '../components/State.tsx';
import useLocationEffect from '../functions/effects/useLocationEffect.ts';

export default function Kardex() {
 const [userCredentials, setUserCredentials] = useState<UserData | null>(null);
 const [subjectName, setSubjectName] = useState("")
 const [semester, setSemester] = useState(0)


 const locationUpdate = useLocationUpdateContext()
 locationUpdate(window.location.pathname)
 const currentLocation = useLocationContext()
 useLocationEffect(currentLocation)

 const userLocal = (useUser().username == "") ? getUser() : useUser()
 const user = useUser()
 const response = useLogin(user)
 let kardex = checkBothCache(response, LSK.Kardex) as Object[]
 const gpa = checkBothCache(response, LSK.GPA)

 const kardexFetch = useKardex(userCredentials)

 let kardexSubjects = State({ fetchedData: kardexFetch, cache: kardex, nameSpace: "kardex", Container: Subject, returnArray: true }) as Object[]

 kardexSubjects = kardex.filter(subject => (((subject as SubjectClass).subject_name.toLowerCase()).includes(subjectName.toLowerCase())))


 kardexSubjects = (semester === 0) ? kardex : kardex.filter(subject => (((subject as SubjectClass).semester) === semester))

 if (!userLocal) return (<h1>Esperate wey</h1>)

 return (
  <main className='kardex-container'>
   {SubjectFilters(semester, setSemester, subjectName, setSubjectName)}
   <div>
    <h1>Promedio general: {gpa.gpa} </h1>
    <button type="button" onClick={() => {
     setUserCredentials(userLocal)
    }} className='refetch kardex'>Actualizar Kardex</button>
   </div>
   <div className='subjects-container'>
    {kardexSubjects.map(subject => (<Subject {...subject as SubjectClass} />))}
   </div>
  </main>
 )
}


// <h1>Busca la materia</h1>
// <select className='semester-filter' value={semester} onChange={(e) => setSemester(parseInt(e.target.value))}>
//  {semesters.map((semesterText, index) => (<option value={index}>{semesterText}</option>))}
// </select>
// <input type="" name="" className='subject-filter' value={subjectName} placeholder='Buscar materia' onChange={(e) => {
//  setSubjectName(e.target.value)
// }} />
