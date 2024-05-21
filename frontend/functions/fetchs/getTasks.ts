import { BASE_PATH } from "../../globals/globals";

const nameSpace = "get-tasks"

export default async function getTasks({ queryKey }) {
 const identifier = queryKey[1]
 if (!identifier) return

 const requestOptions: RequestInit = {
  method: "GET",
  redirect: "follow"
 };

 const response = await fetch(`${BASE_PATH}/${nameSpace}?identifier=${identifier}`, requestOptions)
 return response.json()
}
