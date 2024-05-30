import React from 'react'
import { getCacheOf } from '../functions/store'

export default function DeleteAccount(props) {
 const { setAccount } = props
 return (
  <div id='delete-account-container'>
   <p>¿Deseas borrar tus datos guardados de nuestra base de datos? (Esto solamente borra los datos de nuestro servidor)</p>
   <button id='delete-account-button' type="button" onClick={() => {
    setAccount(getCacheOf("identifier") as string)
    localStorage.clear()
   }}>
    Sí, borrar datos
   </button>
  </div>
 )
}

