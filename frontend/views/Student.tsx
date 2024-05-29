import React, { useState } from 'react'
import { useUser } from '../components/UserContext';
import useLogin from '../functions/useLogin';
import getUser, { LSK, checkBothCache } from '../functions/store';
import StudentCredential from '../components/StudentCredential';
import Student from '../classes/Student';
import SignIn from './SignIn';
import UserData from '../classes/UserData';
import updateCurrentLocation from '../functions/location';
import useCredentials from '../functions/useCredentials';
import State from '../components/State';
import Title from '../components/Title';

export default function Student() {
 const [userCredentials, setUserCredentials] = useState<UserData | null>(null);

 updateCurrentLocation()

 const userLocal = (useUser().username == "") ? getUser() : useUser()

 if (!userLocal) return (<SignIn />)

 const user = useUser()
 const response = useLogin(user)

 const credentialsFetch = useCredentials(userCredentials)
 let userData = checkBothCache(response, LSK.Student)

 let credentialsInfo = State({ fetchedData: credentialsFetch, cache: [userData], nameSpace: "student", Container: StudentCredential, isFiltered: false })

 if (credentialsInfo === null || credentialsInfo.props.name === undefined && !credentialsFetch.isLoading) {
  return (<main>

   <div className='warn'>
    <h1 className='warn-header'>Advertencia</h1>
    <div className='warn-content'>
     <h1>No se cargó correctamente tu credencial,<br /> presiona aquí para hacerlo</h1>
     <button type="button" onClick={() => {
      setUserCredentials(userLocal)
     }} className='refetch kardex'>Cargar credencial</button>
    </div>
   </div>
  </main>)
 }

 return (
  <main className='credential-container'>
   <Title title='Credencial' to='#' />
   {credentialsInfo}
   <button hidden={credentialsFetch.isFetching} style={{ marginBottom: "75px" }} type='button' className="refetch credentials" onClick={() => {
    setUserCredentials(userLocal)
   }}>Actualizar credencial</button>
  </main>
 )
}

