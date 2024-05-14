import UserData from "../../classes/UserData"

const nameSpace = "moodle"

export const refetchMoodle = (userParams: UserData, setUserParams: React.Dispatch<React.SetStateAction<UserData | null>>) => {
 setUserParams(userParams)
 console.log(nameSpace)
}
