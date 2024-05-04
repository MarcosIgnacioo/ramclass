import React, { useState } from 'react'

export default function Credentials() {
  const [username, setUsername] = useState("")
  const [password, setPassword] = useState("")
  return (
    <form method='POST' action='/credentials'>
      <h1>(Scrappeo credencial de alumno)</h1>
      <label htmlFor="username">Nombre de usuario </label>
      <input onChange={(e) => setUsername(e.target.value)} value={username} type="" name="username" />

      <label htmlFor="password">Ingresa tu contraseña</label>
      <input onChange={(e) => setPassword(e.target.value)} value={password} type="password" name="password" />
      <button type="submit">Iniciar sesion</button>
    </form>
  )
}

