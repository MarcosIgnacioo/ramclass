import { useState, useEffect } from "react";
import UserData from "../classes/UserData";
import { logOut } from "./logOut";
import getUser from "./store";

export default function useNavBarEffect(userQuery: UserData | null) {
 const [navBarData, setNavBarData] = useState({});
 useEffect(() => {
  const userLocal = getUser()
  if (!userLocal && (userQuery?.username == "" && userQuery?.password == "")) {
   setNavBarData({
    "sign-in": {
     "text": "Iniciar sesión",
     "function": ""
    },
    "FAQ": {
     "text": "faq",
     "function": ""
    }
   })
  } else {
   setNavBarData({
    "home": {
     "text": "Inicio",
     "function": ""
    },
    "student": {
     "text": "Mi credencial",
     "function": ""
    },
    "my-kardex": {
     "text": "Mi Kardex",
     "function": ""
    },
    "my-curricular-map": {
     "text": "Mapa Curricular",
     "function": ""
    },
    "/": {
     "text": "Cerrar sesión",
     "function": logOut
    },
    "settings": {
     "text": "Configuración",
     "function": ""
    }
   })
  }
 }, [userQuery])
 return navBarData
}
