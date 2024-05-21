import React from 'react'
import TextWobly from './TextWobly'

export default function Message(props) {
 console.log(props.isHidden)
 return (
  <div ref={props.ref} className='success-container' hidden={props.isHidden}>
   <TextWobly title={props.message} class={props.class} />
  </div>
 )
}

