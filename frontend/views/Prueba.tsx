import React from 'react'
import Error from '../components/Error'
import ErrorMinified from '../components/ErrorMinified'
import LoginForm from '../components/LoginForm'
import Title from '../components/Title'

export default function Prueba() {
 return (<main>
  <main className='signin-container'>
   <Title title='Inciar sesión' to='/' />
   <ErrorMinified error="Credenciales incorrectas" />
   <LoginForm setLoginParams={null} />
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
 </main>)
}

