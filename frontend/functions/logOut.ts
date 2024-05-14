export function logOut(): void {
 console.log("haciendo logou")
 localStorage.clear()
 const url = location.href
 console.log(url)
 location.replace("http://localhost:8080/sign-in")
}
