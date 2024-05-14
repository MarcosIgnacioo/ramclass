import React from 'react'
import { useUser } from '../components/UserContext';
import useLogin from '../functions/useLogin';
import getUser, { LSK, checkBothCache } from '../functions/store';
import StudentCredential from '../components/StudentCredential';
import Student from '../classes/Student';

// Esta vista/componente se encarga de obtener al estudiante de la cache (React-Query o LocalStorage)

export default function Student() {
 const userLocal = getUser()
 if (!userLocal) return (<h1>Esperate wey</h1>)
 const user = useUser()
 const response = useLogin(user)
 const userData = checkBothCache(response, LSK.Student)
 return (
  <StudentCredential {...userData as Student} />
 )
}

