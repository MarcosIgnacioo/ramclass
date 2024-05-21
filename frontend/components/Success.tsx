import React from 'react'
import TextWobly from './TextWobly'

export default function Success(props) {
 return (
  <div className='success-container' hidden={props.isHidden}>
   <TextWobly title={props.message} class={"message success"} />
  </div>
 )
}

