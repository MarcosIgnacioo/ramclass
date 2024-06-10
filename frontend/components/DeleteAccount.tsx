import React from 'react'
import deleteAccount from '../functions/fetchs/deleteAccount'
import { getCacheOf } from '../functions/store'

export default function DeleteAccount({setRelocation}) {

 return (
  <div id='delete-account-container'>
   <p>¿Deseas borrar tus datos guardados de nuestra base de datos? (Esto solamente borra los datos de nuestro servidor)</p>
   <button id='delete-account-button' type="button" onClick={ async () => {
    await deleteAccount(getCacheOf("identifier") as string)
    localStorage.clear()
    setRelocation(true)
   }}>
    Sí, borrar datos
   </button>
  </div>
 )
}

