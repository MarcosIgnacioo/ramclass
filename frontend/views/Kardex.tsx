import React from 'react'
import { useUser } from '../components/UserContext';
import useLogin from '../functions/useLogin';
import getUser, { checkBothCache } from '../functions/store';

export default function Kardex() {
 const userLocal = getUser()

 const user = useUser()
 const response = useLogin(user)
 const kardex = checkBothCache(response, "kardex")
 console.log(kardex)

 if (!userLocal) return (<h1>Esperate wey</h1>)
 return (
  <div>Kardex</div>
 )
}

