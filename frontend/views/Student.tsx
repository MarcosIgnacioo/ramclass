import React from 'react'
import { useLocationContext, useLocationUpdateContext, useUser } from '../components/UserContext';
import useLogin from '../functions/useLogin';
import getUser, { LSK, checkBothCache } from '../functions/store';
import StudentCredential from '../components/StudentCredential';
import Student from '../classes/Student';
import useLocationEffect from '../functions/effects/useLocationEffect';

// Esta vista/componente se encarga de obtener al estudiante de la cache (React-Query o LocalStorage)

// const locationUpdate = useLocationUpdateContext()
// locationUpdate(window.location.pathname)
// const currentLocation = useLocationContext()
// useLocationEffect(currentLocation)

export default function Student() {

 const locationUpdate = useLocationUpdateContext()
 locationUpdate(window.location.pathname)
 const currentLocation = useLocationContext()
 useLocationEffect(currentLocation)

 const userLocal = getUser()
 if (!userLocal) return (<h1>Esperate wey</h1>)
 const user = useUser()
 const response = useLogin(user)
 const userData = checkBothCache(response, LSK.Student)
 return (
  <StudentCredential {...userData as Student} />
 )
}

