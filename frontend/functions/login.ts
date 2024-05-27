import { BASE_PATH } from "../globals/globals"

const nameSpace = "login-user"
const login = async ({ queryKey }) => {
 const data = queryKey[1]
 const username = data.username
 const password = data.password
 if (!username || !password) return
 const myHeaders = new Headers();
 myHeaders.append("Content-Type", "application/x-www-form-urlencoded");

 const urlencoded = new URLSearchParams();
 urlencoded.append("username", username);
 urlencoded.append("password", password);

 const requestOptions: RequestInit = {
  method: "POST",
  headers: myHeaders,
  body: urlencoded,
  redirect: "follow"
 };

 const apiResponse = await fetch(`${BASE_PATH}${nameSpace}`, requestOptions)

 if (!apiResponse.ok) {
  throw new Error(`Credentials ${username}, ${password} not okay`);
 }

 return apiResponse.json()
}
export default login;
