import React, { useState } from 'react'
import useLogin from '../functions/useLogin';
import UserData from '../classes/UserData';
import { useUser, useUserUpdate } from '../components/UserContext';
import { setAll } from '../functions/store';
import { useNavigate } from 'react-router-dom';

export default function SignIn() {

 const [loginParams, setLoginParams] = useState<UserData | null>(null);

 const userInfo = useUser()
 const updateUser = useUserUpdate()
 const navigate = useNavigate()

 const response = useLogin(loginParams)

 if (response.isError) return (<div>
  <h2>oh no</h2>
 </div>)

 if (response.isLoading) return (
  <div className="loading-pane">
   <h2 className="loader">o</h2>
  </div>
 )

 if (response.isSuccess) {
  // Guardamos en localstorage los datos scrappeados como las credenciales para iniciar sesion de nuevo la siguiente vez que se abra la pagina 
  const { classroom, curricular_map, kardex, moodle, student } = response.data
  const { username, password } = userInfo
  setAll([moodle, classroom, kardex, curricular_map, student, username, password])
  updateUser(loginParams)
  navigate("/student")
 }

 return (
  <div>
   <form onSubmit={(e) => {
    const formData = new FormData(e.currentTarget)
    const data = {
     username: formData.get("username") ?? "",
     password: formData.get("password") ?? ""
    }
    // Guardar el usuario en la sesion actual en React
    setLoginParams(data)
    // Evitar que se actualice la página
    e.preventDefault()
   }}>
    <label htmlFor="username">Creando AAA el frontend</label>
    <input name="username" />
    <label htmlFor="password">Ingresa tu contraseña</label>
    <input name="password" type="password" />
    <button type="submit">Iniciar sesion</button>
   </form>
  </div>
 )
 // Poner en el boton que no recarge la pagina que namas pues haga el cambio d ruta o asi
}

