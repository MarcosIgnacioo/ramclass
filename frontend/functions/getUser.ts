import UserData from '../classes/UserData'

const localNames = ["moodle", "classroom", "kardex", "curricular_map", "student"]

const getUser = (): UserData | null => {
    const identifier = JSON.parse(localStorage.getItem("identifier") ?? "")
    const password = JSON.parse(localStorage.getItem("password") ?? "")
    if (identifier === "" || password === "") return null
    return new UserData(identifier, password)
}

const getAll = (): Array<String> => {
    const all: string[] = [];
    for (const name of localNames) {
        all.push(JSON.parse(localStorage.getItem(name) ?? ""))
    }
    return all
}

const setAll = (all: string[]) => {
    for (let i = 0; i < localNames.length; i++) {
        const newData = JSON.stringify(all[i])
        localStorage.removeItem(localNames[i])
        localStorage.setItem(localNames[i], newData)
    }
}

const removeAll = () => {
    for (const name of localNames) {
        localStorage.removeItem(name)
    }
}

const getMoodle = (): Array<String> => {
    return JSON.parse(localStorage.getItem("moodle") ?? "")
}


// const setMoodle = (moodle: string) => {
//     moodle = JSON.parse(moodle)
//     localStorage.setItem("moodle", moodle)
// }
//
//
// const setClassroom = (classRoom: string) => {
//     classRoom = JSON.parse(classRoom)
//     localStorage.setItem("classroom", classRoom)
// }
//
// const getClassroom = (): Array<String> => {
//     return JSON.parse(localStorage.getItem("classroom") ?? "")
// }
//
// const setKardex = (kardex: string) => {
//     kardex = JSON.parse(kardex)
//     localStorage.setItem("kardex", kardex)
// }
//
// const getKardex = (): Array<String> => {
//     return JSON.parse(localStorage.getItem("kardex") ?? "")
// }

export default getUser
