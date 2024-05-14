const nameSpace = "moodle"

const moodle = async ({ queryKey }) => {
 const data = queryKey[1]
 const { username, password } = data
 if (!username || !password) return
 // useQuery is in action scenario
 // El usario ha ingresado por primera vez su cuenta
 // Presiona refetch en el moodle o classroom
 // El refetch tendria que actualizar la cache del useQuery, junto a la del local storage
 // actualizar la cache del useQuery
 // Por que el usuario querria hacer esto?
 // Cuando el login por alguna razon no retorne ninguna tarea (que es algo que puede ocurrir y ya ha pasado con mas frecuencia de la que me gustaria)
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
export default moodle;
