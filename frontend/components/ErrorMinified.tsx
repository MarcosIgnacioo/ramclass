import React from 'react'

export default function ErrorMinified(props) {
 const { error } = props
 return (
  <div className='error fade-in-text'>
   {error}
  </div>
 )
}

