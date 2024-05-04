import React, { useState } from 'react'

export default function Curricular() {
  const [username, setUsername] = useState("")
  const [password, setPassword] = useState("")
  return (
    <form method='POST' action='/curricular'>
      <h1>(Scrappeo mapa curricular)</h1>
      <label htmlFor="username">Nombre de usuario </label>
      <input onChange={(e) => setUsername(e.target.value)} value={username} type="" name="username" />

      <label htmlFor="password">Ingresa tu contraseña</label>
      <input onChange={(e) => setPassword(e.target.value)} value={password} type="password" name="password" />
      <button type="submit">Iniciar sesion</button>
    </form>
  )
}

