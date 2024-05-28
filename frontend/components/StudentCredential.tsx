import React from 'react'
import SpanWithCopy from './SpanWithCopy'

// Este componente se encarga de plasmar la credencial del usuario en un container
export default function StudentCredential({ name, control_number, institutional_email, campus, career, period, semester, group, turn, state }) {
 return (
  <div>
   <div className='credential'>
    <div className='upper-credential'>
     <div className='fields-container names'>
      <span className='field-name'>Número de control: </span>
      <span className='field-name'>Nombre:  </span>
      <span className='field-name'>Correo institucuional: </span>
      <span className='field-name'>Carrera: </span>
     </div>
     <div className='fields-container values'>
      {SpanWithCopy(control_number)}
      {SpanWithCopy(name)}
      {SpanWithCopy(institutional_email)}
      <span className='field-value'>{career} </span>
     </div>
    </div>
    <div className='school-data'>
     <div className='career-info'>
      <span className='field-name'>Grupo: </span>
      <span className='field-value'>{group}</span>
      <span className='field-name'> Semestre: </span>
      <span className='field-value'>{semester}</span>
      <span className='field-name'> Turno: </span>
      <span className='field-value'>{turn} </span>
     </div>
     <div className='extra'>
      <span className='field-name'>Campus: </span>
      <span className='field-value'>{campus}</span>
      <span className='field-name'> Estado: </span>
      <span className='field-value'>{state}</span>
      <span className='field-name'> Periodo: </span>
      <span className='field-value'>{period}</span>
     </div>
    </div>
   </div>

  </div>
 )
}

