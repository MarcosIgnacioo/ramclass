import React from 'react'

export default function BackwardArrow(props) {
 function backward() {
  const { thing, setThing, low, high } = props
  if (thing - 1 < low) {
   setThing(high)
   return
  }
  setThing(thing - 1)
 }
 return (
  <span className='left-arrow' onClick={backward}>&lt; </span>
 )
}

