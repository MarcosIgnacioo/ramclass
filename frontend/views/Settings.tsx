import React, { useState } from 'react'
import updateCurrentLocation from '../functions/location'
import Title from '../components/Title'
import { Link } from 'react-router-dom'
import { logOut } from '../functions/logOut'
import { getCacheOf, LSK, storeInLocal, updateQueryCache } from '../functions/store'
import useDeleteAccount from '../functions/useDeleteAccount'
import FloatingWindow from '../components/FloatingWindow'
import DeleteAccount from '../components/DeleteAccount'
import useLogin from '../functions/useLogin'
import { useUser } from '../components/UserContext'

export default function Settings() {
 const userLocal = getCacheOf("identifier") as string
 const user = useUser()
 const response = useLogin(user)
 if (!userLocal) window.location.replace("/")
 updateCurrentLocation()

 let crtPref = true
 if (getCacheOf("crt") === null) {
  storeInLocal(true, "crt")
 } else {
  crtPref = getCacheOf("crt") as boolean
 }
 const [isCrtOn, setIsCrtOn] = useState(crtPref)
 const [classroomUserId, setUserClassroomUserId] = useState(getCacheOf("classroomUserId") as string)
 const [account, setAccount] = useState("")
 const [floatingPopup, setFloatingPopup] = useState<React.JSX.Element>(<div hidden>
 </div>)
 const classroomAssigments = getCacheOf("classroom") as Object[]
 const overlay = document.querySelector(".retro-overlay")
 useDeleteAccount(account)

 // quiero la actualizacion full porque la navbar no se actualiza :(

 if (isCrtOn) overlay?.removeAttribute("hidden")
 else overlay?.setAttribute("hidden", `${!isCrtOn}`)

 return (
  <main>
   {floatingPopup}
   <Title title='Ajustes' to='#' />
   <div className='settings-container'>
    <label className='check-label'><input type="checkbox" checked={isCrtOn as boolean} name="" onChange={(e) => {
     storeInLocal(e.target.checked, "crt")
     setIsCrtOn(e.target.checked)
    }} />Filtro CRT</label>
    <label htmlFor="classroomUserId" className='classroom-id user'>Ingresa el número de usuario de classroom</label>
    <input type="number" min="0" className='classroom-id' name="classroomUserId" value={classroomUserId} pattern="0-9" onChange={(e) => {
     let id = e.target.value
     id = (id.match(/\d+$/g)) ? id : "0"
     storeInLocal(id, "classroomUserId")
     setUserClassroomUserId(id)
     classroomAssigments.map(assigment => {
      assigment["link"] = assigment["link"].replace(/u\/\d+/, "u/" + id)
     })
     storeInLocal(classroomAssigments, "classroom")
     updateQueryCache(response, LSK.Classroom)
     console.log("todo updated")
    }} />
    <Link className='faq settings-button' to={"/faq"}>FAQ</Link>
    <Link className='seeya settings-button' to={"/"} onClick={logOut}>Cerrar sesión</Link>
    <button className="danger settings-button" onClick={() => {
     setFloatingPopup(<FloatingWindow setFloatingPopup={setFloatingPopup} content={<DeleteAccount setAccount={setAccount} />} />)
    }}>Borrar cuenta</button>
   </div>
  </main>
 )
}

