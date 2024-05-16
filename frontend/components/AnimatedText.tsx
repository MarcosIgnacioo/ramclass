import React from 'react'
export default function WoomyText({ title }: { title: string }) {
 const letters = [...title]
 return (
  <span className='loading-top'>
   {letters.map((letter, index) => (
    <span className="wobbly-text" data-v-2f4c1ac1="" style={{ animationDelay: (index * 0.05).toFixed(2) + "s" }}>{letter}</span>
   ))}
  </span>)
}
