import React from 'react'
import ReactDOM from 'react-dom/client'
import "./css/styles.css";
import Title from './components/Title';
import { BrowserRouter, Route, Routes } from 'react-router-dom';
import SignIn from './views/SignIn';
import Prueba from './views/Prueba';
import Class from './views/Class';
import Mood from './views/Mood';
import Kardex from './views/Kardex';
import Curricular from './views/Curricular';
import Credentials from './views/Credentials';
function Application() {
  // Desactivar los botones cuando haga submit a un boton y reactivarlos despues de 10 minutos en caso de q no haya habido un error
  return (
    <div>
      <Title title="Ramclass" />
      <BrowserRouter>
        <Routes>
          <Route path='/' element={<SignIn />} />
          <Route path='/classroom-form' element={<Class />} />
          <Route path='/moodle-form' element={<Mood />} />
          <Route path='/kardex-form' element={<Kardex />} />
          <Route path='/curricular-form' element={<Curricular />} />
          <Route path='/credentials-form' element={<Credentials />} />
          <Route path='/test' element={<Prueba />} />
        </Routes>
      </BrowserRouter>
      <div data-v-e62ee844="" className="retro-overlay screen-h screen-w"></div>
    </div>
  )
}
const root = ReactDOM.createRoot(document.querySelector('#ramses')!)
root.render(<Application />)
