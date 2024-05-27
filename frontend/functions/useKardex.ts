import { useQuery } from '@tanstack/react-query';
import UserData from '../classes/UserData';
import kardex from './fetchs/kardex';

const useKardex = (loginParams: UserData | null | undefined) => {
 return useQuery(
  {
   queryKey: ["kardex", loginParams],
   queryFn: kardex,
   enabled: (loginParams !== null),
  }
 )
}

export default useKardex
