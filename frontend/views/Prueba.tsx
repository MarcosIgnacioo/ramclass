import React from 'react'
import Error from '../components/Error'

export default function Prueba() {
 return (<main>
  <Error error={"Ha ocurrido un error inesperado, puedes volver a intentar lo que querías hacer o actualizar la página."} refreshButton={true} />
 </main>)
}

