import React from 'react'
import updateCurrentLocation from '../functions/location'

export default function Todo() {
 updateCurrentLocation()
 return (
  <h1>Todo</h1>
 )
}

