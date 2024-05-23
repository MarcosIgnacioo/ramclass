import React from 'react'

export default function FowardArrow(props) {
 function foward() {
  const { thing, setThing, low, high } = props
  if (thing + 1 > high) {
   setThing(low)
   return
  }
  setThing(thing + 1)
 }
 return (
  <span className='right-arrow' onClick={foward}>&gt; </span>
 )
}

