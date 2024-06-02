import React, { useState } from 'react'
import useLogin from '../functions/useLogin';
import UserData from '../classes/UserData';
import { useUserUpdate } from '../components/UserContext';
import { setAll, storeInLocal } from '../functions/store';
import { useNavigate } from 'react-router-dom';
import Loading from '../components/Loading';
import updateCurrentLocation from '../functions/location';
import LoginForm from '../components/LoginForm';
import ErrorMinified from '../components/ErrorMinified';
import { BASE_PATH } from '../globals/globals';
import Title from '../components/Title';

export default function SignIn() {

 const [loginParams, setLoginParams] = useState<UserData | null>(null);

 updateCurrentLocation()

 const updateUser = useUserUpdate()
 const navigate = useNavigate()

 const response = useLogin(loginParams)

 if (response.isLoading) {
  return (
   <main>
    <Loading />
   </main>)
 }

 if (response.isError) {
  return (
   <main className='signin-container'>
    <Title title='Inciar meeeh sesión' to='/' />
    <ErrorMinified error="Credenciales incorrectas, favor de intentarlo de nuevo" />
    <LoginForm setLoginParams={setLoginParams} />
    <p>Recuerda, tu identificador es
     <br />
     el inicio de tu dirección de correo
     <br />
     electrónico hasta el @.
     <br />
     <br />
     Ejemplo: <b>pikminc_21</b> @alu.uabcs.mx
    </p>
   </main>
  )
 }

 if (response.isSuccess) {
  // Guardamos en localstorage los datos scrappeados como las credenciales para iniciar sesion de nuevo la siguiente vez que se abra la pagina 
  const { classroom, curricular_map, kardex, moodle, student, gpa, tasks, calendar } = response.data
  const { username, password } = loginParams as UserData
  setAll([moodle, classroom, kardex, curricular_map, student, gpa, tasks, calendar, username, password])
  storeInLocal("0", "classroomUserId")
  updateUser(loginParams)
  navigate("/home")
 }

 return (
  <main className='signin-container'>
   <Title title='Inciar sesión' to='/' />
   <LoginForm setLoginParams={setLoginParams} />
   <p>Recuerda, tu identificador es
    <br />
    el inicio de tu dirección de correo
    <br />
    electrónico hasta el @.
    <br />
    <br />
    Ejemplo: <b>pikminc_21</b> @alu.uabcs.mx
   </p>
  </main>
 )
 // Poner en el boton que no recarge la pagina que namas pues haga el cambio d ruta o asi
}

