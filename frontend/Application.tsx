import React, { createContext } from 'react'
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
import { UserProvider } from './functions/UserContext';


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
            <Title title="Ramclass" />
            <BrowserRouter>
                <QueryClientProvider client={queryClient}>
                    <UserProvider>
                        <Routes >
                            <Route path='/' element={<SignIn />} />
                            <Route path='/classroom-form' element={<Class />} />
                            <Route path='/moodle-form' element={<Mood />} />
                            <Route path='/kardex-form' element={<Kardex />} />
                            <Route path='/curricular-form' element={<Curricular />} />
                            <Route path='/credentials-form' element={<Credentials />} />
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
