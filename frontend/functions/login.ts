const fetchSearch = ({ queryKey }) => {
    const data = queryKey[1]
    const username = data.username
    const password = data.password
    console.log(username)
    console.log(password)
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
    fetch("http://localhost:8080/login-user", requestOptions)
        .then((response) => {
            return response.text()
        })
        .then((result) => console.log(result))
        .catch((error) => console.error(error));
}
export default fetchSearch
