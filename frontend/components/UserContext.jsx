import React, { createContext, useContext, useState } from 'react'
import UserData from '../classes/UserData'

const UserContext = createContext()
const UserUpdateContext = createContext()
const LocationContext = createContext()
const LocationUpdateContext = createContext()

export function useLocationContext() {
 return useContext(LocationContext)
}

export function useLocationUpdateContext() {
 return useContext(LocationUpdateContext)
}

export function useUser() {
 return useContext(UserContext)
}

export function useUserUpdate() {
 return useContext(UserUpdateContext)
}

export function UserProvider({ children }) {
 const [userInfo, setUserInfo] = useState(new UserData("", ""))
 const [location, setLocation] = useState("")

 function updateUser(user) {
  if (user == null) return
  setUserInfo(user)
 }
 function updateLocation(currentLocation) {
  setLocation(currentLocation)
 }

 return (

  <LocationContext.Provider value={location}>
   <LocationUpdateContext.Provider value={updateLocation}>
    <UserContext.Provider value={userInfo}>
     <UserUpdateContext.Provider value={updateUser}>
      {children}
     </UserUpdateContext.Provider>
    </UserContext.Provider>
   </LocationUpdateContext.Provider>
  </LocationContext.Provider>
 )
}

