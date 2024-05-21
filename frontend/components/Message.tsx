import React from 'react'
import TextWobly from './TextWobly'

export default function Message(props, ref) {
 return (
  <div ref={ref} className={`${props.class} fade-animation`} hidden={props.isHidden}>
   <TextWobly title={props.message} class={props.class} />
  </div>
 )
}

