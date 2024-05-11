import React, { useState } from 'react'
import UserData from '../classes/UserData';
import useLogin from '../functions/useLogin';


export default function Prueba() {
    const [test, setTest] = useState<UserData | null>(null);
    const results = useLogin(test)
    console.log(results)
    console.log(results.data)
    const testing = new UserData("fra2_21", "Saw211234566")
    return (
        <div>
            <button type="button" onClick={() => {
                setTest(testing)
            }}>VERGA</button>
            <h1>FUNCTIONA POR FAVOR</h1>
        </div>
    )
    // Poner en el boton que no recarge la pagina que namas pues haga el cambio d ruta o asi
}

