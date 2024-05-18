import React, { useEffect, useState } from 'react'
import { getTerms } from '../functions/fetchs/terms'

export default function Terms(props) {

 const [service, setService] = useState()
 const [privacy, setPrivacy] = useState()

 useEffect(() => {
  getTerms(setService, setPrivacy)
 }, [])

 return (
  <div className='terms-container' hidden={props.isHidden}>
   <div className='terms-text'>
    <h5>Términos y condiciones de Ramclass, Ramtendo.</h5>
    <p>{service}</p>
    <br />
    <p>{privacy}</p>
   </div>
   <div className='accept-fields'>
    <label><input type="checkbox" name="" value="" onChange={(e) => {
     props.setIsAccepting(e.target.checked)
    }} />Acepto los términos y condiciones.</label>
   </div>
  </div>
 )
}

