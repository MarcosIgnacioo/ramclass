import React, { useState } from 'react'
import getUser, { LSK, checkBothCache } from '../functions/store'
import { useUser } from '../components/UserContext'
import useLogin from '../functions/useLogin'
import Subject from '../components/Subject'
import Error from '../components/Error'
import SubjectFilters from '../components/SubjectFilters'
import UserData from '../classes/UserData'
import useCurricularMap from '../functions/useCurricularMap'
import { filterSubjects } from '../functions/filterSubjects'
import State from '../components/State'
import updateCurrentLocation from '../functions/location'
import SignIn from './SignIn'

export const CurricularMap = () => {
 const [userCredentials, setUserCredentials] = useState<UserData | null>(null);
 const [subjectName, setSubjectName] = useState("")
 const [semester, setSemester] = useState(0)

 updateCurrentLocation()

 const userLocal = (useUser().username == "") ? getUser() : useUser()
 if (!userLocal) return (<SignIn />)
 const user = useUser()
 const response = useLogin(user)

 const curricularMapFetch = useCurricularMap(userCredentials)
 let curricularMap = checkBothCache(response, LSK.CurricularMap) as Object[]

 if (curricularMap != null) {
  curricularMap = filterSubjects(curricularMap, subjectName, semester)
 }

 let curricularMapSubjects = State({ fetchedData: curricularMapFetch, cache: curricularMap, nameSpace: "curricular_map", Container: Subject, isFiltered: true })

 let gpa = checkBothCache(response, LSK.GPA)
 gpa = (gpa === null) ? "Error, actualiza tu Kardex" : gpa

 if (curricularMapSubjects == null) return (<main>
  <div className='warn'>
   <h1 className='warn-header'>Advertencia</h1>
   <div className='warn-content'>
    <h1>No se cargó correctamente tu mapa curricular,<br /> presiona aqui para hacerlo</h1>
    <button type="button" onClick={() => {
     setUserCredentials(userLocal)
    }} className='refetch kardex'>Cargar Kardex</button>
   </div>
  </div>
 </main>)

 if (!userLocal) return (<Error />)

 const contentClass = (curricularMapSubjects.props.children !== undefined) ? "subjects-container" : ""
 return (
  <main className='curricular-map-conatainer' >
   <div className='kardex-top'>
    {SubjectFilters(semester, setSemester, subjectName, setSubjectName)}
    <div>
     <h1 className='gpa'>Promedio general: {gpa} </h1>
     <button type="button" onClick={() => {
      setUserCredentials(userLocal)
     }} className='refetch kardex'>Actualizar Mapa Curricular</button>
    </div>
   </div>
   <div className={contentClass}>
    {curricularMapSubjects}
   </div>
  </main>
 )
}
