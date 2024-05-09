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
import { QueryClient, QueryClientProvider } from "@tanstack/react-query"
function Application() {
    // Desactivar los botones cuando haga submit a un boton y reactivarlos despues de 10 minutos en caso de q no haya habido un error

    const queryClient = new QueryClient({
        defaultOptions: {
            queries: {
                // El staleTime si se pone en Infinity lo que le estamos diciendo es que una vez se fetchee algo que no se refetchee, si pusieramos un tiempo se refetchearia cada cierto tiempo esto podria ser util para classmoods pero no se porque servidor no muy gamer
                staleTime: Infinity,
            }
        }
    })
    return (
        <div>
            <Title title="Ramclass" />
            <BrowserRouter>

                <QueryClientProvider client={queryClient}>
                    <Routes >
                        <Route path='/' element={<SignIn />} />
                        <Route path='/classroom-form' element={<Class />} />
                        <Route path='/moodle-form' element={<Mood />} />
                        <Route path='/kardex-form' element={<Kardex />} />
                        <Route path='/curricular-form' element={<Curricular />} />
                        <Route path='/credentials-form' element={<Credentials />} />
                        <Route path='/test' element={<Prueba />} />
                    </Routes>

                </QueryClientProvider>
            </BrowserRouter>
            <div data-v-e62ee844="" className="retro-overlay screen-h screen-w"></div>
        </div>
    )
}
const root = ReactDOM.createRoot(document.querySelector('#ramses')!)
root.render(<Application />)
