import React, { useState } from 'react'
import useLogin from '../functions/useLogin';
import UserData from '../classes/UserData';
import { useUserUpdate } from '../components/UserContext';
import { setAll } from '../functions/store';
import { useNavigate } from 'react-router-dom';
import Error from '../components/Error';
import Loading from '../components/Loading';
import updateCurrentLocation from '../functions/location';
import Terms from '../components/Terms';

export default function SignIn() {

 const [loginParams, setLoginParams] = useState<UserData | null>(null);
 const [isAccepting, setIsAccepting] = useState(true)
 const [isHidden, setIsHidden] = useState(true)
 const buttonClass = (!isAccepting) ? "inactive" : ""

 updateCurrentLocation()

 console.log(isAccepting)

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
   <main>
    <Error />
   </main>
  )

 }

 if (response.isSuccess) {
  // Guardamos en localstorage los datos scrappeados como las credenciales para iniciar sesion de nuevo la siguiente vez que se abra la pagina 
  const { classroom, curricular_map, kardex, moodle, student, gpa } = response.data
  const { username, password } = loginParams as UserData
  setAll([moodle, classroom, kardex, curricular_map, student, gpa, username, password])
  updateUser(loginParams)
  navigate("/student")
 }

 return (
  <main className='signin-container'>
   <form onSubmit={(e) => {
    if (isHidden) {
     console.log("estoy entreando al early return")
     setIsHidden(false)
     setIsAccepting(false)
     e.preventDefault()
     return
    }
    const formData = new FormData(e.currentTarget)
    const data = {
     username: formData.get("username") ?? "",
     password: formData.get("password") ?? ""
    }
    // Guardar el usuario en la sesion actual en React
    setLoginParams(data)
    e.preventDefault()
   }} className='signin-form'>
    <h1>Inciar sesión</h1>
    <label htmlFor="username">Ingrese su identificador</label>
    <input required name="username" />
    <label htmlFor="password">Ingresa tu contraseña</label>
    <input required name="password" type="password" />
    <Terms setIsAccepting={setIsAccepting} isHidden={isHidden} />
    <button className={buttonClass} disabled={!isAccepting} type="submit">Iniciar sesion</button>
   </form>
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

