import { useNavigate } from "react-router-dom"

export function logOut(): void {
 localStorage.clear()
 const navigate = useNavigate()
 navigate("/sign-in")
}
