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
    "home": {
     "text": "Inicio",
     "function": "",
     "icon": "home"
    },
    "faq": {
     "text": "FAQ",
     "function": "",
     "icon": "help"
    },
    "calendar": {
     "text": "Calendario Escolar",
     "function": "",
     "icon": "calendar_month"
    },
   })
  } else {
   setNavBarData({
    "home": {
     "text": "Inicio",
     "function": "",
     "icon": "home"
    },
    "student": {
     "text": "Mi credencial",
     "function": "",
     "icon": "person"
    },
    "my-kardex": {
     "text": "Mi Kardex",
     "function": "",
     "icon": "sheets_rtl"
    },
    "my-curricular-map": {
     "text": "Mapa Curricular",
     "function": "",
     "icon": "mitre"
    },
    "todo": {
     "text": "TODO App",
     "function": "",
     "icon": "splitscreen_bottom"
    },
    "calendar": {
     "text": "Calendario Escolar",
     "function": "",
     "icon": "calendar_month"
    },
    "settings": {
     "text": "Configuración",
     "function": "",
     "icon": "settings"
    }
   })
  }
 }, [userQuery])
 return navBarData
}
