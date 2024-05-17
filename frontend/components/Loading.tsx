import React from 'react'
import WoomyText from './AnimatedText'

export default function Loading() {
 return (
  <div className='loading'>
   <WoomyText title='Espera por favor ...' />
   <span>¡Recuerda!</span>
   <span><span className='ramtendo'>Ramtendo</span> nunca te pedirá tus datos personales,</span>
   <span>porque ya los tiene todos.</span>
  </div>
 )
}

