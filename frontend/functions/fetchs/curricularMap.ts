const nameSpace = "curricular-map"

const curricularMap = async ({ queryKey }) => {
 const data = queryKey[1]
 const { username, password } = data
 if (!username || !password) return

 const headers = new Headers();
 headers.append("Content-Type", "application/x-www-form-urlencoded");

 const urlencoded = new URLSearchParams();
 urlencoded.append("username", username);
 urlencoded.append("password", password);

 const requestOptions: RequestInit = {
  method: "POST",
  headers: headers,
  body: urlencoded,
  redirect: "follow"
 };

 const apiResponse = await fetch(`http://localhost:8080/${nameSpace}`, requestOptions)
 if (!apiResponse.ok) {
  throw new Error(`Credentials ${username}, ${password} not okay`);
 }

 return apiResponse.json()
}

export default curricularMap;
