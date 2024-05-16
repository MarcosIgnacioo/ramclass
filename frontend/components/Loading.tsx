import React from 'react'
import WoomyText from './AnimatedText'

export default function Loading() {
 return (
  <div className='loading'>
   <WoomyText title='Espera por favor ...' />
   <br />
   <span>¡Recuerda!</span>
   <br />
   <span><span className='ramtendo'>Ramtendo</span> nunca te pedirá tus datos personales,</span>
   <br />
   <span>porque ya los tiene todos.</span>
  </div>
 )
}

