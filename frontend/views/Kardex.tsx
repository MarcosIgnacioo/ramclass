import React from 'react'
import { useUser } from '../components/UserContext';
import useLogin from '../functions/useLogin';
import getUser, { LSK, checkBothCache } from '../functions/store';
import SubjectClass from '../classes/Subject.ts';
import Subject from '../components/Subject.tsx';

export default function Kardex() {
 const userLocal = getUser()
 // Obtenemos el user que se genera cada vez que abrimos una pagina, en el caso de que hagamos login este user se actualiza por lo que significa que hemos hecho fetch a de todass las cosas, si abrimos de nuevo la pagina este usuario estara vacio pues signficara que no hemos hecho nada de eso
 const user = useUser()
 const response = useLogin(user)
 console.log("wehasdfads")
 console.log("asdf", response)
 const kardex = checkBothCache(response, LSK.Kardex) as Object[]
 console.log(kardex)

 if (!userLocal) return (<h1>Esperate wey</h1>)

 return (
  <div>{kardex.map(subject => (<Subject {...subject as SubjectClass} />))}</div>
 )
}

