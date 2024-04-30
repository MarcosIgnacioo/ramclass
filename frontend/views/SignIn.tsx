import React, { useState } from 'react'
export default function SignIn() {
  const [username, setUsername] = useState("")
  const [password, setPassword] = useState("")
  return (
    <form method='POST' action='/login-user'>
      <label htmlFor="username">estocabmioIngresa tu nombre de usuario</label>
      <input onChange={(e) => setUsername(e.target.value)} value={username} type="" name="username" />

      <label htmlFor="password">Ingresa tu contraseña</label>
      <input onChange={(e) => setPassword(e.target.value)} value={password} type="password" name="password" />
      <button type="submit">Iniciar sesion</button>
    </form>
  )
  // Poner en el boton que no recarge la pagina que namas pues haga el cambio d ruta o asi
}
