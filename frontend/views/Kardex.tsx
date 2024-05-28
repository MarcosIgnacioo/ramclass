import React, { useState } from 'react'
import { useUser } from '../components/UserContext';
import useLogin from '../functions/useLogin';
import getUser, { LSK, checkBothCache } from '../functions/store';
import Subject from '../components/Subject';
import SubjectFilters from '../components/SubjectFilters';
import UserData from '../classes/UserData';
import useKardex from '../functions/useKardex';
import State from '../components/State';
import { filterSubjects } from '../functions/filterSubjects';
import SignIn from './SignIn';
import updateCurrentLocation from '../functions/location';
import Title from '../components/Title';

export default function Kardex() {
 const [userCredentials, setUserCredentials] = useState<UserData | null>(null);
 const [subjectName, setSubjectName] = useState("")
 const [semester, setSemester] = useState(0)

 updateCurrentLocation()

 const userLocal = (useUser().username == "") ? getUser() : useUser()
 if (!userLocal) return (<SignIn />)

 const user = useUser()

 const response = useLogin(user)

 const kardexFetch = useKardex(userCredentials)
 let kardex = checkBothCache(response, LSK.Kardex) as Object[]

 if (kardex != null) {
  kardex = filterSubjects(kardex, subjectName, semester)
 }

 let kardexSubjects = State({ fetchedData: kardexFetch, cache: kardex, nameSpace: "kardex", Container: Subject, isFiltered: true })

 let gpa = checkBothCache(response, LSK.GPA)

 let gpaState = State({ fetchedData: kardexFetch, cache: gpa, nameSpace: "gpa", Container: null, isFiltered: true })

 gpaState = (gpaState === null) ? "Error, actualiza tu Kardex" : gpaState

 if (kardexSubjects == null) return (<main>
  <div className='warn'>
   <h1 className='warn-header'>Advertencia</h1>
   <div className='warn-content'>
    <h1>No se cargó correctamente tu kardex,<br /> presiona aqui para hacerlo</h1>
    <button type="button" onClick={() => {
     setUserCredentials(userLocal)
    }} className='refetch kardex'>Cargar Kardex</button>
   </div>
  </div>
 </main>)

 const contentClass = (kardexSubjects.props.children !== undefined) ? "subjects-container" : ""


 return (
  <main className='kardex-container'>
   <Title title='Kardex' to='#' />
   <div className='kardex-top'>
    {SubjectFilters(semester, setSemester, subjectName, setSubjectName)}
    <div className='refetch-gpa'>
     <h1 className='gpa'>Promedio general: {gpaState} </h1>
     <button type="button" onClick={() => {
      setUserCredentials(userLocal)
     }} className='refetch kardex'>Actualizar Kardex</button>
    </div>
   </div>
   <div className={contentClass}>
    {kardexSubjects}
   </div>
  </main>
 )
}
