import { useQuery } from '@tanstack/react-query';
import React, { useState } from 'react'
import login from '../functions/login.ts'
import { useNavigate } from "react-router-dom";

export default function SignIn() {

    interface UserData {
        username: string | FormDataEntryValue
        password: string | FormDataEntryValue
    }

    const [loginParams, setLoginParams] = useState<UserData | null>(null);
    const response = useQuery(
        {
            queryKey: ["user", loginParams],
            queryFn: login
        }
    )
    const navigate = useNavigate
    return (
        <form onSubmit={(e) => {
            const formData = new FormData(e.currentTarget)
            const data = {
                username: formData.get("username") ?? "",
                password: formData.get("password") ?? ""
            }
            setLoginParams(data)
            console.log("ou yea diamantes")
            console.log(response)
            e.preventDefault()
        }}>
            <label htmlFor="username">Creando AAA el frontend</label>
            <input name="username" />
            <label htmlFor="password">Ingresa tu contraseña</label>
            <input name="password" type="password" />
            <button type="submit">Iniciar sesion</button>
        </form>
    )
    // Poner en el boton que no recarge la pagina que namas pues haga el cambio d ruta o asi
}
