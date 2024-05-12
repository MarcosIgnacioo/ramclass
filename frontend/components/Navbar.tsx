import { useQueryClient } from '@tanstack/react-query'
import React from 'react'
import { Link } from 'react-router-dom'

export default function Navbar() {
 const queryClient = useQueryClient()
 return (
  <div className='navbar' id='culo'>
   <Link to="/">Home</Link>
   <Link to="/student">student</Link>
   <Link to="/" onClick={() => {
    localStorage.clear()
    queryClient.clear()
    console.log("Logged out")
   }}>Log-out</Link>
  </div>
 )
}

