const API = "http://localhost:8080"
const login = async (userData = { username: "", password: "" }) => {
    const headers = new Headers();
    headers.append("Content-Type", "application/x-www-form-urlencoded");

    const urlencoded = new URLSearchParams();
    urlencoded.append("username", userData.username);
    urlencoded.append("password", userData.password);

    const requestOptions: RequestInit = {
        method: "POST",
        headers: headers,
        body: urlencoded,
        redirect: "follow"
    };
    fetch(`${API}/login-user`, requestOptions)
        .then((response) => response.text())
        .then((result) => console.log(result))
        .catch((error) => console.error(error));
}
