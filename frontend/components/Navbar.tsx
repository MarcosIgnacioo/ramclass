import React, { useState } from 'react'
import { Link } from 'react-router-dom'
import { useUser } from './UserContext'
import useNavBarEffect from '../functions/useNavBarEffect'
import Title from './Title'


export default function Navbar() {
 const userInfo = useUser()
 const navBarData = useNavBarEffect(userInfo);

 const navBarLinks: React.ReactElement[] = []
 for (const link in navBarData) {
  navBarLinks.push(<Link id={link} to={link} onClick={navBarData[link]["function"]}>{navBarData[link]["text"]}</Link>)
 }

 return (
  <div className='navbar'>
   <Title title="RAMCLASS" />
   {navBarLinks}
  </div >
 )
}

