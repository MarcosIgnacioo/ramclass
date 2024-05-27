import { BASE_PATH } from "../../globals/globals";

const nameSpace = "save-tasks"
export default async function saveTasks({ queryKey }) {
 const data = queryKey[1]
 const { tasks, identifier } = data
 if (!tasks || !identifier) return

 const headers = new Headers();
 headers.append("Content-Type", "application/x-www-form-urlencoded");

 const urlencoded = new URLSearchParams();
 const taskToSave = { identifier: identifier, tasks: tasks }
 urlencoded.append("tasks", JSON.stringify(taskToSave));

 const requestOptions: RequestInit = {
  method: "POST",
  headers: headers,
  body: urlencoded,
  redirect: "follow"
 };

 const apiResponse = await fetch(`${BASE_PATH}${nameSpace}`, requestOptions)
 if (!apiResponse.ok) {
  throw new Error(`Credentials ${identifier}, ${tasks} not okay`);
 }
 return apiResponse.json()
}
