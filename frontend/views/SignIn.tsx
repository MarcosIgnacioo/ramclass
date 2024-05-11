import React, { useContext, useState } from 'react'
import useLogin from '../functions/useLogin';
import UserData from '../classes/UserData';
import getInfo from '../functions/aaa';
import { useUser, useUserUpdate } from '../functions/UserContext';

export default function SignIn() {

    const [loginParams, setLoginParams] = useState<UserData | null>(null);
    const userInfo = useUser()
    const updateUser = useUserUpdate()
    updateUser(loginParams)

    console.log("wep?", userInfo)

    const response = useLogin(userInfo)

    if (response.isError) return (<div>
        <h2>oh no</h2>
    </div>)

    if (response.isLoading) return (
        <div className="loading-pane">
            <h2 className="loader">o</h2>
        </div>
    )

    if (response.isSuccess) {
        const { classroom, curricular_map, kardex, moodle, student } = response.data
        localStorage.setItem("classroom", JSON.stringify(classroom))
        localStorage.setItem("moodle", JSON.stringify(moodle))
        localStorage.setItem("kardex", JSON.stringify(kardex))
        localStorage.setItem("curricular_map", JSON.stringify(curricular_map))
        localStorage.setItem("student", JSON.stringify(student))
        localStorage.setItem("identifier", JSON.stringify(loginParams?.username))
        localStorage.setItem("password", JSON.stringify(loginParams?.password))
        return (
            <div className="loading-pane">
                <h2 className="loader">{student.name}</h2>
            </div>
        )
    }

    return (
        <div>
            <button type="button" onClick={updateUser}>como</button>
            <form onSubmit={(e) => {
                const formData = new FormData(e.currentTarget)
                const data = {
                    username: formData.get("username") ?? "",
                    password: formData.get("password") ?? ""
                }
                setLoginParams(data)
                updateUser(loginParams)
                e.preventDefault()
            }}>
                <label htmlFor="username">Creando AAA el frontend</label>
                <input name="username" />
                <label htmlFor="password">Ingresa tu contraseña</label>
                <input name="password" type="password" />
                <button type="submit">Iniciar sesion</button>
            </form>
            <button type="button" onClick={getInfo}>qieuriqweriuo</button>
        </div>
    )
    // Poner en el boton que no recarge la pagina que namas pues haga el cambio d ruta o asi
}

