import React from 'react'
import { useUser } from '../components/UserContext';
import useLogin from '../functions/useLogin';
import { checkBothCache } from '../functions/store';
import StudentCredential from '../components/StudentCredential';
import Student from '../classes/Student';

// Esta vista/componente se encarga de obtener al estudiante de la cache (React-Query o LocalStorage)

export default function Student() {
 const user = useUser()
 const response = useLogin(user)
 const userData = checkBothCache(response, "student")
 return (
  <StudentCredential {...userData as Student} />
 )
}

