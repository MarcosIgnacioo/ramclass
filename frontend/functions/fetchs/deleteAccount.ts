import { BASE_PATH } from "../../globals/globals";

const nameSpace = "student"

export default async function deleteAccount(identifier:string) {
 const headers = new Headers();
 headers.append("Content-Type", "application/x-www-form-urlencoded");

 const urlencoded = new URLSearchParams();
 urlencoded.append("identifier", identifier);

 const requestOptions: RequestInit = {
  method: "DELETE",
  headers: headers,
  body: urlencoded,
  redirect: "follow"
 };

 console.log("QUEVERGA")
 const response = await fetch(`${BASE_PATH}${nameSpace}?identifier=${identifier}`, requestOptions)

 return response.json()
}
