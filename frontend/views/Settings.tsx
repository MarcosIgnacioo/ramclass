import React, { useState } from 'react'
import updateCurrentLocation from '../functions/location'
import Title from '../components/Title'

export default function Settings() {
 const [crtFilter, setCrtFilter] = useState(true)
 updateCurrentLocation()
 if (!crtFilter) {
  document.querySelector("retro-overlay")?.setAttribute("hidden", "")
 }
 return (
  <main>
   <Title title='Ajustes' />
   <div>
    <label><input type="checkbox" name="" value="" onChange={(e) => {
     setCrtFilter(e.target.checked)
    }} />Filtro CRT</label>
   </div>
  </main>
 )
}

