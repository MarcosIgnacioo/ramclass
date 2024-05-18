import { BASE_PATH } from "../../globals/globals";

export async function getTerms(setService: React.Dispatch<React.SetStateAction<undefined>>, setPrivacy: React.Dispatch<React.SetStateAction<undefined>>) {
 const requestOptions: RequestInit = {
  method: "GET",
  redirect: "follow"
 };

 fetch(`${BASE_PATH}/terms-of-privacy`, requestOptions).then((res) => res.json())
  .then((d) => setPrivacy(d))
 fetch(`${BASE_PATH}/terms-of-service`, requestOptions).then((res) => res.json())
  .then((d) => setService(d))

}
