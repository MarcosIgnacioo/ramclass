import { BASE_PATH } from "../../globals/globals";

const nameSpace = "student"

export default async function deleteAccount({ queryKey }) {
 const identifier = queryKey[1]
 if (!identifier) return
 const headers = new Headers();
 headers.append("Content-Type", "application/x-www-form-urlencoded");

 const urlencoded = new URLSearchParams();
 urlencoded.append("identifier", "marcosignc_21");

 const requestOptions: RequestInit = {
  method: "DELETE",
  headers: headers,
  body: urlencoded,
  redirect: "follow"
 };

 const response = await fetch(`${BASE_PATH}${nameSpace}?identifier=${identifier}`, requestOptions)

 return response.json()
}
