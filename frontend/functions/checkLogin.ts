import { useUser } from "../components/UserContext"
import { checkBothCache } from "./store"
import useLogin from "./useLogin"

export const checkCache = (responseType: number) => {
 const user = useUser()
 const response = useLogin(user)
 const cache = checkBothCache(response, responseType)
 if (!cache) return []
 return cache as Object[] | Object
}
