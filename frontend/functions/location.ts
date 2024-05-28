import { useLocationUpdateContext, useLocationContext } from "../components/UserContext"
import useLocationEffect from "./effects/useLocationEffect"

export default function updateCurrentLocation() {
 console.log("wepk")
 const locationUpdate = useLocationUpdateContext()
 locationUpdate(window.location.pathname)
 const currentLocation = useLocationContext()
 useLocationEffect(currentLocation)
}
