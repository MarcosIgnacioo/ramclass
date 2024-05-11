import React, { createContext, useContext, useState } from 'react'
import UserData from '../classes/UserData'

const UserContext = createContext()
const UserUpdateContext = createContext()

export function useUser() {
    return useContext(UserContext)
}

export function useUserUpdate() {
    return useContext(UserUpdateContext)
}

export function UserProvider({ children }) {
    const [userInfo, setUserInfo] = useState(new UserData("", ""))

    function updateUser(user) {
        setUserInfo(user)
    }

    return (
        <UserContext.Provider value={userInfo}>
            <UserUpdateContext.Provider value={updateUser}>
                {children}
            </UserUpdateContext.Provider>
        </UserContext.Provider>
    )
}

