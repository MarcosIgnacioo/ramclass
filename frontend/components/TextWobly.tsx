import React from 'react'

export default function TextWobly(props) {
 const letters = [...props.title]
 return (
  <span className={props.class}>
   {letters.map((letter, index) => (
    <span className="wobbly-text" data-v-2f4c1ac1="" style={{ animationDelay: (index * 0.05).toFixed(2) + "s", color: "#1d1d1d" }}>{letter}</span>
   ))}
  </span>
 )
}

