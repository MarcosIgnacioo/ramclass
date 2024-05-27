import React from 'react'
import Error from '../components/Error'
import ErrorMinified from '../components/ErrorMinified'


export default function Prueba() {
 return (
  <main className='signin-container'>
   <ErrorMinified error="Credenciales incorrectas, favor de intentarlo de nuevo" />
   <form className='signin-form'>
    <h1>Inciar sesión</h1>
    <label htmlFor="username">Ingrese su identificador</label>
    <input required name="username" />
    <label htmlFor="password">Ingresa tu contraseña</label>
    <input required name="password" type="password" />
    <button type="submit" >Iniciar sesion</button>
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
}

