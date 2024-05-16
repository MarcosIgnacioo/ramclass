import React, { useState } from 'react'
import getUser, { LSK, checkBothCache } from '../functions/store'
import { useLocationContext, useLocationUpdateContext, useUser } from '../components/UserContext'
import useLogin from '../functions/useLogin'
import Subject from '../components/Subject'
import SubjectClass from '../classes/Subject'
import Error from '../components/Error'
import SubjectFilters from '../components/SubjectFilters'
import useLocationEffect from '../functions/effects/useLocationEffect'

export const CurricularMap = () => {
 const [subjectName, setSubjectName] = useState("")
 const [semester, setSemester] = useState(0)

 const locationUpdate = useLocationUpdateContext()
 locationUpdate(window.location.pathname)
 const currentLocation = useLocationContext()
 useLocationEffect(currentLocation)

 const userLocal = getUser()
 const user = useUser()
 const response = useLogin(user)
 let curricularMap = checkBothCache(response, LSK.CurricularMap) as Object[]

 curricularMap = curricularMap.filter(subject => (((subject as SubjectClass).subject_name.toLowerCase()).includes(subjectName.toLowerCase())))

 curricularMap = (semester === 0) ? curricularMap : curricularMap.filter(subject => (((subject as SubjectClass).semester) === semester))

 if (!userLocal) return (<Error />)

 return (
  <main className='curricular-map-conatainer' >
   {SubjectFilters(semester, setSemester, subjectName, setSubjectName)}
   <div className='subjects-container'>
    {curricularMap.map(subject => (<Subject{...subject as SubjectClass} />))}
   </div>
  </main>
 )
}
