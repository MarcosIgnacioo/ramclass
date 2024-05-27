import { useQuery } from '@tanstack/react-query';
import UserData from '../classes/UserData';
import credentials from './fetchs/credentials';

const useCredentials = (loginParams: UserData | null | undefined) => {
 console.log("wepppppphola")
 console.log(loginParams)
 return useQuery(
  {
   queryKey: ["credentials", loginParams],
   queryFn: credentials,
   enabled: (loginParams !== null),
  }
 )
}

export default useCredentials
