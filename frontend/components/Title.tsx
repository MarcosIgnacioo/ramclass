import React from 'react'
import { Link } from 'react-router-dom'
export default function Title({ title, to = "#" }: { title: string, to: string }) {
 const letters = [...title]
 return (
  <Link to={to} className="about__title font-pixel text-title">
   {letters.map((letter, index) => (
    <span className="wobbly-text" data-v-2f4c1ac1="" style={{ animationDelay: (index * 0.05).toFixed(2) + "s" }}>{letter}</span>
   ))}
  </Link>)
}
