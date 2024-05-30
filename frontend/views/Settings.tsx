import React, { useState } from 'react'
import updateCurrentLocation from '../functions/location'
import Title from '../components/Title'
import { Link } from 'react-router-dom'
import { logOut } from '../functions/logOut'
import { getCacheOf, storeInLocal } from '../functions/store'
import useDeleteAccount from '../functions/useDeleteAccount'
import FloatingWindow from '../components/FloatingWindow'
import DeleteAccount from '../components/DeleteAccount'

export default function Settings() {
 updateCurrentLocation()

 let crtPref = true
 if (getCacheOf("crt") === null) {
  storeInLocal(true, "crt")
 } else {
  crtPref = getCacheOf("crt") as boolean
 }
 const [isCrtOn, setIsCrtOn] = useState(crtPref)
 const [account, setAccount] = useState("")
 const [floatingPopup, setFloatingPopup] = useState<React.JSX.Element>(<div hidden>
 </div>)
 const overlay = document.querySelector(".retro-overlay")
 useDeleteAccount(account)

 const userLocal = getCacheOf("identifier") as string
 // quiero la actualizacion full porque la navbar no se actualiza :(
 if (!userLocal) window.location.replace("/")

 if (isCrtOn) overlay?.removeAttribute("hidden")
 else overlay?.setAttribute("hidden", `${!isCrtOn}`)

 return (
  <main>
   {floatingPopup}
   <Title title='Ajustes' to='#' />
   <div className='settings-container'>
    <label className='check-label'><input type="checkbox" checked={isCrtOn as boolean} name="" onChange={(e) => {
     console.log("crt:", e.target.checked)
     storeInLocal(e.target.checked, "crt")
     setIsCrtOn(e.target.checked)
    }} />Filtro CRT</label>
    <Link className='faq settings-button' to={"/faq"}>FAQ</Link>
    <Link className='seeya settings-button' to={"/"} onClick={logOut}>Cerrar sesión</Link>
    <button className="danger settings-button" onClick={() => {
     setFloatingPopup(<FloatingWindow setFloatingPopup={setFloatingPopup} content={<DeleteAccount setAccount={setAccount} />} />)
    }}>Borrar cuenta</button>
   </div>
  </main>
 )
}

