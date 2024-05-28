import React, { useRef } from 'react'
import { Link } from 'react-router-dom'
import { useUser } from './UserContext'
import useNavBarEffect from '../functions/useNavBarEffect'
import Title from './Title'


export default function Navbar() {
 const userInfo = useUser()
 const navbarRef = useRef(null)
 const navBarData = useNavBarEffect(userInfo);
 const navBarLinks: React.ReactElement[] = []
 const navBarLinksMobile: React.ReactElement[] = []
 for (const link in navBarData) {
  navBarLinks.push(<Link id={link} to={link} onClick={navBarData[link]["function"]}>{navBarData[link]["text"]}</Link>)
  navBarLinksMobile.push(<Link id={link} to={link} onClick={navBarData[link]["function"]} className="material-symbols-outlined nav__link active-link">{navBarData[link]["icon"]}</Link>)
 }

 return (
  <div>
   <div className='navbar' ref={navbarRef}>
    <div className='head-navbar'>
     <Title title="RAMCLASS" />
    </div>
    {navBarLinks}
   </div>
   <div className='navbar-mobile'>
    {navBarLinksMobile}
   </div>
  </div>
 )
}
// <div className="icon nav-icon-5" ref={navbarRef} id='hamburger' onClick={toggle}>
//     <span></span>
//     <span></span>
//     <span></span>
//    </div>
//
//
