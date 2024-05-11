import { useState } from "react";
import UserData from "../classes/UserData";
import getUser from "./getUser";
import useLogin from "./useLogin";


function getInfo() {
    const [testParams, setTestParams] = useState<UserData | null>(null);
    const response = useLogin(testParams)
    console.log("4")
    console.log(response)
    return response.data
}

export default getInfo
