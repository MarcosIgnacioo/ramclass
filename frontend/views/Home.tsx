import React, { useState } from 'react'
import getUser, { LSK, storeInLocal } from '../functions/store'
import { checkCache } from '../functions/checkLogin'
import Assigment from '../components/Assigment'
import UserData from '../classes/UserData'
import useMoodle from '../functions/useMoodle'
import { useLocationContext, useLocationUpdateContext, useUser } from '../components/UserContext'
import useClassRoom from '../functions/useClassRoom'
import State from '../components/State'
import Title from '../components/Title'
import useLocationEffect from '../functions/effects/useLocationEffect'

export default function Home() {

 const locationUpdate = useLocationUpdateContext()
 locationUpdate(window.location.pathname)
 const currentLocation = useLocationContext()
 useLocationEffect(currentLocation)

 const [userClassRoomParams, setUserClassRoomParams] = useState<UserData | null>(null);
 const [userMoodleParams, setUserMoodleParams] = useState<UserData | null>(null);
 const moodleFetch = useMoodle(userMoodleParams)
 const classRoomFetch = useClassRoom(userClassRoomParams)

 let classRoom = checkCache(LSK.Classroom) as Object[]
 let moodle = checkCache(LSK.Moodle) as Object[]

 const moodleAssigments = State({ fetchedData: moodleFetch, cache: moodle, nameSpace: "moodle", Container: Assigment, returnArray: false })
 const classRoomAssigments = State({ fetchedData: classRoomFetch, cache: classRoom, nameSpace: "classroom", Container: Assigment, returnArray: false })

 const userLocal = (useUser().username == "") ? getUser() : useUser()

 return (
  <main className='home'>
   <Title title='Tareas' />
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
   <div className='classroom'>
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

