const fetchSearch = async ({ queryKey }) => {
    console.log("estamosw los que estamos")
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
    const apiResponse = await fetch("http://localhost:8080/login-user", requestOptions)
    if (!apiResponse.ok) {
        throw new Error(`Credentials ${username}, ${password} not okay`);
    }

    return apiResponse.json()
    // .then((result) => {
    //     console.log("aqui andamos")
    //     console.log(result)
    // })
    // .catch((error) => console.error(error));
}
export default fetchSearch;
