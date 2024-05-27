import { useQuery } from '@tanstack/react-query';
import UserData from '../classes/UserData';
import login from './login';

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
