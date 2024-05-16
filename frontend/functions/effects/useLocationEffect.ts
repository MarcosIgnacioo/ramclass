import { useEffect } from "react";

export default function useLocationEffect(location: string) {
 useEffect(() => {
  location = location.substring(1)
  const oldLocation = document.querySelector(".current-location")
  oldLocation?.classList.remove("current-location")
  const locationInNavBar = document.getElementById(`${location}`)
  if (locationInNavBar !== null) locationInNavBar.classList.add("current-location")
 }, [location])
 return null
}
