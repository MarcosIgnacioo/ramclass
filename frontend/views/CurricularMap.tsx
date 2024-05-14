import React from 'react'
import getUser, { LSK, checkBothCache } from '../functions/store'
import { useUser } from '../components/UserContext'
import useLogin from '../functions/useLogin'
import Subject from '../components/Subject'
import SubjectClass from '../classes/Subject'

export const CurricularMap = () => {
 console.log("wep")
 const userLocal = getUser()
 const user = useUser()
 const response = useLogin(user)
 const curricularMap = checkBothCache(response, LSK.CurricularMap) as Object[]
 console.log("curri", curricularMap)
 if (!userLocal) return (<h1>Esperate wey</h1>)

 return (
  <div>{curricularMap.map(subject => (<Subject{...subject as SubjectClass} />))}</div>
 )
}
