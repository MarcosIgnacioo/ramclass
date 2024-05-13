import { UseQueryResult } from '@tanstack/react-query'
import UserData from '../classes/UserData'

const localNames = ["moodle", "classroom", "kardex", "curricular_map", "student", "identifier", "password"]

const getUser = (): UserData | null => {
 const identifier = (localStorage.getItem("identifier") ?? "")
 const password = (localStorage.getItem("password") ?? "")
 if (identifier === "" || password === "") return null
 return new UserData(identifier, password)
}

export const storeInLocal = (itemsToStore: Array<String> | Object | string, itemName: string) => {
 const curricularMapParsed = JSON.stringify(itemsToStore)
 localStorage.setItem(itemName, curricularMapParsed)
}

export const getCacheOf = (itemName: string): Array<String> | Object | string => {
 return JSON.parse(localStorage.getItem(itemName) ?? "")
}

export const getAll = (): Array<String> | Object | null => {
 const all: string[] = [];
 for (let i = 0; i < localNames.length; i++) {
  const name = localNames[i];
  const item = (localStorage.getItem(name) ?? "")
  if (!item) return null
  // SUPER SEUCURE CHECK
  if (!item.includes("[") && !item.includes("}")) {
   all.push(item)
   continue
  }
  all.push(JSON.parse(item))
 }
 return all
}

export const setAll = (all: string[]) => {
 localStorage.clear()
 for (let i = 0; i < localNames.length; i++) {
  const newData = JSON.stringify(all[i])
  localStorage.removeItem(localNames[i])
  localStorage.setItem(localNames[i], newData)
 }
}

// Checa la cache en el localstorage, si la propiedad data del react-query no tiene nada pues 
// checamos la local
export const checkBothCache = (response: UseQueryResult<any, Error>, cacheName: string) => {
 let userData: Array<String> | Object | null
 if (response.data == undefined) {
  userData = getCacheOf(cacheName)
 } else {
  userData = response.data[cacheName]
 }
 return userData
}

// xd DEPRECATED
const removeAll = () => {
 for (const name of localNames) {
  localStorage.removeItem(name)
 }
}


// xd DEPRECATED
export const setCurrentUser = ({ username, password }) => {
 localStorage.removeItem("identifier")
 localStorage.removeItem("password")
 localStorage.setItem("identifier", username)
 localStorage.setItem("password", password)
}

export default getUser

