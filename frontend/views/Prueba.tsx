import React, { useState } from 'react'
import useLogin from '../functions/useLogin'
import getUser from '../functions/getUser'
import UserData from '../classes/UserData';
export default function Prueba(qc) {
    console.log(qc)

    const [loginParams, setLoginParams] = useState<UserData | null>(null);
    const user = getUser()
    console.log(user)
    const response = useLogin(user)
    console.log(response.data)
    return (
        <h1> hola esto es una prueba</h1 >
    )
}
