import React from 'react'

export default function FloatingWindow(props) {
 const { setFloatingPopup, content } = props
 const close = <div hidden></div>
 return (
  <div className='floating-window'>
   <button id='close-button' type='button' onClick={() =>
    setFloatingPopup(close)}>x</button>
   {content}
  </div>
 )
}

