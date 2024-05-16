import React, { useState } from 'react'
import { useLocationUpdateContext, useLocationContext } from '../components/UserContext'
import useLocationEffect from '../functions/effects/useLocationEffect'

export default function Account() {

 const locationUpdate = useLocationUpdateContext()
 locationUpdate(window.location.pathname)
 const currentLocation = useLocationContext()
 useLocationEffect(currentLocation)

 const [username, setUsername] = useState("")
 const [password, setPassword] = useState("")
 return (
  <form method='POST' action='/classroom'>
   <h1>(Create Forum Account)</h1>
   <label htmlFor="username">Nombre de usuario </label>
   <input onChange={(e) => setUsername(e.target.value)} value={username} type="" name="username" />

   <label htmlFor="password">Ingresa tu contraseña</label>
   <input onChange={(e) => setPassword(e.target.value)} value={password} type="password" name="password" />
   <button type="submit">Iniciar sesion</button>
  </form>
 )
}

