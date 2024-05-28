import React, { useState } from 'react'
import getUser, { LSK } from '../functions/store'
import { checkCache } from '../functions/checkLogin'
import Assigment from '../components/Assigment'
import UserData from '../classes/UserData'
import useMoodle from '../functions/useMoodle'
import { useUser } from '../components/UserContext'
import useClassRoom from '../functions/useClassRoom'
import State from '../components/State'
import updateCurrentLocation from '../functions/location'
import SignIn from './SignIn'
import Title from '../components/Title'

export default function Home() {

 console.log("we are landing home")
 updateCurrentLocation()


 const [userClassRoomParams, setUserClassRoomParams] = useState<UserData | null>(null);
 const [userMoodleParams, setUserMoodleParams] = useState<UserData | null>(null);
 const moodleFetch = useMoodle(userMoodleParams)
 const classRoomFetch = useClassRoom(userClassRoomParams)

 let classRoom = checkCache(LSK.Classroom) as Object[]
 let moodle = checkCache(LSK.Moodle) as Object[]

 // TODO: Por alguna razon cuando le pico 2 veces, la primera si carga el loading pero la segunda no lo hace pero si hace el fetch

 const moodleAssigments = State({ fetchedData: moodleFetch, cache: moodle, nameSpace: "moodle", Container: Assigment, isFiltered: false })
 const classRoomAssigments = State({ fetchedData: classRoomFetch, cache: classRoom, nameSpace: "classroom", Container: Assigment, isFiltered: false })

 const userLocal = (useUser().username == "") ? getUser() : useUser()
 if (!userLocal) return (<SignIn />)

 return (
  <main className='home'>
   <Title title='Tareas' to='#' />
   <div className='moodle'>
    <h1 className='subtitle moodle'>Moodle</h1>
    <br />
    <button className='refetch moodle' onClick={() => {
     setUserMoodleParams(userLocal)
    }} type="button">Actualizar moodle</button>
    <div className='assigments moodle'>
     {moodleAssigments}
    </div>
   </div>
   <br />
   <div className='classroom' style={{ marginBottom: "50px" }}>
    <h1 className='subtitle classroom'>Classroom</h1>
    <br />
    <button className='refetch  classroom' onClick={() => {
     setUserClassRoomParams(userLocal)
    }} type="button">Actualizar classroom</button>
    <div className='assigments classroom'>
     {classRoomAssigments}
    </div>
   </div>

  </main>
 )
}

