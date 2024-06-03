import React from 'react'
import Title from '../components/Title'
import updateCurrentLocation from '../functions/location'

export default function Faq() {
 updateCurrentLocation()

 return (
  <main className='faq'>
   <Title title='FAQ' to='#' />
   <div className='faq-container'>
    <div className='question-container'>
     <h1>¿Qué es el número de usuario de ClassRoom?</h1>
     <h2>
      Es el número del lugar que ocupa tu cuenta institucional en tu navegador, normalmente las personas tienen su correo gmail personal y luego su cuenta perteneciente a la universidad, entonces el lugar que ocupa es el 1; Si al momento de darle click a los enlaces de las tareas de ClassRoom no entra directamente a tu cuenta institucional, signfica que en tu navegador está
     </h2>
    </div>
    <div className='question-container'>
     <h1>¿Cómo tienen acceso a la información?</h1>
     <h2>Usamos las credenciales que proporcionas para iniciar sesión a las páginas correspondientes para obtener tu informacion (siia, moodle y classroom).</h2>
    </div>
    <div className='question-container'>
     <h1>¿Es seguro?</h1>
     <h2>No.</h2>
    </div>
    <div className='question-container'>
     <h1>¿Guardan las contraseñas o alguna otra información sensible?</h1>
     <h2>Solamente guardamos información no sensible en nuestra base de datos, como lo son las actividades que quieres guardar en la aplicacion todo integrada (¡esto solamente se guarda si quieres hacerlo!); Las cosas persisten de manera local.</h2>
    </div>
    <div className='question-container'>
     <h1>¿Por qué se demora tanto?</h1>
     <h2>Porque se realiza por medio de web scrapping, el cual es básicamente la simulación de un usuario entrando a una página, y no se pueden controlar los tiempos de cargas de dichas páginas, sobretodo el de classroom, que es el que tarda más (~10 segundos).</h2>
    </div>
    <div className='question-container'>
     <h1>¿Se podría hacer mas rápido a futuro?</h1>
     <h2>Sí, si google nos deja registrar la aplicación es posible reducir los tiempos de carga hasta un 70%, solamente que por la naturaleza de nuestra aplicación es poco probable que google quiera registrarla.</h2>
    </div>
    <div className='question-container'>
     <h1>¿Ramses?</h1>
     <h2>???</h2>
    </div>
   </div>


  </main>
 )
}

