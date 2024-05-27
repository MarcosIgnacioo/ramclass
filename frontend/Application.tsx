import React from 'react'
import ReactDOM from 'react-dom/client'
import "./css/styles.css";
import { BrowserRouter, Route, Routes } from 'react-router-dom';
import SignIn from './views/SignIn';
import Prueba from './views/Prueba';
import { QueryClient, QueryClientProvider } from "@tanstack/react-query"
import { UserProvider } from './components/UserContext';
import Navbar from './components/Navbar';
import Student from './views/Student';
import Kardex from './views/Kardex';
import { CurricularMap } from './views/CurricularMap';
import Home from './views/Home';
import Settings from './views/Settings';
import Faq from './views/Faq';
import Todo from './views/Todo';
import Calendar from './views/Calendar';

function Application() {


 const queryClient = new QueryClient({
  defaultOptions: {
   queries: {
    staleTime: Infinity,
   }
  }
 })
 return (
  <div>
   <BrowserRouter>
    <QueryClientProvider client={queryClient}>
     <UserProvider>
      <Navbar />
      <Routes >
       <Route path='/' element={<SignIn />} />
       <Route path='/sign-in' element={<SignIn />} />
       <Route path='/home' element={<Home />} />
       <Route path='/student' element={<Student />} />
       <Route path='/my-kardex' element={<Kardex />} />
       <Route path='/my-curricular-map' element={<CurricularMap />} />
       <Route path='/settings' element={<Settings />} />
       <Route path='/todo' element={<Todo />} />
       <Route path='/calendar' element={<Calendar />} />
       <Route path='/faq' element={<Faq />} />
       <Route path='/test' element={<Prueba />} />
      </Routes>
     </UserProvider>
    </QueryClientProvider>
   </BrowserRouter>
   <div data-v-e62ee844="" className="retro-overlay screen-h screen-w"></div>
  </div>
 )
}
const root = ReactDOM.createRoot(document.querySelector('#ramses')!)
root.render(<Application />)
