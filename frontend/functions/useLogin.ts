import { useQuery } from '@tanstack/react-query';
import login from '../functions/login.ts'
import UserData from '../classes/UserData.ts';

const useLogin = (loginParams: UserData | null | undefined) => {
 return useQuery(
  {
   queryKey: ["user", loginParams],
   queryFn: login,
   enabled: ((loginParams?.username !== undefined && loginParams?.password !== undefined) || loginParams !== null),
  }
 )
}


export default useLogin
