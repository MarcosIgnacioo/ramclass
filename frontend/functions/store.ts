import { UseQueryResult } from '@tanstack/react-query'
import UserData from '../classes/UserData'

export enum LSK {
 Moodle,
 Classroom,
 Kardex,
 CurricularMap,
 Student,
 Identifier,
 Password
}

const localNames = ["moodle", "classroom", "kardex", "curricular_map", "student", "identifier", "password"]

const getUser = (): UserData | null => {
 let identifier = (localStorage.getItem("identifier") ?? "")
 let password = (localStorage.getItem("password") ?? "")
 if (!identifier || !password) return null
 identifier = JSON.parse(identifier)
 password = JSON.parse(password)
 if (identifier === "" || password === "") return null
 return new UserData(identifier, password)
}

export const storeInLocal = (itemsToStore: Array<Object> | Array<String> | Object | string, itemName: string) => {
 const curricularMapParsed = JSON.stringify(itemsToStore)
 localStorage.removeItem(itemName)
 localStorage.setItem(itemName, curricularMapParsed)
}

export const getCacheOf = (itemName: string): Array<Object> | Array<String> | Object | string | null => {
 let item = localStorage.getItem(itemName)
 if (!item) return null
 return JSON.parse(item)
}

export const getAll = (): Array<Object> | Array<String> | Object | null => {
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

export const checkContextUser = (user: UserData) => {
 return user.username !== "" && user.password !== ""
}

// Checa la cache en el localstorage, si la propiedad data del react-query no tiene nada pues 
// checamos la local
export const checkBothCache = (response: UseQueryResult<any, Error>, cacheName: number) => {
 const nameSpace = localNames[cacheName]
 const cache = getCacheOf(nameSpace)
 if (response.data !== undefined) {
  if (response.data[nameSpace].length != 0) {
   return response.data[nameSpace]
  }
  response.data[nameSpace] = cache
 }
 return cache
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

