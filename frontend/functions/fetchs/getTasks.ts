import { BASE_PATH } from "../../globals/globals";

const nameSpace = "get-tasks"
export default async function getTasks({ queryKey }) {
 const identifier = queryKey[1]
 if (!identifier) return

 const headers = new Headers();
 headers.append("Content-Type", "application/x-www-form-urlencoded");

 const urlencoded = new URLSearchParams();
 urlencoded.append("identifier", JSON.stringify(identifier));

 const requestOptions: RequestInit = {
  method: "GET",
  headers: headers,
  body: urlencoded,
  redirect: "follow"
 };

 const apiResponse = await fetch(`${BASE_PATH}/${nameSpace}`, requestOptions)
 if (!apiResponse.ok) {
  throw new Error(`Credentials ${identifier}, not okay`);
 }
 return apiResponse.json()
}
