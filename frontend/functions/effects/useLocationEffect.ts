import { useEffect } from "react";

export default function useLocationEffect(location: string) {
 useEffect(() => {
  location = location.substring(1)
  const oldLocations = document.querySelectorAll(".current-location")
  oldLocations.forEach(oldLocation => {
   oldLocation?.classList.remove("current-location")
  })
  location = (location === "") ? "home-location" : location
  const locationInNavBar = document.querySelectorAll(`.${location}-location`)
  locationInNavBar.forEach(navbarElement => {
   navbarElement?.classList.add("current-location")
  })
 }, [location])
 return null
}
