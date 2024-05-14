import React, { useState } from 'react'
import getUser, { LSK, storeInLocal } from '../functions/store'
import { checkCache } from '../functions/checkLogin'
import Assigment from '../components/Assigment'
import AssigmentClass from '../classes/Assigment'
import { refetchMoodle } from '../functions/buttonsFunctions/refetchMoodle'
import UserData from '../classes/UserData'
import useMoodle from '../functions/useMoodle'
import { useUser } from '../components/UserContext'
import useLogin from '../functions/useLogin'
import useClassRoom from '../functions/useClassRoom'

export default function Home() {

 const [userClassRoomParams, setUserClassRoomParams] = useState<UserData | null>(null);
 const [userMoodleParams, setUserMoodleParams] = useState<UserData | null>(null);
 const [userContext, setUserContext] = useState<UserData | null>(null);
 const moodleFetch = useMoodle(userMoodleParams)
 const classRoomFetch = useClassRoom(userClassRoomParams)
 const reactQueryCache = useLogin(userContext).data
 const usr = useUser()


 let classRoom = checkCache(LSK.Classroom) as Object[]
 let moodle = checkCache(LSK.Moodle) as Object[]


 if (moodleFetch.isLoading) {
 }

 if (moodleFetch.isSuccess) {
  setUserContext(usr)
  storeInLocal(moodleFetch.data.moodle, "moodle")
  moodle = moodleFetch.data.moodle
  // Probablemente querramos activar el refetch en estas partes de nuevo
  // setUserParams(null)
 }

 if (classRoomFetch.isLoading) {
 }

 if (classRoomFetch.isSuccess) {
  setUserContext(usr)
  storeInLocal(classRoomFetch.data.classroom, "classroom")
  classRoom = classRoomFetch.data.classroom
  // Probablemente querramos activar el refetch en estas partes de nuevo
  // setUserParams(null)
 }

 if (reactQueryCache !== undefined) {
  reactQueryCache.moodle = moodle
  reactQueryCache.classroom = classRoom
 }

 let userLocal: UserData | null
 userLocal = (useUser().username == "") ? getUser() : useUser()

 if (!userLocal) return (<h1 className='alert'>No has iniciado sesión</h1>)


 // Solucion rapida a lo del refetch, haccer que el boton primero actualice la pagina y luego llene la cache del useQuery
 return (
  <main className='home'>
   <div className='moodle'>
    <button className='refetch moodle' onClick={() => {
     refetchMoodle(userLocal, setUserMoodleParams)
    }} type="button">fetch</button>
    {moodle.map(assigment => (<Assigment{...assigment as AssigmentClass} />))}
   </div>
   <div className='classroom'>
    <button className='refetch  classroom' onClick={() => {
     refetchMoodle(userLocal, setUserClassRoomParams)
    }} type="button">fetch</button>
    {classRoom.map(assigment => (<Assigment{...assigment as AssigmentClass} />))}
   </div>
  </main>
 )
}

