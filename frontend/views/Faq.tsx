import React from 'react'
import { useLocationUpdateContext, useLocationContext } from '../components/UserContext'
import useLocationEffect from '../functions/effects/useLocationEffect'

export default function Faq() {
 const locationUpdate = useLocationUpdateContext()
 locationUpdate(window.location.pathname)
 const currentLocation = useLocationContext()
 useLocationEffect(currentLocation)

 return (
  <div>Faq</div>
 )
}

