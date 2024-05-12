import React from 'react'
import Title from './Title'

// Este componente se encarga de plasmar la credencial del usuario en un container
export default function StudentCredential({ name, control_number, institutional_email, campus, career, period, semester, group, turn, state }) {
 return (
  <div className='credential'>
   <Title title='Credencial del estudiante' />
   <div className='personal-data'>
    <h1>{name}</h1>
    <h2>{control_number}</h2>
    <h2>{institutional_email}</h2>
   </div>
   <div className='school-data'>
    <h1>{career}</h1>
    <h2>{campus}</h2>
    <h2>{semester}</h2>
    <h2>{group}</h2>
    <h2>{turn}</h2>
    <h2>{state}</h2>
    <h2>{period}</h2>
   </div>
  </div>
 )
}

