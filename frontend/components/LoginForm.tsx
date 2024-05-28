import React, { useState } from 'react'
import Terms from './Terms'
import Title from './Title'

export default function LoginForm(props) {

 const [isAccepting, setIsAccepting] = useState(true)
 const [isHidden, setIsHidden] = useState(true)
 const buttonClass = (!isAccepting) ? "" : ""
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
   <Title title='Inciar sesión' to='/' />
   <div className='input-container'>
    <label htmlFor="username">Ingrese su identificador</label>
    <input required name="username" />
   </div>
   <div className='input-container'>
    <label htmlFor="password">Ingresa tu contraseña</label>
    <input required name="password" type="password" />
   </div>
   <Terms setIsAccepting={setIsAccepting} isHidden={isHidden} />
   <button className={buttonClass} type="submit">Iniciar sesion</button>
  </form>
 )
}
// disabled={!isAccepting}

