const getUser = (): { username: string, password: string } => {
    const identifier = JSON.parse(localStorage.getItem("identifier") ?? "")
    const password = JSON.parse(localStorage.getItem("password") ?? "")
    return {
        username: identifier,
        password: password
    }
}
export default getUser
