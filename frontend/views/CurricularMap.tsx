import React, { useState } from 'react'
import getUser, { LSK, checkBothCache } from '../functions/store'
import { useLocationContext, useLocationUpdateContext, useUser } from '../components/UserContext'
import useLogin from '../functions/useLogin'
import Subject from '../components/Subject'
import Error from '../components/Error'
import SubjectFilters from '../components/SubjectFilters'
import useLocationEffect from '../functions/effects/useLocationEffect'
import UserData from '../classes/UserData'
import useCurricularMap from '../functions/useCurricularMap'
import { filterSubjects } from '../functions/filterSubjects'
import State from '../components/State'
import updateCurrentLocation from '../functions/location'

export const CurricularMap = () => {
 const [userCredentials, setUserCredentials] = useState<UserData | null>(null);
 const [subjectName, setSubjectName] = useState("")
 const [semester, setSemester] = useState(0)

 updateCurrentLocation()

 const userLocal = (useUser().username == "") ? getUser() : useUser()
 const user = useUser()
 const response = useLogin(user)

 const curricularMapFetch = useCurricularMap(userCredentials)
 let curricularMap = checkBothCache(response, LSK.CurricularMap) as Object[]

 if (curricularMap != null) {
  curricularMap = filterSubjects(curricularMap, subjectName, semester)
 }

 let curricularMapSubjects = State({ fetchedData: curricularMapFetch, cache: curricularMap, nameSpace: "curricular_map", Container: Subject, isFiltered: true })

 const gpa = checkBothCache(response, LSK.GPA)

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
     <h1>Promedio general: {gpa.gpa} </h1>
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
