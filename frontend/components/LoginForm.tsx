import React, { useState } from 'react'
import Terms from './Terms'

export default function LoginForm(props) {

 const [isAccepting, setIsAccepting] = useState(true)
 const [isHidden, setIsHidden] = useState(true)
 const buttonClass = (!isAccepting) ? "inactive" : ""
 const { setLoginParams } = props

 return (
  <form onSubmit={(e) => {
   if (isHidden) {
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
 )
}

