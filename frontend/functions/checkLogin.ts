import { useUser } from "../components/UserContext"
import { checkBothCache } from "./store"
import useLogin from "./useLogin"

export const checkCache = (responseType: number) => {
 const user = useUser()
 const response = useLogin(user)
 // console.log(response.data.classroom)
 // console.log(response.data.classroom.length)
 const cache = checkBothCache(response, responseType)
 if (!cache) return []
 return cache as Object[] | Object
}
